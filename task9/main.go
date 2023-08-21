package main

import (
    "fmt"
    "sync"
)

type Worker struct{
	wg *sync.WaitGroup
	inputChanel chan int
	outputChanel chan int
}

type WorkerPool struct{
	pool []Worker
}

func CreatePool(n int, wgr *sync.WaitGroup, chaIn chan int,chaOut chan int) *WorkerPool{
	poolOfWorkers := make([]Worker,0,5)
	for i:=0; i< n;i++{
		poolOfWorkers =append(poolOfWorkers,Worker{wg:wgr,inputChanel:chaIn,outputChanel:chaOut})
	}
	result := WorkerPool{pool:poolOfWorkers}
	return &result
}

func (w *Worker) Work(index int){
	defer w.wg.Done()
    for  data  := range w.inputChanel{
		result := data * 2 
		w.outputChanel  <- result
          
}
}
    

func (p *WorkerPool) StartWork(){
	for i,worker:=range p.pool{
		go worker.Work(i)
	}
}

 
 
func main(){
	fmt.Println("Number of workers: ")
	a:=0
	fmt.Scanf("%d",&a)

	 
	arr :=  [15]int {}
	for i:=0;i<15;i++{
		arr[i] = i 
	}

	var wg sync.WaitGroup
	wg.Add(a)

	wg.Add(1)

	chIn := make(chan int)
	chOut := make(chan int)


	p := CreatePool(a,&wg,chIn,chOut)
	 
	p.StartWork()

	 
	go func() {
		defer wg.Done() // пишем в канал
        for i := range arr {
            chIn <- i
            // time.Sleep(time.Second )
        }
		close(chIn)
    }()

	go func() { // читаем из канала
		for i :=range chOut{
			fmt.Println(i)
		}
    }()

		wg.Wait()
		close(chOut)
	 
}