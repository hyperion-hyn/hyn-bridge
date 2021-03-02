package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/viper"
	"log"
	"math/big"
	"os"
	"time"
)

type contractType byte

const (
	rpToken contractType = iota
	hynStaking
	hynStakingMock
	rpHolding
)

func InitConfig() {
	viper.SetConfigName("config") // config file name without extension
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config/") // config file path
	viper.AddConfigPath("../config/") // config file path
	viper.SetEnvPrefix("rp")
	viper.AutomaticEnv() // read value ENV variable
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("fatal error config file: config.yaml \n", err)
		os.Exit(1)
	}
}

func main() {
	InitConfig()
	//deployContract(rpToken)
	//deployContract(hynStaking)
	//deployContract(hynStakingMock)
	deployContract(rpHolding)
}

func GetRPCUrlFromConfig() string {
	env := viper.GetString("env")
	url := viper.GetString(env + ".url")
	fmt.Printf("url: %v \n", url)
	return url
}

func GetPrivateKeyFromConfig() string {
	key := viper.GetString("private_key")
	return key
}

func deployContract(cType contractType) {
	client, err := ethclient.Dial(GetRPCUrlFromConfig())
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA(GetPrivateKeyFromConfig())
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	//gasPrice, err := client.SuggestGasPrice(context.Background())
	//if err != nil {
	//	log.Fatal(err)
	//}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(8000000) // in units
	auth.GasPrice = common.Big1

	var (
		contractAddr common.Address
		tx           *types.Transaction
	)
	switch cType {
	case rpToken:
		contractAddr, tx, _, err = DeployRPToken(auth, client)
	case hynStaking:
		contractAddr, tx, _, err = DeployHYNStaking(auth, client)
	case hynStakingMock:
		contractAddr, tx, _, err = DeployHYNStakingMock(auth, client)
	case rpHolding:
		contractAddr, tx, _, err = DeployRPHolding(auth, client)
	default:
		log.Fatal("invalid contract type")
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("tx hash: %v \n", tx.Hash().Hex())
	fmt.Printf("contract address: %v \n", contractAddr.Hex())

	time.Sleep(12 * time.Second)

	receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {
		panic(err)
	}

	marshaldate, err := json.MarshalIndent(&receipt, "", "  ")
	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Println(string(marshaldate))
}