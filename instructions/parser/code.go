package parser

import (
	"bufio"
	. "esorun/instructions"
	"io"
	"os"
	"strings"
)

const CommentChar = ';'

func ParseCodeFromStream(stream io.Reader) (instructions []*Instruction, err error) {
	scanner := bufio.NewScanner(stream)
	instructions = make([]*Instruction, 0)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimLeft(line, " \t\r")

		if len(strings.TrimRight(line, " \t\r")) == 0 || line[0] == CommentChar {
			continue
		}

		instruction, iErr := ParseInstruction(line)
		if iErr != nil {
			return nil, iErr
		}
		instructions = append(instructions, instruction)
	}
	return instructions, nil
}

func ParseCodeFromFile(filename string) (instructions []*Instruction, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) { _ = file.Close() }(file)
	return ParseCodeFromStream(file)
}

func ParseCodeFromString(s string) (instructions []*Instruction, err error) {
	return ParseCodeFromStream(strings.NewReader(s))
}
