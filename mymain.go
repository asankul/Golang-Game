package main

import (
	"fmt"
)

//Возможности персонажей
type Knight struct {
	name   string
	Damage float64
	HP     float64
}
type Dragon struct {
	name   string
	Damage float64
	HP     float64
}

//Главное меню
var BEGIN [1][2]string
var COND [1][2]string

func conditions() {
	BEGIN[0][0] = "1: Начать игру"
	BEGIN[0][1] = "2: Выйти"
	COND[0][0] = "1: ДА!"
	COND[0][1] = "2: НЕТ!"

}

func begin(x int) {
	for _, b := range BEGIN[x] {
		fmt.Println(b)
	}
}
func begin2(x int) {
	for _, b := range COND[x] {
		fmt.Println(b)
	}
}

// Проверка валидности опции
func check(x, y, z int) int {
	if (x >= y) && (x <= z) {
		return x
	} else {
		fmt.Println("Нет такой опции, введите существующий выбор")
		begin(0)
		var x2 int
		fmt.Scan(&x2)
		return check(x2, y, z)

	}
}
func check2(x, y, z int) int {
	if (x >= y) && (x <= z) {
		return x
	} else {
		fmt.Println("Нет такой опции, введите существующий выбор")
		begin2(0)
		var x2 int
		fmt.Scan(&x2)
		return check2(x2, y, z)

	}
}

//Показать статистику героев
func show(k Knight, d Dragon) {
	fmt.Printf("Рыцарское Имя: %s  Хп: %v Атака: %v\n", k.name, k.HP, k.Damage)
	fmt.Printf("Драконье  Имя: %s  Хп: %v Атака: %v\n", d.name, d.HP, d.Damage)
}

// Лет зе гейм бегин
func game() {
	//Характеристики героя
	var dk = Knight{}
	dk.HP = 500
	dk.Damage = 50
	fmt.Println("Задайте имя рыцаря")
	fmt.Scan(&dk.name)
	//Характеристики дракона
	var dr = Dragon{}
	dr.HP = 500
	dr.Damage = 50
	fmt.Println("Задайте имя дракона")
	fmt.Scan(&dr.name)
	c := 0
	for {
		show(dk, dr)
		if dr.HP <= 0 && dk.HP <= 0 {
			fmt.Printf("Ничья, но вас оплакивают и отпевают\n")
			break
		} else if dk.HP <= 0 {
			fmt.Println("Вы проиграли и мир погряз в хаусе\n")
			break
		} else if dr.HP <= 0 {
			fmt.Println("Вы победили и спасли мир\n", c)
			break
		}
		c++
		fmt.Println("\n\nНапасть на дракона?")
		begin2(0)
		var x, input int
		fmt.Scan(&x)
		input = check2(x, 1, 2)
		switch input {
		case 1:
			{
				fmt.Println("\n\nКрасавчик!")
				dk.HP = dk.HP - dr.Damage
				fmt.Println("Но и дракон тебя атаковал")
				dr.HP = dr.HP - dk.Damage
			}
		case 2:
			{
				fmt.Println("\n\nНу ты добряк!")
				fmt.Println("А дракон тебя атаковал")
				dk.HP = dk.HP - dr.Damage
			}
		}
	}
}

// Игра
func main() {
	conditions()
	begin(0)
	for {
		var input, exit int
		fmt.Scan(&input)
		input = check(input, 1, 2)
		switch input {
		case 1:
			{
				fmt.Println("Начнем игру!")
				game()
				exit = 1
			}
		case 2:
			{
				fmt.Println("Прощай!")
				begin(0)
				exit = 1
			}
		}
		if exit == 1 {
			break
		}
	}

}
