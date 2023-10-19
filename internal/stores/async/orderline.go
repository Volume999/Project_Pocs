package async

import (
	"POCS_Projects/internal/benchmark/databases"
	"POCS_Projects/internal/benchmark/databases/pocsdb"
	"POCS_Projects/internal/models"
	"log"
)

type OrderLineStore struct {
	l  *log.Logger
	db *pocsdb.PocsDB
}

func NewOrderLineStore(l *log.Logger, db *pocsdb.PocsDB) Store[models.OrderLine, models.OrderLinePK] {
	return &OrderLineStore{db: db, l: l}
}

func (o OrderLineStore) Put(ctx *databases.ConnectionContext, value models.OrderLine) <-chan databases.RequestResult {
	return o.db.Put(ctx, models.OrderLine{}, models.OrderLinePK{OrderId: value.OrderId,
		DistrictId: value.DistrictId, WarehouseId: value.WarehouseId, LineNumber: value.LineNumber}, value)
}

func (o OrderLineStore) Get(ctx *databases.ConnectionContext, key models.OrderLinePK) <-chan databases.RequestResult {
	return o.db.Get(ctx, models.OrderLine{}, key)
}

func (o OrderLineStore) Delete(ctx *databases.ConnectionContext, key models.OrderLinePK) <-chan databases.RequestResult {
	return o.db.Delete(ctx, models.OrderLine{}, key)
}
