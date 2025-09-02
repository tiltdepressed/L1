package main

import "fmt"

type Human struct {
	Name   string
	Age    int
	Height float32
	Weight float32
}

func (h *Human) printHuman() {
	fmt.Printf("Name: %s, Age: %d, Height: %.2f, Weight: %.1f\n", h.Name, h.Age, h.Height, h.Weight)
}

type Action struct {
	Human
	Salary int
}

func (a *Action) printAction() {
	a.printHuman()
	fmt.Printf("Salary: %d\n", a.Salary)
}

func main() {
	human := Action{Human: Human{Name: "Alex", Age: 14, Height: 170.0, Weight: 65.0}, Salary: 10000}
	human.printAction()
}
