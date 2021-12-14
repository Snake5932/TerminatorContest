package parser

import (
	"errors"
	"fmt"
	"strconv"
)

type Trs struct {
	Name string // Имя конструктора или переменной
	Args []Trs // конструкторы и переменные в аргументах
	Type int //тип
}

type Rule struct {
	Left Trs  // левая и правая части
	Right Trs
}

type Task struct {
	Input []byte
	Rules []Rule //список правил
	Vars map[string]int //список аргументов, int для альфа преобразования
	Constructors map[string]int //список конструкторов с их арностью
}

func (task *Task)SkipSpaces() {
	for len(task.Input) != 0 && (task.Input[0] == ' ' || task.Input[0] == 9) {
		task.Input = task.Input[1:]
	}
}

func (task *Task)getIdent() (string, error) {
	if len(task.Input) == 0 {
		return "", errors.New("ident missing")
	}
	nSymb := task.Input[0]
	if nSymb < 'A' || (nSymb > 'Z' && nSymb < 'a') || nSymb > 'z'{
		fmt.Println(nSymb)
		return "", errors.New("wrong declaration symbol")
	}
	task.Input = task.Input[1:]
	var buffer []byte
	buffer = append(buffer, nSymb)
	for len(task.Input) != 0 && (task.Input[0] >= '0' && task.Input[0] <= '9' ||
								 task.Input[0] >= 'A' && task.Input[0] <= 'Z' ||
								 task.Input[0] >= 'a' && task.Input[0] <= 'z') {
		buffer = append(buffer, task.Input[0])
		task.Input = task.Input[1:]
		nSymb = task.Input[0]
	}
	return string(buffer), nil
}

func (task *Task)ParseInput() error {
	err := task.parseVars()
	if err != nil {
		return err
	}
	for len(task.Input) != 0 {
		if len(task.Input) >= 1 {
			if (task.Input[0] != '\r' || task.Input[1] != '\n') && task.Input[0] != '\n' {
				err := task.parseRules()
				if err != nil {
					return err
				}
			} else {
				task.SkipSpaces()
				if task.Input[0] == '\n' {
					task.Input = task.Input[1:]
				} else {
					task.Input = task.Input[2:]
				}
				task.SkipSpaces()
			}
		} else {
			return errors.New("unexpected character at the end of the line")
		}
	}
	return nil
}

func (task *Task)parseVars() error {
	task.SkipSpaces()
	if len(task.Input) <= 1 {
		return errors.New("vars list is absent")
	}
	for (task.Input[0] == '\r' && task.Input[1] == '\n') || task.Input[0] == '\n' {
		if task.Input[0] == '\n' {
			task.Input = task.Input[1:]
		} else {
			task.Input = task.Input[2:]
		}
		task.SkipSpaces()
	}
	if task.Input[0] != '[' {
		return errors.New("wrong vars list opening")
	}
	task.Input = task.Input[1:]
	if task.Input[0] != ']' {
		err := task.parseVarsList()
		if err != nil {
			return err
		}
	}
	if task.Input[0] != ']' {
		return errors.New("wrong vars list closing")
	}
	task.Input = task.Input[1:]
	task.SkipSpaces()
	return nil
}

func (task *Task)parseVarsList() error {
	task.SkipSpaces()
	ident, err := task.getIdent()
	if err != nil {
		return err
	}
	task.Vars[ident] = 0
	task.SkipSpaces()
	for len(task.Input) != 0 && task.Input[0] == ',' {
		task.Input = task.Input[1:]
		ident, err = task.getIdent()
		if err != nil {
			return err
		}
		if _, ok := task.Vars[ident]; ok {
			return errors.New("var redeclaration")
		}
		task.Vars[ident] = 0
		task.SkipSpaces()
	}
	return nil
}

func (task *Task)parseRules() error {
	if len(task.Input) == 0 {
		return nil
	}
	trs1, err := task.parseRule()
	if err != nil {
		return err
	}
	if len(task.Input) <= 1 || task.Input[0] != '-' || task.Input[1] != '>' {
		return errors.New("-> is absent")
	}
	task.Input = task.Input[2:]
	trs2, err := task.parseRule()
	if err != nil {
		return err
	}
	task.Rules = append(task.Rules, Rule{Left: trs1,
										Right: trs2})
	return nil
}

func (task *Task)parseRule() (Trs, error) {
	task.SkipSpaces()
	nTermS, err := task.getIdent()
	if err != nil {
		return Trs{}, err
	}
	task.SkipSpaces()
	if len(task.Input) == 0 || task.Input[0] != '(' {
		if _, ok := task.Vars[nTermS]; ok {
			task.Vars[nTermS] += 1
			alpha := strconv.Itoa(task.Vars[nTermS])
			return Trs{Name: nTermS + alpha,
					   Type: 0}, nil
		}
		if _, ok := task.Constructors[nTermS]; ok {
			if task.Constructors[nTermS] != 0 {
				return Trs{}, errors.New("polymorphism: " + nTermS)
			}
		} else {
			task.Constructors[nTermS] = 0
		}
		return Trs{Name: nTermS,
				   Type: 1}, nil
	}
	task.Input = task.Input[1:]
	task.SkipSpaces()
	rule := Trs{Name: nTermS,
				Type: 1}
	arity := 0
	trs, err := task.parseRule()
	if err != nil {
		return Trs{}, err
	}
	rule.Args = append(rule.Args, trs)
	arity += 1
	task.SkipSpaces()
	for len(task.Input) != 0 && task.Input[0] == ',' {
		task.Input = task.Input[1:]
		trs, err = task.parseRule()
		if err != nil {
			return Trs{}, err
		}
		rule.Args = append(rule.Args, trs)
		arity += 1
		task.SkipSpaces()
	}
	if task.Input[0] != ')' {
		return Trs{}, errors.New("wrong rules args closing")
	}
	task.Input = task.Input[1:]
	if _, ok := task.Constructors[nTermS]; ok {
		if task.Constructors[nTermS] != arity {
			return Trs{}, errors.New("polymorphism: " + nTermS)
		}
	} else {
		task.Constructors[nTermS] = arity
	}
	task.SkipSpaces()
	return rule, nil
}
