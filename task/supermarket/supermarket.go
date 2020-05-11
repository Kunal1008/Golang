package supermarket

import (
	"fmt"
)

var cost = map[string]interface{}{
	"Pen":  5.5,
	"Surf": 10.2,
}

// Gives the items asked for
func Get(item string) {
	value, ok := cost[item]
	if ok {
		fmt.Println("Cost of", item, "is", value)
	} else {
		fmt.Println("Data not found")
	}
}

//Add items to the list if not available
func Post(item string, val interface{}) {
	value, ok := cost[item]
	if ok {
		fmt.Println("Item already exists", value)
	} else {
		cost[item] = val
		//fmt.Println(item, "added to the list with value ", val)
	}
}

//Updates the price of the value
func Put(item string, val interface{}) {
	value, ok := cost[item]
	if ok {
		cost[item] = val
		// fmt.Println(cost[item])
	} else {
		fmt.Println("Item", item, "with the value", value, "is not present")
	}
}

//Delete the item
func Delete(item string) {
	value, ok := cost[item]
	if ok {
		fmt.Println("Item with the value", value, "is present")
		delete(cost, item)
		fmt.Println(cost)
	} else {
		fmt.Println("Item is not present to be deleted")
	}
}
