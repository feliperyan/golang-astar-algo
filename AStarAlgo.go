// If I have a function that accepts a pointer to a struct
// instead of just the struct type, is it more efficient? Yes.
// If I pass the struct am I then copying the whole thing? Yes.

package main

import (
	"container/heap"
	"fmt"
	"math"
)

type MapElement struct {
	pos_x    int
	pos_y    int
	name     string
	passable bool
}

type Map2d struct {
	x     int
	y     int
	two_d [][]MapElement
}

func new_map(width, height int) Map2d {
	a_map := Map2d{width, height, make([][]MapElement, width)}

	for i := 0; i < width; i++ {
		a_map.two_d[i] = make([]MapElement, height)
		for j := 0; j < height; j++ {
			a_map.two_d[i][j] = MapElement{i, j, ".", true}
		}
	}

	return a_map
}

func print_map(a_map *Map2d) string {
	printed := "    "

	for i := 0; i < a_map.x; i++ {
		printed = fmt.Sprintf("%s %2d ", printed, i)
	}

	for n := 0; n < a_map.y; n++ {
		row := fmt.Sprintf(" %2d", n)
		for m := 0; m < a_map.x; m++ {
			row = fmt.Sprintf("%s   %s", row, a_map.two_d[m][n].name)
		}
		printed = fmt.Sprintf("%s\n%s", printed, row)
	}

	return fmt.Sprintf("%s\n\n", printed)
}

func get_passable_neighbours(x, y int, a_map *Map2d) []*MapElement {
	l := make([]*MapElement, 0)

	// left
	if x-1 >= 0 && x-1 < a_map.x && y >= 0 && y < a_map.y {
		if a_map.two_d[x-1][y].passable {
			l = append(l, &a_map.two_d[x-1][y])
		}
	}
	// top
	if x >= 0 && x < a_map.x && y-1 >= 0 && y-1 < a_map.y {
		if a_map.two_d[x][y-1].passable {
			l = append(l, &a_map.two_d[x][y-1])
		}
	}
	// right
	if x+1 >= 0 && x+1 < a_map.x && y >= 0 && y < a_map.y {
		if a_map.two_d[x+1][y].passable {
			l = append(l, &a_map.two_d[x+1][y])
		}
	}
	// bottom
	if x >= 0 && x < a_map.x && y+1 >= 0 && y+1 < a_map.y {
		if a_map.two_d[x][y+1].passable {
			l = append(l, &a_map.two_d[x][y+1])
		}
	}

	return l
}

func getPathToGoal(cameFrom map[*MapElement]*MapElement, start, goal *MapElement) []*MapElement {
	nodes := make([]*MapElement, 0)

	for e := goal; e != start; {
		// prepend trick so I don't have to reverse the list later
		nodes = append([]*MapElement{e}, nodes...) // this is prob inneficient.
		e = cameFrom[e]
	}
	return nodes
}

func paintThem(elements []*MapElement) {
	for _, e := range elements {
		e.name = "+"
	}
}

func heuristic(a, b *MapElement) int {

	return int(math.Abs(float64(a.pos_x-b.pos_x)) + math.Abs(float64(a.pos_y-b.pos_y)))
}

func aStarAlgorithm(a_map *Map2d, start *MapElement, goal *MapElement) []*MapElement {
	frontier := make(PriorityQueue, 0)
	heap.Init(&frontier)
	heap.Push(&frontier, &Item{value: start, priority: 0})

	cameFrom := make(map[*MapElement]*MapElement)
	costSoFar := make(map[*MapElement]int)
	cameFrom[start] = nil
	costSoFar[start] = 0

	finished := false

	for i := len(frontier); i > 0; {
		topItem := heap.Pop(&frontier).(*Item)
		current := topItem.value

		if current.pos_x == goal.pos_x && current.pos_y == goal.pos_y {
			finished = true
			break
		}

		neighbours := get_passable_neighbours(current.pos_x, current.pos_y, a_map)
		// paintThem(neighbours) // debug help
		// current.name = "o"
		// fmt.Printf(print_map(a_map))
		// fmt.Printf("%v\n", current)

		for _, nextElement := range neighbours {
			newCost := costSoFar[current] + 1 // static cost for now.
			_, exists := costSoFar[nextElement]

			if (!exists) || (newCost < costSoFar[nextElement]) {
				costSoFar[nextElement] = newCost
				priority := newCost + heuristic(goal, nextElement)
				heap.Push(&frontier, &Item{value: nextElement, priority: priority})
				cameFrom[nextElement] = current
			}
		}

		i = len(frontier)
	}

	if finished {
		path := getPathToGoal(cameFrom, start, goal)
		return path
	} else {
		path := make([]*MapElement, 0)
		return path
	}
}

func main() {
	fmt.Print("\n")
	a_map := new_map(6, 6)

	bob := MapElement{0, 2, "b", false}
	a_map.two_d[0][2] = bob

	a_map.two_d[1][0] = MapElement{1, 0, "#", false}
	a_map.two_d[2][0] = MapElement{2, 0, "#", false}
	a_map.two_d[2][1] = MapElement{2, 1, "#", false}
	a_map.two_d[2][2] = MapElement{2, 2, "#", false}
	a_map.two_d[2][3] = MapElement{2, 3, "#", false}
	a_map.two_d[1][3] = MapElement{1, 3, "#", false}

	a_map.two_d[4][1] = MapElement{4, 1, "*", true}

	fmt.Printf(print_map(&a_map))

	results := aStarAlgorithm(&a_map, &bob, &a_map.two_d[4][1])
	paintThem(results)

	fmt.Printf(print_map(&a_map))

}
