# day 3 notes

## Requirements

- add up all part numbers in the engine schematic
- any number agacent to a symbol (also diagonally) is a "part number"
- periods, do not count as symbols

## Tactic

- essentially we have a 2d grid.
- part numbers that have a symbol on any cell surrounding them are a valid part number.
- fill a 2d array with rowsN and colsN of the input
- find numbers and their indicie per row.
- iterate through the neighbouring rows and cell +/- 1 of the same row and check if there are
  symbols
- if true, the digit, is a valid part number.