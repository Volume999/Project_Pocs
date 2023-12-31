package async

import (
	"POCS_Projects/internal/cmd/benchmark/databases"
	"POCS_Projects/internal/cmd/benchmark/databases/pocsdb"
	"POCS_Projects/internal/models"
	"log"
)

type CustomerStore struct {
	l  *log.Logger
	db *pocsdb.PocsDB
}

func (c *CustomerStore) Put(ctx *pocsdb.ConnectionContext, value models.Customer) <-chan databases.RequestResult {
	return c.db.Put(ctx, models.Customer{}, models.CustomerPK{ID: value.ID}, value)
}

func (c *CustomerStore) Get(ctx *pocsdb.ConnectionContext, key models.CustomerPK) <-chan databases.RequestResult {
	return c.db.Get(ctx, models.Customer{}, key)
}

func (c *CustomerStore) Delete(ctx *pocsdb.ConnectionContext, key models.CustomerPK) <-chan databases.RequestResult {
	return c.db.Delete(ctx, models.Customer{}, key)
}

func NewCustomerStore(l *log.Logger, db *pocsdb.PocsDB) Store[models.Customer, models.CustomerPK] {
	return &CustomerStore{db: db, l: l}
}
