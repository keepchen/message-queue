// package main

// import (
// 	"fmt"
// 	"message-queue/queue"
// 	"sync"
// 	"time"
// )

// func main() {
// 	// instance := queue.GetDBInstance("db1")
// 	// instance.RPop()
// 	// instance.LPush("a")
// 	// instance.RPush(0)
// 	// instance.LPop()
// 	// instance.RPush(1)
// 	// instance.RPop()
// 	// instance.RPop()
// 	// instance.RPush(2)
// 	// instance.LPush(3)
// 	// instance.DisplayQueue()

// 	benchmark(1000000)
// }

// func benchmark(ops int) {
// 	benchmarkLPush(ops)
// 	benchmarkRPush(ops)
// 	benchmarkLPop(ops)
// 	benchmarkRPop(ops)

// 	// benchmarkPushAndPop(ops)
// }

// func benchmarkLPush(ops int) {
// 	start := time.Now().UnixNano()
// 	instance := queue.GetDBInstance("db0")
// 	for i := 0; i < ops; i++ {
// 		instance.LPush(i)
// 	}
// 	end := time.Now().UnixNano()
// 	cost := float64(end-start) / 1e6

// 	fmt.Printf("------------------\n<benchmark>\n%d次[LPush]性能测试\n------------------\n-->操作次数: %d 次\n-->耗时: %f 毫秒\n-->队列长度: %d\n",
// 		ops, ops, cost, instance.GetSize())
// }

// func benchmarkRPush(ops int) {
// 	start := time.Now().UnixNano()
// 	instance := queue.GetDBInstance("db1")
// 	for i := 0; i < ops; i++ {
// 		instance.RPush(i)
// 	}
// 	end := time.Now().UnixNano()
// 	cost := float64(end-start) / 1e6

// 	fmt.Printf("------------------\n<benchmark>\n%d次[RPush]性能测试\n------------------\n-->操作次数: %d 次\n-->耗时: %f 毫秒\n-->队列长度: %d\n",
// 		ops, ops, cost, instance.GetSize())
// }

// func benchmarkLPop(ops int) {
// 	instance := queue.GetDBInstance("db2")
// 	for i := 0; i < ops; i++ {
// 		instance.LPush(i)
// 	}
// 	start := time.Now().UnixNano()
// 	for i := 0; i < ops; i++ {
// 		instance.LPop()
// 	}
// 	end := time.Now().UnixNano()
// 	cost := float64(end-start) / 1e6

// 	fmt.Printf("------------------\n<benchmark>\n%d次[LPop]性能测试\n------------------\n-->操作次数: %d 次\n-->耗时: %f 毫秒\n-->队列长度: %d\n",
// 		ops, ops, cost, instance.GetSize())
// }

// func benchmarkRPop(ops int) {
// 	instance := queue.GetDBInstance("db3")
// 	for i := 0; i < ops; i++ {
// 		instance.LPush(i)
// 	}
// 	start := time.Now().UnixNano()
// 	for i := 0; i < ops; i++ {
// 		instance.RPop()
// 	}
// 	end := time.Now().UnixNano()
// 	cost := float64(end-start) / 1e6

// 	fmt.Printf("------------------\n<benchmark>\n%d次[RPop]性能测试\n------------------\n-->操作次数: %d 次\n-->耗时: %f 毫秒\n-->队列长度: %d\n",
// 		ops, ops, cost, instance.GetSize())
// }

// func benchmarkPushAndPop(ops int) {
// 	var wg sync.WaitGroup
// 	start := time.Now().UnixNano()
// 	instance := queue.GetDBInstance("db4")
// 	wg.Add(2)
// 	go func() {
// 		for i := 0; i < ops; i++ {
// 			instance.RPush(i)
// 		}
// 		wg.Done()
// 	}()

// 	go func(instance *queue.Instance) {
// 		var times int
// 		for {
// 			if times == ops {
// 				break
// 			}
// 			if instance.GetSize() == 0 {
// 				continue
// 			}
// 			_, e := instance.RPop()
// 			if e == nil {
// 				times++
// 			}
// 		}
// 		wg.Done()
// 	}(instance)

// 	wg.Wait()

// 	end := time.Now().UnixNano()
// 	cost := float64(end-start) / 1e6
// 	fmt.Printf("------------------\n<benchmark>\n%d次同时[RPush|RPop]性能测试\n------------------\n-->操作次数: %d 次\n-->耗时: %f 毫秒\n-->队列长度: %d\n",
// 		ops, ops*2, cost, instance.GetSize())

// 	instance.FlushDB()

// 	// instance.LPush("x")
// }
