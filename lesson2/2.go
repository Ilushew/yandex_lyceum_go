package main

import (
	"fmt"
)

func main() {
	var place1, place2, place3, place4, place5 string
	for {
		var operation string
		fmt.Scan(&operation)
		if operation == "конец"{
			if place1 == ""{
				fmt.Println("1. -")
			} else {
				fmt.Printf("1. %s \n", place1)
			}
			if place2 == ""{
				fmt.Println("2. -")
			} else {
				fmt.Printf("2. %s \n", place2)
			}
			if place3 == ""{
				fmt.Println("3. -")
			} else {
				fmt.Printf("3. %s \n", place3)
			}
			if place4 == ""{
				fmt.Println("4. -")
			} else {
				fmt.Printf("4. %s \n", place4)
			}
			if place5 == ""{
				fmt.Println("5. -")
			} else {
				fmt.Printf("5. %s \n", place5)
			}
			return
		}
		if operation == "очередь"{
			if place1 == ""{
				fmt.Println("1. -")
			} else {
				fmt.Printf("1. %s \n", place1)
			}
			if place2 == ""{
				fmt.Println("2. -")
			} else {
				fmt.Printf("2. %s \n", place2)
			}
			if place3 == ""{
				fmt.Println("3. -")
			} else {
				fmt.Printf("3. %s \n", place3)
			}
			if place4 == ""{
				fmt.Println("4. -")
			} else {
				fmt.Printf("4. %s \n", place4)
			}
			if place5 == ""{
				fmt.Println("5. -")
			} else {
				fmt.Printf("5. %s \n", place5)
			}
			continue
		}
		if operation == "количество"{
			var count = 0
			if place1 != ""{
				count++
			}
			if place2 != ""{
				count++
			}
			if place3 != ""{
				count++
			}
			if place4 != ""{
				count++
			}
			if place5 !=""{
				count++
			}
			fmt.Println("Осталось свободных мест:", 5 - count)
			fmt.Println("Всего человек в очереди:", count)
			continue
		}
		name := operation
		var num int
		fmt.Scan(&num)
		if num < 1 || num > 5{
			fmt.Printf("Запись на место номер %d невозможна: некорректный ввод \n", num)
			continue
		}
		var count = 0
		if place1 != ""{
			count++
		}
		if place2 != ""{
			count++
		}
		if place3 != ""{
			count++
		}
		if place4 != ""{
			count++
		}
		if place5 !=""{
			count++
		}
		if count == 5{
			fmt.Printf("Запись на место номер %d невозможна: очередь переполнена \n", num)
			continue
		}
		switch num{
		case 1:
			if place1 == ""{
				place1 = name
			} else {
				fmt.Printf("Запись на место номер %d невозможна: место уже занято \n", num)
			}
		case 2:
			if place2 == ""{
				place2 = name
			} else {
				fmt.Printf("Запись на место номер %d невозможна: место уже занято \n", num)
			}
		case 3:
			if place3 == ""{
				place3 = name
			} else {
				fmt.Printf("Запись на место номер %d невозможна: место уже занято \n", num)
			}
		case 4:
			if place4 == ""{
				place4 = name
			} else {
				fmt.Printf("Запись на место номер %d невозможна: место уже занято \n", num)
			}
		case 5:
			if place5 == ""{
				place5 = name
			} else {
				fmt.Printf("Запись на место номер %d невозможна: место уже занято \n", num)
			}
		}
	}
}