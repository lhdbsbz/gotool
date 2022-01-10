package gotool

import (
	"fmt"

	"github.com/robfig/cron/v3"
)

type MemCacheType string

var TimedClear MemCacheType = "定时清除"

type MemCache struct {
	Type MemCacheType
	Corn string
	kv   map[string]interface{}
	c    *cron.Cron
}

type MemCacheinterface interface {
	Init(*MemCache) *MemCache
	Info() string
	Size() int
	Put(key string, value interface{})
	Get(key string) (value interface{})
	Remove(key string) (value interface{})
	Clear()
}

func (item *MemCache) Info() string {
	return fmt.Sprintf("缓存类型为:%v corn表达式为%v 当前Size为%v", item.Type, item.Corn, item.Size())
}

func (item *MemCache) Size() int {
	return len(item.kv)
}

func (item *MemCache) Put(key string, value interface{}) {
	item.kv[key] = value
}

func (item *MemCache) Get(key string) (value interface{}) {
	return item.kv[key]
}

func (item *MemCache) Remove(key string) (value interface{}) {
	value = item.kv[key]
	delete(item.kv, key)
	return
}

func Init(item *MemCache) *MemCache {
	item.kv = make(map[string]interface{})
	if IsNotEmpty(item.Corn) {
		item.c = cron.New()
		_, _ = item.c.AddFunc(item.Corn, item.doCorn)
	}
	if item.c != nil {
		item.c.Start()
	}
	return item
}

func (item *MemCache) Clear() {
	item.kv = make(map[string]interface{})
}

func (item *MemCache) doCorn() {
	switch item.Type {
	case TimedClear:
		item.Clear()
	}
}
