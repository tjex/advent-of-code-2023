# day 3 notes

## Requirements

- add up all part numbers in the engine schematic
- any number agacent to a symbol (also diagonally) is a "part number"
- periods, do not count as symbols

## Tactic


### 2d array

- essentially we have a 2d grid.
- part numbers that have a symbol on any cell surrounding them are a valid part number.
- fill a 2d array with rowsN and colsN of the input
- find numbers and their indicie per row.
- iterate through the neighbouring rows and cell +/- 1 of the same row and check if there are
  symbols
- if true, the digit, is a valid part number.

### 1d array

- essentially we have a 2d grid.
- part numbers that have a symbol on any cell surrounding them are a valid part number.
- but do we need a 2d array?
    - as before:
    - fill an array with lines of input file
    - find numbers and their indicie per row.
    - iterate through the neighbouring rows and cell +/- 1 of the same row and check if there are
      symbols
    - if true, the digit, is a valid part number.

## Review

- clear your array within the for loop before appending again you twit!
- learn to use a debugger more efficiently
- what caused me the biggest issue was not considering that when using the index range of a found
  number, that the end range already represented the cell that I should check until when searching
  for a symbol on lines above and below. 
    - This was hidden from me because my printDataAtRange() function was correctly printing the
      range of text, but my function to actually check for symbols was incorrect.
