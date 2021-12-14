package prover

import (
	"TFL/Contest/parser"
	"fmt"
)

func HandleTask(input []byte) {
	task := parser.Task{}
	task.Vars = make(map[string]int)
	task.Constructors = make(map[string]int)
	task.Input = input
	err := task.ParseInput()
	if err != nil {
		fmt.Println("error while parsing: " + err.Error())
	}
	fmt.Println(task)
}
