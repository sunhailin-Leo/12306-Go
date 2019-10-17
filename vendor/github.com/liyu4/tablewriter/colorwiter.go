// +build !windows

package tablewriter

import "io"

type Table struct {
	out            io.Writer
	mode           outputMode
	rows           [][]string
	lines          [][][]string
	cs             map[int]int
	rs             map[int]int
	headers        []string
	footers        []string
	autoFmt        bool
	autoWrap       bool
	mW             int
	pCenter        string
	pRow           string
	pColumn        string
	tColumn        int
	tRow           int
	hAlign         int
	fAlign         int
	align          int
	newLine        string
	rowLine        bool
	autoMergeCells bool
	hdrLine        bool
	borders        Border
	colSize        int
}

func (t *Table) Write(p []byte) (int, error) {
	return t.out.Write(p)
}
