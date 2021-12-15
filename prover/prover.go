package prover

import (
	"TFL/Contest/parser"
	"errors"
	"fmt"
)

const (
	VAR = iota
	CTR
	CST
)

func unify(trs1, trs2 parser.Trs) (parser.Trs, error) {
	fmt.Println(trs1.Name + " : " + trs2.Name)
	if trs1.Type == VAR && (trs2.Type == CTR || trs2.Type == CST) {
		return trs2, nil
	}
	if trs2.Type == VAR && (trs1.Type == CTR || trs1.Type == CST) {
		return trs1, nil
	}
	if trs1.Type == VAR && trs2.Type == VAR ||
		trs1.Type == CST && trs2.Type == CST && trs1.Name == trs2.Name {
		return trs1, nil
	}
	if trs1.Type == CTR && trs2.Type == CTR &&
		trs2.Name == trs1.Name && len(trs2.Args) == len(trs1.Args) {
		trs := parser.Trs{
			Name: trs1.Name,
			Type: CTR,
		}
		for i, val := range trs1.Args {
			res, err := unify(val, trs2.Args[i])
			if err != nil {
				return res, err
			}
			trs.Args = append(trs.Args, res)
		}
		return trs, nil
	}
	return parser.Trs{}, errors.New("unification not possible")
}

func HandleTask(input []byte) {
	task := parser.Task{}
	task.Vars = make(map[string]int)
	task.Constructors = make(map[string]int)
	task.Input = input
	err := task.ParseInput()
	if err != nil {
		fmt.Println("error while parsing: " + err.Error())
	}
	fmt.Println(task.Rules)
	fmt.Println(unify(task.Rules[0].Left, task.Rules[0].Right))
}
