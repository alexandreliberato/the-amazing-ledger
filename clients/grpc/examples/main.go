package main

import (
	"github.com/sirupsen/logrus"

	"github.com/stone-co/the-amazing-ledger/clients/grpc/ledger"
)

func main() {
	logrus := logrus.New()
	log := logrus.WithField("ClientSDK", "Test")

	log.Println("Server example starting...")
	defer log.Println("Server example finishing...")

	// Connect to the Ledger gRPC server
	host := "localhost"
	port := 50051
	conn, err := ledger.Connect(host, port)
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	//	DefineTransactionWithThreeEntries(log, conn)
	//	invalidTransactionsTests(log, conn)

	getSyntheticReportFullPath(log, conn)
	getSyntheticReportFullPathDoubleEntry(log, conn)

	getSyntheticReportFullPathDebit(log, conn)

	getSyntheticReportSubgroup(log, conn)
	getSyntheticReportGroup(log, conn)

	//	getAccountBalance(log, conn)
	//	getAccountBalanceWithMoreEntries(log, conn)
	//	getAccountBalanceNotFoundAccount(log, conn)
}
