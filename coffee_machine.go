package main

import "fmt"

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

func (c *CoffeeMachine) MakeCoffee(coffee Coffee) bool {
	if c.water < coffee.water {
		fmt.Println("Sorry, not enough water!")
		return false
	}
	if c.milk < coffee.milk {
		fmt.Println("Sorry, not enough milk!")
		return false
	}
	if c.beans < coffee.beans {
		fmt.Println("Sorry, not enough coffee beans!")
		return false
	}
	if c.cups < 1 {
		fmt.Println("Sorry, not enough disposable cups!")
		return false
	}

	fmt.Printf("I have enough resources, making you a %s!\n", coffee.name)
	c.water -= coffee.water
	c.milk -= coffee.milk
	c.beans -= coffee.beans
	c.cups--
	c.money += coffee.cost

	return true
}

func (c *CoffeeMachine) AddWater(amount int) {
	c.water += amount
}

func (c *CoffeeMachine) AddMilk(amount int) {
	c.milk += amount
}

func (c *CoffeeMachine) AddBeans(amount int) {
	c.beans += amount
}

func (c *CoffeeMachine) AddCups(amount int) {
	c.cups += amount
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
