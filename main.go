package main

import (
	"coffee-machine-in-go/model"
	"errors"
	"fmt"
	"strconv"
)

const (
	BuyAction       = "buy"
	FillAction      = "fill"
	TakeAction      = "take"
	RemainingAction = "remaining"
	ExitAction      = "exit"
)

const (
	ErrInvalidChoice = "invalid choice"
	ErrFailedToScan  = "failed to scan input"
	ErrInvalidInput  = "invalid input"
)

func main() {
	// initialize coffee machine with starting values
	coffeeMachine := model.NewCoffeeMachine(400, 540, 120, 9, 550, 500, 100, 100)

	for {
		// prompt user for action
		action, err := promptForAction()
		if err != nil {
			fmt.Printf("Error: %v. Please try again.\n", err)
			continue
		}

		switch action {
		case BuyAction:
			handleBuyAction(coffeeMachine)
		case FillAction:
			handleFillAction(coffeeMachine)
		case TakeAction:
			handleTakeAction(coffeeMachine)
		case RemainingAction:
			handleRemainingAction(coffeeMachine)
		case ExitAction:
			// exit the program
			return

		default:
			fmt.Println("Invalid action")
			continue
		}
	}
}

func handleBuyAction(coffeeMachine *model.CoffeeMachine) {
	// prompt user for type of coffee to buy
	choice, err := promptForCoffeeChoice()
	if err != nil {
		fmt.Printf("Error: %v. Please try again.\n", err)
		return
	}

	if choice == "back" {
		// return to main menu
		return
	}

	// convert the user input to an integer
	i, err := strconv.Atoi(choice)
	if err != nil || i < 1 || i > 5 {
		fmt.Println(ErrInvalidChoice)
		return
	}

	// prompt user for decorations
	fmt.Println("Do you want to add decorations? (1 - sugar, 2 - chocolate, 3 - cinnamon):")
	decorationsChoice, err := scanInt()
	if err != nil || decorationsChoice < 0 || decorationsChoice > 3 {
		fmt.Printf("Error: %v. Please try again.\n", ErrInvalidInput)
		return
	}

	// set the sugar to 0 initially
	sugar := 0

	// prompt user for level of sugar addition
	if decorationsChoice == 1 {
		fmt.Println("Choose level of sugar addition (1 - low, 2 - medium, 3 - high):")
		sugarLevelChoice, err := scanInt()
		if err != nil || sugarLevelChoice < 1 || sugarLevelChoice > 3 {
			fmt.Printf("Error: %v. Please try again.\n", ErrInvalidInput)
			return
		}
		sugar = sugarLevelChoice
	}

	// make the coffee with decorations
	coffee := model.CoffeeFactory(i)
	if decorationsChoice == 1 {
		coffee = model.AddSugar(coffee, sugar)
	} else if decorationsChoice == 2 {
		coffee = model.AddChocolate(coffee)
	} else if decorationsChoice == 3 {
		coffee = model.AddCinnamon(coffee)
	}

	isCoffeeMade := coffeeMachine.MakeCoffee(coffee)

	if !isCoffeeMade {
		return
	}
}

func handleFillAction(coffeeMachine *model.CoffeeMachine) {
	addWater, addMilk, addBeans, addCups := promptForAddingSupplies()
	// add supplies to coffee machine
	coffeeMachine.Fill(addWater, addMilk, addBeans, addCups)
}

func handleTakeAction(coffeeMachine *model.CoffeeMachine) {
	// give the money to the user and reset the coffee machine's money
	coffeeMachine.TakeMoney()
}

func handleRemainingAction(coffeeMachine *model.CoffeeMachine) {
	// display remaining stock levels
	coffeeMachine.GetStock()
}

func promptForAction() (string, error) {
	fmt.Println("\nWrite action (buy, fill, take, remaining, exit):")
	var action string
	_, err := fmt.Scan(&action)
	if err != nil {
		return "", errors.New(ErrFailedToScan)
	}
	return action, nil
}

func promptForCoffeeChoice() (string, error) {
	fmt.Println(model.GetCoffeeMenu())
	var choice string
	_, err := fmt.Scan(&choice)
	if err != nil {
		return "", errors.New(ErrFailedToScan)
	}
	return choice, nil
}

func promptForAddingSupplies() (int, int, int, int) {
	fmt.Println("Write how many ml of water you want to add:")
	addWater, err := scanInt()
	if err != nil {
		fmt.Printf("Error: %v. Please try again.\n", err)
		return promptForAddingSupplies()
	}

	fmt.Println("Write how many ml of milk you want to add:")
	addMilk, err := scanInt()
	if err != nil {
		fmt.Printf("Error: %v. Please try again.\n", err)
		return promptForAddingSupplies()
	}

	fmt.Println("Write how many grams of coffee beans you want to add:")
	addBeans, err := scanInt()
	if err != nil {
		fmt.Printf("Error: %v. Please try again.\n", err)
		return promptForAddingSupplies()
	}

	fmt.Println("Write how many disposable cups you want to add:")
	addCups, err := scanInt()
	if err != nil {
		fmt.Printf("Error: %v. Please try again.\n", err)
		return promptForAddingSupplies()
	}

	return addWater, addMilk, addBeans, addCups
}

func scanInt() (int, error) {
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		return 0, errors.New(ErrFailedToScan)
	}
	if n < 0 {
		return 0, errors.New(ErrInvalidInput)
	}
	return n, nil
}
