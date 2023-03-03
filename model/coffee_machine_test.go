package model

import (
	"testing"
)

func TestMakeCoffee(t *testing.T) {
	machine := NewCoffeeMachine(1000, 1000, 1000, 10, 0)
	espresso := CoffeeFactory(1)
	latte := CoffeeFactory(2)

	// Make sure the coffee machine can make an espresso
	if !machine.MakeCoffee(espresso) {
		t.Errorf("Failed to make an espresso")
	}
	if machine.water != 750 {
		t.Errorf("Incorrect water amount after making an espresso")
	}
	if machine.beans != 984 {
		t.Errorf("Incorrect coffee beans amount after making an espresso")
	}
	if machine.cups != 9 {
		t.Errorf("Incorrect cups amount after making an espresso")
	}
	if machine.money != 4 {
		t.Errorf("Incorrect money amount after making an espresso")
	}

	// Make sure the coffee machine can make a latte
	if !machine.MakeCoffee(latte) {
		t.Errorf("Failed to make a latte")
	}
	if machine.water != 400 {
		t.Errorf("Incorrect water amount after making a latte")
	}
	if machine.milk != 925 {
		t.Errorf("Incorrect milk amount after making a latte")
	}
	if machine.beans != 964 {
		t.Errorf("Incorrect coffee beans amount after making a latte")
	}
	if machine.cups != 8 {
		t.Errorf("Incorrect cups amount after making a latte")
	}
	if machine.money != 11 {
		t.Errorf("Incorrect money amount after making a latte")
	}

	// Make sure the coffee machine cannot make a latte due to lack of milk
	machine.milk = 0
	if machine.MakeCoffee(latte) {
		t.Errorf("Unexpectedly made a latte without enough milk")
	}
	if machine.money != 11 {
		t.Errorf("Money should not have been updated after failed coffee making attempt")
	}
}

func TestFill(t *testing.T) {
	coffeeMachine := NewCoffeeMachine(0, 0, 0, 0, 0)

	// Test filling the machine with valid amounts
	coffeeMachine.Fill(100, 200, 300, 2)
	if coffeeMachine.water != 100 || coffeeMachine.milk != 200 || coffeeMachine.beans != 300 || coffeeMachine.cups != 2 {
		t.Error("Failed to fill coffee machine with valid amounts")
	}

}

func TestTakeMoney(t *testing.T) {
	coffeeMachine := NewCoffeeMachine(0, 0, 0, 0, 500)

	// Test taking money from the machine
	if money := coffeeMachine.TakeMoney(); money != 500 {
		t.Errorf("Expected to take $500 from coffee machine, but took $%d", money)
	}
	if coffeeMachine.money != 0 {
		t.Error("Money not correctly deducted from coffee machine")
	}
}

func TestGetStock(t *testing.T) {
	coffeeMachine := NewCoffeeMachine(500, 400, 300, 5, 0)

	// Call GetStock and check returned values
	water, milk, beans, cups, money := coffeeMachine.GetStock()
	if water != 500 || milk != 400 || beans != 300 || cups != 5 || money != 0 {
		t.Errorf("Expected (500, 400, 300, 5, 0) but got (%d, %d, %d, %d, %d)", water, milk, beans, cups, money)
	}
}
