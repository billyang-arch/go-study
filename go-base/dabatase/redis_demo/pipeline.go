package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

// pipeline 一次性执行多个命令，减少网络开销，即批量执行
func pipeline_demo1(rdb *redis.Client) {
	ctx := context.Background()
	//方式一
	//pipe := rdb.Pipeline()

	//incr := pipe.Incr(ctx, "pipeline_counter")
	//pipe.Expire(ctx, "pipeline_counter", time.Hour)
	//
	//_, err := pipe.Exec(ctx)
	//if err != nil {
	//	panic(err)
	//}
	//
	//// 在执行pipe.Exec之后才能获取到结果
	//fmt.Println(incr.Val())

	//方式二
	var incr *redis.IntCmd

	_, err := rdb.Pipelined(ctx, func(pipe redis.Pipeliner) error {
		incr = pipe.Incr(ctx, "pipelined_counter")
		pipe.Expire(ctx, "pipelined_counter", time.Hour)
		return nil
	})
	if err != nil {
		panic(err)
	}

	// 在pipeline执行后获取到结果
	fmt.Println(incr.Val())

}

func pipeline_demo2(rdb *redis.Client) {
	ctx := context.Background()
	cmds, err := rdb.Pipelined(ctx, func(pipe redis.Pipeliner) error {
		for i := 1; i < 4; i++ {
			pipe.Get(ctx, fmt.Sprintf("key%d", i))
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

	for _, cmd := range cmds {
		fmt.Println(cmd.(*redis.StringCmd).Val())
	}

}
