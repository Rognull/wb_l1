package main

import (
    "fmt"
    "os"
    "os/signal"
    "sync"
    "syscall"
    "time"
)

type Worker struct{
	wg *sync.WaitGroup
	ch chan int
}

type WorkerPool struct{
	pool []Worker
}

func CreatePool(n int, wgr *sync.WaitGroup, cha chan int) *WorkerPool{
	poolOfWorkers := make([]Worker,0,5)
	for i:=0; i< n;i++{
		poolOfWorkers =append(poolOfWorkers,Worker{wg:wgr,ch:cha})
	}
	result := WorkerPool{pool:poolOfWorkers}
	return &result
}

func (w *Worker) Work(index int){
	defer w.wg.Done()
    for  data  := range w.ch{
            fmt.Printf("Worker#%d got data - %d\n",index,data)
}
}
    

func (p *WorkerPool) StartWork(){
	for i,worker:=range p.pool{
		go worker.Work(i)
	}
}

func GracefulShutdown(ch chan int, wg *sync.WaitGroup) {
	interruptChannel := make(chan os.Signal, 1)
	signal.Notify(interruptChannel, syscall.SIGINT, syscall.SIGTERM)
	<-interruptChannel

	fmt.Println("Exiting")

	close(ch)
	wg.Wait()

	fmt.Println("Exit done.")
}
 
func main(){
	fmt.Println("Number of workers: ")
	a:=0
	fmt.Scanf("%d",&a)

	var wg sync.WaitGroup
	wg.Add(a)

	ch := make(chan int)

	p := CreatePool(a,&wg,ch)
	 
	p.StartWork()

	go func() {
        for i := 0; ; i++ {
            ch <- i
            time.Sleep(time.Second  )
        }
    }()

	GracefulShutdown(ch,&wg)

}