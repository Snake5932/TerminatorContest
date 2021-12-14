package parser

import (
	"errors"
	"fmt"
)

type Trs struct {
	name string // Имя конструктора или переменной
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

func (task *Task)ParseInput() error {
	err := task.parseRules()
	if err != nil {
		return err
	}
	for len(task.Input) != 0 {
		if len(task.Input) >= 1 {
			if (task.Input[0] != '\r' || task.Input[1] != '\n') && task.Input[0] != '\n' {
				return errors.New("unexpected character at the end of the line")
			}
			if task.Input[0] == '\n' {
				task.Input = task.Input[1:]
			} else {
				task.Input = task.Input[2:]
			}
			err = task.parseRules()
			if err != nil {
				return err
			}
		} else {
			return errors.New("unexpected character at the end of the line")
		}
	}
	return nil
}

func (task *Task)parseRules() error {
	task.SkipSpaces()
	if len(task.Input) == 0 || (len(task.Input) > 0 && (task.Input[0] == '\r' && task.Input[1] == '\n' || task.Input[0] == '\n')) {
		return nil
	}
	err := task.parseRule()
	if err != nil {
		return err
	}

}

func (task *Task)parseRule() error {
	task.SkipSpaces()
	nSymb := task.Input[0]
	if nSymb < 'A' || (nSymb > 'Z' && nSymb < 'a') || nSymb > 'z'{
		return errors.New("wrong declaration symbol")
	}
	task.Input = task.Input[1:]
	var buffer []byte
	buffer = append(buffer, nSymb)
	nSymb = task.Input[0]
	for len(task.Input) != 0 && (nSymb >= '0' && nSymb <= '9' ||
		nSymb >= 'A' && nSymb <= 'Z' ||
		nSymb >= 'a' && nSymb <= 'z') {
		buffer = append(buffer, nSymb)
		task.Input = task.Input[1:]
		nSymb = task.Input[0]
	}
	nTermS := string(buffer)

	task.SkipSpaces()
	return nil
}
