package utils

import (
	"fmt"
	"sync"
	"time"
	"github.com/garyburd/redigo/redis"
)

func GetRedisConnWithoutPool(ip, port, password string) (redis.Conn, error) {
	var conn redis.Conn
	var err error
	if len(password) > 0 {
		option := redis.DialPassword(password)
		conn, err = redis.Dial("tcp", fmt.Sprintf("%s:%s", ip, port), option)
	} else {
		conn, err = redis.Dial("tcp", fmt.Sprintf("%s:%s", ip, port))
	}
	return conn, err
}

var mutex = new(sync.Mutex)
var pool = make(map[string]*redis.Pool)

func GetRedisConnFromPool(maxidle, timeout int, ip, port, password string) (*redis.Pool, error) {
	mutex.Lock()
	// mutex是互斥锁，只有Lock和Unlock两个方法，在这两个方法之间的代码不能被多个goroutins同时调用
	defer mutex.Unlock()
	key := fmt.Sprintf("%s-%s", ip, port)
	if v, ok := pool[key]; ok {
		return v, nil
	} else {
		p, err := NewRedisPool(maxidle, timeout, ip, port, password)
		if err != nil {
			return nil, err
		}
		pool[key] = p
		return p, nil
	}
}

func NewRedisPool(maxIdle, timeout int, ip, port, password string) (*redis.Pool, error) {
	server := fmt.Sprintf("%s:%s", ip, port)
	pool := &redis.Pool{
		MaxIdle:     maxIdle,
		MaxActive:   12000,
		IdleTimeout: time.Duration(int64(timeout)) * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			if len(password) > 0 {
				if _, err := c.Do("AUTH", password); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}
