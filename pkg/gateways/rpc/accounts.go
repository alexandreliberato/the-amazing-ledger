package rpc

import (
	"context"

	"github.com/sirupsen/logrus"
	proto "github.com/stone-co/the-amazing-ledger/gen/ledger"
	"github.com/stone-co/the-amazing-ledger/pkg/command-handler/domain/ledger/entities"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a *API) GetAccountBalance(ctx context.Context, request *proto.GetAccountBalanceRequest) (*proto.GetAccountBalanceResponse, error) {
	log := a.log.WithFields(logrus.Fields{
		"handler": "GetAccountBalance",
	})

	accountName, err := entities.NewAccountName(request.AccountPath)
	if err != nil {
		log.WithError(err).Error("can't create account name")
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	accountBalance, err := a.UseCase.GetAccountBalance(ctx, *accountName)
	if err != nil {
		if err == entities.ErrAccountNotFound {
			log.WithError(err).Error("account name does not exist")
			return nil, status.Error(codes.NotFound, err.Error())
		}

		log.WithError(err).Error("can't get account")
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &proto.GetAccountBalanceResponse{
		AccountPath:    accountBalance.AccountName.Name(),
		CurrentVersion: accountBalance.CurrentVersion.ToUInt64(),
		TotalCredit:    int64(accountBalance.TotalCredit),
		TotalDebit:     int64(accountBalance.TotalDebit),
		Balance:        int64(accountBalance.Balance()),
	}, nil
}
