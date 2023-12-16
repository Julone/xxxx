package main

import (
	"container/list"
	time "github.com/jinzhu/now"
	jsoniter "github.com/json-iterator/go"
	"strings"
	"sync"
	"testing"
)

type SharedResource struct {
	mu    sync.Mutex
	value int
}

func (r *SharedResource) Increment() {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.value++
}

func (r *SharedResource) GetValue() int {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.value
}

func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[0:len(prefix)] == prefix
}

func TestNddame(t *testing.T) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	s0 := "sadfadsfjasdiofoasdjfas"
	s := strings.NewReplacer("as", "屁股")
	s1 := s.Replace(s0)
	n := list.New()
	n.PushBack(3)
	n.PushBack(4)
	firstNode := n.Front()
	for i := 0; i < n.Len(); i++ {
		t.Logf("%v asdflasjdfklasjd ", firstNode.Value)
		firstNode = firstNode.Next()
	}
	t.Logf("%v %v", n, 0)
	t.Logf("%v, %v", s, s1)

}
