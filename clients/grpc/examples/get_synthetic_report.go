package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/stone-co/the-amazing-ledger/clients/grpc/ledger"
	"github.com/stone-co/the-amazing-ledger/pkg/command-handler/domain/ledger/entities"
)

func getSyntheticReport(log *logrus.Entry, conn *ledger.Connection) {
	log.Println("starting GetSyntheticReport")
	defer log.Println("finishing GetSyntheticReport")

	// expectedBalance := 1000
	accountPathOne := "liability:stone:clients:" + uuid.New().String()
	accountPathTwo := "liability:stone:clients:" + uuid.New().String()

	log.Println("starting GetSyntheticReport 2")
	// Define a new transaction with 2 entries
	t := conn.NewTransaction(uuid.New())
	t.AddEntry(uuid.New(), accountPathOne, entities.NewAccountVersion, entities.CreditOperation, 1000)
	t.AddEntry(uuid.New(), accountPathTwo, entities.NewAccountVersion, entities.DebitOperation, 1000)
	err := conn.SaveTransaction(context.Background(), t)
	AssertEqual(nil, err)

	log.Println("starting GetSyntheticReport 3")
	now := time.Now().UnixNano()

	report, err := conn.GetSyntheticReport(context.Background(), accountPathOne, now, now)

	fmt.Printf("> report: %v\n\n", report)

	AssertTrue(report != nil)

	//AssertEqual(accountPathOne, accountBalance.AccountName().Name())
	//AssertEqual(expectedBalance, accountBalance.Balance())
	//AssertEqual(nil, err)
}
