package repository

import "github.com/roistat/go-clickhouse"

type ClickHouseLog interface {
	SaveCreateUserLog(logMessage string) error
}

type ClickHouse struct {
	ClickHouseLog
}

func NewClickHouseRepository(con *clickhouse.Conn) *ClickHouse {
	return &ClickHouse{
		NewClickHouseLogs(con),
	}
}
