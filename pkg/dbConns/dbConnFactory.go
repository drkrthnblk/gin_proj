package dbconns

import (
	"gin__proj/pkg/dbConns/postgresConn"
)

type DbConnSelectorFactory struct {}

func NewDbConnSelectorFactory() *DbConnSelectorFactory {
	return *DbConnSelectorFactory{}
}

func (dbf DbConnSelectorFactory) GetDbConn(dbtType string) (DbConnSelector,error) {
	switch dbtType {
	case DbConnType.POSTGRESCONN:
		return postgresConn.NewPostgresClient(), nil
	default:
		return nil, error.Error("unknown DB conn type")
	}
}