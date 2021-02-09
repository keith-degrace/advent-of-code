package p18_2

import (
	"au"
	"container/list"
	"fmt"
	"strconv"
	"strings"
)

type Value struct {
	isConstant bool
	constant   int
	register   string
}

func NewValue(value string) *Value {
	v := new(Value)

	numberValue, err := strconv.Atoi(value)
	if err != nil {
		v.register = value
		v.isConstant = false
	} else {
		v.constant = numberValue
		v.isConstant = true
	}

	return v
}

func (v *Value) get(registers map[string]int) int {
	if v.isConstant {
		return v.constant
	}

	registerValue, ok := registers[v.register]
	if ok {
		return registerValue
	} else {
		return 0
	}
}

const snd = 1
const set = 2
const add = 3
const mul = 4
const mod = 5
const rcv = 6
const jgz = 7

type Instruction struct {
	opcode int
	param1 *Value
	param2 *Value
}

func loadInstructions(input []string) []Instruction {
	instructions := []Instruction{}

	for _, line := range input {
		tokens := strings.Split(line, " ")

		instruction := Instruction{}

		if tokens[0] == "snd" {
			instruction.opcode = snd
		} else if tokens[0] == "set" {
			instruction.opcode = set
		} else if tokens[0] == "add" {
			instruction.opcode = add
		} else if tokens[0] == "mul" {
			instruction.opcode = mul
		} else if tokens[0] == "mod" {
			instruction.opcode = mod
		} else if tokens[0] == "rcv" {
			instruction.opcode = rcv
		} else if tokens[0] == "jgz" {
			instruction.opcode = jgz
		}

		instruction.param1 = NewValue(tokens[1])
		if len(tokens) == 3 {
			instruction.param2 = NewValue(tokens[2])
		}

		instructions = append(instructions, instruction)
	}

	return instructions
}

type Queue struct {
	values *list.List
}

func NewQueue() *Queue {
	queue := new(Queue)
	queue.values = list.New()
	return queue
}

func (q *Queue) push(value int) {
	q.values.PushBack(value)
}

func (q *Queue) pop() (int, bool) {
	if q.values.Len() == 0 {
		return 0, false
	}

	front := q.values.Front()
	value := front.Value.(int)
	q.values.Remove(front)

	return value, true
}

type Program struct {
	id           int
	instructions []Instruction
	current      int
	registers    map[string]int
	queue        *Queue
	sendCount    int
}

func NewProgram(id int, instructions []Instruction) *Program {
	program := new(Program)
	program.id = id
	program.instructions = instructions
	program.current = 0
	program.registers = make(map[string]int)
	program.registers["p"] = id
	program.queue = NewQueue()
	program.sendCount = 0
	return program
}

func (p *Program) isTerminated() bool {
	return p.current >= len(p.instructions)
}

func (p *Program) step(sendQueue *Queue) bool {
	if p.isTerminated() {
		return false
	}

	instruction := p.instructions[p.current]

	if instruction.opcode == snd {
		x := instruction.param1.get(p.registers)
		sendQueue.push(x)
		p.sendCount++
	} else if instruction.opcode == set {
		x := instruction.param1.register
		y := instruction.param2.get(p.registers)

		p.registers[x] = y
	} else if instruction.opcode == add {
		x := instruction.param1.register
		y := instruction.param2.get(p.registers)

		p.registers[x] += y
	} else if instruction.opcode == mul {
		x := instruction.param1.register
		y := instruction.param2.get(p.registers)

		p.registers[x] *= y
	} else if instruction.opcode == mod {
		x := instruction.param1.register
		y := instruction.param2.get(p.registers)

		p.registers[x] %= y
	} else if instruction.opcode == rcv {
		value, ok := p.queue.pop()
		if !ok {
			// Nothing yet, try again later
			return false
		}

		x := instruction.param1.register
		p.registers[x] = value
	} else if instruction.opcode == jgz {
		x := instruction.param1.get(p.registers)
		y := instruction.param2.get(p.registers)

		if x > 0 {
			// Minus one because the current will increment below.
			p.current += (y - 1)
		}
	}

	p.current++

	return true
}

func Solve() {
	input := au.ReadInputAsStringArray("18")

	instructions := loadInstructions(input)

	program0 := NewProgram(0, instructions)
	program1 := NewProgram(1, instructions)

	for {
		stepped0 := program0.step(program1.queue)
		stepped1 := program1.step(program0.queue)

		if !stepped0 && !stepped1 {
			break
		}

		//	time.Sleep(100)
	}

	fmt.Println(program1.sendCount)
}
