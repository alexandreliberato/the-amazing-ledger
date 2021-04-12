package rpc

import (
	"context"

	"github.com/sirupsen/logrus"
	proto "github.com/stone-co/the-amazing-ledger/gen/ledger"
	"github.com/stone-co/the-amazing-ledger/pkg/command-handler/domain/ledger/entities"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"time"
)

func (a *API) GetSyntheticReport(ctx context.Context, request *proto.GetSyntheticReportRequest) (*proto.GetSyntheticReportResponse, error) {
	log := a.log.WithFields(logrus.Fields{
		"handler": "GetSyntheticReport",
	})

	startTime := time.Unix(0, request.StartTime)
	endTime := time.Unix(0, request.EndTime)

	syntheticReport, err := a.UseCase.GetSyntheticReport(ctx, request.AccountPath, startTime, endTime)
	if err != nil {
		if err == entities.ErrAccountNotFound {
			log.WithError(err).Error("account name does not exist")
			return nil, status.Error(codes.NotFound, err.Error())
		}

		log.WithError(err).Error("can't get account")
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &proto.GetSyntheticReportResponse{
		CurrentVersion: uint64(syntheticReport.CurrentVersion),
		TotalCredit:    int64(syntheticReport.TotalCredit),
		TotalDebit:     int64(syntheticReport.TotalDebit),
		Paths:          toProto(syntheticReport.Paths),
	}, nil
}

func toProto(paths []entities.Path) []*proto.Path {
	protoPaths := []*proto.Path{}

	for _, element := range paths {
		protoPaths = append(protoPaths, &proto.Path{
			Account: element.Account,
			Credit:  int64(element.Credit),
			Debit:   int64(element.Debit),
		})
	}

	return protoPaths
}
