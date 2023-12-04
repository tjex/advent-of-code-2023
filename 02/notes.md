# Notes Day 2

## Requirements

- cubes are red, green or blue
- for each game played, a secret number of cubes of each color will be hidden
- goal is to discover "information about the number of cubes" (I guess the amount of cubes in the
  bag?)
- to get info, after a bag is loaded with cubes, a handful of random cubes is shown and put back.
    - this happens a few times per game.
- I play several games and record the info from each game (the puzzle input)
- each game is listed with its id number "Game: 11" = id 11
    - the listing is followed by semicolon-separated list of the subsets of cubes that are revealed


### Example

For example, the record of a few games might look like this:

Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green

In game 1, three sets of cubes are revealed from the bag (and then put back again). The first set is
3 blue cubes and 4 red cubes; the second set is 1 red cube, 2 green cubes, and 6 blue cubes; the
third set is only 2 green cubes.

### Question

Which games would have been possible if the bag contained **only 12 red cubes, 13 green cubes and 14
blue cubes?**

Add the ids of each game that would be possible given this requirement.

### Tactic

- make a struct for a game
    - id (int)
    - red (int) - total cubes revealed for the game
    - green (int) - //
    - blue (int) - //

- iterate through the input and generate structs for each game
- filter the structs based on requirements
    - 12 red
    - 13 green
    - 14 blue

- return the ids of the filtered games
- sum the returned ids (ints)
