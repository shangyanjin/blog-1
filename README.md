# Blog

A simple example of interfacing Go with a database, MySQL in this case.

## Schema

The schema for the database is defined in schema.sql.

## Test Database

I copy and paste schema.sql into the mysql REPL to load the examples,
but there is probably a better way to do it.

## Usage

run `go run main.go`. This shows an example of printing all entries
from all blogs and then printing all the titles of entries in general.
