// If I have a function that accepts a pointer to a struct
// instead of just the struct type, is it more efficient? Yes.
// If I pass the struct am I then copying the whole thing? Yes.

package main

import (
	"container/heap"
	"errors"
	"fmt"
	"math"
	"math/rand"
	"time"
)

type MapElement struct {
	pos_x    int
	pos_y    int
	name     string
	passable bool
	path     []*MapElement
}

type Map2d struct {
	x     int
	y     int
	two_d [][]*MapElement
}

func get_map_element(x int, y int, name string, passable bool) MapElement {
	return MapElement{x, y, name, passable, nil}
}

func new_map(width, height int) Map2d {
	a_map := Map2d{width, height, make([][]*MapElement, width)}

	for i := 0; i < width; i++ {
		a_map.two_d[i] = make([]*MapElement, height)
		for j := 0; j < height; j++ {
			a_map.two_d[i][j] = &MapElement{i, j, ".", true, nil}
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
			l = append(l, a_map.two_d[x-1][y])
		}
	}
	// top
	if x >= 0 && x < a_map.x && y-1 >= 0 && y-1 < a_map.y {
		if a_map.two_d[x][y-1].passable {
			l = append(l, a_map.two_d[x][y-1])
		}
	}
	// right
	if x+1 >= 0 && x+1 < a_map.x && y >= 0 && y < a_map.y {
		if a_map.two_d[x+1][y].passable {
			l = append(l, a_map.two_d[x+1][y])
		}
	}
	// bottom
	if x >= 0 && x < a_map.x && y+1 >= 0 && y+1 < a_map.y {
		if a_map.two_d[x][y+1].passable {
			l = append(l, a_map.two_d[x][y+1])
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

func putElementinMap2d(a_map *Map2d, name string, passable bool, x, y int) (*MapElement, error) {
	if x < 0 || y < 0 || x > (a_map.x-1) || y > (a_map.y-1) {
		return nil, errors.New("Out of bounds of the map2d")
	}

	// fmt.Printf("%v %v | ", x, y)

	element := MapElement{x, y, name, passable, nil}
	a_map.two_d[x][y] = &element

	return a_map.two_d[x][y], nil
}

func generateDungeon(width int, height int, maxTunnels int, maxLength int) Map2d {
	dungeon := Map2d{width, height, make([][]*MapElement, width)}

	for i := 0; i < width; i++ {
		dungeon.two_d[i] = make([]*MapElement, height)
		for j := 0; j < height; j++ {
			dungeon.two_d[i][j] = &MapElement{i, j, "#", false, nil}
		}
	}

	rand.Seed(time.Now().UnixNano())
	// a couple of good examples of how to deal with 2d array declarations:
	startingPoint := []int{rand.Intn(dungeon.x - 1), rand.Intn(dungeon.y - 1)}
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	randomDirection := rand.Intn(len(directions))
	lastDirection := randomDirection

	for i := 0; i < maxTunnels; i++ {

		if lastDirection <= 1 {
			randomDirection = rand.Intn(2) + 2
		} else {
			randomDirection = rand.Intn(2)
		}

		thisTunnelLength := rand.Intn(maxLength) + 1

		// fmt.Printf("Pos: %v  |  length: %v  |  direction:%v\n", startingPoint, thisTunnelLength, directions[randomDirection])

		for j := 0; j <= thisTunnelLength; j++ {

			x := startingPoint[0] + directions[randomDirection][0]
			y := startingPoint[1] + directions[randomDirection][1]

			if ele, err := putElementinMap2d(&dungeon, ".", true, x, y); err != nil {
				//hit edge of map.
				break
			} else {
				startingPoint[0] = ele.pos_x
				startingPoint[1] = ele.pos_y
			}
		}
		lastDirection = randomDirection

	}

	//fmt.Printf("%s | %s", startingPoint, directions)
	return dungeon
}

func testPath() {
	fmt.Print("\n")
	a_map := new_map(6, 6)

	// bob := MapElement{0, 2, "b", false, nil}
	// a_map.two_d[0][2] = bob

	// goal := MapElement{4, 1, "*", true, nil}
	// a_map.two_d[4][1] = goal

	bob, _ := putElementinMap2d(&a_map, "b", false, 0, 2)
	goal, _ := putElementinMap2d(&a_map, "*", true, 4, 1)

	a_map.two_d[1][0] = &MapElement{1, 0, "#", false, nil}
	a_map.two_d[2][0] = &MapElement{2, 0, "#", false, nil}
	a_map.two_d[2][1] = &MapElement{2, 1, "#", false, nil}
	a_map.two_d[2][2] = &MapElement{2, 2, "#", false, nil}
	a_map.two_d[2][3] = &MapElement{2, 3, "#", false, nil}
	a_map.two_d[1][3] = &MapElement{1, 3, "#", false, nil}

	fmt.Printf(print_map(&a_map))

	// results := aStarAlgorithm(&a_map, &bob, &a_map.two_d[4][1])
	// results := aStarAlgorithm(&a_map, &bob, &goal)
	results := aStarAlgorithm(&a_map, bob, goal)

	paintThem(results)

	fmt.Printf(print_map(&a_map))

	// paths := make(chan []*MapElement)

	// // Passing values into a channel via a go routine, this feels messy
	// // kinda repeating myself but it seems to work.
	// go func(a_map *Map2d, start *MapElement, goal *MapElement) {
	// 	paths <- aStarAlgorithm(a_map, start, goal)
	// }(&a_map, &bob, &a_map.two_d[4][1])

	// result := <-paths
	// fmt.Printf("\n Channel Value:\n%s", result)
}

func testDungeon() {
	fmt.Println("Generating dungeon...")
	dungeon := generateDungeon(10000, 10000, 8000, 3000)

	var bob *MapElement
	var goal *MapElement

	rand.Seed(time.Now().UnixNano())
	for {
		bob = dungeon.two_d[rand.Intn(dungeon.x-1)][rand.Intn(dungeon.y-1)]
		if bob.passable {
			// bob.name = "b"
			// bob.passable = false

			// I think putElement is screwing up.
			bob, _ = putElementinMap2d(&dungeon, "b", false, bob.pos_x, bob.pos_y)

			break
		}
	}
	for {
		goal = dungeon.two_d[rand.Intn(dungeon.x-1)][rand.Intn(dungeon.y-1)]
		if goal.passable {
			// goal.name = "*"
			goal, _ = putElementinMap2d(&dungeon, "0", true, goal.pos_x, goal.pos_y)
			break
		}
	}

	// fmt.Printf(print_map(&dungeon))

	fmt.Println("Done. Algo running...")

	way := aStarAlgorithm(&dungeon, bob, goal)

	fmt.Println("Algo done, len: ", len(way))

	// paintThem(way)

	// fmt.Printf(print_map(&dungeon))

}

func main() {

	fmt.Println("\nRunning main... \n")
	testDungeon()
}
