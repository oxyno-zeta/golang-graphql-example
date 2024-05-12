package database

import "database/sql"

type TransactionOption func(cfg *TransactionOptionsConfig)

type TransactionOptionsConfig struct {
	ReadTransaction bool
	IsolationLevel  sql.IsolationLevel
}

var WithReadTransactionOpt TransactionOption = func(cfg *TransactionOptionsConfig) { cfg.ReadTransaction = true }
var WithWriteTransactionOpt TransactionOption = func(cfg *TransactionOptionsConfig) { cfg.ReadTransaction = false }
