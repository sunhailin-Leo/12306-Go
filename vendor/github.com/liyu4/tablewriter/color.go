package tablewriter

import (
	"io"
)

type outputMode int

// DiscardNonColorEscSeq supports the divided color escape sequence.
// But non-color escape sequence is not output.
// Please use the OutputNonColorEscSeq If you want to output a non-color
// escape sequences such as ncurses. However, it does not support the divided
// color escape sequence.
const (
	_ outputMode = iota
	DiscardNonColorEscSeq
	OutputNonColorEscSeq
)

// NewColorWriter creates and initializes a new ansiColorWriter
// using io.Writer w as its initial contents.
// In the console of Windows, which change the foreground and background
// colors of the text by the escape sequence.
// In the console of other systems, which writes to w all text.
func NewColorWriter(w io.Writer) *Table {
	return NewModeColorWriter(w, OutputNonColorEscSeq)
}

func NewColorIO(w io.Writer) io.Writer {
	return NewModeColorIO(w, OutputNonColorEscSeq)
}

// NewModeColorWriter create and initializes a new ansiColorWriter
// by specifying the outputMode.
func NewModeColorWriter(w io.Writer, mode outputMode) *Table {
	if _, ok := w.(*Table); !ok {
		return &Table{
			out:      w,
			mode:     mode,
			rows:     [][]string{},
			lines:    [][][]string{},
			cs:       make(map[int]int),
			rs:       make(map[int]int),
			headers:  []string{},
			footers:  []string{},
			autoFmt:  true,
			autoWrap: true,
			mW:       MAX_ROW_WIDTH,
			pCenter:  CENTER,
			pRow:     ROW,
			pColumn:  COLUMN,
			tColumn:  -1,
			tRow:     -1,
			hAlign:   ALIGN_DEFAULT,
			fAlign:   ALIGN_DEFAULT,
			align:    ALIGN_DEFAULT,
			newLine:  NEWLINE,
			rowLine:  false,
			hdrLine:  true,
			borders:  Border{Left: true, Right: true, Bottom: true, Top: true},
			colSize:  -1}

	}
	return w.(*Table)
}

func NewModeColorIO(w io.Writer, mode outputMode) io.Writer {
	if _, ok := w.(*Table); !ok {
		return &Table{
			out:      w,
			mode:     mode,
			rows:     [][]string{},
			lines:    [][][]string{},
			cs:       make(map[int]int),
			rs:       make(map[int]int),
			headers:  []string{},
			footers:  []string{},
			autoFmt:  true,
			autoWrap: true,
			mW:       MAX_ROW_WIDTH,
			pCenter:  CENTER,
			pRow:     ROW,
			pColumn:  COLUMN,
			tColumn:  -1,
			tRow:     -1,
			hAlign:   ALIGN_DEFAULT,
			fAlign:   ALIGN_DEFAULT,
			align:    ALIGN_DEFAULT,
			newLine:  NEWLINE,
			rowLine:  false,
			hdrLine:  true,
			borders:  Border{Left: true, Right: true, Bottom: true, Top: true},
			colSize:  -1}

	}
	return w
}
