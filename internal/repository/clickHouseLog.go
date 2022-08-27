package repository

import "github.com/roistat/go-clickhouse"

type ClickHouseLogs struct {
	con *clickhouse.Conn
}

func NewClickHouseLogs(con *clickhouse.Conn) *ClickHouseLogs {
	return &ClickHouseLogs{con: con}
}

func (c *ClickHouseLogs) SaveCreatedUserId(logMessage string) error {
	q := clickhouse.NewQuery("INSERT INTO default.logs VALUES (toDate(now()), ?, ?)", 1, logMessage)
	err := q.Exec(c.con)

	if err != nil {
		return err
	}

	return nil
}
