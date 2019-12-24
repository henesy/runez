# Runez

An absolutely awful compression format for utf-8 text.

This was implemented as a proof of concept to an idea discussed over coffee with a friend.

See [the spec](./spec.md) for implementation details, or read the source ☺.

## Build

	; go build

## Usage

	; go build && ./txtz < mac.txt | wc -c
	'л' has 12 locations
	'н' has 3 locations
	'ј' has 3 locations
	'г' has 3 locations
	'у' has 3 locations
	'М' has 3 locations
	'о' has 24 locations
	'а' has 6 locations
	'\x12' has 1 locations
	' ' has 18 locations
	'ч' has 3 locations
	'в' has 3 locations
	'и' has 6 locations
	'\n' has 3 locations
	'е' has 9 locations
	'т' has 6 locations
	'к' has 3 locations
	'з' has 3 locations
	'п' has 3 locations
	'с' has 3 locations
	195
	; wc -c mac.txt
	214 mac.txt
	;
