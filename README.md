# What is this?

I'm just messing around and using the 2022 advent of code to learn a new language.
I'm thinking about getting more into go because it's always interested me.

[Link](https://adventofcode.com/2022/)

To run any of these days, simply run the following:

```bash
$ go run advent5.go
```

# Notes

## advent5.go

This was my first start on the advent calendar this year and I did it in probably
the most backwards way possible. I didn't realize until I had written out all code
that the site would give you some input to parse and enter an answer into. I didn't feel like spending too much time on fixing my code so that it could read some input from a file, so I just used some multicursor editing in my text editor to create the necessary function calls that the input required.

Part 1 solution: `VRWBSFZWM`
Part 2 solution: `RBTWJWMCF`

I got both stars for the day, but I'll probably go back and fix this by adding some actual input parsing.


## advent6.go

Okay, I read in an input file easily enough. I didn't use any structs here, although
I probably could have for the data stream processing. Instead I just went for
some simple functions that iterate over the buffer of characters a character at
a time while I attempted to detect the marker packets.

For the uniqueness check, found in `allUniqueCharacters()` I initially was looking
for a "set" data structure in go, because that's typically how I handle it in
python (awesome built-in type). Then all I have to do is put the packet into a set
and compare its cardinality to the original packet size to detect if all the characters
are unique.  Unfortunately it didn't look like there was any built-in "set" data structure in go, so I thought about a slightly different solution. Ultimately,
what I settled on would be add the characters, one by one, into a hashmap. While
adding them the hashmap I would check for membership and this would allow me to
fail fast if a packet contained repeated characters. This was simple enough to
implement with built-in map data structure in go, so I was happy with it.

For additional challenges here, if I would redo this challenge, I think I could
do a couple things:

* better argument parsing (files vs command line input)
* data stream struct for processing
* Create a simple "set" struct with the operations I needed.

Part 1 solution: `1757`
Part 2 solution: `2950`
