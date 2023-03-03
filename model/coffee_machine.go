package model

import (
	"fmt"
)

type CoffeeMachine struct {
	water int
	milk  int
	beans int
	cups  int
	money int
}

func NewCoffeeMachine(water, milk, beans, cups, money int) *CoffeeMachine {
	return &CoffeeMachine{
		water: water,
		milk:  milk,
		beans: beans,
		cups:  cups,
		money: money,
	}
}

func (c *CoffeeMachine) MakeCoffee(coffee *Coffee) bool {
	if c.water < coffee.Water {
		fmt.Println("Sorry, not enough water!")
		return false
	}
	if c.milk < coffee.Milk {
		fmt.Println("Sorry, not enough milk!")
		return false
	}
	if c.beans < coffee.Beans {
		fmt.Println("Sorry, not enough coffee beans!")
		return false
	}
	if c.cups < 1 {
		fmt.Println("Sorry, not enough disposable cups!")
		return false
	}

	c.water -= coffee.Water
	c.milk -= coffee.Milk
	c.beans -= coffee.Beans
	c.cups--
	c.money += coffee.Cost
	fmt.Println("Enjoy your coffee!")
	return true
}

func (c *CoffeeMachine) Fill(water int, milk int, beans int, cups int) {
	c.water += water
	c.milk += milk
	c.beans += beans
	c.cups += cups
	fmt.Println("Supplies added successfully.")
}

func (c *CoffeeMachine) TakeMoney() int {
	fmt.Printf("I gave you $%d\n", c.money)
	money := c.money
	c.money = 0
	return money
}

func (c *CoffeeMachine) GetStock() (int, int, int, int, int) {
	fmt.Printf("\nThe coffee machine has:\n%d ml of water\n%d ml of milk\n%d g of coffee beans\n%d disposable cups\n$%d of money\n", c.water, c.milk, c.beans, c.cups, c.money)
	return c.water, c.milk, c.beans, c.cups, c.money
}
