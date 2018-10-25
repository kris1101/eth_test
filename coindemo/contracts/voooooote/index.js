//npm install express
var express = require('express');
var app = express();
var server = require('http').createServer(app);
var Web3 = require("web3");

app.use(express.static('.'));

let web3 = new Web3(new Web3.providers.HttpProvider("http://localhost:8545"));

let fs = require('fs');
let code = fs.readFileSync('Voting.sol').toString();
let solc = require('solc');

let compiledCode = solc.compile(code);

let abi = JSON.parse(compiledCode.contracts[':Voting'].interface);

//let abi = JSON.parse('[{"constant":true,"inputs":[{"name":"candidate","type":"bytes32"}],"name":"totalVotesFor","outputs":[{"name":"","type":"uint8"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"candidate","type":"bytes32"}],"name":"validCandidate","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"","type":"bytes32"}],"name":"votesReceived","outputs":[{"name":"","type":"uint8"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"","type":"uint256"}],"name":"candidateList","outputs":[{"name":"","type":"bytes32"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"candidate","type":"bytes32"}],"name":"voteForCandidate","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"inputs":[{"name":"candidateNames","type":"bytes32[]"}],"payable":false,"stateMutability":"nonpayable","type":"constructor"}]')
let VotingContract = web3.eth.contract(abi);
//这里要替换成你自己的地址
let contractInstance = VotingContract.at('0xC9D4D622936BeaDc39Fbf38cc1c2E866875D226A');

app.get("/totalVotesFor", function(req, res) {
	var voteName = req.query.voteName;
	var vote_num = contractInstance.totalVotesFor.call(voteName).toString();
	console.log(vote_num);
	res.send(vote_num);
});

app.get("/voteForCandidate", function(req, res) {
	var voteName = req.query.voteName;
	contractInstance.voteForCandidate(voteName, {from: web3.eth.accounts[0]});
	var vote_num = contractInstance.totalVotesFor.call(voteName).toString();
	res.send(vote_num);
});

server.listen(3000);

// 控制台会输出以下信息
console.log('Server running at http://127.0.0.1:3000/index.html');
