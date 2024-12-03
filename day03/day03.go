package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

type Parser struct {
	buf string
	loc int
	len int
}

func parser_next(parser *Parser) byte {
	if parser.loc < parser.len {
		character := parser.buf[parser.loc]
		parser.loc += 1
		return character
	}

	return 0
}

func parser_rest(parser *Parser) string {
	return parser.buf[parser.loc:]
}

func parser_remaining(parser *Parser) int {
	return parser.len - parser.loc
}

func parser_revert(parser *Parser, loc int) {
	parser.loc = loc
}

func parse_mul(parser *Parser) bool {
	mark := parser.loc
	match := (parser_next(parser) == 'm' &&
		parser_next(parser) == 'u' &&
		parser_next(parser) == 'l')

	if !match {
		parser.loc = mark
	}

	return match
}

func parse_do(parser *Parser) bool {
	mark := parser.loc
	match := (parser_next(parser) == 'd' &&
		parser_next(parser) == 'o' &&
		parser_next(parser) == '(' &&
		parser_next(parser) == ')')

	if !match {
		parser.loc = mark
	}

	return match
}

func parse_dont(parser *Parser) bool {
	mark := parser.loc
	match := (parser_next(parser) == 'd' &&
		parser_next(parser) == 'o' &&
		parser_next(parser) == 'n' &&
		parser_next(parser) == '\'' &&
		parser_next(parser) == 't' &&
		parser_next(parser) == '(' &&
		parser_next(parser) == ')')

	if !match {
		parser.loc = mark
	}

	return match
}

func parse_num(parser *Parser, max_len int) (int, error) {
	mark := parser.loc

	for parser.loc-mark < max_len {
		c := parser_next(parser)
		if !unicode.IsDigit(rune(c)) {
			parser.loc = parser.loc - 1
			break
		}
	}

	token := parser.buf[mark:parser.loc]

	return strconv.Atoi(token)
}

type MulArgs struct {
	left  int
	right int
}

func parse_arguments(parser *Parser) (MulArgs, error) {
	empty := MulArgs{}

	if parser_next(parser) != '(' {
		return empty, errors.New("MissingLeftParens")
	}

	leftNum, err := parse_num(parser, 3)

	if err != nil {
		return empty, errors.New("LeftArgNotNumber")
	}

	if parser_next(parser) != ',' {
		return empty, errors.New("MissingComma")
	}

	rightNum, err := parse_num(parser, 3)

	if err != nil {
		return empty, errors.New("RightArgNotNumber")
	}

	if parser_next(parser) != ')' {
		return empty, errors.New("MissingRightParens")
	}

	return MulArgs{leftNum, rightNum}, nil
}

func part_one(parser Parser) int {
	result := 0

	for parser_remaining(&parser) > 0 {
		mark := parser.loc

		if parse_mul(&parser) {
			args, err := parse_arguments(&parser)

			if err != nil {
				parser_revert(&parser, mark+1)
			}

			result += args.left * args.right
		} else {
			parser_revert(&parser, mark+1)
		}
	}

	return result
}

type State struct {
	mulEnabled bool
}

func part_two(parser Parser, state *State) int {
	result := 0

	for parser_remaining(&parser) > 0 {
		mark := parser.loc

		if parse_mul(&parser) {
			mark = parser.loc

			args, err := parse_arguments(&parser)

			if err != nil {
				parser_revert(&parser, mark)
			}

			if state.mulEnabled {
				result += args.left * args.right
			}
		} else if parse_do(&parser) {
			state.mulEnabled = true
		} else if parse_dont(&parser) {
			state.mulEnabled = false
		} else {
			parser_revert(&parser, mark+1)
		}
	}

	return result
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	result1 := 0
	result2 := 0

	state := State{true}

	for scanner.Scan() {
		input := scanner.Text()

		parser1 := Parser{
			buf: input,
			loc: 0,
			len: len(input),
		}

		parser2 := Parser{
			buf: input,
			loc: 0,
			len: len(input),
		}

		result1 += part_one(parser1)
		result2 += part_two(parser2, &state)
	}

	fmt.Println("Part One:", result1)
	fmt.Println("Part Two:", result2)
}
