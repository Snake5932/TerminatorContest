package prover

import (
	"TFL/Contest/parser"
	"TFL/Contest/util"
	"fmt"
)

const (
	VAR = iota
	CTR
	CST
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
	//fmt.Println(task.Rules)
	//fmt.Println(util.Unify(task.Rules[0].Left, task.Rules[0].Right))
	fmt.Println(util.DFS(task))
}
