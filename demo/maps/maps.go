package main

import "fmt"

func main() {
	shoppingList := make(map[string]int)

	shoppingList["eggs"] = 10
	shoppingList["milk"] = 2
	shoppingList["bread"] += 1

	shoppingList["eggs"] += 2

	fmt.Println(shoppingList)

	delete(shoppingList, "bread")

	fmt.Println(shoppingList)

	fmt.Println("need", shoppingList["eggs"], "eggs")

	cereal, found := shoppingList["cereal"]

	if !found {
		fmt.Println("no cereal on the list")
	} else {
		fmt.Println("need", cereal, "cereal")
	}

	totalItem := 0

	for _, amount := range shoppingList {
		totalItem += amount
	}
	fmt.Println("total item", totalItem)
}
