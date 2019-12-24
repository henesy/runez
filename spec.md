# Runez Specification

The file extension for a runez archive is `.rz`.

## Archive format

The archive file is little-endian binary and consists of an index of utf-8 characters (runes) and their positions starting from 0 packed in the form:

	[null][uint8 # positions][rune][uint8 position(s)…]

such as:

	\0 2 a 0 3
	\0 3 Z 9 212 9087
	\0 2 ß 31 123121
	\0 4 д 2 32 57 86545

The runes are valid utf-8 and the positions are uint8 integers.

## Algorithm

The general conversion looks like:

	αβξαβξ

to

	\02α03
	\02β14
	\02ξ25

## Restrictions

We assume:

- The whole file is read into memory
- There are no more than `^uint8(0)` runes
