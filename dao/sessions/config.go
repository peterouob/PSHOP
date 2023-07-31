package sessions

import (
	"fmt"
	"net"
	"regexp"
	"strings"
	"time"
)

type store uint8

const (
	ram store = iota //儲存在內存
	rds store = iota //儲存在redis
	def

	prefix   = "gws_id"                          //默認session前綴,rds儲存時使用
	lifetime = time.Duration(1800) * time.Second //默認session的lifeTime
)

var (
	defaultOptions = option{
		LifeTime:   lifetime,
		CookieName: prefix,
		Domain:     "",
		Path:       "/",
		HttpOnly:   true,
		Secure:     true,
	}
	NewRDSOptions = func(ip string, port uint16, password string, opts ...func(*RDSOption)) *RDSOption {
		var rdsopt RDSOption
		rdsopt.option = defaultOptions
		rdsopt.Index = 0
		rdsopt.Prefix = prefix
		rdsopt.PoolSize = 10
		rdsopt.Password = password
		rdsopt.Address = fmt.Sprintf("%s:%v", ip, port)
		//這段程式碼的目的是在opts這個函式切片中，依次調用每個函式，並將rdsopt的指標作為參數傳入。
		//這樣可以讓每個函式都對rdsopt進行特定的設定，最終達到對rdsopt進行多個設定的目的。
		for _, opt := range opts {
			opt(&rdsopt)
		}
		return &rdsopt
	}
	//柯里化函數，縮減程式碼量
	WithIndex = func(index int) func(*RDSOption) {
		return func(r *RDSOption) {
			r.Index = index
		}
	}
	WithPool = func(size int) func(*RDSOption) {
		return func(r *RDSOption) {
			r.PoolSize = size
		}
	}
	WithPrefix = func(prefix string) func(*RDSOption) {
		return func(r *RDSOption) {
			r.Prefix = prefix
		}
	}
	WithOpt = func(opt Options) func(*RDSOption) {
		return func(r *RDSOption) {
			r.option = opt.option
		}
	}
)

type Options struct {
	option
}

type option struct {
	LifeTime   time.Duration `json:"lifeTime"`
	CookieName string        `json:"cookieName"`
	HttpOnly   bool          `json:"httpOnly"`
	Path       string        `json:"path"`
	Secure     bool          `json:"secure"` //安全
	Domain     string        `json:"domain"` //範圍
}

var (
	// WithLifeTime 是設定session生命週期,其中返回一個含option指針的函數
	WithLifeTime = func(d time.Duration) func(o *option) {
		return func(o *option) {
			o.LifeTime = d
		}
	}
	WithCookieName = func(n string) func(o *option) {
		return func(o *option) {
			o.CookieName = n
		}
	}
	WithHttpOnly = func(b bool) func(o *option) {
		return func(o *option) {
			o.HttpOnly = b
		}
	}
	WithPath = func(p string) func(o *option) {
		return func(o *option) {
			o.Path = p
		}
	}
	WithSecure = func(s bool) func(o *option) {
		return func(o *option) {
			o.Secure = s
		}
	}
	WithDomain = func(domain string) func(o *option) {
		return func(o *option) {
			o.Domain = domain
		}
	}
)

type RDSOption struct {
	option
	Index    int    `json:"db_index"`
	Prefix   string `json:"prefix"`
	Address  string `json:"address"`
	PoolSize int    `json:"pool_size"`
	Password string `json:"password"`
}

type Config struct {
	store `json:"store,omitempty"`
	*RDSOption
}

// Configure 配置session
type Configure interface {
	Parse() (cfg *Config)
}

func (opt *Options) Parse() (cfg *Config) {
	cfg = new(Config)
	cfg.RDSOption = new(RDSOption)
	cfg.RDSOption.option = opt.option
	return verifyConfig(cfg)
}

func (opt *RDSOption) Parse() (cfg *Config) {
	cfg = new(Config)
	cfg.store = rds
	cfg.RDSOption = opt
	return verifyConfig(cfg)
}

// 判斷data
func verifyConfig(cfg *Config) *Config {
	if cfg.CookieName == "" {
		panic("cookie name cannot be empty")
	}
	if cfg.Path == "" {
		panic("path cannot be empty")
	}
	if cfg.LifeTime <= lifetime {
		cfg.LifeTime = lifetime
	}
	if cfg.Index > 16 {
		cfg.Index = 0
	}
	if cfg.PoolSize == 0 {
		cfg.PoolSize = 10
	}
	if cfg.Prefix == "" {
		cfg.Prefix = prefix
	}
	//特殊的net請求
	if net.ParseIP(strings.Split(cfg.Address, ":")[0]) == nil {
		panic("remote ip address illegal")
	}
	if match, err := regexp.MatchString("^[0-9]*$", strings.Split(cfg.Address, ":")[1]); err == nil {
		if !match {
			panic("remote port illegal")
		}
	} else {
		panic("regexp error ")
	}
	return cfg
}
