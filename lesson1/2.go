package main

import (
	"fmt"
	"time"
)

func main() {

	var date, surname, name, patronymic string
	var fm, sm, tm float64

	fmt.Scanln(&date)
	fmt.Scanln(&name)
	fmt.Scanln(&surname)
	fmt.Scanln(&patronymic)
	fmt.Scanln(&fm)
	fmt.Scanln(&sm)
	fmt.Scanln(&tm)

	startDate, _ := time.Parse("02.01.2006", date)
	signDate := startDate.AddDate(0, 0, 15)
	signDateStr := signDate.Format("02.01.2006")
	total := fm + sm + tm
	rubles := int(total)
	kopecks := int((total - float64(rubles)) * 100)
  message := fmt.Sprintf(`Уважаемый, %s %s %s, доводим до вашего сведения, что бухгалтерия сформировала документы по факту выполненной вами работы. 
Дата подписания договора: %s. Просим вас подойти в офис в любое удобное для вас время в этот день. 
Общая сумма выплат составит %d руб. %d коп.

С уважением,
Гл. бух. Иванов А.Е.`, surname, name, patronymic, signDateStr, rubles, kopecks)
    
    fmt.Println(message)
}
