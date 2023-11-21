package main

import "fmt"

// класс собака
type Dog struct{}

// реакция собаки
func (dog *Dog) WoofWoof() {
	fmt.Println("Гав-Гав. Хозяин, дай пожрать")
}

// класс кошка
type Cat struct{}

// реакция кошки, она немного посложнее и если ее не позвать - она молчит
func (cat *Cat) MeowMeow(isCall bool) {
	if isCall {
		fmt.Println("Где моя еда, раб? Ну так уж и быть... Мяу-мяу")
	}
}

// целевой интерфейс - Target
type AnimalAdapter interface {
	Reaction()
}

// адаптер для собаки
type DogAdapter struct {
	*Dog
}

// реакция собаки
func (adapter *DogAdapter) Reaction() {
	adapter.WoofWoof()
}

// конструктор адаптера для собаки
func NewDogAdapter(dog *Dog) AnimalAdapter {
	return &DogAdapter{dog}
}

// адаптер для кошки
type CatAdapter struct {
	*Cat
}

// реакция кошки
func (adapter *CatAdapter) Reaction() {
	// адаптер автоматически зовет кошку isCall = true
	adapter.MeowMeow(true)
}

// конструктор адаптера для кота
func NewCatAdapter(cat *Cat) AnimalAdapter {
	return &CatAdapter{cat}
}
func main() {
	fmt.Println("\nВы останавливаетесь перед дверью и вставляете в ухо адаптер с двумя чипами\n")
	myFamily := [2]AnimalAdapter{NewDogAdapter(&Dog{}), NewCatAdapter(&Cat{})}
	//
	fmt.Println("Открываете дверь и заходите домой\n")
	for _, member := range myFamily {
		member.Reaction()
	}
}
