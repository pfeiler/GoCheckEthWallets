package main

import(
	"log"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"context"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"encoding/hex"
	"os"
	"fmt"
)

// yeah, its a infura access token, use it as you want or replace with yours
const EthServiceUrl = "https://mainnet.infura.io/6pgDwimM6ABLKOxTxa2j"

func main() {
	log.Printf("Start checking addresses ...")
	count := 0
	for {
		checkRandomAddress()
		count = count + 1
		if count == 100 {
			log.Printf("Checked 100 addresses")
			count = 0
		}
	}
}

func checkRandomAddress() {
	address, key := getRandomAddressAndKey()
	balance := checkBalance(address)
	// log.Printf("address=%s key=%s balance=%v", address.Hex(), key, balance)
	if balance.Sign() != 0 {
		alertBalance(balance, address, key)
	}
}

func getRandomAddressAndKey() (common.Address, string) {
	key, _ := crypto.GenerateKey()
	keyString := hex.EncodeToString(key.D.Bytes())
	address := crypto.PubkeyToAddress(key.PublicKey)
	return address, keyString
}

func checkBalance(address common.Address) *big.Int {
	// create an IPC based RPC connection to a remote node
	conn, err := ethclient.Dial(EthServiceUrl)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
		os.Exit(-1)
	}
	// instantiate the contract and display its name
	// var ctx, _ = context.WithTimeout(context.Background(), 10000)
	balance, err := conn.BalanceAt(context.TODO(), address, nil)
	if err != nil {
		log.Fatalf("Error while getting a balance: %v", err)
		os.Exit(-1)
	}

	return balance
}

func alertBalance(balance *big.Int, address common.Address, key string) {
	info := fmt.Sprintf("Found balance %s with address %s and key %s", balance, address.Hex(), key)
	log.Printf("################################################")
	log.Printf(info)
	log.Printf("################################################")

	f, _ := os.OpenFile("balances.txt", os.O_APPEND|os.O_WRONLY, 0755)
	fmt.Fprintf(f, "%s\n", info)
	defer f.Close()

}

