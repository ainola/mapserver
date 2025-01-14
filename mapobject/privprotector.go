package mapobject

import (
	"mapserver/coords"
	"mapserver/mapobjectdb"

	"github.com/minetest-go/mapparser"
)

type PrivProtectorBlock struct{}

func (this *PrivProtectorBlock) onMapObject(mbpos *coords.MapBlockCoords, x, y, z int, block *mapparser.MapBlock) *mapobjectdb.MapObject {
	md := block.Metadata.GetMetadata(x, y, z)

	o := mapobjectdb.NewMapObject(mbpos, x, y, z, "privprotector")
	o.Attributes["owner"] = md["owner"]
	o.Attributes["priv"] = md["priv"]

	return o
}
