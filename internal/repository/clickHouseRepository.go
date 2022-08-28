package repository

import "github.com/roistat/go-clickhouse"

type ClickHouseLog interface {
	SaveCreatedUserLog(logMessage string) error
}

type ClickHouse struct {
	ClickHouseLog
}

func NewClickHouseRepository(con *clickhouse.Conn) *ClickHouse {
	return &ClickHouse{
		ClickHouseLog: NewClickHouseLogs(con),
	}
}
