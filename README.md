## Parse

#This is a file parser written in Go.

The original properties.txt file has windows carriage returns, so you can use the linux command:

`tr '\r' '\n' <properties.txt> properties-unix.txt`

...to make it less monstrous (and thus parsable by the bufio library).

Some assumptions:
- Although highly unlikely, it is possible that there may be a collision on [streetAddress, saleDate] so I added the town to strengthen the key. Edge casey I know.
- The first three tests asked for the list to be simply printed, and not returned in any way, so I left it as a map.
- I assumed that I/O is the most costly operation, so only parse the file once and pass around a slice of rawdata.

Usage of this program:

- compile using: go build parse.go
- run using: ./parse <filename> 
