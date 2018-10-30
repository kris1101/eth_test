pragma solidity ^0.4.23;

contract ShrimpFarmer{
    uint256 public EGGS_TO_HATCH_1SHRIMP = 86400;// 通过一整天的时间（以秒为单位）计算孵化虾籽的情况
    uint256 public STARTING_SHRIMP = 300; // 开始给一个人的虾的数量
    uint256 PSN = 10000;
    uint256 PSNH = 5000;
    bool public initialized = false; // 是否完成初始化
    address public ceoAddress; // CEO的地址
    mapping (address => uint256) public hatcheryShrimp; // 正在下籽的虾的数量
    mapping (address => uint256) public claimedEggs; // 保存用户购买的虾籽和推荐得到的虾籽总数，但是不包含池塘中的虾产下来的虾籽
    mapping (address => uint256) public lastHatch; // 最后一次操作的时间
    mapping (address => address) public referrals; // 推荐人
    uint256 public marketEggs; // 市场虾籽数的评估指标

    constructor () public{
        ceoAddress = msg.sender;
    }

    function hatchEggs(address ref) public{
        require(initialized); // 需要完成平台初始化的过程
        if(referrals[msg.sender] == 0 && referrals[msg.sender] != msg.sender){ // 当自己没有推荐人，并且自己保存的推荐人不是自己，进行更新；QY，代码有可能写错了，应该是ref!=msg.sender
            referrals[msg.sender] = ref; // 设置推荐人
        }
        uint256 eggsUsed = getMyEggs(); // 得到自己的虾籽的数量
        uint256 newShrimp = SafeMath.div(eggsUsed, EGGS_TO_HATCH_1SHRIMP); // 一天的秒数作为除数，计算孵化出来的虾的数量
        hatcheryShrimp[msg.sender] = SafeMath.add(hatcheryShrimp[msg.sender], newShrimp); //
        claimedEggs[msg.sender] = 0;
        lastHatch[msg.sender] = now; // 记录最后一次孵化虾籽的时间为当前时间

        //send referral eggs
        claimedEggs[referrals[msg.sender]] = SafeMath.add(claimedEggs[referrals[msg.sender]], SafeMath.div(eggsUsed, 5)); // 推荐者获得使用的虾籽的20%

        //boost market to nerf shrimp hoarding
        marketEggs = SafeMath.add(marketEggs, SafeMath.div(eggsUsed,10)); // 增加市场上虾籽的数量，增加消耗掉的虾籽数量除以10
    }

    function getMyEggs() public view returns(uint256){ // 得到调用者虾籽的数量，由两部分组成，自己可以claim的虾的数量和距离上次新生产出的虾籽的数量
        return SafeMath.add(claimedEggs[msg.sender], getEggsSinceLastHatch(msg.sender));
    }

    function getEggsSinceLastHatch(address adr) public view returns(uint256){
        uint256 secondsPassed = min(EGGS_TO_HATCH_1SHRIMP, SafeMath.sub(now, lastHatch[adr])); // 记录经过了多久的虾的统计时间
        return SafeMath.mul(secondsPassed, hatcheryShrimp[adr]); // 每一秒都能产生一个新的虾籽
    }

    function min(uint256 a, uint256 b) private pure returns (uint256) { // 比较两个数大小
        return a < b ? a : b;
    }

}

library SafeMath {

  function mul(uint256 a, uint256 b) internal pure returns (uint256) {
    if (a == 0) {
      return 0;
    }
    uint256 c = a * b;
    assert(c / a == b);
    return c;
  }

  function div(uint256 a, uint256 b) internal pure returns (uint256) {
    uint256 c = a / b;
    return c;
  }

  function sub(uint256 a, uint256 b) internal pure returns (uint256) {
    assert(b <= a);
    return a - b;
  }

  function add(uint256 a, uint256 b) internal pure returns (uint256) {
    uint256 c = a + b;
    assert(c >= a);
    return c;
  }
}
