pragma solidity ^0.4.23;

contract Set {
      uint private result = 0;

      constructor () public {

      }

      function assign(uint x, uint y) public returns (uint) {
          result = x + y;
      }

      function get_result() public view returns (uint){
            return result;
      }
}

//每步计算需要消耗gas;
//每步计算后需要开启挖矿更新
//函数直接调用需要初始地址账户，.call后缀调用不需要（本地运行环境）