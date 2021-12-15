package main

import (
	"TFL/Contest/prover"
	"fmt"
	"io/ioutil"
)

func main() {
	input, err := ioutil.ReadFile("test.trs") //test.trs
	if err != nil {
		fmt.Println(err)
		return
	}
	prover.HandleTask(input)

	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(arr)
	indexes := make([]int, len(arr))
	i := 0
	for i < len(arr) {
		if indexes[i] < i {
			var in int
			if i%2 == 0 {
				in = 0
			} else {
				in = indexes[i]
			}
			tmp := arr[in]
			arr[in] = arr[i]
			arr[i] = tmp
			fmt.Println(arr)
			indexes[i]++
			i = 0
		} else {
			indexes[i] = 0
			i++
		}
	}
}
