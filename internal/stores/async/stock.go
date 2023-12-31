package async

import (
	"POCS_Projects/internal/cmd/benchmark/databases"
	"POCS_Projects/internal/cmd/benchmark/databases/pocsdb"
	"POCS_Projects/internal/models"
	"log"
)

type StockStore struct {
	l  *log.Logger
	db *pocsdb.PocsDB
}

func NewStockStore(l *log.Logger, db *pocsdb.PocsDB) Store[models.Stock, models.StockPK] {
	return &StockStore{db: db, l: l}
}

func (s StockStore) Put(ctx *pocsdb.ConnectionContext, value models.Stock) <-chan databases.RequestResult {
	return s.db.Put(ctx, models.Stock{}, models.StockPK{ItemId: value.ItemId, WarehouseId: value.WarehouseId}, value)
}

func (s StockStore) Get(ctx *pocsdb.ConnectionContext, key models.StockPK) <-chan databases.RequestResult {
	return s.db.Get(ctx, models.Stock{}, key)
}

func (s StockStore) Delete(ctx *pocsdb.ConnectionContext, key models.StockPK) <-chan databases.RequestResult {
	return s.db.Delete(ctx, models.Stock{}, key)
}
