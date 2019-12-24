package main

import (
	"fmt"
	"bufio"
	"os"
	"io"
	"flag"
	"container/list"
//	"sort"
	"encoding/binary"
)


// Naive text compression/decompression program
func main() {
	c := flag.Bool("c", false, "Explicit compress mode")
	d := flag.Bool("d", false, "De-compress mode")
	flag.Parse()

	// Compress by default
	if (! *c) && (! *d) {
		*c = true
	}

	if *c == *d {
		fatal("err: choose one explicit mode")
	}

	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	// Choose mode operation
	switch {
	case *c:
		Compress(in, out)
	case *d:
		Decompress(in, out)
	}
}


// Compress text to the archive format
func Compress(r *bufio.Reader, w *bufio.Writer) {
	dict := make(map[rune]*list.List)

	// Build table
	for i := 0; ; i++ {
		r, _, err := r.ReadRune()
		if err != nil {
			if err == io.EOF {
				break
			}

			fatal("err: could not read rune - ", err)
		}

		if dict[r] == nil {
			dict[r] = list.New()
		}

		// Check for character count overflow
		if max := int(^uint8(0)); i > max {
			fatal("err: too many characters to compress, stopped at: ", max)
		}

		dict[r].PushFront(uint8(i))
	}

	// Iterate dict to emit output file format
	for r, l := range dict {
		// DEBUG
		fmt.Fprintf(os.Stderr, "%q has %v locations\n", r, l.Len())

		// Null byte preamble
		_, err := w.Write([]byte{0})
		if err != nil {
			fatal("err: could not write null byte - ", err)
		}

		// Position count
		pc := byte(uint8(l.Len()))

		err = binary.Write(w, binary.LittleEndian, pc)
		if err != nil {
			fatal("err: could not write position count - ", err)
		}

		// Rune
		err = binary.Write(w, binary.LittleEndian, []byte(string(r)))
		if err != nil {
			fatal("err: could not write rune - ", err)
		}

		// Positions
		for p := l.Front(); p != nil; p = p.Next() {
			err := binary.Write(w, binary.LittleEndian, byte(p.Value.(uint8)))
			if err != nil {
				fatal("err: could not write position - ", err)
			}
		}
	}

	w.Flush()
}

// Decompress the archive format to text
func Decompress(r *bufio.Reader, w *bufio.Writer) {
	//dict := make(map[rune]*list.List)

	// Turn lists → slice of {rune, uint8}

	// Sort slice using "sort" for uint8

	// Emit file in order of uint8 positions
}

// Fatal - end program with an error message and newline
func fatal(s ...interface{}) {
	fmt.Fprintln(os.Stderr, s...)
	os.Exit(1)
}
