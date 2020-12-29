package main

import (
	"fmt"
	"log"
)

func main() {
	// Feel free to use the main function for testing your functions
	world := struct {
		English string
		Spanish string
		French  string
	}{
		"world",
		"mundo",
		"monde",
	}
	fmt.Printf("Hello, %s/%s/%s!\n", world.English, world.Spanish, world.French)
	fmt.Println(Prices["eggs"].String())
	c := Cart{Items: []string{"eggs", "milk"}, TotalPrice: 510}
	c2 := Cart{Items: []string{"eggs", "malk"}, TotalPrice: 510}
	fmt.Println(c.hasMilk())
	fmt.Println(c2.hasMilk())
	c.Checkout()
}

// Price is the cost of something in US cents.
type Price int64

// String is the string representation of a Price
// These should be represented in US Dollars
// Example: 2595 cents => $25.95
func (p Price) String() string {
	var dollars int
	if p > 100 {
		dollars = int((p - p%100) / 100)
	} else {
		dollars = int(p)
	}
	var cents int = int(p % 100)
	return fmt.Sprintf("$%d.%d", dollars, cents)
}

// Prices is a map from an item to its price.
var Prices = map[string]Price{
	"eggs":          219,
	"bread":         199,
	"milk":          295,
	"peanut butter": 445,
	"chocolate":     150,
}

// RegisterItem adds the new item in the prices map.
// If the item is already in the prices map, a warning should be displayed to the user,
// but the value should be overwritten.
// Bonus (1pt) - Use the "log" package to print the error to the user
func RegisterItem(prices map[string]Price, item string, price Price) {
	// check if item already exists
	_, ok := prices[item]
	if !ok {
		prices[item] = price
	} else {
		log.Println("Already exists in prices! Overriding...")
		prices[item] = price
	}
}

// Cart is a struct representing a shopping cart of items.
type Cart struct {
	Items      []string
	TotalPrice Price
}

// hasMilk returns whether the shopping cart has "milk".
func (c *Cart) hasMilk() bool {
	for _, item := range c.Items {
		if item == "milk" {
			return true
		}
	}
	return false
}

// HasItem returns whether the shopping cart has the provided item name.
func (c *Cart) HasItem(item string) bool {
	for _, cartItem := range c.Items {
		if cartItem == item {
			return true
		}
	}
	return false
}

// AddItem adds the provided item to the cart and update the cart balance.
// If item is not found in the prices map, then do not add it and print an error.
// Bonus (1pt) - Use the "log" package to print the error to the user
func (c *Cart) AddItem(item string) {
	price, ok := Prices[item]
	if !ok {
		log.Println("Item has no found price.")
	} else {
		c.Items = append(c.Items, item)
		c.TotalPrice += price
	}
}

// Checkout displays the final cart balance and clears the cart completely.
func (c *Cart) Checkout() {
	// Display final cart balance
	fmt.Printf("Final cart balance of: %s\n", fmt.Sprint(c.TotalPrice))
	// Clear cart
	c.Items = make([]string, 0)
	c.TotalPrice = Price(0)
}
