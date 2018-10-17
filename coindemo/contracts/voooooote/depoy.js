//npm install solc
//npm install web3@^0.19.0 --save
//注意:web3版本

let Web3 = require('web3');

let web3 = new Web3(new Web3.providers.HttpProvider("http://127.0.0.1:8545"));

let fs = require('fs');
let code = fs.readFileSync('Voting.sol').toString();
let solc = require('solc');

let compiledCode = solc.compile(code);

let abiDefinition = JSON.parse(compiledCode.contracts[':Voting'].interface);
//写入ABI文件至本地文件目录
fs.writeFile('Voting.json',JSON.stringify(abiDefinition), {}, function(err) {
    console.log('write ABI file [Voting.json] complete . ');
});

let VotingContract = web3.eth.contract(abiDefinition);
let byteCode = '0x' + compiledCode.contracts[':Voting'].bytecode;
console.log(byteCode);
//调用VotingContract对象的new()方法来将投票合约部署到区块链。new()方法参数列表应当与合约的 构造函数要求相一致。对于投票合约而言，new()方法的第一个参数是候选人名单。
let deployedContract = VotingContract.new(['Rama','Nick','Jose'],{data: byteCode, from: web3.eth.accounts[0], gas: 4700000});

let contractInstance = VotingContract.at('0x99bdfb1f4c5d0c227d6cd98cf7a254bfc27c35cc');

contractInstance.voteForCandidate('Rama', {from: web3.eth.accounts[0]});
contractInstance.voteForCandidate('Nick', {from: web3.eth.accounts[0]});
contractInstance.voteForCandidate('Jose', {from: web3.eth.accounts[0]});

console.log("--------------");
let RamaVote=contractInstance.totalVotesFor.call('Rama');
let NickVote=contractInstance.totalVotesFor.call('Nick');
let JoseVote=contractInstance.totalVotesFor.call('Jose');
console.log("Rama's vote is "+RamaVote);
console.log("Nick's vote is "+NickVote);
console.log("Jose's vote is "+JoseVote);

