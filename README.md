# gitgo
Gitgo is a simple CLI app built with Golang that prints information about specified Github users to the command line

## Installation
go get github.com/nodejss/gitgo

## Usage
In order to view info on a single user. Use the -u or --user flag like so:
```
gitgo -u nodejss
```
To query for multiple users:
```
gitgo -u nodejss, <Another Username>,<Another Username>
```
