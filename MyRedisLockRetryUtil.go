package gotool

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

type MyRedisLockRetryUtil struct {
	client        *redis.Client
	prefix        string
	retryTimes    int64
	retryDuration time.Duration
}

func InitMyRedisLockRetryUtil(client *redis.Client, prefix string, times int64) *MyRedisLockRetryUtil {
	if times <= 0 {
		times = 40
	}
	return &MyRedisLockRetryUtil{
		client:        client,
		prefix:        prefix,
		retryTimes:    times,
		retryDuration: 25 * time.Millisecond,
	}
}

func (m *MyRedisLockRetryUtil) LockWithRetry(ctx context.Context, key string, duration time.Duration, waitUntilHold bool) (locked bool) {
	var times int64 = 0
	for {
		nx := m.client.SetNX(ctx, m.formatKey(key), key, duration)
		locked, _ = nx.Result()
		if locked {
			break
		}
		times++
		if !waitUntilHold {
			if times > m.retryTimes {
				break
			}
		}
		time.Sleep(m.retryDuration)
	}
	return
}

func (m *MyRedisLockRetryUtil) LockJob(ctx context.Context, key string, duration time.Duration) (locked bool) {
	nx := m.client.SetNX(ctx, m.formatKey(key), key, duration)
	locked, _ = nx.Result()
	return
}

func (m *MyRedisLockRetryUtil) Unlock(key string) {
	_, _ = m.client.Del(context.Background(), m.formatKey(key)).Result()
	return
}

func (m *MyRedisLockRetryUtil) formatKey(key string) string {
	return fmt.Sprintf("%v:%v", m.prefix, key)
}

func (m *MyRedisLockRetryUtil) NewSetByKey(ctx context.Context, key string, value string, expiration time.Duration) (err error) {
	err = m.client.Set(ctx, m.formatKey(key), value, expiration).Err()
	return
}

func (m *MyRedisLockRetryUtil) NewGetByKey(ctx context.Context, key string) (value string, err error) {
	var exists int64
	exists, err = m.client.Exists(ctx, m.formatKey(key)).Result()
	if err != nil {
		return
	}
	if exists == 0 {
		value = ""
		return
	}
	value, err = m.client.Get(ctx, m.formatKey(key)).Result()
	return
}
