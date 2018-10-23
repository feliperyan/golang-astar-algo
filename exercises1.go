// If I have a function that accepts a pointer to a struct
// instead of just the struct type, is it more efficient?
// If I pass the struct am I then copying the whole thing?

// YES.

package main

import (
	"fmt"
)

type map_element struct {
	pos_x    int
	pos_y    int
	name     string
	passable bool
}

type the_map struct {
	x     int
	y     int
	two_d [][]map_element
}

func new_map(width, height int) the_map {
	a_map := the_map{width, height, make([][]map_element, width)}

	for i := 0; i < width; i++ {
		a_map.two_d[i] = make([]map_element, height)
		for j := 0; j < height; j++ {
			a_map.two_d[i][j] = map_element{i, j, ".", true}
		}
	}

	return a_map
}

func print_map(a_map *the_map) string {
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

func get_passable_neighbours(x, y int, a_map *the_map) []map_element {
	l := make([]map_element, 0)

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

func main() {
	fmt.Print("Aloha\n")
	a_map := new_map(6, 6)

	bob := map_element{1, 1, "b", false}
	a_map.two_d[1][1] = bob

	fmt.Printf(print_map(&a_map))

	// hood := get_passable_neighbours(1, 1, &a_map)
	// for _, homie := range hood {
	// 	h := &a_map.two_d[homie.pos_x][homie.pos_y]
	// 	h.name = "@"
	// }

	a_map.two_d[2][0] = map_element{2, 0, "#", false}
	a_map.two_d[2][1] = map_element{2, 1, "#", false}
	a_map.two_d[2][2] = map_element{2, 2, "#", false}
	a_map.two_d[1][2] = map_element{1, 2, "#", false}
	a_map.two_d[4][1] = map_element{4, 1, "*", false}

	fmt.Printf(print_map(&a_map))

}
