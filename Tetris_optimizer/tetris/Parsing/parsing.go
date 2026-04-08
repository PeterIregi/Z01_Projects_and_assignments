package tetris


import(
	model "tetris/tetris/model"
	//"fmt"
)

func ParseTetrominoes(lines []string) ([]model.Tetromino, bool) {
	var pieces []model.Tetromino
	letter := 'A'
	i := 0

	for i < len(lines) {

		// must have 4 lines available
		if i+3 >= len(lines) {
			return nil, false
		}

		var blocks []model.Point
		count := 0

		// read 4x4 grid
		for y := 0; y < 4; y++ {
			line := lines[i+y]

			if len(line) != 4 {
				return nil, false
			}

			for x := 0; x < 4; x++ {
				if line[x] != '.' && line[x] != '#' {
					return nil, false
				}

				if line[x] == '#' {
					blocks = append(blocks, model.Point{X: x, Y: y})
					count++
				}
			}
		}

		// must contain exactly 4 blocks
		if count != 4 || !isValidTetromino(blocks) {
			return nil, false
		}

		t := model.Tetromino{
			Blocks: blocks,
			Letter: letter,
		}

		Normalize(&t)
		pieces = append(pieces, t)
		letter++

		i += 4

		// AFTER piece: expect ONE empty line or EOF
		if i < len(lines) {
			if lines[i] != "" {
				return nil, false
			}
			i++ // skip separator
		}
	}

	// must have at least one piece
	if len(pieces) == 0 {
		return nil, false
	}

	return pieces, true
}


func Normalize(t *model.Tetromino){
	minX :=t.Blocks[0].X
	minY :=t.Blocks[0].Y

	for  _, b := range t.Blocks {
		if b.X < minX {
			minX = b.X
		}
		if b.Y < minY {
			minY = b.Y
		}
	}

	for i := range t.Blocks {
		t.Blocks[i].X -=minX
		t.Blocks[i].Y -=minY
		
	}
}

func isValidTetromino(blocks []model.Point) bool {
	touches := 0

	for i := 0; i < len(blocks); i++ {
		for j := i + 1; j < len(blocks); j++ {

			dx := blocks[i].X - blocks[j].X
			if dx < 0 {
				dx = -dx
			}

			dy := blocks[i].Y - blocks[j].Y
			if dy < 0 {
				dy = -dy
			}

			// side neighbors only
			if dx+dy == 1 {
				touches++
			}
		}
	}

	// valid tetromino rule
	return touches == 3 || touches == 4
}
