package prover

import (
	"TFL/Contest/parser"
	"TFL/Contest/util"
	"fmt"
	"os"
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
	file, _ := os.Create("result")
	defer file.Close()
	if err != nil {
		file.WriteString("Syntax error")
		//fmt.Println("Syntax error")
	} else {
		//for _, rule := range task.Rules {

		//}
		//fmt.Println(task.Rules)
		//fmt.Println(util.Unify(task.Rules[0].Left, task.Rules[0].Right))
		//fmt.Println(util.DFS(task))
		res := util.DFS(task)
		if res {
			file.WriteString("False")
		} else {
			file.WriteString("Unknown")
		}
	}
}
