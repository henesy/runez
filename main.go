package main

import (
	"bufio"
	"os"
	"io"
	"log"
	"flag"
	"container/list"
//	"sort"
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
		log.Fatal("err: choose one explicit mode")
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
	for i := uint8(0); ; i++ {
		r, _, err := r.ReadRune()
		if err != nil {
			if err == io.EOF {
				break
			}

			log.Fatal("err: could not read rune - ", err)
		}

		if dict[r] == nil {
			dict[r] = list.New()
		}

		dict[r].PushFront(i)
	}

	// Iterate dict to emit output file format
	for r, l := range dict {
		log.Printf("%c has %v locations", r, l.Len())
	}

	log.Println(dict)
}

// Decompress the archive format to text
func Decompress(r *bufio.Reader, w *bufio.Writer) {
	//dict := make(map[rune]*list.List)

	// Turn lists → slice of {rune, uint8}

	// Sort slice using "sort" for uint8

	// Emit file in order of uint8 positions
}
