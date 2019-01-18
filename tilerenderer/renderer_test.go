package tilerenderer

import (
	"bytes"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"mapserver/colormapping"
	"mapserver/coords"
	"mapserver/db"
	"mapserver/layerconfig"
	"mapserver/mapblockaccessor"
	"mapserver/mapblockrenderer"
	"mapserver/testutils"
	"mapserver/tiledb"
	"os"
	"testing"
)

func TestTileRender(t *testing.T) {
	logrus.SetLevel(logrus.DebugLevel)

	tmpfile, err := ioutil.TempFile("", "TestTileRender.*.sqlite")
	if err != nil {
		panic(err)
	}

	defer os.Remove(tmpfile.Name())
	testutils.CreateTestDatabase(tmpfile.Name())

	a, err := db.NewSqliteAccessor(tmpfile.Name())
	if err != nil {
		panic(err)
	}

	err = a.Migrate()
	if err != nil {
		panic(err)
	}

	cache := mapblockaccessor.NewMapBlockAccessor(a)
	c := colormapping.NewColorMapping()
	err = c.LoadVFSColors(false, "/colors.txt")
	if err != nil {
		t.Fatal(err)
	}

	r := mapblockrenderer.NewMapBlockRenderer(cache, c)

	tiletmpfile, err := ioutil.TempFile("", "TestTileRenderTiles.*.sqlite")
	defer os.Remove(tiletmpfile.Name())

	tdb, _ := tiledb.NewSqliteAccessor(tiletmpfile.Name())
	tdb.Migrate()

	tr := NewTileRenderer(r, tdb, a, layerconfig.DefaultLayers)

	if tr == nil {
		panic("no renderer")
	}

	coord := coords.NewTileCoords(0, 0, 12, 0)

	data, err := tr.Render(coord)

	if err != nil {
		panic(err)
	}

	if data == nil {
		panic("no data")
	}

	f, _ := os.Create("../output/0_0_12.png")
	bytes.NewReader(data).WriteTo(f)
}
