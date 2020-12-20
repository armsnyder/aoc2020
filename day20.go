package main

import (
	"io"
	"math"
	"strconv"

	"github.com/armsnyder/aoc2020/aocutil"
)

var _ = declareDay(20, func(part2 bool, inputReader io.Reader) interface{} {
	if part2 {
		return day20Part2(inputReader)
	}
	return day20Part1(inputReader)
})

func day20Part1(inputReader io.Reader) interface{} {
	result := 1
	jigsaw := day20NewJigsaw(inputReader)
	for _, id := range jigsaw.cornerTileIDs() {
		result *= id
	}
	return result
}

func day20Part2(inputReader io.Reader) interface{} {
	jigsaw := day20NewJigsaw(inputReader)
	layout := jigsaw.solve()
	image := layout.image()
	image.eraseFigure(day20NewImage([]string{
		"                  # ",
		"#    ##    ##    ###",
		" #  #  #  #  #  #   ",
	}))
	return image.countFilledInCells()
}

type day20Tile struct {
	id    int
	image day20Image
}

func day20NewTile(ss []string) day20Tile {
	var tile day20Tile
	var err error

	tile.id, err = strconv.Atoi(ss[0][5:9])
	if err != nil {
		panic(err)
	}

	tile.image = day20NewImage(ss[1:])

	return tile
}

func (t day20Tile) edges() [4]day20Edge {
	var edges [4]day20Edge

	for i := 0; i < 10; i++ {
		if t.image[0][i] {
			edges[0] |= 1 << i
		}
		if t.image[i][9] {
			edges[1] |= 1 << i
		}
		if t.image[9][9-i] {
			edges[2] |= 1 << i
		}
		if t.image[9-i][0] {
			edges[3] |= 1 << i
		}
	}

	return edges
}

func (t day20Tile) reorient(fromEdge day20Edge, fromEdgeIndex int) day20Tile {
	tile := t
	for i, edge := range tile.edges() {
		if edge == fromEdge {
			if i%2 == 0 {
				tile.image.flipHorizontal()
			} else {
				tile.image.flipVertical()
			}
			break
		}
	}

	for i, edge := range tile.edges() {
		if edge.flip() == fromEdge {
			for i != (fromEdgeIndex+2)%4 {
				tile.image.rotate90()
				i = (i + 1) % 4
			}
			break
		}
	}

	return tile
}

type day20Edge uint16

func (e day20Edge) flip() day20Edge {
	var flippedEdge day20Edge
	for i := 0; i < 10; i++ {
		flippedEdge |= ((e >> i) & 1) << (9 - i)
	}
	return flippedEdge
}

type day20Jigsaw struct {
	tilesByID   map[int]day20Tile
	tilesByEdge map[day20Edge][]day20Tile
}

func day20NewJigsaw(inputReader io.Reader) day20Jigsaw {
	jigsaw := day20Jigsaw{
		tilesByID:   make(map[int]day20Tile),
		tilesByEdge: make(map[day20Edge][]day20Tile),
	}

	aocutil.VisitStringGroups(inputReader, func(ss []string) {
		tile := day20NewTile(ss)
		jigsaw.tilesByID[tile.id] = tile
		for _, edge := range tile.edges() {
			jigsaw.tilesByEdge[edge] = append(jigsaw.tilesByEdge[edge], tile)
			flippedEdge := edge.flip()
			jigsaw.tilesByEdge[flippedEdge] = append(jigsaw.tilesByEdge[flippedEdge], tile)
		}
	})

	return jigsaw
}

func (g day20Jigsaw) cornerTileIDs() []int {
	var cornerTileIDs []int

	for id, tile := range g.tilesByID {
		neighbors := 0

		for _, edge := range tile.edges() {
			if len(g.tilesByEdge[edge]) == 2 {
				neighbors++
			}
		}

		if neighbors == 2 {
			cornerTileIDs = append(cornerTileIDs, id)
		}
	}

	return cornerTileIDs
}

func (g day20Jigsaw) solve() day20JigsawLayout {
	type placement struct {
		tile  day20Tile
		coord day20Coord
	}

	seedTile := func() day20Tile {
		for _, tile := range g.tilesByID {
			return tile
		}
		panic("no tiles")
	}()

	layout := make(day20JigsawLayout)
	visitedTileIDs := make(map[int]bool)
	queue := []placement{{tile: seedTile, coord: day20Coord{}}}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		if visitedTileIDs[cur.tile.id] {
			continue
		}
		visitedTileIDs[cur.tile.id] = true

		layout[cur.coord] = cur.tile

		for i, edge := range cur.tile.edges() {
			tilesSharingEdge := g.tilesByEdge[edge]
			if len(tilesSharingEdge) != 2 {
				continue
			}
			var neighbor day20Tile
			for _, tile := range tilesSharingEdge {
				if tile.id != cur.tile.id {
					neighbor = tile
				}
			}
			neighbor = neighbor.reorient(edge, i)
			var deltaCoord day20Coord
			switch i {
			case 0:
				deltaCoord = day20Coord{i: -1}
			case 1:
				deltaCoord = day20Coord{j: 1}
			case 2:
				deltaCoord = day20Coord{i: 1}
			case 3:
				deltaCoord = day20Coord{j: -1}
			}
			neighborCoord := day20Coord{
				i: cur.coord.i + deltaCoord.i,
				j: cur.coord.j + deltaCoord.j,
			}
			queue = append(queue, placement{tile: neighbor, coord: neighborCoord})
		}
	}

	return layout
}

type day20Coord struct{ i, j int }

type day20JigsawLayout map[day20Coord]day20Tile

func (a day20JigsawLayout) image() day20Image {
	minI, minJ := math.MaxInt32, math.MaxInt32
	maxI, maxJ := math.MinInt32, math.MinInt32

	for coord := range a {
		if coord.i < minI {
			minI = coord.i
		} else if coord.i > maxI {
			maxI = coord.i
		}

		if coord.j < minJ {
			minJ = coord.j
		} else if coord.j > maxJ {
			maxJ = coord.j
		}
	}

	height := (maxI - minI + 1) * 8
	width := (maxJ - minJ + 1) * 8

	image := make(day20Image, height)
	for h := 0; h < height; h++ {
		image[h] = make([]bool, width)
	}

	for coordI := minI; coordI <= maxI; coordI++ {
		rootI := (coordI - minI) * 8
		for coordJ := minJ; coordJ <= maxJ; coordJ++ {
			rootJ := (coordJ - minJ) * 8

			coord := day20Coord{i: coordI, j: coordJ}
			tile := a[coord]

			for ti, row := range tile.image[1:9] {
				for tj, v := range row[1:9] {
					image[rootI+ti][rootJ+tj] = v
				}
			}
		}
	}

	return image
}

type day20Image [][]bool

func day20NewImage(ss []string) day20Image {
	image := make(day20Image, len(ss))
	for i, row := range ss {
		image[i] = make([]bool, len(row))
		for j, ch := range row {
			if ch == '#' {
				image[i][j] = true
			}
		}
	}
	return image
}

func (m *day20Image) eraseFigure(figure day20Image) {
	m.eraseFigureStrictFlipLooseRotation(figure)

	figure.flipHorizontal()

	m.eraseFigureStrictFlipLooseRotation(figure)

	figure.flipHorizontal()
	figure.flipVertical()

	m.eraseFigureStrictFlipLooseRotation(figure)
}

func (m *day20Image) eraseFigureStrictFlipLooseRotation(figure day20Image) {
	m.eraseFigureStrictOrientation(figure)
	for i := 0; i < 3; i++ {
		figure.rotate90()
		m.eraseFigureStrictOrientation(figure)
	}
}

func (m *day20Image) eraseFigureStrictOrientation(figure day20Image) {
	for imageRootI := 0; imageRootI < len(*m)-len(figure); imageRootI++ {
	nextImageLocation:
		for imageRootJ := 0; imageRootJ < len((*m)[0])-len(figure[0]); imageRootJ++ {
			for figureI := 0; figureI < len(figure); figureI++ {
				for figureJ := 0; figureJ < len(figure[0]); figureJ++ {
					if figure[figureI][figureJ] {
						imageI := imageRootI + figureI
						imageJ := imageRootJ + figureJ
						if !(*m)[imageI][imageJ] {
							continue nextImageLocation
						}
					}
				}
			}
			for figureI := 0; figureI < len(figure); figureI++ {
				for figureJ := 0; figureJ < len(figure[0]); figureJ++ {
					if figure[figureI][figureJ] {
						imageI := imageRootI + figureI
						imageJ := imageRootJ + figureJ
						(*m)[imageI][imageJ] = false
					}
				}
			}
		}
	}
}

func (m day20Image) countFilledInCells() int {
	total := 0
	for _, row := range m {
		for _, v := range row {
			if v {
				total++
			}
		}
	}
	return total
}

func (m *day20Image) flipHorizontal() {
	result := make(day20Image, len(*m))
	for i, row := range *m {
		result[i] = make([]bool, len(row))
	}

	for i, row := range *m {
		for j, v := range row {
			result[i][len(row)-1-j] = v
		}
	}

	*m = result
}

func (m *day20Image) flipVertical() {
	result := make(day20Image, len(*m))
	for i, row := range *m {
		result[i] = make([]bool, len(row))
	}

	for i, row := range *m {
		for j, v := range row {
			result[len(*m)-1-i][j] = v
		}
	}

	*m = result
}

func (m *day20Image) rotate90() {
	result := make(day20Image, len((*m)[0]))
	for i := range (*m)[0] {
		result[i] = make([]bool, len(*m))
	}

	for i, row := range *m {
		for j, v := range row {
			result[j][len(*m)-1-i] = v
		}
	}

	*m = result
}
