package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"sync"
	"time"
)

// 事务demo TxPipeline相当于使用MUTI和EXEC把多个命令包起来
// 这里的事务是为了保证多个命令执行的时候，不会有其他客户端的命令插进来
func tx_demo1(rdb *redis.Client) {
	ctx := context.Background()
	// TxPipeline demo
	pipe := rdb.TxPipeline()
	incr := pipe.Incr(ctx, "tx_pipeline_counter")
	pipe.Expire(ctx, "tx_pipeline_counter", time.Hour)
	_, err := pipe.Exec(ctx)
	fmt.Println(incr.Val(), err)

	// TxPipelined demo
	var incr2 *redis.IntCmd
	_, err = rdb.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
		incr2 = pipe.Incr(ctx, "tx_pipeline_counter")
		pipe.Expire(ctx, "tx_pipeline_counter", time.Hour)
		return nil
	})
	fmt.Println(incr2.Val(), err)

}

// watchDemo 在key值不变的情况下将其值+1
// watch的作用是监控一个key的变化，如果key的值没有发生变化，那么就执行事务，否则该事务操作失败，返回redis: transaction failed错误。
func watchDemo(ctx context.Context, key string, rdb *redis.Client) error {
	return rdb.Watch(ctx, func(tx *redis.Tx) error {
		n, err := tx.Get(ctx, key).Int()
		if err != nil && err != redis.Nil {
			return err
		}
		// 假设操作耗时5秒
		// 5秒内我们通过其他的客户端修改key，当前事务就会失败
		time.Sleep(5 * time.Second)
		_, err = tx.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
			pipe.Set(ctx, key, n+1, time.Hour)
			return nil
		})
		return err
	}, key)
}

// 使用 GET + SET + WATCH 实现，类似 INCR
func watch_demo1(rdb *redis.Client) {
	// 此处rdb为初始化的redis连接客户端
	const routineCount = 100

	// 设置5秒超时
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// increment 是一个自定义对key进行递增（+1）的函数
	// 使用 GET + SET + WATCH 实现，类似 INCR
	increment := func(key string) error {
		txf := func(tx *redis.Tx) error {
			// 获得当前值或零值
			n, err := tx.Get(ctx, key).Int()
			if err != nil && err != redis.Nil {
				return err
			}

			// 实际操作（乐观锁定中的本地操作）
			n++

			// 仅在监视的Key保持不变的情况下运行
			_, err = tx.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
				// pipe 处理错误情况
				pipe.Set(ctx, key, n, 0)
				return nil
			})
			return err
		}

		// 最多重试100次
		for retries := routineCount; retries > 0; retries-- {
			err := rdb.Watch(ctx, txf, key)
			log.Println(err)
			if err != redis.TxFailedErr {
				return err
			}
			// 乐观锁丢失
		}
		return errors.New("increment reached maximum number of retries")
	}

	// 开启100个goroutine并发调用increment
	// 相当于对key执行100次递增
	var wg sync.WaitGroup
	wg.Add(routineCount)
	for i := 0; i < routineCount; i++ {
		go func() {
			defer wg.Done()

			if err := increment("counter3"); err != nil {
				fmt.Println("increment error:", err)
			}
		}()
	}
	wg.Wait()

	n, err := rdb.Get(ctx, "counter3").Int()
	fmt.Println("最终结果：", n, err)

}
