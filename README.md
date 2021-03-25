# TRIE
A quick implementation of a trie in Golang

## What is done?
* Insert
* Find
* Tests for the above two functions

## What needs to be done?
* Delete (and tests)

## Choices
I chose to use a map to store children. The other option was to allocate a 128 byte array which seemed wasteful. Either solution decouples time completxity for inserts and lookups from the number of items in the data structure. 
