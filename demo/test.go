package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	//WaitGroup
	//var wg sync.WaitGroup
	//wg.Add(2)
	//go func() {
	//	time.Sleep(2 * time.Second)
	//	fmt.Println("1号完成")
	//	wg.Done()
	//}()
	//go func() {
	//	time.Sleep(2 * time.Second)
	//	fmt.Println("2号完成")
	//	wg.Done()
	//}()
	//wg.Wait()
	//fmt.Println("好了，大家都干完了，放工")
	////chan
	//stop := make(chan bool)
	//go func() {
	//	for {
	//		select {
	//		case <-stop:
	//			fmt.Println("监控退出，停止了...")
	//			return
	//		default:
	//			fmt.Println("goroutine监控中...")
	//			time.Sleep(2 * time.Second)
	//		}
	//	}
	//}()
	//time.Sleep(10 * time.Second)
	//fmt.Println("可以了，通知监控停止")
	//stop <- true
	////为了检测监控过是否停止，如果没有监控输出，就表示停止了
	//time.Sleep(5 * time.Second)
	////context
	//ctx, cancel := context.WithCancel(context.Background())
	//go watch(ctx, "【监控1】")
	//go watch(ctx, "【监控2】")
	//go watch(ctx, "【监控3】")
	//time.Sleep(10 * time.Second)
	//fmt.Println("可以了，通知监控停止")
	//cancel()
	////为了检测监控过是否停止，如果没有监控输出，就表示停止了
	//time.Sleep(5 * time.Second)

	tickDemo()
}

func watch(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "监控退出，停止了...")
			return
		default:
			fmt.Println(name, "goroutine监控中...")
			time.Sleep(2 * time.Second)
		}
	}
}

func tickDemo() {
	ticker := time.Tick(5 * time.Second)
	for i := range ticker {
		fmt.Println(i)
	}
}
