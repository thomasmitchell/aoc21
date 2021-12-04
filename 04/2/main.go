package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("../input.txt")
	if err != nil {
		panic("Failed to read input: " + err.Error())
	}

	lines := strings.Split(strings.TrimSpace(string(input)), "\n")

	sequence := strings.Split(lines[0], ",")
	boards := []*Board{}

	for i := 1; i < len(lines[1:]); i += 6 {
		boards = append(boards, newBoard(lines[i+1:i+6]))
	}

	var drawn string
	var winningBoards int
	var board *Board
drawLoop:
	for _, drawn = range sequence {
		for _, board = range boards {
			if board.Won() {
				continue
			}
			winner := board.Mark(drawn)
			if winner {
				winningBoards++
			}
			if len(boards)-winningBoards == 0 {
				break drawLoop
			}
		}
	}

	sum := board.sumUnmarked()
	drawnParsed, err := strconv.ParseInt(drawn, 10, 64)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse string as int `%s': %s", drawn, err))
	}
	fmt.Printf("sum: %d\n", sum)
	fmt.Printf("drawn: %d\n", drawnParsed)
	fmt.Println("board")
	fmt.Println(board.String())
	fmt.Println(sum * drawnParsed)
}

type Board struct {
	//first idx is row. second is col.
	nums  [5][5]string
	rowXs [5]int8
	won   bool
}

func newBoard(in []string) *Board {
	if len(in) != 5 {
		panic(fmt.Sprintf("too many lines given to newBoard: %d", len(in)))
	}

	ret := &Board{}

	for i, row := range in {
		for j := 0; j < len(row); j += 3 {
			ret.nums[i][j/3] = strings.TrimSpace(row[j : j+2])
		}
	}

	return ret
}

//returns true if this board has won
func (b *Board) Mark(num string) bool {
	checkRow, checkCol := -1, -1
outerLoop:
	for i, row := range b.nums {
		for j, space := range row {
			if space == num {
				b.rowXs[i] |= int8(1 << (4 - j))
				checkRow, checkCol = i, j
				break outerLoop
			}
		}
	}

	if checkRow >= 0 {
		b.won = b.rowWins(checkRow) || b.colWins(checkCol)
	}

	return b.won
}

func (b *Board) rowWins(row int) bool {
	return b.rowXs[row] == 0b11111
}

func (b *Board) colWins(col int) bool {
	for i := range b.rowXs {
		if !b.isMarked(i, col) {
			return false
		}
	}

	return true
}

func (b *Board) isMarked(row, col int) bool {
	return b.rowXs[row]&(1<<(4-col)) != 0
}

func (b *Board) sumUnmarked() int64 {
	var sum int64
	for i, row := range b.nums {
		for j, space := range row {
			if !b.isMarked(i, j) {
				parsed, err := strconv.ParseInt(space, 10, 64)
				if err != nil {
					panic(fmt.Sprintf("Failed to parse string as int `%s': %s", space, err))
				}

				sum += parsed
			}
		}
	}

	return sum
}

func (b *Board) Won() bool {
	return b.won
}

func (b *Board) String() string {
	var ret []byte
	for _, row := range b.nums {
		for _, space := range row {
			if len(space) == 1 {
				ret = append(ret, ' ')
			}
			ret = append(ret, space...)
			ret = append(ret, ' ')
		}
		ret = append(ret, '\n')
	}

	return string(ret)
}
