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
		var alpha1 bool
		var alpha2 bool
		for _, rule := range task.Rules {
			alpha1 = util.CheckAlpha(rule.Left)
			alpha2 = util.CheckAlpha(rule.Right)
			if !alpha1 || !alpha2 {
				break
			}
		}
		if !alpha1 || !alpha2 {
			file.WriteString("Unknown")
			//fmt.Println("Unknown alpha")
		} else if util.DFS(task) {
			file.WriteString("False")
			//fmt.Println("False")
		} else {
			file.WriteString("Unknown")
			//fmt.Println("Unknown")
		}
	}
}
