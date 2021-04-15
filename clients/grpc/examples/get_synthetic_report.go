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

func getSyntheticReportFullPath(log *logrus.Entry, conn *ledger.Connection) {
	log.Println("starting GetSyntheticReport")
	defer log.Println("finishing GetSyntheticReport")

	// expectedBalance := 1000
	accountPathOne := "liability:stone:clients:" + uuid.New().String()

	// Define a new transaction with 2 entries
	t := conn.NewTransaction(uuid.New())
	t.AddEntry(uuid.New(), accountPathOne, entities.NewAccountVersion, entities.CreditOperation, 1000)
	err := conn.SaveTransaction(context.Background(), t)
	AssertEqual(nil, err)

	now := time.Now().UnixNano()

	report, err := conn.GetSyntheticReport(context.Background(), accountPathOne, now, now)

	fmt.Printf("> reportt: %v\n\n", report)

	AssertTrue(report != nil)

	paths := report.Paths()

	AssertTrue(paths != nil)

	AssertEqual(accountPathOne, paths[0].Account)
	AssertEqual(int64(1000), paths[0].Credit)
	AssertEqual(int64(0), paths[0].Debit)
}

func getSyntheticReportFullPathDoubleEntry(log *logrus.Entry, conn *ledger.Connection) {
	log.Println("starting GetSyntheticReport")
	defer log.Println("finishing GetSyntheticReport")

	// expectedBalance := 1000
	accountPathOne := "liability:stone:clients:" + uuid.New().String()

	// Define a new transaction with 2 entries
	t := conn.NewTransaction(uuid.New())
	t.AddEntry(uuid.New(), accountPathOne, entities.NewAccountVersion, entities.CreditOperation, 1000)
	t.AddEntry(uuid.New(), accountPathOne, entities.NewAccountVersion, entities.CreditOperation, 1000)
	err := conn.SaveTransaction(context.Background(), t)
	AssertEqual(nil, err)

	now := time.Now().UnixNano()

	report, err := conn.GetSyntheticReport(context.Background(), accountPathOne, now, now)

	fmt.Printf("> reportt: %v\n\n", report)

	AssertTrue(report != nil)

	paths := report.Paths()

	AssertTrue(paths != nil)

	AssertEqual(accountPathOne, paths[0].Account)
	AssertEqual(int64(2000), paths[0].Credit)
	AssertEqual(int64(0), paths[0].Debit)
}

func getSyntheticReportFullPathDebit(log *logrus.Entry, conn *ledger.Connection) {
	log.Println("starting GetSyntheticReport")
	defer log.Println("finishing GetSyntheticReport")

	// expectedBalance := 1000
	accountPathOne := "liability:stone:clients:" + uuid.New().String()

	// Define a new transaction with 2 entries
	t := conn.NewTransaction(uuid.New())
	t.AddEntry(uuid.New(), accountPathOne, entities.NewAccountVersion, entities.DebitOperation, 1000)
	err := conn.SaveTransaction(context.Background(), t)
	AssertEqual(nil, err)

	now := time.Now().UnixNano()

	report, err := conn.GetSyntheticReport(context.Background(), accountPathOne, now, now)

	fmt.Printf("> reportt: %v\n\n", report)

	AssertTrue(report != nil)

	paths := report.Paths()

	AssertTrue(paths != nil)

	AssertEqual(accountPathOne, paths[0].Account)
	AssertEqual(int64(1000), paths[0].Debit)
	AssertEqual(int64(0), paths[0].Credit)
}

func getSyntheticReportSubgroup(log *logrus.Entry, conn *ledger.Connection) {
	log.Println("starting GetSyntheticReport Subgroup")
	defer log.Println("finishing GetSyntheticReport Subgroup")

	// expectedBalance := 1000
	accountPathOne := "liability:stone:clients"

	// Define a new transaction with 2 entries
	t := conn.NewTransaction(uuid.New())
	t.AddEntry(uuid.New(), accountPathOne, entities.NewAccountVersion, entities.CreditOperation, 1000)
	err := conn.SaveTransaction(context.Background(), t)
	AssertEqual(nil, err)

	now := time.Now().UnixNano()

	report, err := conn.GetSyntheticReport(context.Background(), accountPathOne, now, now)

	fmt.Printf("> reportt: %v\n\n", report)

	AssertTrue(report != nil)

	paths := report.Paths()

	AssertTrue(paths != nil)

	AssertEqual(accountPathOne, paths[0].Account)
	AssertEqual(int64(1000), paths[0].Credit)
	AssertEqual(int64(0), paths[0].Debit)
}

func getSyntheticReportGroup(log *logrus.Entry, conn *ledger.Connection) {
	log.Println("starting GetSyntheticReport Subgroup")
	defer log.Println("finishing GetSyntheticReport Subgroup")

	// expectedBalance := 1000
	accountPathOne := "liability:stone"
	accountPathTwo := "liability:xpto"

	// Define a new transaction with 2 entries
	t := conn.NewTransaction(uuid.New())
	t.AddEntry(uuid.New(), accountPathOne, entities.NewAccountVersion, entities.CreditOperation, 1000)
	t.AddEntry(uuid.New(), accountPathTwo, entities.NewAccountVersion, entities.DebitOperation, 1000)
	err := conn.SaveTransaction(context.Background(), t)
	AssertEqual(nil, err)

	now := time.Now().UnixNano()

	report, err := conn.GetSyntheticReport(context.Background(), accountPathOne, now, now)

	fmt.Printf("> reportt: %v\n\n", report)

	AssertTrue(report != nil)

	paths := report.Paths()

	AssertTrue(paths != nil)

	AssertEqual(accountPathOne, paths[0].Account)
	AssertEqual(int64(1000), paths[0].Credit)
	AssertEqual(int64(0), paths[0].Debit)
}
