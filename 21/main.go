package main

import "fmt"

// Предположим, что это интерфейс, который нельзя менять

type Animal interface {
	Speak()
}

// Основная функция импортированного пакета

func MakeAnimalSound(a Animal) {
	fmt.Println("Система просит животное издать звук:")
	a.Speak()
	fmt.Println("---")
}

// Предположим, что Dog и Cat - наши структуры,
// которые мы не хотим изменять
// У них есть свои методы Bark и Meow

type Dog struct{}

func (d *Dog) Bark() {
	fmt.Println("Гав!")
}

type Cat struct{}

func (c *Cat) Meow() {
	fmt.Println("Мяу!")
}

// Адаптеры, которые решают проблему несовместимости:
// Их методы реализуют требуемый метод Speak()

type DogAdapter struct {
	dog *Dog
}

func (da *DogAdapter) Speak() {
	// fmt.Println("(Адаптер переводит 'Speak' в 'Bark')")
	da.dog.Bark()
}

type CatAdapter struct {
	cat *Cat
}

func (ca *CatAdapter) Speak() {
	// fmt.Println("(Адаптер переводит 'Speak' в 'Meow')")
	ca.cat.Meow()
}

func main() {
	dog := &Dog{}
	cat := &Cat{}
	adaptedDog := &DogAdapter{dog: dog}
	adaptedCat := &CatAdapter{cat: cat}

	// fmt.Println("Работаем с адаптированной собакой:")
	MakeAnimalSound(adaptedDog)

	// fmt.Println("Работаем с адаптированной кошкой:")
	MakeAnimalSound(adaptedCat)
}

/*
Применимость:
Когда нужно заставить работать вместе объекты с несовместимыми интерфейсами.

Плюсы:
1. Позволяет повторно использовать существующие классы.
2. Изолирует и скрывает детали преобразования интерфейсов от основного кода.

Минусы:
1. Увеличивает общее количество кода в проекте за счет создания классов-адаптеров.

Реальные примеры:
1. Работа с разными СУБД (PostgreSQL, MySQL) через единый интерфейс в приложении.
2. Чтение данных из разных форматов (XML, JSON, CSV) и представление их в едином виде.
*/
