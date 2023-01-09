package dbconns

import (
	"fmt"
	"gin_proj/pkg/dbConns/postgresConn"
)

type DbConnSelectorFactory struct{}

func NewDbConnSelectorFactory() *DbConnSelectorFactory {
	return &DbConnSelectorFactory{}
}

func (dbf DbConnSelectorFactory) GetDbConn(dbtType DbConnType) (interface{}, error) {
	switch dbtType {
	case POSTGRESCONN:
		return postgresConn.NewPostgresClient(), nil
	default:
		return nil, fmt.Errorf("unknown DB conn type")
	}
}
