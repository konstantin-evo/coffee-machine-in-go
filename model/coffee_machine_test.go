package model

import "testing"

func TestNewCoffeeMachine(t *testing.T) {
	coffeeMachine := NewCoffeeMachine(100, 200, 300, 400, 500, 600, 700, 800)
	if coffeeMachine == nil {
		t.Error("Expected a non-nil CoffeeMachine instance, got nil")
	}
	if coffeeMachine.water != 100 {
		t.Errorf("Expected water to be 100, got %d", coffeeMachine.water)
	}
	if coffeeMachine.milk != 200 {
		t.Errorf("Expected milk to be 200, got %d", coffeeMachine.milk)
	}
	if coffeeMachine.beans != 300 {
		t.Errorf("Expected beans to be 300, got %d", coffeeMachine.beans)
	}
	if coffeeMachine.cups != 400 {
		t.Errorf("Expected cups to be 400, got %d", coffeeMachine.cups)
	}
	if coffeeMachine.money != 500 {
		t.Errorf("Expected money to be 500, got %d", coffeeMachine.money)
	}
	if coffeeMachine.sugar != 600 {
		t.Errorf("Expected sugar to be 600, got %d", coffeeMachine.sugar)
	}
	if coffeeMachine.chocolate != 700 {
		t.Errorf("Expected chocolate to be 700, got %d", coffeeMachine.chocolate)
	}
	if coffeeMachine.cinnamon != 800 {
		t.Errorf("Expected cinnamon to be 800, got %d", coffeeMachine.cinnamon)
	}
}

func TestCoffeeMachine_MakeCoffee(t *testing.T) {
	coffeeMachine := NewCoffeeMachine(500, 500, 500, 10, 0, 500, 500, 500)
	coffee := &Coffee{
		Name:      "Espresso",
		Water:     50,
		Milk:      0,
		Beans:     18,
		Cost:      60,
		Sugar:     0,
		Chocolate: 0,
		Cinnamon:  0,
	}
	coffeeMachine.MakeCoffee(coffee)
	if coffeeMachine.water != 450 {
		t.Errorf("Expected water to be 450, got %d", coffeeMachine.water)
	}
	if coffeeMachine.milk != 500 {
		t.Errorf("Expected milk to be 500, got %d", coffeeMachine.milk)
	}
	if coffeeMachine.beans != 482 {
		t.Errorf("Expected beans to be 482, got %d", coffeeMachine.beans)
	}
	if coffeeMachine.cups != 9 {
		t.Errorf("Expected cups to be 9, got %d", coffeeMachine.cups)
	}
	if coffeeMachine.money != 60 {
		t.Errorf("Expected money to be 60, got %d", coffeeMachine.money)
	}
	if coffeeMachine.sugar != 500 {
		t.Errorf("Expected sugar to be 500, got %d", coffeeMachine.sugar)
	}
	if coffeeMachine.chocolate != 500 {
		t.Errorf("Expected chocolate to be 500, got %d", coffeeMachine.chocolate)
	}
	if coffeeMachine.cinnamon != 500 {
		t.Errorf("Expected cinnamon to be 500, got %d", coffeeMachine.cinnamon)
	}
}

func TestCoffeeMachine_Fill(t *testing.T) {
	coffeeMachine := NewCoffeeMachine(0, 0, 0, 0, 0, 0, 0, 0)

	// Fill with some values
	coffeeMachine.Fill(200, 100, 50, 10)

	// Check if values are correct after filling
	if coffeeMachine.water != 200 {
		t.Errorf("Expected water to be 200, but got %d", coffeeMachine.water)
	}
	if coffeeMachine.milk != 100 {
		t.Errorf("Expected milk to be 100, but got %d", coffeeMachine.milk)
	}
	if coffeeMachine.beans != 50 {
		t.Errorf("Expected beans to be 50, but got %d", coffeeMachine.beans)
	}
	if coffeeMachine.cups != 10 {
		t.Errorf("Expected cups to be 10, but got %d", coffeeMachine.cups)
	}
}
