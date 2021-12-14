package parser

type Trs struct {
	name string // Имя конструктора или переменной
	Args []Trs  // конструкторы и переменные в аргументах
}

type Rule struct {
	left  []Trs // левая часть правила
	right []Trs // правая часть правила
}

type Task struct {
	Input        []byte
	Rules        []Rule         //список правил
	Vars         map[string]int //список аргументов, int для альфа преобразования
	Constructors map[string]int //список конструкторов с их арностью
}
