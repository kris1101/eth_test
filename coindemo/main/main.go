package main

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"gowallet3/mypxc"
	//"coindemo/mypxc"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/core/types"
	"strings"
	"os"
	"io/ioutil"
	"math/big"
	"fmt"
	"log"
)

// 密钥
const KEYDIR = "/Users/Weelin/private_ent/eth01/node1/keystore/"

// 账户密码
const SECRET = "abcdef"

// 合约部署后的地址
const CDDR = "0x6eb5E517a6D0Ea77664168C456ccd91574809D26"

const URL = "/Users/Weelin/private_ent/eth01/node1/geth.ipc"
const URL1 = "http://localhost:8545"

var (
	owner   string = "0x740eec88db4e2fb8cc50eae456d108775aa3a33f"
	u1      string = "0x740eec88db4e2fb8cc50eae456d108775aa3a33f"
	u2      string = "0xc0ad31f8d2e1e9532a9a3966b4c5553229c4230e"
	u3      string = "0xd7fb2ebd2037ff0959f6513343661772d279a7f5"
	u4      string = "0x5ab6a67381a913fd071bd477148aed23b3c02c7a"
	session *mypxc.PxcSession
)

func init() {
	cli, err := GetCli(URL)
	if err != nil {
		log.Fatalf("Failed during init ,cannot get cli %v", err)
		panic(err)
	}
	token, err := mypxc.NewPxc(common.HexToAddress(CDDR), cli)

	keyFilePath, err := GetKeyFile(owner, KEYDIR)
	if err != nil {
		log.Fatalf("Failed during init ,cannot get keyfilepath %v", err)
		panic(err)
	}
	fp, err := os.Open(keyFilePath)
	if err != nil {
		log.Fatalf("Failed during init ,cannot open keyfile %v", err)
		panic(err)
	}
	auth, err := bind.NewTransactor(fp, SECRET)
	if err != nil {
		panic(err)
		log.Fatalf("Failed during init ,cannot create auth %v", err)
	}
	session = &mypxc.PxcSession{
		Contract: token,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
		TransactOpts: bind.TransactOpts{
			From:   auth.From,
			Signer: auth.Signer,
			//GasLimit: 3141592,
		},
	}
	// 代币名称
	//fmt.Println(session.Name())
}

/**
根据账户获取私钥文件
 */
func GetKeyFile(addr, keydir string) (string, error) {

	rds, err := ioutil.ReadDir(keydir)
	for _, rd := range rds {
		if !rd.IsDir() && strings.HasSuffix(rd.Name(), strings.ToLower(string(addr[2:]))) {
			return KEYDIR + rd.Name(), err
		}
	}
	return "", err
}

/**
打印交易信息
 */
func printTx(tx *types.Transaction, params ...interface{}) {
	if len(params) > 0 {
		fmt.Printf("transfer pending,交易信息如下:\n执行合约: %s, 交易hash: %s, 花费(wei):%d, gas数量:%d, gas单价:%d,nonce:%d, value(amount): %d \n", tx.To(), tx.Hash().Hex(), tx.Cost().Uint64(), tx.Gas(), tx.GasPrice().Uint64(), tx.Nonce(), tx.Value())
	} else {
		fmt.Printf("transfer pending,交易信息如下:\n执行合约: %s, 交易hash: %s, 花费(wei):%d, gas数量:%d, gas单价:%d,nonce:%d, value(amount): %d \n", tx.To().Hex(), tx.Hash().Hex(), tx.Cost().Uint64(), tx.Gas(), tx.GasPrice().Uint64(), tx.Nonce(), tx.Value())
	}
}

/**
获取连接
 */
func GetCli(url string) (cli *ethclient.Client, err error) {
	// 建立连接
	cli, err = ethclient.Dial(url)
	return cli, err
}

/**
部署合约
 */
func deployContract(deployAddr, passwd string, params ...interface{}) (common.Address, *types.Transaction, *mypxc.Pxc, error) {
	keyfilenamepath, err := GetKeyFile(owner, KEYDIR)
	fp, err := os.Open(keyfilenamepath)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	defer fp.Close()

	auth, err := bind.NewTransactor(fp, passwd)
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	backend, err := GetCli(URL)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	defer backend.Close()
	addr, tx, pxc, _ := mypxc.DeployPxc(auth, backend, big.NewInt(20000000000), common.HexToAddress(deployAddr))
	log.Println("合约地址: ", addr.Hex())
	printTx(tx, "deploy contract")
	return addr, tx, pxc, nil
}

/**
获取总的币的数量(可以使用session实现)
 */
func getTotalSupply(contractAddr string) (int64, error) {
	cli, err := GetCli(URL)
	if err != nil {
		return 0, err
	}
	defer cli.Close()

	pxcCaller, err := mypxc.NewPxcCaller(common.HexToAddress(contractAddr), cli)
	if err != nil {
		return 0, err
	}

	total, err := pxcCaller.TotalSupply(nil)
	if err != nil {
		return 0, err
	}
	return total.Int64(), nil

	// session实现
	//total, err := session.TotalSupply()
	//if err != nil {
	//	return 0, err
	//}
	//return total.Int64(), nil
}

/**
获取指定账户的balance(可以使用session实现)
 */
func getAddrBalanceOfCDDR(contractAddr string, addr string) (int64, error) {
	cli, err := GetCli(URL)
	if err != nil {
		return 0, err
	}
	defer cli.Close()

	pxcCaller, err := mypxc.NewPxcCaller(common.HexToAddress(contractAddr), cli)
	if err != nil {
		return 0, err
	}

	var callOpts *bind.CallOpts
	//callOpts = &bind.CallOpts{
	//false, //
	//common.HexToAddress(u1),
	//nil,
	//}
	bigT, err := pxcCaller.BalanceOf(callOpts, common.HexToAddress(addr))
	if err != nil {
		return 0, err
	}
	return bigT.Int64(), nil
}

/**
A->B转账(可以用session实现)
 */
func Transfer(contractAddr, ownerAddr, fromPass, toAddr string, count int64) (tx *types.Transaction, err error) {
	tAddr := common.HexToAddress(toAddr)
	tcount := big.NewInt(count)
	cAddr := common.HexToAddress(contractAddr)
	defer func() {
		if err != nil {
			log.Fatalf("Failed to transfer from %s to %s, because %v", owner, toAddr, err)
		}
	}()
	cli, err := GetCli(URL)
	if err != nil {
		return nil, err
	}
	defer cli.Close()
	keyFilePath, err := GetKeyFile(ownerAddr, KEYDIR)
	if err != nil {
		return nil, err
	}
	fp, err := os.Open(keyFilePath)
	if err != nil {
		return nil, err
	}

	auth, err := bind.NewTransactor(fp, fromPass)
	if err != nil {
		return nil, err
	}
	// 构建一个BoundContract,(cli实现了接口)
	token, err := mypxc.NewPxcTransactor(cAddr, cli)
	if err != nil {
		return nil, err
	}
	tx, err = token.Transfer(auth, tAddr, tcount)
	if err != nil {
		return nil, err
	}
	return tx, err
}

/**
A 把 B 的token转给 C(不可用session)
 */
func TransferFromTo(contractAddr, spenderAddr, spenderPass, fromAddr, toAddr string, count int64) (tx *types.Transaction, err error) {
	cAddr := common.HexToAddress(contractAddr)
	tAddr := common.HexToAddress(toAddr)
	fAddr := common.HexToAddress(fromAddr)
	tcount := big.NewInt(count)
	defer func() {
		if err != nil {
			log.Fatalf("Failed when spender %s transfer from %s to %s, because %v", spenderAddr, fromAddr, toAddr, err)
		}
	}()
	cli, err := GetCli(URL)
	if err != nil {
		return nil, err
	}
	defer cli.Close()
	keyFilePath, err := GetKeyFile(spenderAddr, KEYDIR)
	if err != nil {
		return nil, err
	}
	fp, err := os.Open(keyFilePath)
	if err != nil {
		return nil, err
	}

	auth, err := bind.NewTransactor(fp, spenderPass)
	if err != nil {
		return nil, err
	}

	// 构建一个BoundContract,(cli实现了接口)
	token, err := mypxc.NewPxcTransactor(cAddr, cli)
	if err != nil {
		return nil, err
	}
	//, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	tx, err = token.TransferFrom(auth, fAddr, tAddr, tcount)
	if err != nil {
		return nil, err
	}
	return tx, err
}

/**
授权A 可以把 B 的token 转发给 C(不可用session)
 */
func approveToSpenderCount(contractAddr, ownerAddr, ownerPass, spenderAddr string, count int64) (tx *types.Transaction, err error) {
	cAddr := common.HexToAddress(contractAddr)
	sAddr := common.HexToAddress(spenderAddr)
	scount := big.NewInt(count)
	defer func() {
		if err != nil {
			log.Fatalf("Failed when owner: %s approve spender: %s count: %d\n, because %v", ownerAddr, spenderAddr, count, err)
		}
	}()
	cli, err := GetCli(URL)
	if err != nil {
		return nil, err
	}
	defer cli.Close()
	keyFilePath, err := GetKeyFile(ownerAddr, KEYDIR)
	if err != nil {
		return nil, err
	}
	fp, err := os.Open(keyFilePath)
	if err != nil {
		return nil, err
	}

	auth, err := bind.NewTransactor(fp, ownerPass)
	if err != nil {
		return nil, err
	}

	// 构建一个BoundContract,(cli实现了接口)
	token, err := mypxc.NewPxcTransactor(cAddr, cli)
	if err != nil {
		return nil, err
	}
	tx, err = token.Approve(auth, sAddr, scount)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

/**
查看授权,查看owner允许spender转给其它账户的数目(不可用session)
 */
func getallowanceOfAToB(contractAddr string, ownerAddr, spenderAddr string) (int64, error) {
	oAddr := common.HexToAddress(ownerAddr)
	sAddr := common.HexToAddress(spenderAddr)
	cli, err := GetCli(URL)
	if err != nil {
		return 0, err
	}
	defer cli.Close()

	pxcCaller, err := mypxc.NewPxcCaller(common.HexToAddress(contractAddr), cli)
	if err != nil {
		return 0, err
	}

	allowCount, err := pxcCaller.Allowance(nil, oAddr, sAddr)
	if err != nil {
		return 0, err
	}
	return allowCount.Int64(), nil
}

/**
获取所有账户信息
 */
func showBalanceOfAccounts(contractAddr string) {
	total, err := getTotalSupply(contractAddr)
	if err != nil {
		panic(err)
	}
	fmt.Printf("合约 %s, totalSupply: %d\n", contractAddr, total)
	users := []string{u1, u2, u3, u4}
	for index, u := range users {
		bigT, err := session.BalanceOf(common.HexToAddress(u))
		if err != nil {
			panic(err)
		} else {
			fmt.Printf("u%d balanceof %s: %d\n", index+1, contractAddr, bigT)
		}
	}
}

func main() {
	showBalanceOfAccounts(CDDR)
	// u1-->u3, 使用 &mypxc.PxcSession
	//tx, err := session.Transfer(common.HexToAddress(u3), big.NewInt(3000))
	//if err != nil {
	//	panic(err)
	//}
	//printTx(tx)
	//
	//lastNonce:=tx.Nonce()
	//session.TransactOpts.Nonce = big.NewInt(int64(lastNonce+1))
	//tx, err = session.Transfer(common.HexToAddress(u2), big.NewInt(3000))
	//if err != nil {
	//	panic(err)
	//}
	//printTx(tx)

	// u1-->u4, 使用 *mypxc.pxcTransactor
	//tx, err:=Transfer(CDDR, u1, SECRET, u4, 1000)
	//if err != nil {
	//	panic(err)
	//}
	//printTx(tx)

	//u3向u1授权300的量
	//tx, err:=approveToSpenderCount(CDDR, u3, SECRET, u1, 300)
	//if err!=nil{
	//	panic(err)
	//}
	//printTx(tx)
	//
	// 获取u3向u1的授权量
	//allowCount, err:=getallowanceOfAToB(CDDR, u3, u1)
	//if err!=nil{
	// panic(err)
	//}else{
	//	fmt.Println(allowCount)
	//}

	// A transfer B's token to C
	//tx, err := TransferFromTo(CDDR, u1, SECRET, u3, u4, 200)
	//if err != nil {
	// panic(err)
	//}
	//printTx(tx)
}
