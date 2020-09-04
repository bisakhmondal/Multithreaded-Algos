package sort

import (
	"runtime"
	"fmt"
)

 
//MaxParallelism function to help with inquiring go runtime.
func MaxParallelism() (int,int) {
    maxProcs := runtime.GOMAXPROCS(0)
	numCPU := runtime.NumCPU()
	fmt.Println("Max concurrent threads supported: ",numCPU)
	// runtime.GOMAXPROCS(procs)
	fmt.Printf("Running on threads: %d\n",maxProcs)
    return numCPU,maxProcs
}
/*
SetProcsCount to change number of concurrent threads for multithreaded system. 

@param procs Change number of logical Processor
*/
func SetProcsCount(procs int){
	// numCPU := runtime.NumCPU()
	
	// if procs<=numCPU{
		runtime.GOMAXPROCS(procs)
	// }else{
	// 	runtime.GOMAXPROCS(numCPU)
	// }
}