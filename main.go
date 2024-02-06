package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func attack(playerName, gamePath string) string {
	if gamePath == "warrior" {
		return fmt.Sprintf("%s нанес урон противнику равный %d.", playerName, 5+randint(3, 5))
	}

	if gamePath == "mage" {
		return fmt.Sprintf("%s нанес урон противнику равный %d.", playerName, 5+randint(5, 10))
	}

	if gamePath == "healer" {
		return fmt.Sprintf("%s нанес урон противнику равный %d.", playerName, 5+randint(-3, -1))
	}
	return "неизвестный класс персонажа"
}

// обратите внимание на "if else" и на "else"
func defence(playerName, gamePath string) string {
	if gamePath == "warrior" {
		return fmt.Sprintf("%s блокировал %d урона.", playerName, 10+randint(5, 10))
	}
	if gamePath == "mage" {
		return fmt.Sprintf("%s блокировал %d урона.", playerName, 10+randint(-2, 2))
	}
	if gamePath == "healer" {
		return fmt.Sprintf("%s блокировал %d урона.", playerName, 10+randint(2, 5))
	} 
	return "неизвестный класс персонажа"
}

// обратите внимание на "if else" и на "else"
func special(playerName, gamePath string) string {
	if gamePath == "warrior" {
		return fmt.Sprintf("%s применил специальное умение `Выносливость %d`", playerName, 80+25)
	} 
	if gamePath == "mage" {
		return fmt.Sprintf("%s применил специальное умение `Атака %d`", playerName, 5+40)
	} 
	if gamePath == "healer" {
		return fmt.Sprintf("%s применил специальное умение `Защита %d`", playerName, 10+30)
	} 
	return "неизвестный класс персонажа"
}

// здесь обратите внимание на имена параметров
func startTraining(playerName, gamePath string) string {
	if gamePath == "warrior" {
		fmt.Printf("%s, ты Воитель - отличный боец ближнего боя.\n", playerName)
	}

	if gamePath == "mage" {
		fmt.Printf("%s, ты Маг - превосходный укротитель стихий.\n", playerName)
	}

	if gamePath == "healer" {
		fmt.Printf("%s, ты Лекарь - чародей, способный исцелять раны.\n", playerName)
	}

	fmt.Println("Потренируйся управлять своими навыками.")
	fmt.Println("Введи одну из команд: attack — чтобы атаковать противника,")
	fmt.Println("defence — чтобы блокировать атаку противника,")
	fmt.Println("special — чтобы использовать свою суперсилу.")
	fmt.Println("Если не хочешь тренироваться, введи команду skip.")

	var cmd string
	for cmd != "skip" {
		fmt.Print("Введи команду: ")
		fmt.Scanf("%s\n", &cmd)

		if cmd == "attack" {
			fmt.Println(attack(playerName, gamePath))
		}

		if cmd == "defence" {
			fmt.Println(defence(playerName, gamePath))
		}

		if cmd == "special" {
			fmt.Println(special(playerName, gamePath))
		}
	}

	return "тренировка окончена"
}

// обратите внимание на имя функции и имена переменных
func choiceGamePath() string {
	var approveChoice string
	var gamePath string

	for approveChoice != "y" {
		fmt.Print("Введи название персонажа, за которого хочешь играть: Воитель — warrior, Маг — mage, Лекарь — healer: ")
		fmt.Scanf("%s\n", &gamePath)
		if gamePath == "warrior" {
			fmt.Println("Воитель — дерзкий воин ближнего боя. Сильный, выносливый и отважный.")
		} else if gamePath == "mage" {
			fmt.Println("Маг — находчивый воин дальнего боя. Обладает высоким интеллектом.")
		} else if gamePath == "healer" {
			fmt.Println("Лекарь — могущественный заклинатель. Черпает силы из природы, веры и духов.")
		}
		fmt.Print("Нажми (Y), чтобы подтвердить выбор, или любую другую кнопку, чтобы выбрать другого персонажа: ")
		fmt.Scanf("%s\n", &approveChoice)
		approveChoice = strings.ToLower(approveChoice)
	}
	return gamePath
}

// обратите внимание на имена переменных
func main() {
	fmt.Println("Приветствую тебя, искатель приключений!")
	fmt.Println("Прежде чем начать игру...")

	var playerName string
	fmt.Print("...назови себя: ")
	fmt.Scanf("%s\n", &playerName)

	fmt.Printf("Здравствуй, %s\n", playerName)
	fmt.Println("Сейчас твоя выносливость — 80, атака — 5 и защита — 10.")
	fmt.Println("Ты можешь выбрать один из трёх путей силы:")
	fmt.Println("Воитель, Маг, Лекарь")

	gamePath := choiceGamePath()

	fmt.Println(startTraining(playerName, gamePath))
}

func randint(min, max int) int {
	return rand.Intn(max-min) + min
}
