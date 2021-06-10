package main

import (
	"context"
	"fmt"
	"runtime"
	"strconv"
	"time"
)

func main() {

	// c := make(chan string)
	// stopCh := make(chan error)
	// fmt.Println("运行的groutine个数->" + strconv.Itoa(runtime.NumGoroutine()))
	// for i := 0; i < 10; i++ {
	// 	go do1(c)
	// }

	// time.Sleep(1 * time.Second)
	// fmt.Println("运行的groutine个数->" + strconv.Itoa(runtime.NumGoroutine()))
	// go func(t2 chan string) {

	// 	dasd := <-t2
	// 	fmt.Println("dsad->" + dasd) //go routine执行完成了
	// }(c)

	// for i := 0; i < 5; i++ {
	// 	go func(tempc chan string, sp chan error) {
	// 		defer func() {
	// 			fmt.Println("退出了goroutine方法了")
	// 		}()

	// 		// for item := range tempc {
	// 		// 	println("item->" + item)
	// 		// }

	// 		for {

	// 			select {

	// 			case item, ok := <-tempc:
	// 				fmt.Println("item->" + item)
	// 				if !ok {
	// 					return
	// 				}
	// 			case <-sp:
	// 				return
	// 			}
	// 		}
	// 	}(c, stopCh)
	// }

	// time.Sleep(5 * time.Second)

	// //再往c1里面去发送数据
	// c <- "单独发送的"
	// println("执行了单独返送")
	// time.Sleep(2 * time.Second)
	// fmt.Println("运行的groutine个数->" + strconv.Itoa(runtime.NumGoroutine()))
	// close(c)
	// time.Sleep(1 * time.Second)
	// fmt.Println("运行的groutine个数->" + strconv.Itoa(runtime.NumGoroutine()))
	// time.Sleep(10 * time.Second)
	// c := make(chan string)
	// cx, cancel := context.WithCancel(context.Background())

	// go docontext(cx, c)

	// time.Sleep(1 * time.Second)
	// c <- "测试数据"

	// go func(doc context.CancelFunc) {
	// 	time.Sleep(2 * time.Second)
	// 	doc()
	// }(cancel)
	// // time.Sleep(1 * time.Second)
	// // cancel()

	// time.Sleep(4 * time.Second)

	// recv := make(chan int)
	// pass := make(chan int)

	// go doAdd(recv, pass)
	// go doMulti(recv, pass)
	// go doDivi(recv, pass)
	// go resultGo(pass)

	// time.Sleep(1 * time.Second)

	// for i := 0; i < 4; i++ {
	// 	recv <- i
	// }

	// time.Sleep(2 * time.Second)
	// ochan := make(chan string)
	// innerchan := make(chan string)
	// go outgo(innerchan, ochan)
	// time.Sleep(1 * time.Second)
	// fmt.Println("现存的gorotine->" + strconv.Itoa(runtime.NumGoroutine()))
	// time.Sleep(1 * time.Second)
	// ochan <- "外部结束"
	// // innerchan <- "内部结束"
	// time.Sleep(1 * time.Second)
	// fmt.Println("最后有几个gorotine->" + strconv.Itoa(runtime.NumGoroutine()))
	// time.Sleep(3 * time.Second)

	// var input string
	// fmt.Scanln(&input)

	//针对单一的groutine进行控制的时候可以使用chan
	ch1 := make(chan string, 3)

	ctx, cancel := context.WithCancel(context.Background())

	for i := 0; i < 3; i++ {
		go commGro(ch1, ctx)
	}
	time.Sleep(1 * time.Second)
	fmt.Println("运行的goroutine数量是1111->" + strconv.Itoa(runtime.NumGoroutine()))
	ch1 <- "测试"
	ch1 <- "测试2"
	ch1 <- "测试3"
	ch1 <- "测试4"
	time.Sleep(time.Second)
	fmt.Println("channel里面个数->" + strconv.Itoa(len(ch1)) + "    " + strconv.Itoa(cap(ch1)))

	time.Sleep(time.Second)
	fmt.Println("运行的goroutine数量是2222->" + strconv.Itoa(runtime.NumGoroutine()))

	time.Sleep(time.Second)
	cancel()
	time.Sleep(time.Second)
	fmt.Println("运行的goroutine数量是3333->" + strconv.Itoa(runtime.NumGoroutine()))
	var input string
	fmt.Scan(&input)
}

//多个groutine用了同一个channel，向公用的channel里面发送数据，是所有groutine都执行还是只有一个去执行

func commGro(c1 chan string, ctx context.Context) {

	//则塞执行
	for {

		select {
		case <-ctx.Done():
			fmt.Println("group执行了done")
			return
		case tempstr := <-c1:
			fmt.Println("运行参数->" + tempstr)
			return
		}
	}
}

func do1(c1 chan string) {

	time.Sleep(2 * time.Second)
	c1 <- "c1"
}

func do2(c2 chan string) {
	time.Sleep(3 * time.Second)
	c2 <- "c2"
}

//使用context去取消

func docontext(ctx context.Context, c1 chan string) {

	defer func() {
		println("docontext ->  结束")
	}()

	for {

		select {
		case <-ctx.Done():
			fmt.Println("执行done")
			return
		case str := <-c1:
			fmt.Println("执行了c1->" + str)
		}
	}
}

//例子，一个goroutine执行数字+1  一个goroutine执行数字*3   一个go routine执行数字除以2
func doAdd(recv chan int, pass chan int) {

	for {
		number, ok := <-recv
		if !ok {
			fmt.Println("接收数据通道关闭")
			return
		}
		number += 1
		fmt.Println("doAdd:" + strconv.Itoa(number))
		pass <- number
	}
}

func doMulti(recv chan int, pass chan int) {

	for {
		number, ok := <-recv
		if !ok {
			fmt.Println("接收数据通道关闭")
			return
		}
		number *= 3
		fmt.Println("doMulti:" + strconv.Itoa(number))
		pass <- number
	}
}

func doDivi(recv chan int, pass chan int) {

	for {
		number, ok := <-recv
		if !ok {
			fmt.Println("接收数据通道关闭")
			return
		}
		number /= 2
		fmt.Println("doDivi:" + strconv.Itoa(number))
		pass <- number
	}

}

func resultGo(recv chan int) {
	var total int = 0
	for {
		number, ok := <-recv
		if !ok {
			return
		}
		total += number
		fmt.Println("计算的到的总数据是->" + strconv.Itoa(total))
	}
}

//一个gorountine里面有个goroutine
func outgo(innerch chan string, outerCh chan string) {

	defer func() {
		fmt.Println("outgo 结尾的部分")
	}()
	go innergo(innerch)
	<-outerCh
}

func innergo(sp chan string) {

	defer func() {

		fmt.Println("innergo 内部执行完成")
	}()

	// for {

	str := <-sp
	fmt.Println("str->" + str)
	// }
}
