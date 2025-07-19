package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // 密码
		DB:       0,  // 数据库
		PoolSize: 20, // 连接池大小
	})
	//执行原生命令
	doCommand(rdb)
	fmt.Println("-----------------------------")
	//执行自定义命令
	doDemo(rdb)
	fmt.Println("-----------------------------")
	//zsetDemo zset操作示例
	zsetDemo(rdb)
	fmt.Println("-----------------------------")
	// scanKeysDemo1 按前缀查找所有key示例
	scanKeysDemo1(rdb)
	fmt.Println("-----------------------------")
	// scanKeysDemo2 按前缀查找所有key示例
	scanKeysDemo2(rdb)
	fmt.Println("-----------------------------")
	//pipeline demo
	pipeline_demo1(rdb)
	pipeline_demo2(rdb)
	fmt.Println("-----------------------------")
	watch_demo1(rdb)
}

// doCommand go-redis基本使用示例
func doCommand(rdb *redis.Client) {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	// 执行命令获取结果
	val, err := rdb.Get(ctx, "key").Result()
	fmt.Println("val", val, "err", err)

	// 先获取到命令对象
	cmder := rdb.Get(ctx, "key")
	fmt.Println("val", cmder.Val()) // 获取值
	fmt.Println("err", cmder.Err()) // 获取错误

	// 直接执行命令获取错误
	err = rdb.Set(ctx, "key", 10, time.Hour).Err()

	// 直接执行命令获取值
	value := rdb.Get(ctx, "key").Val()
	fmt.Println(value)
}

// doDemo rdb.Do 方法使用示例
func doDemo(rdb *redis.Client) {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	// 直接执行命令获取错误
	err := rdb.Do(ctx, "set", "v1", 10, "NX", "EX", 3600).Err()
	fmt.Println(err)

	// 执行命令获取结果
	val, err := rdb.Do(ctx, "get", "v1").Result()
	fmt.Println(val, err)
}

// zsetDemo 操作zset示例
func zsetDemo(rdb *redis.Client) {
	// key
	zsetKey := "language_rank"
	// value
	// 注意：v8版本使用[]*redis.Z；此处为v9版本使用[]redis.Z
	languages := []redis.Z{
		{Score: 90.0, Member: "Golang"},
		{Score: 98.0, Member: "Java"},
		{Score: 95.0, Member: "Python"},
		{Score: 97.0, Member: "JavaScript"},
		{Score: 99.0, Member: "C/C++"},
	}
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	// ZADD
	err := rdb.ZAdd(ctx, zsetKey, languages...).Err()
	if err != nil {
		fmt.Printf("zadd failed, err:%v\n", err)
		return
	}
	fmt.Println("zadd success")

	// 把Golang的分数加10
	newScore, err := rdb.ZIncrBy(ctx, zsetKey, 10.0, "Golang").Result()
	if err != nil {
		fmt.Printf("zincrby failed, err:%v\n", err)
		return
	}
	fmt.Printf("Golang's score is %f now.\n", newScore)

	// 取分数最高的3个
	ret := rdb.ZRevRangeWithScores(ctx, zsetKey, 0, 2).Val()
	for _, z := range ret {
		fmt.Println(z.Member, z.Score)
	}

	// 取95~100分的
	op := &redis.ZRangeBy{
		Min: "95",
		Max: "100",
	}
	ret, err = rdb.ZRangeByScoreWithScores(ctx, zsetKey, op).Result()
	if err != nil {
		fmt.Printf("zrangebyscore failed, err:%v\n", err)
		return
	}
	for _, z := range ret {
		fmt.Println(z.Member, z.Score)
	}
}

// vals, err := rdb.Keys(ctx, "user:*").Result() 这种方式性能较差
// scanKeysDemo1 按前缀查找所有key示例
func scanKeysDemo1(rdb *redis.Client) {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	var cursor uint64
	for {
		var keys []string
		var err error
		// 将redis中所有以prefix:为前缀的key都扫描出来
		keys, cursor, err = rdb.Scan(ctx, cursor, "user:*", 0).Result()
		if err != nil {
			panic(err)
		}

		for _, key := range keys {
			fmt.Println("key", key)
		}

		if cursor == 0 { // no more keys
			break
		}
	}
}

// scanKeysDemo2 按前缀扫描key示例
func scanKeysDemo2(rdb *redis.Client) {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	// 按前缀扫描key
	iter := rdb.Scan(ctx, 0, "user:*", 0).Iterator()
	for iter.Next(ctx) {
		fmt.Println("keys", iter.Val())
	}
	if err := iter.Err(); err != nil {
		panic(err)
	}
}
