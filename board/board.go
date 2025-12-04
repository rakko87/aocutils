package board

import (
	"cmp"
	"fmt"
	"slices"
)

// Board collects the 2D input puzzle as a map of complex numbers
// board[Z(real,imag)] gives the value as the given position Z(real,imag)
type Board map[complex128]rune

// Populate board from strings
// Put origo to bottom left corner (south-west corner = (0,0))
// Top left corner becomes (column0, max(row))
func (m *Board) Populate(strings []string) {
	for ii, row := range strings {
		for jj, c := range row {
			// Move x-axis to bottom row. Make north positive direction
			z := complex(float64(jj), float64(len(strings)-ii))
			(*m)[z] = c
		}
	}
}

// Max complex for real and imaginary part
func (m *Board) Max() (maxReal, maxImag complex128) {
	positions := make([]complex128, len(*m))
	i := 0
	for pos, c := range *m {
		positions[i] = pos
		_ = c
		i++
	}
	maxReal = slices.MaxFunc(positions, func(a, b complex128) int {
		return cmp.Compare(real(a), real(b))
	})
	maxImag = slices.MaxFunc(positions, func(a, b complex128) int {
		return cmp.Compare(imag(a), imag(b))
	})
	return maxReal, maxImag
}

// Print board/map
func (mPtr *Board) Print() {
	maxReal, maxImag := mPtr.Max()
	m := *mPtr

	rows := make([][]rune, int(imag(maxImag))+1)
	for i := range rows {
		rows[i] = make([]rune, int(real(maxReal))+1)
	}
	for pos, c := range m {
		row, col := m.ZtoRowColumn(pos)
		// row, col := ZtoRowCol(pos, len(rows))
		rows[row][col] = c
	}
	for _, r := range rows {
		p(string(r))
	}
}

func (m *Board) LocateSymbol(c rune) complex128 {
	for pos, v := range *m {
		if v == c {
			return pos
		}
	}
	return -1 - 1i
}

// Complex z to (row, column)
func (m *Board) ZtoRowColumn(z complex128) (row, column int) {
	_, maxImag := m.Max()
	maxRows := int(imag(maxImag))
	row = maxRows - int(imag(z))
	column = int(real(z))
	return row, column
}

var p = fmt.Println

/*
           M x N

     (0,0)┌────N───┐(N,0)
 │        │        │
 │        │        │
Im, i     M        │
 │        │        │
 │        │        │
 ▼   (0,M)└────────┘(N,M)

          ──Re,─j──►

i ∈ [0,M]
j ∈ [0,N]

for j := range row

z = complex(j, i)

func matrix2coord()
return complex(j, M-i)

func NewZ2m( m map[(j,i)] func (complex) (int,int))
    max(map.keys(m))
    return func (z complex) (j,i int)


func coord2matrix(0,0) -> (0, M)
z2m(0,0)
*/
