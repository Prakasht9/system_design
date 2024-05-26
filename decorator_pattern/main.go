package main

import "fmt"

func main() {
	pizza := &veggieMania{}

	pizzaWithCheese := &CheeseTopping{
		pizza: pizza,
	}

	pizzaWithCheeseAndTomato := &TomatoTopping{
		pizza: pizzaWithCheese,
	}

	fmt.Printf("Price of veggeMania with tomato and cheese topping is %d\n", pizzaWithCheeseAndTomato.getPrice())

}
