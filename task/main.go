package main

import (
	market "task/supermarket"
)

// main function
func main() {
	market.Get("Pen")
	market.Post("Grapes", "kunal") //The value of the new item can be anything int, float or string
	market.Put("Pen", 8.9)
	market.Delete("Pen")
}
