package main
import "testing"
import "math"
func TestCreateSquares( t *testing.T) {
	n1 := 5
	n2 := 10
	d := 1
	grid := createGrid(n1,n2)
	if len(grid) != n2 {
		t.Errorf("expected height to be %d, got %d",n2,len(grid));
	}
	if ( len(grid[0]) != n1 ) {
		t.Errorf("expected width to be %d, got %d",n1,len(grid[0]));
	}
	grid = assignMines(grid,d)
	mines := getMines(grid)
	nd := float64(d) / 10
	xy := float64(len(grid) * len(grid[0]))
	l := int(math.Floor( float64(xy * nd )))

	if  len(mines) != l  {
		t.Errorf("expected 10 percent of mines to be %d got %d",l,len(mines))
	}
}
func TestThirtyPercent(t *testing.T) {
	n1 := 6
	n2 := 6
	d := 3
	grid := createGrid(n1,n2)
	grid = assignMines(grid,d)
	mines := getMines(grid)
	nd := float64(d) / 10
	xy := float64(len(grid) * len(grid[0]))
	l := int(math.Floor( float64(xy * nd )))

	if  len(mines) != l  {
		t.Errorf("expected 30 percent of mines to be %d got %d",l,len(mines))
	}
	//renderGrid(grid)
}
func TestGetConnectedMines(t *testing.T) {
	n1 := 6
	n2 := 6
	grid := createGrid(n1,n2)
	grid[0][0].IsMine = true
	grid[0][2].IsMine = true
	mines := getMinesAround(grid,1,1)
	if len(mines) != 2 {
		t.Errorf("Expected 2 mines, got %d\n", len(mines))
	}
	grid[2][0].IsMine = true
	grid[2][2].IsMine = true
	mines = getMinesAround(grid,1,1)
	if len(mines) != 4 {
		t.Errorf("Expected 4 mines, got %d\n", len(mines))
	}
	grid[1][0].IsMine = true
	grid[1][2].IsMine = true
	mines = getMinesAround(grid,1,1)
	if len(mines) != 6 {
		t.Errorf("Expected 6 mines, got %d\n", len(mines))
	}
	safes := getSafeAround(grid,1,1)
	if len(safes) > 2 {
		t.Errorf("there are only 2 safe squares! got %d", len(safes))
	}
}
