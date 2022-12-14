package repository

import (
	"fmt"
	"github.com/roistat/go-clickhouse"
)

type ConfigClickHouse struct {
	Host string
	Port string
}

func NewClickHouseDB(cfg ConfigClickHouse) (*clickhouse.Conn, error) {
	transport := clickhouse.NewHttpTransport()
	conn := clickhouse.NewConn(fmt.Sprintf("#{cfg.Host}:#{cfg.Port}"), transport)
	err := conn.Ping()
	if err != nil {
		return nil, err
	}

	return conn, nil

}
