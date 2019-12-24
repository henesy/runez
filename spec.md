# Txtz Specification

## Usage

	txtz [-c] [-d] < [input] > [output]

## Archive format

The archive file is little-endian binary and consists of an index of characters and their positions starting from 0 packed in the form:

	[null][uint8 # positions][rune][uint8 position(s)…]

such as:

	\0 2 a 0 3
	\0 3 Z 9 212 9087
	\0 2 ß 31 123121
	\0 4 д 2 32 57 86545

The runes are valid utf-8 and the positions are uint8 integers.

## Algorithm

The general conversion looks like:

	The quick brown fox

to

	==> TODO

## Restrictions

We assume:

- The whole file is read into memory
- There are no more than uint8 characters

