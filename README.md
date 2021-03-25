# TRIE
A quick implementation of a trie in Golang

## Running tests
From the project root:
```
go test ./... --coverprofile=coverage.out
```

Example output:
```
ok  	github.com/absolutscottie/trie/pkg/trie	0.134s	coverage: 100.0% of statements
```

## What is done?
* Insert
* Find
* Tests for the above two functions

## What needs to be done?
* Delete (and tests)

## Choices
I chose to use a map to store children. The other option was to allocate a 128 byte array which seemed wasteful. Either solution decouples time completxity for inserts and lookups from the number of items in the data structure. 

