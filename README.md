# Runez

An absolutely awful, but (probably) lossless, compression format for utf-8 text.

This was implemented as a proof of concept to an idea discussed over coffee with a friend.

See [the spec](./spec.md) for implementation details, or read the source ☺.

## Build

	; go build

## Usage

	runez [-D] [-c] [-d] < [input] > [output]

## Example

Compressing some text in Macedonian:

	; cat mac.txt
	Моето летачко возило е полно со јагули
	Моето летачко возило е полно со јагули
	Моето летачко возило е полно со јагули
	Моето летачко возило е полно со јагули
	Моето летачко возило е полно со јагули
	; wc -c mac.txt # Byte count
	355 mac.txt
	; ./runez < mac.txt > out.rz
	; wc -c out.rz
	309 out.rz
	; ./runez -d < out.rz > newmac.txt
	; diff mac.txt newmac.txt
	;

Pass runez output through itself:

	; ./runez -c < mac.txt | ./runez -d
	Моето летачко возило е полно со јагули
	Моето летачко возило е полно со јагули
	Моето летачко возило е полно со јагули
	Моето летачко возило е полно со јагули
	Моето летачко возило е полно со јагули
	;
