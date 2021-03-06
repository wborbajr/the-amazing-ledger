package main

import (
	"github.com/sirupsen/logrus"

	"github.com/stone-co/the-amazing-ledger/pkg/command-handler/domain/ledger/usecase"
	"github.com/stone-co/the-amazing-ledger/pkg/common/configuration"
	"github.com/stone-co/the-amazing-ledger/pkg/gateways/grpc"
	"github.com/stone-co/the-amazing-ledger/pkg/gateways/grpc/transactions"
)

func grpcAPIStart(config configuration.GRPCConfig, log *logrus.Logger, useCase *usecase.LedgerUseCase) {
	transactionsHandler := transactions.NewHandler(log, useCase)
	api := grpc.NewAPI(log, transactionsHandler)
	api.Start(config)
}
