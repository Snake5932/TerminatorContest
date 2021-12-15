package util

import (
	"TFL/Contest/parser"
	"errors"
	"reflect"
)

type Void struct{}

var Member Void

const (
	VAR = iota
	CTR
	CST
)

func Unify(trs1, trs2 parser.Trs) (parser.Trs, error) {
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
			res, err := Unify(val, trs2.Args[i])
			if err != nil {
				return res, err
			}
			trs.Args = append(trs.Args, res)
		}
		return trs, nil
	}
	return parser.Trs{}, errors.New("unification not possible")
}

func DFS(task parser.Task) bool {
	for _, rule := range task.Rules {
		used := make(map[string]Void)
		if visitVertex(rule, &used, task) {
			return true
		}
	}
	return false
}

func visitVertex(rule parser.Rule, used *map[string]Void, task parser.Task) bool {
	(*used)[makeHash(rule.Left)+makeHash(rule.Right)] = Member
	ts := false
	for _, ruleN := range task.Rules {
		_, err := Unify(rule.Right, ruleN.Left)
		if err == nil {
			ts = true
			if _, ok := (*used)[makeHash(ruleN.Left)+makeHash(ruleN.Right)]; !ok {
				if !visitVertex(ruleN, used, task) {
					return false
				}
			}
		}
	}
	return ts
}

func DFS2(task parser.Task) bool {
	for _, rule := range task.Rules {
		used := make(map[string]Void)
		if visitVertex2(rule, &used, task) {
			return true
		}
	}
	return false
}

func visitVertex2(rule parser.Rule, used *map[string]Void, task parser.Task) bool {
	(*used)[makeHash(rule.Left) + makeHash(rule.Right)] = Member
	var ts bool
	for _, arg := range rule.Right.Args {
		ts = false
		for _, ruleN := range task.Rules {
			_, err := Unify(arg, ruleN.Left)
			if err == nil {
				ts = true
				if _, ok := (*used)[makeHash(ruleN.Left) + makeHash(ruleN.Right)]; !ok {
					if !visitVertex(ruleN, used, task) {
						return false
					}
				}
			}
		}
	}
	return ts
}

func CheckAlpha(rule parser.Trs) bool {
	for i, arg := range rule.Args {
		for j := i + 1; j < len(rule.Args); j++ {
			if reflect.DeepEqual(arg, rule.Args[j]) {
				//fmt.Println(arg.Name + " : " + rule.Args[j].Name)
				//if arg.Name == rule.Args[j].Name {
				return false
			}
		}
	}
	return true
}

func makeHash(trs parser.Trs) string {
	res := trs.Name + "("
	for _, arg := range trs.Args {
		res = res + makeHash(arg)
	}
	return res + ")"
}
