package Config

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"time"
)

func InitRedis() {
	RedisPool = &redis.Pool{
		MaxIdle:     1000,
		MaxActive:   0,
		IdleTimeout: 240 * time.Second,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			con, err := redis.Dial("tcp", App.GetString("redis.host")+":"+App.GetString("redis.port"),
				//redis.DialPassword(conf["Password"].(string)),
				redis.DialDatabase(App.GetInt("redis.database")),
				redis.DialConnectTimeout(240*time.Second),
				redis.DialReadTimeout(240*time.Second),
				redis.DialWriteTimeout(240*time.Second))
			if err != nil {
				return nil, err
			}
			return con, nil
		},
	}
	err := RedisPool.Get().Err()
	if err != nil {
		panic("redis连接失败: " + err.Error())
	}
	fmt.Printf("redis连接成功: %s:%s/%s \r\n",
		App.GetString("redis.host"),
		App.GetString("redis.port"),
		App.GetInt("redis.database"),
	)

}
