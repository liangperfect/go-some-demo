package goroutine_demo

import (
	"errors"
	"fmt"
)

type GoDemo struct {
	Name   string
	StopCh chan error
}

func (d *GoDemo) StartServer() {

	go d.initServer()
}

func (d *GoDemo) initServer() {

	go func() {

		// tchan := time.Tick(2 * time.Second)
		// for {
		// 	fmt.Println("开始执行for循环")
		// 	select {
		// 	case t1 := <-d.StopCh:
		// 		fmt.Println("停止的stop测试->" + t1.Error())

		// 		// default:
		// 		// 	c := <-tchan
		// 		// 	fmt.Println("打印出数据:" + c.Format("2006-01-02 15:04:05"))
		// 	}
		// }
		fmt.Println("等待接收到的信息")
		t1, err := <-d.StopCh
		if !err {
			fmt.Println("通道已关闭")
			return
		}
		fmt.Println("停止的stop测试->" + t1.Error())
	}()

	temp := <-d.StopCh
	fmt.Println("接收到的数据是->" + temp.Error())
}

func (d *GoDemo) StopServer() {

	d.StopCh <- errors.New("停止了某个channel")
}
