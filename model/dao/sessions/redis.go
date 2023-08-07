package sessions

import (
	"PSHOP/model/dao/user"
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"sync"
	"time"
)

type RdsStore struct {
	rw    sync.RWMutex
	store *redis.Client
}

func NewRdsStore() *RdsStore {
	return &RdsStore{
		rw: sync.RWMutex{},
		store: redis.NewClient(&redis.Options{
			Addr:     globalConfig.Address,
			Password: "",
			DB:       globalConfig.Index,
			PoolSize: globalConfig.PoolSize,
		}),
	}
}

func (r *RdsStore) Read(s *Session) (err error) {
	timeout, cancel := timeoutCtx()
	r.rw.RLock()
	defer func() {
		cancel()
		r.rw.RUnlock()
	}()
	var val []byte
	if val, err = r.store.Get(timeout, formatPrefix(s.id)).Bytes(); err != nil {
		return err
	}
	return json.Unmarshal(val, s)
}

func (r *RdsStore) Write(s *Session) (err error) {
	data, err := json.Marshal(s)
	if err != nil {
		return err
	}
	timeout, cancel := timeoutCtx()
	r.rw.Lock()
	defer func() {
		cancel()
		r.rw.Unlock()
	}()
	prename := formatPrefix(s.id)
	user.C <- prename
	return r.store.Set(timeout, prename, data, expire(s.ExpireAt)).Err()
}

func (r *RdsStore) Remove(s *Session) (err error) {
	timeout, cancel := timeoutCtx()
	r.rw.Lock()
	defer func() {
		cancel()
		r.rw.Unlock()
	}()
	name := <-user.C
	return r.store.Del(timeout, name).Err()
}

func formatPrefix(sid string) string {
	return fmt.Sprintf("%s:%s", globalConfig.Prefix, sid)
}

func timeoutCtx() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), time.Duration(3)*time.Second)
}

func expire(t time.Time) time.Duration {
	return time.Until(t)
}
