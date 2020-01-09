package test

import (
	"sync"
	"testing"
)

func TestInterfaces(t *testing.T) {
	var (
		s   = make(chan string)
		ret []string
		wg  sync.WaitGroup
		wg2 sync.WaitGroup
	)

	go func(retPtr *[]string) {
		*retPtr = append(*retPtr, <-s)
	}(&ret)

	// Add WaitGroups to ensure "hello" executes first.
	wg.Add(1)
	wg2.Add(1)

	go func() {
		s <- "goodbye"
	}()

	go func() {
		s <- "hello"
		wg.Done()
		wg2.Done()
	}()

	wg.Wait()
	wg2.Wait()

	if len(ret) == 0 {
		t.Error("nothing in ret")
		return
	}

	if ret[0] != "hello" {
		t.Error("hello is not", ret[0])
	}
}
