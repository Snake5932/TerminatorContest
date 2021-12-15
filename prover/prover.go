package prover

import (
	"TFL/Contest/parser"
	"TFL/Contest/util"
	"os"
	"reflect"
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

		if util.DFS(task) || util.DFS2(task) {
			file.WriteString("False")
		} else if !alpha1 || !alpha2 {
			file.WriteString("Unknown")
		} else if lexicographic(task) {
			file.WriteString("True")
		} else {
			file.WriteString("Unknown")
		}
	}
}

func lexicographic(task parser.Task) bool {
	var arr []string
	for name := range task.Constructors {
		arr = append(arr, name)
	}

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

			succ := true
			for _, rule := range task.Rules {
				succ = lexicRules(rule.Left, rule.Right, arr) && succ
			}
			if succ {
				return true
			}

			indexes[i]++
			i = 0
		} else {
			indexes[i] = 0
			i++
		}
	}
	return false
}

func lexicRules(trs1, trs2 parser.Trs, permTable []string) bool {
	return lexicRule1(trs1, trs2) ||
		lexicRule2(trs1, trs2, permTable) ||
		lexicRule3(trs1, trs2, permTable) ||
		lexicRule4(trs1, trs2, permTable)
}

func lexicRule1(trs1, trs2 parser.Trs) bool {
	succ := false
	for _, arg := range trs1.Args {
		succ = succ || reflect.DeepEqual(arg, trs2)
	}
	return succ
}

func lexicRule2(trs1, trs2 parser.Trs, permTable []string) bool {
	succ := false
	for _, arg := range trs1.Args {
		succ = succ || lexicRules(arg, trs2, permTable)
	}
	return succ
}

func lexicRule3(trs1, trs2 parser.Trs, permTable []string) bool {
	if !more(trs1.Name, trs2.Name, permTable) {
		return false
	}
	succ := true
	for _, arg := range trs2.Args {
		succ = lexicRules(trs1, arg, permTable) && succ
	}
	return succ
}

func lexicRule4(trs1, trs2 parser.Trs, permTable []string) bool {
	if trs1.Name == trs2.Name {
		succ := true
		for _, arg := range trs2.Args {
			succ = lexicRules(trs1, arg, permTable) && succ
		}
		j := 0
		n := len(trs1.Args)
		for j < n && reflect.DeepEqual(trs1.Args[j], trs2.Args[j]) {
			j++
		}
		if j < n {
			return succ && lexicRules(trs1.Args[j], trs2.Args[j], permTable)
		}
		return false
	} else {
		return false
	}
}

func more(name1, name2 string, permTable []string) bool {
	for _, name := range permTable {
		if name == name1 {
			return true
		} else if name == name2 {
			return false
		}
	}
	return true
}
