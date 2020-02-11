package main

import (
	"fmt"
	"github.com/appliedgocourses/bank"
	"os"
	"strings"
)

func usage() {
	fmt.Println(`Usage:
	bank create <name>						Create an account.
	bank list								List all accounts.
	bank update <name> <amount>				Deposit or withdraw  money
	bank transfer <name> <name> <amount>	Transfer money between two accounts.
	bank history <name>						Show an account's Transaction history.
	
	`)
}

func main () {
	if len(os.Args) < 2 {
		fmt.Println("No command found")
		usage()
		return
	}

	switch os.Args[1] {

	case "create":
		// check if the name argument was passed
		if len(os.Args) <= 2 {
			fmt.Println("<name> argument not passed")
			usage()
			return
		}
		accountName := os.Args[2]
		// load the data
		err := bank.Load()
		if err != nil {
			fmt.Println("Unable to load data set")
			return
		}
		// Get all list of accounts
		accountsList := bank.ListAccounts()
		// compare if accountName exist
		if strings.Contains(accountsList, accountName) {
			fmt.Println("Oops, Account name already exist")
			usage()
			return
		}
		// create a newAccount
		_ = bank.NewAccount(accountName)
		errs := bank.Save()
		if errs != nil {
			fmt.Println("Oops, Unable to save ", errs)
		}

	case "list":
		println("list all account")
	case "update":
		println("deposit or withdraw money")
	case "transfer":
		println("transfer money between two accounts")
	case "history":
		println("show an account's transaction history")
	default:
		fmt.Println("Invalid command")
		usage()

	}
}


