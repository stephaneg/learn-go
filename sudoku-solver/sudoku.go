package main

import (
	"fmt"
	"strconv"
)

/*
SIZE is the SIZE of the sudoku matrix
CHECKSUM is the sum of 1 to 9 elements
*/
const SIZE int8 = 9
const CHECKSUM int8 = 45

/*
Position is the linked list that store path to fullfill sudoku grid
*/
type position struct {
	i    int8
	j    int8
	next *position
}

/*
grid is the structure to model a sudoku matrix
*/
type grid struct {
	data       [SIZE][SIZE]int8
	isOnRow    [SIZE][9]bool
	isOnCol    [SIZE][9]bool
	isOnSquare [SIZE][9]bool
}

func (g *grid) printSudoku() {
	for i, v := range g.data {
		var s [SIZE]string
		for j, v2 := range v {
			if v2 != 0 {
				s[j] = strconv.Itoa(int(v2))
			} else {
				s[j] = " "
			}
		}

		fmt.Printf("| %s %s %s | %s %s %s | %s %s %s|\n", s[0], s[1], s[2], s[3], s[4], s[5], s[6], s[7], s[8])
		if (i+1)%3 == 0 {
			fmt.Println("------------------------")
		}
	}
	//fmt.Println(g.isOnRow)
	//fmt.Println(g.isOnCol)
	//fmt.Println(g.isOnSquare)
}

func (g *grid) getRow(row int8) [SIZE]int8 {
	return g.data[row]
}

func (g *grid) getCol(col int8) [SIZE]int8 {
	var ret [SIZE]int8
	for i, v := range g.data {
		ret[i] = v[col]
	}
	return ret
}

func getSquareId(row int8, col int8) int8 {
	return (3*(row/3) + (col / 3))
}

/* somme les elements d'un vecteur */
func sum(vec [SIZE]int8) int8 {
	s := int8(0)
	for _, v := range vec {
		s = s + v
	}
	return s
}

func isValid(grid *grid, list *position) bool {
	if list == nil {
		return true
	}

	i := list.i
	j := list.j

	for k := 0; k < 9; k++ {
		if (grid.isOnRow[i][k] == false) && (grid.isOnCol[j][k] == false) && (grid.isOnSquare[getSquareId(i, j)][k] == false) {
			grid.isOnRow[i][k] = true
			grid.isOnCol[j][k] = true
			grid.isOnSquare[getSquareId(i, j)][k] = true

			if isValid(grid, list.next) {
				grid.data[i][j] = int8(k + 1)
				return true
			}
			grid.isOnRow[i][k] = false
			grid.isOnCol[j][k] = false
			grid.isOnSquare[getSquareId(i, j)][k] = false
		}
	}
	return false
}

/*
main
*/
func main() {

	/* init the sudoku grid */
	d := [SIZE][SIZE]int8{
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 7},
		{0, 0, 0, 4, 0, 5, 8, 0, 0},
		{0, 0, 2, 0, 0, 0, 0, 0, 0},
		{0, 0, 9, 0, 0, 0, 0, 1, 5},
		{0, 8, 0, 0, 6, 1, 9, 0, 0},
		{0, 4, 0, 0, 9, 0, 0, 0, 0},
		{0, 0, 7, 0, 0, 6, 1, 2, 0},
		{0, 1, 8, 0, 5, 7, 0, 3, 9},
	}

	var ri [SIZE][9]bool
	var ci [SIZE][9]bool
	var si [SIZE][9]bool
	grid := grid{d, ri, ci, si}

	// init flags vectors
	for i, vi := range grid.data {
		for j, vj := range vi {
			if vj != 0 {
				grid.isOnRow[i][vj-1] = true
				grid.isOnCol[j][vj-1] = true
				grid.isOnSquare[getSquareId(int8(i), int8(j))][vj-1] = true
			}
		}
	}

	// start sudoku resolution
	grid.printSudoku()

	// construction de la liste de cellules à évaluer et de la liste chainee position
	// on démarre avec une stratégie stupide
	var plist *position
	var current *position

	for i, vi := range grid.data {
		for j, vj := range vi {
			if vj == 0 {
				newPosition := position{int8(i), int8(j), nil}

				if plist == nil {
					current = &newPosition
					plist = current
				} else {
					current.next = &newPosition
					current = current.next
				}
			}
		}
	}

	// print the list
	/*fmt.Println("Printing the list")
	for p := plist; p != nil; p = p.next {
		fmt.Println(p)
	}*/

	fmt.Println("soliving the sudoku")
	isValid(&grid, plist)

	grid.printSudoku()
}
