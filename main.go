package main

import (
	"bufio"
	"container/list"
	"encoding/binary"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
)

// Encodes a rune and a position for decompression
type Pair struct {
	R rune
	P uint8
}

// For "sort"
type ByPosition []Pair

func (a ByPosition) Len() int           { return len(a) }
func (a ByPosition) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByPosition) Less(i, j int) bool { return a[i].P < a[j].P }

var chatty bool

// Naive utf-8 text compression/decompression program
func main() {
	c := flag.Bool("c", false, "Explicit compress mode")
	d := flag.Bool("d", false, "De-compress mode")
	flag.BoolVar(&chatty, "D", false, "Chatty debug mode")
	flag.Parse()

	// Compress by default
	if (!*c) && (!*d) {
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
		if chatty {
			fmt.Fprintf(os.Stderr, "%q has %v locations\n", r, l.Len())
		}

		// Position count
		pc := byte(uint8(l.Len()))

		err := binary.Write(w, binary.LittleEndian, pc)
		if err != nil {
			fatal("err: could not write position count - ", err)
		}

		// Rune
		err = binary.Write(w, binary.LittleEndian, r)
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
	dict := make(map[rune][]uint8)
	sum := 0

	// Populate lists by reading all definitions
	for {
		// Read out position count
		var pc uint8

		err := binary.Read(r, binary.LittleEndian, &pc)
		if err != nil {
			if err == io.EOF {
				break
			}

			fatal("err: could not read position count - ", err)
		}

		// Read out rune
		var ru rune

		err = binary.Read(r, binary.LittleEndian, &ru)
		if err != nil {
			fatal("err: could not read rune - ", err)
		}

		dict[ru] = make([]uint8, pc)
		sum += int(pc)

		if chatty {
			fmt.Fprintf(os.Stderr, "%q has %v locations\n", ru, pc)
		}

		// Read positions into slice
		for i := uint8(0); i < pc; i++ {
			var p uint8

			err = binary.Read(r, binary.LittleEndian, &p)
			if err != nil {
				fatal("err: could not read position #", i, "-", err)
			}

			dict[ru][i] = p
		}
	}

	// Merge slices
	master := make([]Pair, 0, sum)

	for ru, s := range dict {
		for _, p := range s {
			master = append(master, Pair{ru, p})
		}
	}

	// Sort super slice using "sort" for uint8
	sort.Sort(ByPosition(master))

	// Emit file in order of uint8 positions
	for _, pair := range master {
		w.Write([]byte(string(pair.R)))
	}

	w.Flush()
}

// Fatal - end program with an error message and newline
func fatal(s ...interface{}) {
	fmt.Fprintln(os.Stderr, s...)
	os.Exit(1)
}
