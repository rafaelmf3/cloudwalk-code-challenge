# This is the solution for the CloudWalk code Challenge
## a quake game log parser with two kinds of report

# Run Locally

```
    go run main.go {cli command} {file}
```

cli commands:
- `parse` is used to create a structured file for the quake log, you need add as a argument the filename

- `gamereport` is used for generate the report for the games, with total kills and a player ranking, you need run `parse` first

- `deathcausereport` is user for generate the report for kills by meaning for each game, you need run `parse` first

# How to build

```
go build -o parser main.go 
```
after run this command you can use `./parser {cli command} {filename}`