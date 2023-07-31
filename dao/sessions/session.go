package sessions

import (
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"sync"
	"time"
)

var (
	globalStorage Storage
	globalConfig  *Config
	migrateMux    sync.Mutex
)

type Session struct {
	id       string
	CreateAt time.Time
	ExpireAt time.Time
	Values   map[string]interface{}
}

func NewSession() *Session {
	nowTime := time.Now()
	return &Session{
		id:       uuid73(),
		CreateAt: nowTime,
		ExpireAt: nowTime.Add(lifetime),
		Values:   make(map[string]interface{}),
	}
}

func NewCookie() *http.Cookie {
	return &http.Cookie{
		Name:     globalConfig.CookieName,
		Path:     globalConfig.Path,
		Secure:   globalConfig.Secure,
		HttpOnly: globalConfig.HttpOnly,
		Domain:   globalConfig.Domain,
	}
}

func GetSession(w http.ResponseWriter, r *http.Request) (*Session, error) {
	var session Session
	cookie, err := r.Cookie(globalConfig.CookieName)
	if cookie == nil && err != nil {
		return CreateSession(w, cookie)
	}
	if len(cookie.Value) >= 73 {
		session.id = cookie.Value
		if err := globalStorage.Read(&session); err != nil {
			return CreateSession(w, cookie)
		}
	}
	return &session, nil
}

func CreateSession(w http.ResponseWriter, cookie *http.Cookie) (*Session, error) {
	session := NewSession()
	if cookie == nil {
		cookie = NewCookie()
	}
	cookie.Value = session.id
	cookie.MaxAge = int(globalConfig.LifeTime) / 1e9
	if err := globalStorage.Write(session); err != nil {
		return nil, err
	}
	http.SetCookie(w, cookie)
	return session, nil
}
func uuid73() string {
	return fmt.Sprintf("%s-%s", uuid.New().String(), uuid.New().String())
}

func Open(opt Configure) {
	globalConfig = opt.Parse()
	switch globalConfig.store {
	case rds:
		rdb := NewRdsStore()
		timeout, cancel := timeoutCtx()
		defer cancel()
		if err := rdb.store.Ping(timeout).Err(); err != nil {
			panic(err.Error())
		}
		globalStorage = rdb
	default:
		fmt.Errorf("unsupported store: %v", globalConfig.store)
	}
}
func (s *Session) Sync() error {
	return globalStorage.Write(s)
}
