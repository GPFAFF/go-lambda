package main

import read "github.com/GPFAFF/go-lambda/read"

func main() {
	read.File("one_active.csv")
	read.File("two.csv")
}
