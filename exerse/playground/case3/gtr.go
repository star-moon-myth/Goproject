package main

import (
	"fmt"
	"sync"
	"time"
)
// 记得关闭通道啊，不然就会一直等待通道里的东西
// 或者等待等待通道里面的东西出来
// 另外，控制goroutinue的sync，waitgroup也要记得关闭（wait阻塞当前线程，直到计数器归零）
func worker(i int ,wg *sync.WaitGroup,result chan string,jobs chan int){
	defer wg.Done()
	fmt.Printf("worker%d开始执行\n",i)
	for job:=range jobs{
		fmt.Printf("worker%d任务%d开始执行\n",i,job)
		time.Sleep(time.Millisecond*100)
		result<-fmt.Sprintf("任务%d执行完毕\n",job)
	}
	fmt.Printf("worker %d 执行完毕\n",i)

}

func main() {
	numjobs := 10
	numworker := 5
	var wg sync.WaitGroup
	jobs:=make(chan int,numjobs)
	result:=make(chan string,numjobs)

	for w:=0;w<numworker;w++{
		wg.Add(1)
		go worker(w,&wg,result,jobs)
	}
	fmt.Println("任务开始分发！")
	for i:=0;i<numjobs;i++{
		jobs<-i
	}

	fmt.Println("任务分发完毕！")
	close(jobs)
	wg.Wait()
	close(result)
	for i:=range result{
		fmt.Println(i)
	}
	fmt.Println("🎉 All jobs done!")
}