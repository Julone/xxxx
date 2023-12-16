package main

import (
	"encoding/gob"
	"fmt"
	"github.com/dgraph-io/ristretto"
	"os"
	"strconv"
	"testing"
)

type Sfas struct {
	id int
}
type CommonMap map[string]interface{}

func TestGetAllTodos(t *testing.T) {
	cache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e7,     // number of keys to track frequency of (10M).
		MaxCost:     1 << 30, // maximum cost of cache (1GB).
		BufferItems: 64,      // number of keys per Get buffer.
	})

	if err != nil {
		panic(err)
	}
	//mu := sync.RWMutex{}

	w, err := os.Create("./saving.txt")
	enc := gob.NewEncoder(w)
	var tmpdata = make(CommonMap)

	gob.Register(Sfas{})
	for i := 1; i < 10; i++ {
		item := Sfas{id: i}
		cache.Set("key"+strconv.Itoa(i), item, 1)
		cache.Wait()
		tmpdata["key"+strconv.Itoa(i)] = item
	}
	enc.Encode(&tmpdata)
}

func TestName(t *testing.T) {
	w, _ := os.Open("./saving.txt")
	enc := gob.NewDecoder(w)
	gob.Register(Sfas{})
	items := CommonMap{}
	_ = enc.Decode(&items)
	fmt.Printf("%v", items)

	for s, i := range items {
		fmt.Printf("%v %v", s, i)

	}

}
