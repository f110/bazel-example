package testsharding

import (
	"testing"
	"time"
)

func Test1(t *testing.T) {
	t.Log(t.Name())
	time.Sleep(10 * time.Second)
}

func Test2(t *testing.T) {
	t.Log(t.Name())
	time.Sleep(10 * time.Second)
}
