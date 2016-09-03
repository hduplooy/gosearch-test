// 8queensdepth.go
// Author: Hannes du Plooy
// Revision Date: 3 Sep 2016
// Implements DepthFirstSearch of hduplooy/gosearch on the 8 queens problem
// Place 8 queens on a normal 8X8 chess board without any one queen able to capture another
package main

import (
	"fmt"
	"strconv"
	"strings"

	src "github.com/hduplooy/gosearch"
)

// The state of the board is represented by a slice of integers with each entry indicating that rank on the board
// The file is indicated by the value in the slice entry
type Board []int

// A simple integer abs function
func iabs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// Descendants return all valid entries on the next rank based on the current position
func (brd Board) Descendants() []src.SearchF {
	// Make the slice holding all the descendants
	tmp := make([]src.SearchF, 0, 8)
	sz := len(brd)
	// Go through all the files (columns) looking for valid moves
	for i := 0; i < 8; i++ {
		fine := true
		// For all previous ranks
		for j := 0; j < sz; j++ {
			// If it is in the same file (column) or diagonally they are the same then it is not a valid move
			if i == brd[j] || iabs(brd[j]-i) == sz-j {
				fine = false
				break
			}
		}
		// If still fine going through all previous ranks then it must be a valid descendant
		if fine {
			// The slice for the descendant
			tmp2 := make([]int, sz+1, 8)
			// Copy state up to previous
			copy(tmp2, brd)
			// Set new rank+file
			tmp2[sz] = i
			// Add to slice of descendants
			tmp = append(tmp, Board(tmp2))
		}
	}
	return tmp
}

// Done check if done and this is the case if the state has already 8 ranks
func (brd Board) Done() bool {
	return len(brd) == 8
}

// Cost is not used
func (brd Board) Cost() float64 { return 0.0 }

// Away is not used
func (brd Board) Away() float64 { return 0.0 }

// Key returns a unique key describing the state
// It is easy in that it is just all the files for the ranks placed in a string
func (brd Board) Key() string {
	tmp := make([]string, len(brd))
	for i, val := range brd {
		tmp[i] = strconv.Itoa(val)
	}
	return strings.Join(tmp, "")
}

func main() {
	// Call DepthFirstSearch with an empty state and we don't want history
	cnt, ans, _ := src.DepthFirstSearch(Board(make([]int, 0, 8)), false)
	fmt.Printf("Done in %d steps\n", cnt)
	// Print the resulting board
	fmt.Println("+-+-+-+-+-+-+-+-+")
	brd := ans.(Board)
	for _, val := range brd {
		fmt.Print("|")
		for i := 0; i < 8; i++ {
			if i == val {
				fmt.Print("Q|")
			} else {
				fmt.Print(" |")
			}
		}
		fmt.Println()
		fmt.Println("+-+-+-+-+-+-+-+-+")
	}
}
