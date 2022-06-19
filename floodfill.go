package main

import "fmt"

type Pair struct {
	x int
	y int
}

func validCoord(x, y, n, m int) bool {
	if x < 0 || y < 0 {
		return false
	}
	if x >= n || y >= m {
		return false
	}
	return true
}
func floodFillRecursion(arr [][]uint8, x, y int) {
	cntRow := len(arr)
	cntColumn := len(arr[0])

	if !validCoord(x, y, cntColumn, cntRow) {
		return
	}
	if arr[y][x] == 1 {
		return
	}
	arr[y][x] = 1
	floodFillRecursion(arr, x+1, y) // right
	floodFillRecursion(arr, x, y+1) // down
	floodFillRecursion(arr, x-1, y) // left
	floodFillRecursion(arr, x, y-1) // up
}
func floodFillStack(arr [][]uint8, x, y int) {
	var stack []Pair
	cntRow := len(arr)
	cntColumn := len(arr[0])

	stack = append(stack, Pair{x, y})
	lenStack := len(stack)
	for lenStack > 0 {
		// Dequeue the front node
		currPixel := stack[lenStack-1]
		stack = stack[:lenStack-1]

		// right
		if validCoord(currPixel.x+1, currPixel.y, cntColumn, cntRow) && arr[currPixel.y][currPixel.x+1] != 1 {
			arr[currPixel.y][currPixel.x] = 1
			stack = append(stack, Pair{currPixel.x + 1, currPixel.y})
		}
		// down
		if validCoord(currPixel.x, currPixel.y+1, cntColumn, cntRow) && arr[currPixel.y+1][currPixel.x] != 1 {
			arr[currPixel.y+1][currPixel.x] = 1
			stack = append(stack, Pair{currPixel.x, currPixel.y + 1})
		}
		// left
		if validCoord(currPixel.x-1, currPixel.y, cntColumn, cntRow) && arr[currPixel.y][currPixel.x-1] != 1 {
			arr[currPixel.y][currPixel.x-1] = 1
			stack = append(stack, Pair{currPixel.x - 1, currPixel.y})
		}
		// up
		if validCoord(currPixel.x, currPixel.y-1, cntColumn, cntRow) && arr[currPixel.y-1][currPixel.x] != 1 {
			arr[currPixel.y-1][currPixel.x] = 1
			stack = append(stack, Pair{currPixel.x, currPixel.y - 1})
		}
		lenStack = len(stack)
	}

}
func main() {
	matrix := [][]uint8{
		{1, 1, 1, 1, 1, 1, 1, 1},
		{0, 1, 0, 0, 0, 0, 0, 0},
		{0, 0, 1, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 0, 0, 0, 0},
		{1, 1, 1, 1, 1, 1, 1, 1},
	}
	fmt.Println("Before:")
	fmt.Println(matrix)
	fmt.Println("After (x=4, y=2):")
	//floodFillRecursion(matrix, 4, 2)
	floodFillStack(matrix, 4, 2)
	fmt.Println(matrix)
}
