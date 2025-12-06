package main

import (
	"aoc2025/lib"
	"fmt"
)

func main() {
	lines := lib.ReadInputAsLines(6)

	values := make([][]string, 0, len(lines))
	for _, line := range lines[:len(lines)-1] {
		buffer := ""
		onValue := false
		var lineValues []string
		for _, char := range line {
			if char == '\n' {
				break
			}

			if char == ' ' {
				if onValue {
					lineValues = append(lineValues, buffer)
					buffer = ""
					onValue = false
				}
				continue
			}

			onValue = true
			buffer += string(char)
		}
		if buffer != "" {
			lineValues = append(lineValues, buffer)
		}
		values = append(values, lineValues)
	}

	total := 0
	for j := 0; j < len(values[0]); j++ {
		operation := Add
		if values[len(values)-1][j] == "*" {
			operation = Multiply
		}

		result := lib.MustParseToInt(values[0][j])
		for i := 1; i < len(values)-2; i++ {
			result = operation(result, lib.MustParseToInt(values[i][j]))
		}
		total += result
	}
	fmt.Println(total)

	//~~~~~Part 2~~~~~
	type Task struct {
		operation func(int, int) int
		values    []int
	}
	var tasks []Task

	//go through all lines
	task := Task{}
	for j := 0; j < len(lines[0]); j++ {
		allSpaces := true
		buffer := ""
		for i := 0; i < len(lines)-1; i++ {
			switch lines[i][j] {
			case '*':
				task.operation = Multiply
				allSpaces = false
			case '+':
				task.operation = Add
				allSpaces = false
			case ' ':
				continue
			default:
				buffer += string(lines[i][j])
				allSpaces = false
			}
		}
		if allSpaces {
			tasks = append(tasks, task)
			task = Task{}
			continue
		}

		task.values = append(task.values, lib.MustParseToInt(buffer))
	}
	// add the final task
	tasks = append(tasks, task)

	total = 0
	for _, t := range tasks {
		taskTotal := t.values[0]
		for _, value := range t.values[1:] {
			taskTotal = t.operation(taskTotal, value)
		}
		total += taskTotal
	}
	fmt.Println(total)
}

func Add(a, b int) int {
	return a + b
}

func Multiply(a, b int) int {
	return a * b
}
