package sessions

import (
	H "PSHOP/http"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
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

type Values map[string]interface{}

type Session struct {
	id       string
	CreateAt time.Time
	ExpireAt time.Time
	Values
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

func GetSession(c *gin.Context, r *http.Request) (*Session, error) {
	var session Session
	cookie, err := r.Cookie(globalConfig.CookieName)
	if cookie == nil && err != nil {
		return CreateSession(c, cookie)
	}
	if len(cookie.Value) >= 73 {
		session.id = cookie.Value
		if err := globalStorage.Read(&session); err != nil {
			return CreateSession(c, cookie)
		}
	}
	return &session, nil
}

func CreateSession(c *gin.Context, cookie *http.Cookie) (*Session, error) {
	session := NewSession()
	if cookie == nil {
		cookie = NewCookie()
	}
	cookie.Value = session.id
	if err := globalStorage.Write(session); err != nil {
		return nil, err
	}
	//http.SetCookie(w, cookie)
	H.SetCookie(c, "session_cookie", cookie.Value)
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
		err := fmt.Errorf("unsupported store: %v", globalConfig.store)
		fmt.Println(err.Error())
	}
}

func Migrate(w http.ResponseWriter, oldSession *Session) (*Session, error) {
	var (
		ns     = NewSession()
		cookie = NewCookie()
	)
	migrateMux.Lock()
	defer migrateMux.Unlock()
	ns.Values = oldSession.Values
	cookie.Value = ns.id
	cookie.MaxAge = int(globalConfig.LifeTime) / 1e9
	return ns,
		func() error {
			if ns.Sync() != nil {
				return errors.New("error migrating session")
			}
			if globalStorage.Remove(oldSession) != nil {
				return errors.New("error removing session")
			}
			http.SetCookie(w, cookie)
			return nil
		}()
}

func (s *Session) Sync() error {
	return globalStorage.Write(s)
}
func (s *Session) Remove() error {
	return globalStorage.Remove(s)
}
