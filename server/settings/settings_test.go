package settings

import (
  "testing"
  "os"
  "io/ioutil"
  "mapserver/mapobjectdb/sqlite"
)

func TestStrings(t *testing.T){
  tmpfile, err := ioutil.TempFile("", "TileDBTest.*.sqlite")
	if err != nil {
		panic(err)
	}
	defer os.Remove(tmpfile.Name())

  db, err := sqlite.New(tmpfile.Name())
  if err != nil {
    panic(err)
  }

  err = db.Migrate()
  if err != nil {
    panic(err)
  }

  s := New(db)
  s.SetString("k", "v")
  str := s.GetString("k", "v2")
  if str != "v" {
    t.Fatal("getstring failed: " + str)
  }

  if s.GetString("k2", "v3") != "v3" {
    t.Fatal("getstring with default failed")
  }

  s.SetInt("i", 123)
  i := s.GetInt("i", 456)
  if i != 123 {
    t.Fatal("getint failed")
  }

  if s.GetInt("i2", 111) != 111 {
    t.Fatal("getint with default failed")
  }

  s.SetInt64("i", 1230000012300056)
  i2 := s.GetInt64("i", 456)
  if i2 != 1230000012300056 {
    t.Fatal("getint64 failed")
  }

  if s.GetInt64("i2", 12300000123000564) != 12300000123000564 {
    t.Fatal("getint with default failed")
  }

}