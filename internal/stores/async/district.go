package async

import (
	"POCS_Projects/internal/benchmark/databases"
	"POCS_Projects/internal/benchmark/databases/pocsdb"
	"POCS_Projects/internal/models"
	"log"
)

type DisctrictStore struct {
	l  *log.Logger
	db *pocsdb.PocsDB
}

func NewDiscrictStore(l *log.Logger, db *pocsdb.PocsDB) Store[models.District, models.DistrictPK] {
	return &DisctrictStore{db: db, l: l}
}

func (d *DisctrictStore) Put(ctx *databases.ConnectionContext, value models.District) <-chan databases.RequestResult {
	return d.db.Put(ctx, models.District{}, models.DistrictPK{Id: value.Id}, value)
}

func (d *DisctrictStore) Get(ctx *databases.ConnectionContext, key models.DistrictPK) <-chan databases.RequestResult {
	return d.db.Get(ctx, models.District{}, key)
}

func (d *DisctrictStore) Delete(ctx *databases.ConnectionContext, key models.DistrictPK) <-chan databases.RequestResult {
	return d.db.Delete(ctx, models.District{}, key)
}
