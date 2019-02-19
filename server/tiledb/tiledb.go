package tiledb

import (
	"fmt"
	"mapserver/coords"

	"github.com/dgraph-io/badger"
	"github.com/prometheus/client_golang/prometheus"
)

func New(path string) (*TileDB, error) {
	opts := badger.DefaultOptions
	opts.Dir = path
	opts.ValueDir = path
	db, err := badger.Open(opts)

	if err != nil {
		return nil, err
	}

	return &TileDB{
		db: db,
	}, nil
}

type TileDB struct {
	db *badger.DB
}

func getKey(pos *coords.TileCoords) []byte {
	return []byte(fmt.Sprintf("%d/%d/%d/%d", pos.X, pos.Y, pos.Zoom, pos.LayerId))
}

func (this *TileDB) GC() {
	for {
		err := this.db.RunValueLogGC(0.25)
		if err != nil {
			return
		}
	}
}

func (this *TileDB) GetTile(pos *coords.TileCoords) ([]byte, error) {
	timer := prometheus.NewTimer(getDuration)
	defer timer.ObserveDuration()

	var tile []byte
	err := this.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(getKey(pos))
		if item != nil {
			tile, err = item.ValueCopy(nil)
		}
		return err
	})
	if err != nil {
		return nil, nil
	}

	return tile, err
}

func (this *TileDB) SetTile(pos *coords.TileCoords, tile []byte) error {
	timer := prometheus.NewTimer(setDuration)
	defer timer.ObserveDuration()

	err := this.db.Update(func(txn *badger.Txn) error {
		err := txn.Set(getKey(pos), tile)
		return err
	})

	return err
}

func (this *TileDB) RemoveTile(pos *coords.TileCoords) error {
	timer := prometheus.NewTimer(removeDuration)
	defer timer.ObserveDuration()

	err := this.db.Update(func(txn *badger.Txn) error {
		err := txn.Delete(getKey(pos))
		return err
	})

	return err
}