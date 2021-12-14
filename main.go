package main

import (
	"TFL/Contest/prover"
	"fmt"
	"io/ioutil"
)

func main() {
	input, err := ioutil.ReadFile("./resources/trs.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	prover.HandleTask(input)
}