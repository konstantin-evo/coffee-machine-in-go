package model

import (
	"fmt"
)

type CoffeeMachine struct {
	water     int
	milk      int
	beans     int
	cups      int
	money     int
	sugar     int
	chocolate int
	cinnamon  int
}

func NewCoffeeMachine(water, milk, beans, cups, money, sugar, chocolate, cinnamon int) *CoffeeMachine {
	return &CoffeeMachine{
		water:     water,
		milk:      milk,
		beans:     beans,
		cups:      cups,
		money:     money,
		sugar:     sugar,
		chocolate: chocolate,
		cinnamon:  cinnamon,
	}
}

func (coffeeMachine *CoffeeMachine) MakeCoffee(coffee *Coffee) bool {
	if coffeeMachine.water < coffee.Water {
		fmt.Println("Sorry, not enough water!")
		return false
	}
	if coffeeMachine.milk < coffee.Milk {
		fmt.Println("Sorry, not enough milk!")
		return false
	}
	if coffeeMachine.beans < coffee.Beans {
		fmt.Println("Sorry, not enough coffee beans!")
		return false
	}
	if coffeeMachine.cups < 1 {
		fmt.Println("Sorry, not enough disposable cups!")
		return false
	}
	if coffeeMachine.sugar < coffee.Sugar {
		fmt.Println("Sorry, not enough sugar!")
		return false
	}
	if coffeeMachine.chocolate < coffee.Chocolate {
		fmt.Println("Sorry, not enough sugar!")
		return false
	}
	if coffeeMachine.cinnamon < coffee.Cinnamon {
		fmt.Println("Sorry, not enough sugar!")
		return false
	}

	coffeeMachine.water -= coffee.Water
	coffeeMachine.milk -= coffee.Milk
	coffeeMachine.beans -= coffee.Beans
	coffeeMachine.sugar -= coffee.Sugar
	coffeeMachine.chocolate -= coffee.Chocolate
	coffeeMachine.cinnamon -= coffee.Cinnamon
	coffeeMachine.cups--
	coffeeMachine.money += coffee.Cost
	fmt.Println("Enjoy your coffee!")
	return true
}

func (coffeeMachine *CoffeeMachine) Fill(water int, milk int, beans int, cups int) {
	coffeeMachine.water += water
	coffeeMachine.milk += milk
	coffeeMachine.beans += beans
	coffeeMachine.cups += cups
	fmt.Println("Supplies added successfully.")
}

func (coffeeMachine *CoffeeMachine) TakeMoney() int {
	fmt.Printf("I gave you $%d\n", coffeeMachine.money)
	money := coffeeMachine.money
	coffeeMachine.money = 0
	return money
}

func (coffeeMachine *CoffeeMachine) GetStock() (int, int, int, int, int, int, int, int) {
	fmt.Printf("\nThe coffee machine has:\n%d ml of water\n%d ml of milk\n%d g of coffee beans\n%d disposable cups\n$%d of money\n%d g of sugar\n%d g of chocolate\n%d g of cinnamon\n", coffeeMachine.water, coffeeMachine.milk, coffeeMachine.beans, coffeeMachine.cups, coffeeMachine.money, coffeeMachine.sugar, coffeeMachine.chocolate, coffeeMachine.cinnamon)
	return coffeeMachine.water, coffeeMachine.milk, coffeeMachine.beans, coffeeMachine.cups, coffeeMachine.money, coffeeMachine.sugar, coffeeMachine.chocolate, coffeeMachine.cinnamon
}
