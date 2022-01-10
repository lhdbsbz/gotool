package gotool

import (
	"fmt"
	"testing"
	"time"
)

func TestInitMemCache(t *testing.T) {
	m := InitMemCache(MemCacheTypeTimedClear, "*/5 * * * * ?")
	m.Put("abc", "abc")
	fmt.Println(m.Get("abc"))
	time.Sleep(time.Second * 30)
	fmt.Println(m.Get("abc"))
	fmt.Println(m.Info())
}
