package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
	"golang.org/x/text/currency"
)

func main() {
	if err := runMain(); err != nil {
		log.Fatalln("Error while executing runMain(): ", err)
		os.Exit(1)
	}
}

func runMain() error {
	fmt.Println("YNAB baby budjet")
	date := time.Now().UTC()
	fmt.Printf("x: %v\n", date)
	return nil
}

type Budjet interface {
	NewAccount() error
	DeleteAccount() error
	RenameAccount() error
	DisplayAccount() error
	UpdateAccount() error
}
type Transaction struct {
	uuid               uuid.UUID
	date               time.Time
	outflow            Money
	inflow             Money
	amount             Money
	transactionType    TransactionType
	time               Time
	party              Party
	memoTransfer       string
	account            BudjetAccount
	memo               string
	status             StatusTransaction
	fromcategory       Category
	isCategoryTransfer bool
	tocategory         Category
}

type Status struct {
	pending      bool
	consolidated bool
	checked      bool
}
type StatusTransaction struct{ status Status }

type Party struct {
	uuid uuid.UUID
	name string
	bank Bank
}
type User struct {
	incomeMonthly   Money
	avalaibleBudget Money
}

type Money struct{ amount currency.Amount }
type BudjetBalance struct {
	actual  Money
	pending Money
	settled Money
}

type BudjetAccount struct {
	uuid     uuid.UUID
	name     string
	bank     Bank
	balance  Balance
	time     Time
	currency currency.Unit
	group    []Group
}
type Group []Category
type Category struct {
	uuid     uuid.UUID
	id       int64
	name     string
	time     Time
	target   Target
	balance  currency.Amount
	isHidden bool
}
type Target struct {
	amount   currency.Amount
	isActive bool
	time     Time
}
type Corpus struct {
	Assets    int64
	Liability int64
	Accounts  []BudjetAccount
}
type Bank struct {
	uuid          uuid.NullUUID
	name          string
	accountNumber int64
	accountType   AccountType
}
type AccountType struct {
	savings    bool
	checking   bool
	creditCard bool
	cash       bool
}

type Time struct {
	created  time.Time
	modified time.Time
}
type Balance struct {
	opening int64
	current int64
	closing int64
}
type TransactionType struct {
	credit int64
	debit  int64
}

func (a *BudjetAccount) NewAccount() *BudjetAccount {
	return func() *BudjetAccount { return a }()
}
