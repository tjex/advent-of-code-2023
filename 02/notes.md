Things learnt. 

- fmt.Sprintf is great. It makes creating new digits/strings/etc easy.
- When iterating and appending to a new array, filter out empty / invalid objects before appending

This is the correct way to iterate through a string to grep:

```go
	re := regexp.MustCompile(pattern)

		for i := range line {
			found := re.FindString(line[i:])
			// Don't want to append empty strings to our slice
			if found != "" {
				currentLine = append(currentLine, found)
			}
		}

```

Got burnt by incorrect grep pattern.
Given this string: "prrvrjlpgxpjdxfchqonepchqbhqxx9nbrvh"

```go
	const pattern = `^(one)|(two)|(three)|(four)|(five)|(six)|(seven)|(eight)|(nine)|(zero)|(1)|(2)|(3)|(4)|(5)|(6)|(7)|(8)|(9)|(0)`
````
gives 

```text
prrvrjlpgxpjdxfchqonepchqbhqxx9nbrvh
[one]
rrvrjlpgxpjdxfchqonepchqbhqxx9nbrvh
[one one]
rvrjlpgxpjdxfchqonepchqbhqxx9nbrvh
[one one one]
vrjlpgxpjdxfchqonepchqbhqxx9nbrvh
// ... //
nepchqbhqxx9nbrvh
[one one one one one one one one one one one one one one one one one one one 9]

```
it's finding "one" and returning it for every iteration through the loop (see ./main.go)

```go
	const pattern = `^((one)|(two)|(three)|(four)|(five)|(six)|(seven)|(eight)|(nine)|(zero)|(1)|(2)|(3)|(4)|(5)|(6)|(7)|(8)|(9)|(0))`
```

gives 

```text

prrvrjlpgxpjdxfchqonepchqbhqxx9nbrvh
[]
rrvrjlpgxpjdxfchqonepchqbhqxx9nbrvh
[]
rvrjlpgxpjdxfchqonepchqbhqxx9nbrvh
[]
qonepchqbhqxx9nbrvh
[]
onepchqbhqxx9nbrvh
[one]
```

It only returns the search term *on the iteration it is found*. 
