package main

import (
	read "github.com/GPFAFF/go-lambda/read"
	send "github.com/GPFAFF/go-lambda/sqs/send"
)

func main() {
	// v := read.File("one_active.csv")
	x := read.File("two.csv")

	send.Message(x)

}
