package freshstorage_test

import (
	"fmt"
	"os"
	"strings"
	"testing"
	"wiki-knogen-zero/service/freshstorage"
)

var (
	freshS *freshstorage.FreshStorage
)

func TestMain(m *testing.M) {
	fmt.Println("pgconn", os.Getenv("TestPgConnUri"))
	freshS = freshstorage.NewFreshStorage(
		strings.ReplaceAll(os.Getenv("TestPgConnUri"), "\"", ""),
	)
	m.Run()

}

func TestStorePut(t *testing.T) {
	ret := freshS.PutRecord("test", "1", "string")
	t.Log(ret)

	ret = freshS.PutRecord("test", "1", "basic")
	t.Log(ret)
}

func TestStoreGet(t *testing.T) {
	ret := freshS.GetLastestRecord("test", "1")
	t.Log(ret)

	ret = freshS.GetLastestRecord("test", "2")
	t.Log(ret)
}
