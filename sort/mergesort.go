package sort

// LessThan : A comparator function to compare generic interface type for sorting.
// @params a,b of type interface{} and
// @returns either true if a<b else false
type LessThan func(a, b interface{}) bool

type sortHandler struct {
	data []interface{}

	comp LessThan
}

const (
	BufferSize int = 128
)

/* deal with it later though it may beunstable
// func split(slice []interface{}, lt LessThan,){

// }

// MergeSort ... The O(Nlog(N)) sorting using multithreading
// func MergeSort(slice []interface{}, lt LessThan) []interface{} {
// 	MaxParallelism()

// 	mergeHandler := new(sortHandler)
// 	mergeHandler.data = slice
// 	mergeHandler.comp = lt

// 	ch := make(chan []interface{},)

// }
*/

//Internal function to merge two halves of slices into one.
func merge(l, r []interface{}, lt LessThan, c chan []interface{}) {

	//create an slice of interface to store the slices after merging.
	result := make([]interface{}, len(l)+len(r))

	i, j, idx := 0, 0, 0
	for ; idx < cap(result) && i < len(l) && j < len(r); idx++ {
		if lt(l[i], r[j]) {
			result[idx] = l[i]
			i++
		} else {
			result[idx] = r[j]
			j++
		}
	}
	for i < len(l) {
		result[idx] = l[i]
		i++
		idx++
	}
	for j < len(r) {
		result[idx] = r[j]
		j++
		idx++
	}

	//passing into channel
	c <- result
}

// Internal Function to split int slice into two half and call mergesort on the same.
func sort(data []interface{}, lt LessThan, c chan []interface{}) {
	
	//base case
	if len(data) == 1 {
		c <- data
		return
	}
	
	mid := len(data) / 2
	
	//left and right channel for mainining stable sorting.
	lchan := make(chan []interface{})
	rchan := make(chan []interface{})

	// call two goroutines for left and right part concurrently.
	go sort(data[:mid], lt, lchan)
	go sort(data[mid:], lt, rchan)

	//receive data from the channels to merge them further.
	ldata := <-lchan
	rdata := <-rchan

	// closing the channel to mark that the use is done.
	close(lchan)
	close(rchan)

	//creating another goroutine to merge two slices with channel as argument. 
	go merge(ldata, rdata, lt, c)

}

// ToInterface creates array of interface from any slice.
func ToInterface(data ...interface{}) []interface{} {

	result := make([]interface{}, len(data))

	for idx, dat := range data {
		result[idx] = dat
	}

	return result
}

//StableMergeSort ...
func StableMergeSort(data []interface{}, lt LessThan) []interface{} {
	
	MaxParallelism()
	
	//channel to receive final output
	outC := make(chan []interface{})

	go sort(data, lt, outC)

	output := <-outC

	return output
}

// Providing Most used API for Int, Float32 & Strings data type
/*
MergeSortInts attaches StableMergeSort for integer slices with custom comparator.
 @param data  integer slice
*/
func MergeSortInts(data []int) []int {

	comparator := func(a, b interface{}) bool { return a.(int) < b.(int) }

	intf := make([]interface{}, len(data))

	//converting integer slice to generic interface
	for idx, dat := range data {
		intf[idx] = dat
	}

	output := StableMergeSort(intf, comparator)

	//Making inplace operation.
	for idx,op := range output{
		data[idx] = op.(int)
	}
	return data
}

/*
MergeSortFloats attaches StableMergeSort for integer slices with custom comparator.
 @param data  float32 slice
*/
func MergeSortFloats(data []float32) []float32 {

	comparator := func(a, b interface{}) bool { return a.(float32) < b.(float32) }

	intf := make([]interface{}, len(data))

	//converting float32 slice to generic interface
	for idx, dat := range data {
		intf[idx] = dat
	}

	output := StableMergeSort(intf, comparator)

	//Making inplace operation.
	for idx, op := range output {
		data[idx] = op.(float32)
	}
	return data
}


/*
MergeSortStrings attaches StableMergeSort for integer slices with custom comparator.
 @param data  float32 slice
*/
func MergeSortStrings(data []string) []string {

	comparator := func(a, b interface{}) bool { return a.(string) < b.(string) }

	intf := make([]interface{}, len(data))

	//converting string slice to generic interface
	for idx, dat := range data {
		intf[idx] = dat
	}

	output := StableMergeSort(intf, comparator)

	//Making inplace operation.
	for idx, op := range output {
		data[idx] = op.(string)
	}
	return data
}