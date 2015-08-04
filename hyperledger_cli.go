package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Printf("Hyperledger CLI\n\n")
	fmt.Printf("Available commands: \n")
	fmt.Printf("hyperledger asset register \t[assetId]\t[websiteURL]\n")

	// register account
	fmt.Printf("hyperledger createuser\t[ledgerId]\t[public key]\n")

	// transfer units
	fmt.Printf("hyperledger transfer\t[amount]\t[source]\t[destination]\n")

	// help
	fmt.Printf("hyperledger help \t[command]\n")

	// this loop runs the cli engine.
	cliLoop()
}

func cliLoop() {
	var inputReader *bufio.Reader
	var input string
	var err error
	for {
		fmt.Printf(">")
		inputReader = bufio.NewReader(os.Stdin)
		input, err = inputReader.ReadString('\n')
		var results []string = strings.Fields(input)

		if err == nil {
			analyzeInput(results)
		} else {
			os.Exit(1)
		}
	}
}

func analyzeInput(results []string) {
	if results[0] == "hyperledger" {
		var request string = results[1]
		var params []string = results[2:]
		switch request {
		case "asset":
			requestAsset(params)
		case "issue":
			requestIssue(params)
		case "account":
			requestAccount(params)
		case "transfer":
			requestTransfer(params)
		case "help":
			printHelp()
		}
	} else if results[0] == "exit" {
		os.Exit(0)
	} else {
		fmt.Println("Invalid input.\n")
	}
}

// Register Asset should have two parmas
// register asset (ledger)
//   desc 'register <hash>', 'Registers a new asset.
// 		The hash must be an unique hex encoded SHA265 hash.'
// hyperledger register asset nameOfAsset www.somewebpage.org
// fmt.Printf("hyperledger asset register \t[assetId]\t[websiteURL]\n")
func requestAsset(params []string) {
	switch params[0] {
	case "list":
		fmt.Println("# List all known assets")
		// get "#{options[:server]}/assets"
		// todo: api call for list of all assets.
	case "register":
		if len(params) > 2 {
			var assetId string = params[1]
			var websiteURL string = params[2]
			fmt.Println(assetId, websiteURL)
			// todo serialize
		} else {
			printUsageAsset()
		}
	default:
		printUsageAsset()
	}
}
func printUsageAsset() {
	fmt.Println("Commands:")
	fmt.Println("  hyperledger asset help [COMMAND]   # Describe subcommands or one specific subcommand")
	fmt.Println("  hyperledger asset register <hash>  # Registers a new asset. The hash must be an unique hex encoded SHA265 hash.")
	fmt.Println("  hyperledger asset register <legerID> <websiteURL>  # Registers a new asset.")
	fmt.Println("  hyperledger asset list             # List all known assets.")
	// todo : fix the case with argument <hash>!
}

// Create account?
// fmt.Println("hyperledger issue \t[amount]\t[asset]\t[public key]\n")
func requestIssue(params []string) {
	if len(params) < 3 {
		printUsageIssue()
	} else {
		var amount string = params[0]
		var asset string = params[1]
		var publicKey string = params[2]
		fmt.Println(amount, asset, publicKey) // todo remove
		// todo serialize
	}
}
func printUsageIssue() {
	fmt.Println(" Usage:")
	fmt.Println("   hyperledger issue <amount> <asset> <public key>")
	fmt.Println(" Options:")
	fmt.Println("   --server=SERVER")
	fmt.Println("Issue <amount> new units of <asset> authorised by <public key>.")
}

func requestAccount(params []string) {
	if len(params) < 1 {
		fmt.Println("createuser\t[ledgerId]\t[public key]\n")
		printUsageAccount()
	} else {
		var ledgerId string = params[0]
		var publicKey string = params[1]
		fmt.Println(ledgerId, publicKey)
		fmt.Println("Successfully created a user.\n")

		// todo serialize
	}
}
func printUsageAccount() {
	fmt.Println("Commands:")
	fmt.Println("	   hyperledger account help [COMMAND]    # Describe subcommands or one specific subcommand")
	fmt.Println("	   hyperledger account list              # List all known accounts")
	fmt.Println("	   hyperledger account register <asset>  # Registers a new account to hold <asset>.")
}

func requestTransfer(params []string) {
	if len(params) < 3 {
		printUsageTransfer()
	} else {
		//amount, source, destination := params[0],params[1],params[2]
		var amount string = params[0]
		var source string = params[1]
		var destination string = params[2]
		fmt.Println(amount, source, destination)

		// todo serialize
	}
}
func printUsageTransfer() {

	fmt.Println(" Commands:")
	fmt.Println(" Usage:")
	fmt.Println("   hyperledger transfer <amount> <source> <destination>")
	fmt.Println(" Options:")
	fmt.Println("   --server=SERVER")
	fmt.Println(" Transfer <amount> of units from <source> to <destination>")
}

// Documentation for the usage of the cli
func printHelp() {
	fmt.Println("Commands:")
	fmt.Println("	hyperledger asset SUBCOMMAND                          # Subcommands relating to ledgers.")
	fmt.Println("	hyperledger issue <amount> <asset> <public key>       # Issue <amount> new units of <asset> authorised by <public key>.")
	fmt.Println(" 	hyperledger account SUBCOMMAND                        # Subcommands relating to accounts.")
	fmt.Println(" 	hyperledger transfer <amount> <source> <destination>  # Transfer <amount> of units from <source> to <destination>")
	fmt.Println("	hyperledger help [COMMAND]                            # Describe available commands or one specific command")
}
