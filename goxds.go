package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/urfave/cli"
)

func createCliApp() *cli.App {
	app := cli.NewApp()
	app.Name = "GoXDS"
	app.Usage = "Fullnode CLI"
	app.Authors = []*cli.Author{
		&cli.Author{
			Name:  "xandronus",
			Email: "xandronus@gmail.com",
		},
	}
	app.Version = "0.9.2"
	app.Compiled = time.Now()
	app.ArgsUsage = "[walletname]"
	return app
}

func jsonPrettyPrint(in string) string {
	var out bytes.Buffer
	err := json.Indent(&out, []byte(in), "", "  ")
	if err != nil {
		return in
	}
	return out.String()
}

func getWalletName(c *cli.Context) string {
	walletName := "default"
	if c.Args().First() != "" {
		walletName = c.Args().First()
	}
	return walletName
}

func getHostName() string {
	return "http://localhost:48334"
}

func httpGetStakingInfo() string {
	url := getHostName() + "/api/Staking/getstakinginfo"
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return string(body)
}

func httpReceiveAddr(walletname string) string {
	url := getHostName() + "/api/Wallet/unusedaddress?WalletName=" + walletname + "&AccountName=account%200&Segwit=true"
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var address string
	json.Unmarshal(body, &address)
	return address
}

func httpStartStaking(walletname string, password string) string {
	url := getHostName() + "/api/Staking/startstaking"
	fmt.Println(url)

	requestBody, err := json.Marshal(map[string]string{
		"password": password,
		"name":     walletname,
	})
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return string(body)
}

func httpCreateWallet(walletname string, mneumonic string, password string) string {
	url := getHostName() + "/api/wallet/create"
	fmt.Println(url)

	requestBody, err := json.Marshal(map[string]string{
		"mnemonic":   mneumonic,
		"password":   password,
		"passphrase": password,
		"name":       walletname,
	})
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		fmt.Println("OK")
	} else {
		fmt.Println("Failed")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return string(body)
}

func httpStopStaking() string {
	url := getHostName() + "/api/Staking/stopstaking"
	fmt.Println(url)

	requestBody, err := json.Marshal(true)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return string(body)
}

func httpCreatePrivateKey() string {
	url := getHostName() + "/api/Wallet/mnemonic?language=English&wordCount=12"
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return string(body)
}

func httpAddNode(ip string) string {
	url := getHostName() + "/api/ConnectionManager/addnode?endpoint=" + ip + "&command=add"
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return string(body)
}

type recipientRequest struct {
	DestinationAddress string `json:"destinationAddress"`
	Amount             string `json:"amount"`
}

type buildTransactionRequest struct {
	FeeAmount           string             `json:"feeAmount"`
	Password            string             `json:"password"`
	SegwitChangeAddress bool               `json:"segwitChangeAddress"`
	WalletName          string             `json:"walletName"`
	Recipients          []recipientRequest `json:"recipients"`
}

type buildTransactionResponse struct {
	Fee           int64  `json:"fee"`
	Hex           string `json:"hex"`
	TransactionId string `json:"transactionId`
}

func httpBuildTransaction(walletname string, pwd string, sendAddress string, coins string, fee string) string {
	url := getHostName() + "/api/wallet/build-transaction"
	fmt.Println(url)

	transRequest := buildTransactionRequest{
		FeeAmount:           fee,
		Password:            pwd,
		SegwitChangeAddress: true,
		WalletName:          walletname,
		Recipients: []recipientRequest{
			recipientRequest{
				DestinationAddress: sendAddress,
				Amount:             coins,
			},
		},
	}

	requestBody, err := json.Marshal(transRequest)
	fmt.Println("-- Request --")
	fmt.Println(jsonPrettyPrint(string(requestBody)))
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		fmt.Println("OK")
	} else {
		fmt.Println("Failed")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return string(body)
}

func httpSendTransaction(hex string) string {
	url := getHostName() + "/api/wallet/send-transaction"
	fmt.Println(url)

	requestBody, err := json.Marshal(map[string]string{
		"hex": hex,
	})

	fmt.Println("-- Request --")
	fmt.Println(jsonPrettyPrint(string(requestBody)))

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		fmt.Println("OK")
	} else {
		fmt.Println("Failed")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return string(body)
}

func createWalletCommand() *cli.Command {
	return &cli.Command{
		Name:    "wallet-create",
		Aliases: []string{"wc"},
		Usage:   "creates a new wallet",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "password, p",
				Usage:    "wallet password",
				Required: true,
			},
		},
		Category: "Wallet",
		Action: func(c *cli.Context) error {
			walletname := getWalletName(c)
			fmt.Println("Executing 'wallet-create' for walletname:", walletname)
			fmt.Println("Return wallet mnemonic")
			fmt.Println("password:", c.String("password"))
			mneumonic := httpCreatePrivateKey()
			fmt.Println(httpCreateWallet(walletname, strings.Trim(mneumonic, "\""), c.String("password")))
			fmt.Println("-------------------")
			fmt.Println("Record this info")
			fmt.Println("-------------------")
			fmt.Println("WalletName:", walletname)
			fmt.Println("Backup Phrase:", mneumonic)
			fmt.Println("Password:", c.String("password"))
			fmt.Println("Passphrase:", c.String("password"))
			fmt.Println("Creation Date:", time.Now())
			return nil
		},
	}
}

func startStakingCommand() *cli.Command {
	return &cli.Command{
		Name:    "staking-start",
		Aliases: []string{"ss"},
		Usage:   "starts staking",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "password, p",
				Usage:    "wallet password",
				Required: true,
			},
		},
		Category: "Mining",
		Action: func(c *cli.Context) error {
			walletname := getWalletName(c)
			fmt.Println("Executing 'staking-start' for walletname:", walletname)
			httpStartStaking(walletname, c.String("password"))
			fmt.Println("Starting to stake (will take approx 1 min)...")
			time.Sleep(60 * time.Second)
			fmt.Println(jsonPrettyPrint(httpGetStakingInfo()))
			return nil
		},
	}
}

func stopStakingCommand() *cli.Command {
	return &cli.Command{
		Name:     "staking-quit",
		Aliases:  []string{"sq"},
		Usage:    "stops staking",
		Category: "Mining",
		Action: func(c *cli.Context) error {
			fmt.Println("Executing 'staking-quit'")
			httpStopStaking()
			fmt.Println(jsonPrettyPrint(httpGetStakingInfo()))
			return nil
		},
	}
}

func stakingInfoCommand() *cli.Command {
	return &cli.Command{
		Name:     "staking-info",
		Aliases:  []string{"si"},
		Usage:    "get staking info",
		Category: "Mining",
		Action: func(c *cli.Context) error {
			fmt.Println("Executing 'staking-info'")
			fmt.Println(jsonPrettyPrint(httpGetStakingInfo()))
			return nil
		},
	}
}

func receiveWalletCommand() *cli.Command {
	return &cli.Command{
		Name:     "wallet-receive",
		Aliases:  []string{"wr"},
		Usage:    "gets a wallet address to receive funds",
		Category: "Wallet",
		Action: func(c *cli.Context) error {
			walletName := getWalletName(c)
			fmt.Println("Executing 'wallet-receive' address for walletname:", walletName)
			address := httpReceiveAddr(walletName)
			fmt.Println(address)
			return nil
		},
	}
}

func addNodeCommand() *cli.Command {
	return &cli.Command{
		Name:     "node-add",
		Aliases:  []string{"na"},
		Usage:    "adds a peer",
		Category: "Node",
		Action: func(c *cli.Context) error {
			fmt.Println("Executing 'node-add'")
			if c.Args().First() != "" {
				ip := c.Args().First()
				fmt.Println(jsonPrettyPrint(httpAddNode(ip)))
			} else {
				fmt.Println("Error. Argument [ip] is required.")
			}

			return nil
		},
	}
}

func sendWalletCommand() *cli.Command {
	return &cli.Command{
		Name:     "wallet-send",
		Aliases:  []string{"ws"},
		Usage:    "sends coins from wallet to address",
		Category: "Wallet",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "password, p",
				Usage:    "wallet password",
				Required: true,
			},
			&cli.StringFlag{
				Name:  "fee, f",
				Usage: "transaction fee",
			},
		},
		Action: func(c *cli.Context) error {
			fmt.Println("Executing 'wallet-send'")
			walletName := "default"
			sendAddress := ""
			amount := "0.0"
			fee := "0.01"
			if c.NArg() > 2 {
				walletName = c.Args().Get(0)
				amount = c.Args().Get(1)
				sendAddress = c.Args().Get(2)
			} else if c.NArg() > 1 {
				amount = c.Args().Get(0)
				sendAddress = c.Args().Get(1)
			} else {
				fmt.Println("Error. Arguments [sendAddress] [amount] are required.")
				return nil
			}

			if c.String("fee") != "" {
				fee = c.String("fee")
			}

			transRespText := httpBuildTransaction(walletName, c.String("password"), sendAddress, amount, fee)
			fmt.Println("-- Response --")
			fmt.Println(jsonPrettyPrint(transRespText))
			var transResp buildTransactionResponse
			json.Unmarshal([]byte(transRespText), &transResp)

			sendTransRespText := httpSendTransaction(transResp.Hex)
			fmt.Println("-- Response --")
			fmt.Println(jsonPrettyPrint(sendTransRespText))

			return nil
		},
	}
}

func addCommands(app *cli.App) {
	app.Commands = []*cli.Command{
		createWalletCommand(),
		receiveWalletCommand(),
		sendWalletCommand(),
		startStakingCommand(),
		stopStakingCommand(),
		stakingInfoCommand(),
		addNodeCommand(),
	}
}

func main() {
	app := createCliApp()
	addCommands(app)
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
