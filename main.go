package main

import (
	"fmt"
	"strconv"
)

type Coffee struct {
	name  string
	water int
	milk  int
	beans int
	cost  int
}

func main() {
	// initialize coffee machine with starting values
	water := 400
	milk := 540
	beans := 120
	cups := 9
	money := 550

	espresso := Coffee{"Espresso", 250, 0, 16, 4}
	latte := Coffee{"Latte", 350, 75, 20, 7}
	cappuccino := Coffee{"Cappuccino", 200, 100, 12, 6}

	for {
		// prompt user for action
		fmt.Println("\nWrite action (buy, fill, take, remaining, exit):")
		var action string
		fmt.Scan(&action)

		switch action {
		case "buy":
			// prompt user for type of coffee to buy
			fmt.Println("What do you want to buy? 1 - espresso, 2 - latte, 3 - cappuccino, back - to main menu:")
			var choice string
			fmt.Scan(&choice)

			if choice == "back" {
				// return to main menu
				continue
			}

			// convert the user input to an integer
			i, err := strconv.Atoi(choice)
			if err != nil {
				fmt.Println("Invalid choice")
				continue
			}

			// select the appropriate coffee object
			var coffee Coffee
			switch i {
			case 1:
				coffee = espresso
			case 2:
				coffee = latte
			case 3:
				coffee = cappuccino
			default:
				fmt.Println("Invalid choice")
				continue
			}

			// check if there are enough ingredients to make the coffee
			if water < coffee.water {
				fmt.Println("Sorry, not enough water!")
				continue
			} else if milk < coffee.milk {
				fmt.Println("Sorry, not enough milk!")
				continue
			} else if beans < coffee.beans {
				fmt.Println("Sorry, not enough coffee beans!")
				continue
			} else if cups < 1 {
				fmt.Println("Sorry, not enough disposable cups!")
				continue
			}

			// make the coffee
			fmt.Printf("I have enough resources, making you a %s!\n", coffee.name)
			water -= coffee.water
			milk -= coffee.milk
			beans -= coffee.beans
			cups -= 1
			money += coffee.cost

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
			water += addWater
			milk += addMilk
			beans += addBeans
			cups += addCups

		case "take":
			// give the money to the user and reset the coffee machine's money
			fmt.Printf("I gave you $%d\n", money)
			money = 0

		case "remaining":
			// display remaining stock levels
			fmt.Printf("\nThe coffee machine has:\n%d ml of water\n%d ml of milk\n%d g of coffee beans\n%d disposable cups\n$%d of money\n", water, milk, beans, cups, money)

		case "exit":
			// exit the program
			return

		default:
			fmt.Println("Invalid action")
			continue
		}
	}
}
