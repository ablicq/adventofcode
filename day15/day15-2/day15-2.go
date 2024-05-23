package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ablicq/adventofcode/utils"
)

const (
	REMOVE string = "-"
	ADD    string = "="
)

type Operation struct {
	label          string
	operation_type string
	focal_length   int
}

type Lens struct {
	label        string
	focal_length int
}

func hash(s string) int {
	h := 0
	for _, c := range s {
		h = (h + int(c)) * 17 % 256
	}
	return h
}

func parseOperation(s string) Operation {
	if strings.Contains(s, REMOVE) {
		return Operation{s[:len(s)-1], REMOVE, -1}
	}
	focal_length, _ := strconv.ParseInt(string(s[len(s)-1]), 10, 10)
	return Operation{s[:len(s)-2], ADD, int(focal_length)}
}

func applyOperation(op Operation, boxes [][]Lens) [][]Lens {
	box_id := hash(op.label)
	box := boxes[box_id]
	if op.operation_type != REMOVE {
		for i, l := range box {
			if l.label == op.label {
				box[i] = Lens{op.label, op.focal_length}
				return boxes
			}
		}
		box = append(box, Lens{op.label, op.focal_length})
		boxes[box_id] = box
		return boxes
	}

	for i, l := range box {
		if l.label == op.label {
			boxes[box_id] = append(box[:i], box[i+1:]...)
			break
		}
	}
	return boxes
}

func eval_boxes(boxes [][]Lens) int {
	s := 0
	for i, box := range boxes {
		for j, lens := range box {
			s += (i + 1) * (j + 1) * lens.focal_length
		}
	}
	return s
}

func main() {
	input := strings.Split(utils.ParseInput("../input.txt")[0], ",")

	boxes := make([][]Lens, 256)
	for i := range boxes {
		boxes[i] = make([]Lens, 0)
	}

	for _, op_s := range input {
		op := parseOperation(op_s)
		applyOperation(op, boxes)
	}

	fmt.Println(eval_boxes(boxes))
}
