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
}
