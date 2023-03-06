package model

import "fmt"

type CoffeeName string

const (
	Espresso   CoffeeName = "Espresso"
	Latte      CoffeeName = "Latte"
	Cappuccino CoffeeName = "Cappuccino"
	Mocha      CoffeeName = "Mocha"
	Americano  CoffeeName = "Americano"
)

type Coffee struct {
	Name      CoffeeName
	Water     int
	Milk      int
	Beans     int
	Cost      int
	Sugar     int
	Chocolate int
	Cinnamon  int
}

func CoffeeFactory(coffeeNumber int) *Coffee {
	switch coffeeNumber {
	case 1:
		return &Coffee{Name: Espresso, Water: 250, Milk: 0, Beans: 16, Cost: 4}
	case 2:
		return &Coffee{Name: Latte, Water: 350, Milk: 75, Beans: 20, Cost: 7}
	case 3:
		return &Coffee{Name: Cappuccino, Water: 200, Milk: 100, Beans: 12, Cost: 6}
	case 4:
		return &Coffee{Name: Mocha, Water: 200, Milk: 100, Beans: 16, Cost: 8}
	case 5:
		return &Coffee{Name: Americano, Water: 300, Milk: 0, Beans: 20, Cost: 5}
	default:
		return nil
	}
}

func AddSugar(c *Coffee, amount int) *Coffee {
	c.Sugar = amount
	return c
}

func AddChocolate(c *Coffee) *Coffee {
	c.Chocolate = 1
	return c
}

func AddCinnamon(c *Coffee) *Coffee {
	c.Cinnamon = 1
	return c
}

func GetCoffeeMenu() string {
	coffeeTypes := []CoffeeName{Espresso, Latte, Cappuccino, Mocha, Americano}
	var menuString string
	for i, name := range coffeeTypes {
		menuString += fmt.Sprintf("%d - %s", i+1, name)
		if i < len(coffeeTypes)-1 {
			menuString += ", "
		}
	}
	return fmt.Sprintf("What do you want to buy? %s, back - to main menu:", menuString)
}
