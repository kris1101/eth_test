//npm install solc
//npm install web3@^0.19.0 --save
//注意web3版本

let Web3 = require('web3');

let web3 = new Web3(new Web3.providers.HttpProvider("http://127.0.0.1:8545"));

console.log('Initialization web3 complete,the first account is '+ web3.eth.accounts[0]);
let fs = require('fs');
let code = fs.readFileSync('../tokenerc20.sol').toString();
let solc = require('solc');

let compiledCode = solc.compile(code);

console.log('Compile TokenERC20.sol complete');
//部署合约至区块链节点
let abiDefinition = JSON.parse(compiledCode.contracts[':TokenERC20'].interface);
//写入ABI文件至本地文件目录
fs.writeFile('TokenERC20.json',JSON.stringify(abiDefinition), {}, function(err) {
    console.log('write ABI file [TokenERC20.json] complete . ');
});

let VotingContract = web3.eth.contract(abiDefinition);
let byteCode = '0x' + compiledCode.contracts[':TokenERC20'].bytecode;

let deployedContract = VotingContract.new(100, "test", "symbol", {data: byteCode, from: web3.eth.accounts[0], gas: 3600000});
//输出合约 地址，如果此处没有返回地址，可以在Ganache日志中查看到
console.log('deploy complete, deploy address is '+ deployedContract.address);
//let contractInstance = VotingContract.at(deployedContract.address);
let contractInstance = VotingContract.at('0x0D10Dd04A08459C2C3BDd8b7cE5CF26FE14d1E4e');

//测试合约调用
contractInstance.transfer(web3.eth.accounts[1], 100*10**18, {from: web3.eth.accounts[0]});
contractInstance.burn(10*10**18, {from: web3.eth.accounts[0]});
//console.log(contractInstance.totalSupply);
console.log("--------------finish----------------");


