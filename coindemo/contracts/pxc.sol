pragma solidity^0.4.21;

contract ERC20 {
      function totalSupply() constant returns (uint totalSupply);
      function balanceOf(address _owner) constant returns (uint balance);
      function transfer(address _to, uint _value) returns (bool success);
      function transferFrom(address _from, address _to, uint _value) returns (bool success);
      function approve(address _spender, uint _value) returns (bool success);
      function allowance(address _owner, address _spender) constant returns (uint remaining);
      event Transfer(address indexed _from, address indexed _to, uint _value);
      event Approval(address indexed _owner, address indexed _spender, uint _value);
    }

contract pdjtoken is ERC20 {
    string public name = "pdjcoin";
    string public sym = "pxc";
    uint _totalSupply;
    address public foundation;
    mapping(address => uint) balances;
    mapping(address => mapping(address=>uint)) allowed;
    function pdjtoken( uint totalSupply, address addr ) public {
        _totalSupply = totalSupply;
        foundation   = addr;
        balances[foundation] = totalSupply;
    }
    function totalSupply() constant returns (uint totalSupply) {
        totalSupply = _totalSupply;
        return totalSupply;
    }
    function balanceOf(address _owner) constant returns (uint balance) {
        balance = balances[_owner];
    }
    function transfer(address _to, uint _value) returns (bool success) {
        if( balances[msg.sender] >= _value && 
            _value >0 &&
            balances[_to] + _value > balances[_to] ) {
            balances[_to] += _value;
            balances[msg.sender] -= _value;
            emit Transfer(msg.sender, _to, _value);
            return true;
        } else {
            return false;
        }
    }
    function transferFrom(address _from, address _to, uint _value) returns (bool success) {
        if( allowed[_from][_to] >= _value && 
            _value >0 &&
            balances[_to] + _value > balances[_to] 
        ) {
            balances[_to] += _value;
            balances[_from] -= _value;
            allowed[_from][_to] -= _value;
            emit Transfer(_from, _to, _value);
            return true;
        } else {
            return false;
        }
    }
    function approve(address _spender, uint _value) returns (bool success) {
        if( balances[msg.sender] >= _value &&
            _value > 0 &&
            allowed[msg.sender][_spender] + _value > 0 
        ) {
            allowed[msg.sender][_spender] = _value;
            emit Approval(msg.sender, _spender, _value);
            return true;
        } else {
            return false;
        }
    }
    function allowance(address _owner, address _spender) constant returns (uint remaining) {
        remaining = allowed[_owner][_spender];
    }
}
