# Txtz Specification

## Usage

	txt

## Archive format

The archive file is binary and consists of an index of characters and their positions starting from 0 packed in the form:

	[null][null][uint8 # positions][rune][uint8 position(s)…]

such as:

	2 a 0 3
	3 Z 9 212 9087
	2 ß 31 123121
	4 д 2 32 57 86545

The runes are valid utf-8 and the positions are uint8 integers.

Two nulls are used so null characters may be compressed.

When a null is encountered, the decompressor should read ahead one rune to check for a second null and start a new rune entry if a second null is found.

## Algorithm

The general conversion looks like:

	The quick brown fox

to

	==> TODO

## Restrictions

We assume:

- The whole file is read into memory
- There are no more than uint8 characters

