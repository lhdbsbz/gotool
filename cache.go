package gotool

import (
	"fmt"

	"github.com/robfig/cron"
)

type MemCacheType string

var MemCacheTypeTimedClear MemCacheType = "定时清除"

type MemCache struct {
	Type MemCacheType
	Corn string
	kv   map[string]string
	c    *cron.Cron
}

type MemCacheinterface interface {
	Info() string
	Size() int
	Put(key string, value string)
	Get(key string) (value string)
	Remove(key string) (value string)
	Clear()
}

func (item *MemCache) Info() string {
	return fmt.Sprintf("缓存类型为:%v corn表达式为%v 当前Size为%v", item.Type, item.Corn, item.Size())
}

func (item *MemCache) Size() int {
	return len(item.kv)
}

func (item *MemCache) Put(key string, value string) {
	item.kv[key] = value
}

func (item *MemCache) Get(key string) (value string) {
	return item.kv[key]
}

func (item *MemCache) Remove(key string) (value string) {
	value = item.kv[key]
	delete(item.kv, key)
	return
}

func InitMemCache(memCacheType MemCacheType, corn string) *MemCache {
	item := &MemCache{
		Type: memCacheType,
		Corn: corn,
	}
	item.kv = make(map[string]string)
	if IsNotEmpty(item.Corn) {
		item.c = cron.New()
		if err := item.c.AddFunc(item.Corn, item.doCorn); err == nil {
			item.c.Start()
		} else {
			panic(err)
		}
	}
	return item
}

func (item *MemCache) Clear() {
	item.kv = make(map[string]string)
}

func (item *MemCache) doCorn() {
	fmt.Println("开始执行corn")
	switch item.Type {
	case MemCacheTypeTimedClear:
		item.Clear()
	}
}
