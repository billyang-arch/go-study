package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

func lua_demo(rdb *redis.Client) {
	ctx := context.Background()
	// 将 Lua 脚本定义为字符串
	luaScript := `
		local current = redis.call('get', KEYS[1])
		if current and tonumber(current) >= tonumber(ARGV[2]) then
			return 0
		end
		
		local count = redis.call('incr', KEYS[1])
		
		if count == 1 then
			redis.call('expire', KEYS[1], ARGV[1])
		end
		
		return 1
	`

	// 定义 key 和参数
	key := "ip:127.0.0.1"
	expire := "60" // 60s
	limit := "3"   // 3次

	// 执行脚本
	// func (c *Client) Eval(ctx context.Context, script string, keys []string, args ...interface{}) *Cmd
	result, err := rdb.Eval(ctx, luaScript, []string{key}, expire, limit).Result()
	if err != nil {
		panic(err)
	}

	fmt.Printf("执行结果: %v, 类型: %T\n", result, result)

	// 根据返回结果判断是否允许访问
	if result.(int64) == 1 {
		fmt.Println("访问成功！")
	} else {
		fmt.Println("访问被拒绝，已达到访问上限！")
	}
}

// 我们将脚本的定义和业务逻辑分离开来。limiterScript 可以在包级别定义，随处复用。go-redis 在底层为我们处理了 SHA1 的缓存和 EVALSHA 的调用，非常优雅。
func lua_demo2(rdb *redis.Client) {
	ctx := context.Background()

	//go-redis 提供了 redis.NewScript 函数，它会返回一个 *redis.Script 对象。这个对象非常智能，它的 Run 方法会优先尝试使用 EVALSHA 执行脚本。如果 Redis 返回 NOSCRIPT 错误（表示脚本缓存中不存在该 SHA1 对应的脚本），它会自动回退（fallback），
	//改用 EVAL 命令执行脚本原文，并将脚本重新加载到缓存中。后续的调用又会回到 EVALSHA 的轨道上。
	// 使用 redis.NewScript 创建一个脚本对象
	limiterScript := redis.NewScript(`
		local current = redis.call('get', KEYS[1])
		if current and tonumber(current) >= tonumber(ARGV[2]) then
			return 0
		end
		local count = redis.call('incr', KEYS[1])
		if count == 1 then
			redis.call('expire', KEYS[1], ARGV[1])
		end
		return 1
	`)

	key := "ip:127.0.0.1"
	expire := 60 // 60s
	limit := 3   // 3次

	for i := 0; i < 5; i++ {
		// 使用脚本对象的 Run 方法执行
		// func (s *Script) Run(ctx context.Context, c Scripter, keys []string, args ...interface{}) *Cmd
		result, err := limiterScript.Run(ctx, rdb, []string{key}, expire, limit).Result()
		if err != nil && err != redis.Nil {
			panic(err)
		}

		if result.(int64) == 1 {
			fmt.Printf("第 %d 次访问: 成功！\n", i+1)
		} else {
			fmt.Printf("第 %d 次访问: 被拒绝！\n", i+1)
		}
	}
}

func lua_demo3(rdb *redis.Client, limiterScript *redis.Script) {
	ctx := context.Background()
	key := "ip:192.168.1.1" // 换个IP测试
	expire := 60
	limit := 3

	fmt.Println("使用 embed 方式加载的 Lua 脚本:")
	fmt.Println("--------------------------------")
	fmt.Println(limiterScriptSource)
	fmt.Println("--------------------------------")

	for i := 0; i < 5; i++ {
		result, err := limiterScript.Run(ctx, rdb, []string{key}, expire, limit).Result()
		if err != nil && err != redis.Nil {
			panic(err)
		}

		if result.(int64) == 1 {
			fmt.Printf("第 %d 次访问: 成功！\n", i+1)
		} else {
			fmt.Printf("第 %d 次访问: 被拒绝！\n", i+1)
		}
	}
}
