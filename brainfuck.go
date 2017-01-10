package brainfuck

import (
	"fmt"
)

const (
	left      = '<'
	right     = '>'
	increase  = '+'
	decrease  = '-'
	input     = ','
	output    = '.'
	loopBegin = '['
	loopEnd   = ']'
)

const MaxSize = 3000

func Exec(program string) (result []uint8, err error) {
	vm := make([]uint8, MaxSize)
	ptr := 0

	for i := 0; i < len(program); i++ {
		cmd := program[i]
		switch cmd {
		case left:
			ptr--
		case right:
			ptr++
		case increase:
			vm[ptr]++
		case decrease:
			vm[ptr]--
		case input:
			var in byte
			_, err := fmt.Scanf("%c", &in)
			if err != nil {
				return nil, fmt.Errorf("Unable to parse input")
			}
			vm[ptr] = in
		case output:
			result = append(result, vm[ptr])
		case loopBegin:
			if vm[ptr] == 0 {
				for j := i; j < len(program); j++ {
					if program[j] == loopEnd {
						i = j
						break
					}
					if j == len(program)-1 {
						return nil, fmt.Errorf("Couldn't find loop's end")
					}
				}
			}
		case loopEnd:
			if vm[ptr] != 0 {
				for j := i; j > 0; j-- {
					if program[j] == loopBegin {
						i = j
						break
					}
					if j == 0 {
						return nil, fmt.Errorf("Couldn't find loop's beginning")
					}
				}
			}
		}
	}
	return
}
