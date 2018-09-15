package main
import (
	"fmt"
	"bufio"
	"os"
	"regexp"
	"math"
	"math/rand"
	"strconv"
)

type Square struct {
	IsMine bool
	AdjacentCount int
	Hidden bool
	Id string
	X,Y int
	Flagged bool
	Display string
}

func readInput() string {
	r := bufio.NewReader(os.Stdin)
	fmt.Print("Cmd: ");
	text,_ := r.ReadString('\n');
	return text
}
func displayHelp() {
	lines := []string{
		"start n n, starts a game of N x N squares",
		"start n n [1-3], starts a game of N x N squares,\n\t\t setting the difficulty.",
		"n n, pick a square",
	};
	for i := range lines {
		fmt.Printf("%s\n",lines[i]);
	}
}
func createGrid(x,y int) [][]Square {
	var rows [][]Square;
	for i := 0; i < y; i++ {
		cols := make([]Square,x)
		rows = append(rows,cols)
		for j := range cols {
			cols[j].Hidden = true
			cols[j].Id = fmt.Sprintf("%d_%d",i,j)
			cols[j].IsMine = false
			cols[j].X = j
			cols[j].Y = i
			cols[j].Display = " *|"
		}
	}
	return rows
}
func assignMines(grid [][]Square,d int) [][]Square {
	nd := float64(d) / 10
	xy := float64(len(grid) * len(grid[0]))
	l := int(math.Floor( float64(xy * nd )))
	var assigned [][]int
	in_array := func (a [][]int,v[]int) bool {
		for i := range a {
			if a[i][0] == v[0] && a[i][1] == v[1] {
				return true
			}
		}
		return false
	}
	var x,y int
	for len(assigned) < l {
		x = rand.Intn(len(grid[0]))
		y = rand.Intn(len(grid))
		point := []int{x,y}
		if ! in_array(assigned,point) {
			assigned = append(assigned,point)
			grid[y][x].IsMine = true
		}
	}
	return grid
}
func getMines(grid [][]Square) []Square {
	var mines []Square
	for y := range grid {
		for x := range grid[y] {
			if grid[y][x].IsMine {
				mines = append(mines,grid[y][x])
			}
		}
	}
	return mines
}
func renderGrid(grid [][]Square) {
	for y := range grid {
		for x := range grid[y] {
			s := grid[y][x]
			fmt.Print(s.Display);
		}
		fmt.Print("\n");
	}
}
func getMinesAround(grid [][]Square,y,x int) []Square {
	var mines []Square
	yn := y - 1
	if yn >= 0 {
		for i := -1; i < 2; i++ {
			if (x+i) >= 0 && (x+i) < len(grid[yn]) && grid[yn][x+i].IsMine {
				mines = append(mines,grid[yn][x+i])
			}
		}
	}
	yn = y + 1
	if yn < len(grid) {
		for i := -1; i < 2; i++ {
			if (x+i) >= 0 && (x+i) < len(grid[yn]) && grid[yn][x+i].IsMine {
				mines = append(mines,grid[yn][x+i])
			}
		}
	}
	if (x-1) >= 0 && grid[y][x-1].IsMine {
		mines = append(mines,grid[y][x-1])
	}
	if (x+1) < len(grid[y]) && grid[y][x+1].IsMine {
		mines = append(mines,grid[y][x+1])
	}
	return mines
}
func getSafeAround(grid [][]Square,y,x int) []Square {
	is_safe := func(sq Square) bool {
		if sq.Hidden == true && sq.IsMine == false {
			return true
		}
		return false
	}
	var safes []Square
	yn := y - 1
	if yn >= 0 {
		for i := -1; i < 2; i++ {
			if (x+i) >= 0 && (x+i) < len(grid[yn]) && is_safe(grid[yn][x+i]) {
				safes = append(safes,grid[yn][x+i])
			}
		}
	}
	yn = y + 1
	if yn < len(grid) {
		for i := -1; i < 2; i++ {
			if (x+i) >= 0 && (x+i) < len(grid[yn]) && is_safe(grid[yn][x+i]) {
				safes = append(safes,grid[yn][x+i])
			}
		}
	}
	if (x-1) >= 0 && is_safe(grid[y][x-1]) {
		safes = append(safes,grid[y][x-1])
	}
	if (x+1) < len(grid[y]) && is_safe(grid[y][x+1]) {
		safes = append(safes,grid[y][x+1])
	}
	return safes
}
func main() {
	var text string
	fmt.Print("Use ? for help\n")
	startMatcher := regexp.MustCompile(`start (\d+) (\d+)\s+$`)
	startMatcherAlt := regexp.MustCompile(`start (\d+) (\d+) ([1-3])\s+$`)
	moveMatcher := regexp.MustCompile(`^(\d+) (\d+)\s+$`)
	gameOn := false
	minesSet := false
	difficulty := 1
	var grid [][]Square
	for {
		if gameOn {
			renderGrid(grid)
		}
		text = readInput()
		if text == "q\n" {
			return;
		}
		if text == "?\n" {
			displayHelp();
		}
		s1 := startMatcher.FindAllStringSubmatch(text,-1)
		s2 := startMatcherAlt.FindAllStringSubmatch(text,-1)
		m := moveMatcher.FindAllStringSubmatch(text,-1)
		if ( len(s1) > 0 ) {
			n1,_ := strconv.Atoi(s1[0][1])
			n2,_ := strconv.Atoi(s1[0][2])
			fmt.Printf("n1=%d,n2=%d\n",n1,n2)
			grid = createGrid(n1,n2)
			gameOn = true
			minesSet = false
		} 
		if ( len(s2) > 0 ) {
			n1,_ := strconv.Atoi(s2[0][1])
			n2,_ := strconv.Atoi(s2[0][2])
			difficulty,_ = strconv.Atoi(s2[0][3])
			grid = createGrid(n1,n2)
			gameOn = true
			minesSet = false
		}
		if len(m) > 0 && gameOn {
			y,_ := strconv.Atoi(m[0][1])
			x,_ := strconv.Atoi(m[0][2])
			if ! minesSet {
				grid = assignMines(grid,difficulty)
				if grid[y][x].IsMine {
					grid[y][x].IsMine = false
				}
				minesSet = true
			}
			if grid[y][x].IsMine {
				fmt.Printf("Game Over!\n")
				gameOn = false
			}
			mines := getMinesAround(grid,y,x)
			grid[y][x].Display = fmt.Sprintf("%2d|",len(mines))
			safes := getSafeAround(grid,y,x)
			for i := range safes {
				s := safes[i]
				mines = getMinesAround(grid,s.Y,s.X)
				grid[s.Y][s.X].Display = fmt.Sprintf("%2d|",len(mines))
				grid[s.Y][s.X].Hidden = false
			}
		}
	}
}
