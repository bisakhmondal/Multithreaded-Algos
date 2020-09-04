package main

import "fmt"

//insertion sort
func sort(ar []int, c chan []int){
	for i :=0;i<len(ar);i++{
		
		key := ar[i]
		j := i-1
		
		for j>=0 && ar[j]>key{
			
			ar[j+1] = ar[j]
			j--
		}
		ar[j+1]=key

	}

	c<- ar
}

//merge operation to merge the small slices.
func merge(a []int,b []int) []int{
	i,j:=0,0
	var c []int
	for i<len(a) && j<len(b){
		if a[i]<b[j]{
			c = append(c,a[i])
			i++
		}else{
			c = append(c,b[j])
			j++
		}
	}
	for i<len(a) {
		c = append(c,a[i])
		i++
	}
	for j<len(b){
		c = append(c,b[j])
		j++
	}
	return c
}

func main(){
	
	//taking inputs
	var n int
	fmt.Println("Enter Number of elements: ")
	
	fmt.Scan(&n)
	arr := make([]int, n)
	for idx,i := range arr{
		fmt.Scan(&i)
		arr[idx]=i
	}
	
	// creating channels with buffer length 4.
	c := make(chan []int,4)
	
	//pulling 4 threads for insertion sort on quadrant of slice.
	go sort(arr[:n/4],c)
	go sort(arr[n/4:n/2],c)
	go sort(arr[n/2:3*n/4],c)
	go sort(arr[3*n/4:n],c)
	
	var op []int
	
	for i:=0;i<4;{
		select{
		case cop := <-c:
			//fmt.Println(cop)
			op = merge(op,cop)
			i++
			//fmt.Println(op)
		}
	}
	
	//final slice
	fmt.Println(op)
}
