package main

import (
	"coffee-machine-in-go/model"
	"fmt"
	"strconv"
)

func main() {
	// initialize coffee machine with starting values
	coffeeMachine := model.NewCoffeeMachine(400, 540, 120, 9, 550)

	for {
		// prompt user for action
		fmt.Println("\nWrite action (buy, fill, take, remaining, exit):")
		var action string
		fmt.Scan(&action)

		switch action {
		case "buy":
			// prompt user for type of coffee to buy
			fmt.Println(model.GetCoffeeMenu())
			var choice string
			fmt.Scan(&choice)

			if choice == "back" {
				// return to main menu
				continue
			}

			// convert the user input to an integer
			i, err := strconv.Atoi(choice)
			if err != nil || i < 1 || i > 5 {
				fmt.Println("Invalid choice")
				continue
			}

			// make the coffee
			coffee := model.CoffeeFactory(i)
			isCoffeeMade := coffeeMachine.MakeCoffee(coffee)

			if !isCoffeeMade {
				continue
			}

		case "fill":
			// prompt user for amounts to fill
			fmt.Println("Write how many ml of water you want to add:")
			var addWater int
			fmt.Scan(&addWater)

			fmt.Println("Write how many ml of milk you want to add:")
			var addMilk int
			fmt.Scan(&addMilk)

			fmt.Println("Write how many grams of coffee beans you want to add:")
			var addBeans int
			fmt.Scan(&addBeans)

			fmt.Println("Write how many disposable cups you want to add:")
			var addCups int
			fmt.Scan(&addCups)

			// add supplies to coffee machine
			coffeeMachine.Fill(addWater, addMilk, addBeans, addCups)

		case "take":
			// give the money to the user and reset the coffee machine's money
			coffeeMachine.TakeMoney()

		case "remaining":
			// display remaining stock levels
			coffeeMachine.GetStock()

		case "exit":
			// exit the program
			return

		default:
			fmt.Println("Invalid action")
			continue
		}
	}
}
