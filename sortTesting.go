package main

import (
	"Multithreaded-Algos/sort"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

func ConvertInts(s []string) ([]int, error) {
	sli := make([]int, len(s))

	for idx, st := range s {
	
		i, j := strconv.Atoi(st)
	
		if j != nil {
	
			return sli, j
		}
	
		sli[idx] = i
	}

	return sli, nil
}
func ConvertFloats(s []string) ([]float32, error) {
	
	slf := make([]float32, len(s))

	for idx, st := range s {
	
		i, j := strconv.ParseFloat(st, 32)
	
		if j != nil {
	
			return slf, j
		}
	
		slf[idx] = float32(i)
	}
	
	return slf, nil
}
func main() {
	sort.SetProcsCount(4)
	
	fmt.Println("Reading Filestream")
	
	barr, err := ioutil.ReadFile("./tests/nums1e6.txt")
	fmt.Println("reading Done")
	
	if err == nil {
		arr := strings.Split(string(barr), " ")
	
		sli, _ := ConvertInts(arr)

		start := time.Now()
		sort.MergeSortInts(sli)
		
		fmt.Println("time taken: ",time.Since(start))
		
		fmt.Println("sorting done")

	} else {
		fmt.Println(err)
		return
	}
	
}
