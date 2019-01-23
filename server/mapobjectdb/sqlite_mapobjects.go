package mapobjectdb

import (
	"mapserver/coords"
)

func (db *Sqlite3Accessor) GetMapData(q SearchQuery) ([]MapObject, error) {
	return nil, nil
}

const removeMapDataQuery = `
delete from objects where posx = ? and posy = ? and posz = ?
`

func (db *Sqlite3Accessor) RemoveMapData(pos coords.MapBlockCoords) error {
	_, err := db.db.Exec(removeMapDataQuery, pos.X, pos.Y, pos.Z)
	return err
}

const addMapDataQuery = `
insert into
objects(x,y,z,posx,posy,posz,type,mtime)
values(?, ?, ?, ?, ?, ?, ?, ?)
returning id
`

const addMapDataAttributeQuery = `
insert into
object_attributes(objectid, key, value)
values(?, ?, ?)
`

func (db *Sqlite3Accessor) AddMapData(data MapObject) error {
	_, err := db.db.Exec(addMapDataQuery,
			data.X, data.Y, data.Z,
			data.MBPos.X, data.MBPos.Y, data.MBPos.Z,
			data.Type, data.Mtime)

	if err != nil {
		return err
	}

	//TODO
	id := 1

	for k, v := range data.Attributes {
		_, err := db.db.Exec(addMapDataAttributeQuery, id, k, v)

			if err != nil {
				return err
			}
	}

	return nil
}
