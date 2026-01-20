package main

import (
	"strings"
	"time"
	"unicode/utf8"
	"errors"
)


func currentDayOfTheWeek() string{
	now := TimeNow()
	switch now.Weekday(){
	case time.Monday:
		return "Понедельник"
	case time.Tuesday:
		return "Вторник"
	case time.Wednesday:
		return "Среда"
	case time.Thursday:
		return "Четверг"
	case time.Friday:
		return "Пятница"
	case time.Saturday:
		return "Суббота"
	case time.Sunday:
		return "Воскресенье"
	default:
		return ""
	}
}

func dayOrNight() string{
	now := TimeNow()
	hour := now.Hour()
	if hour >= 10 && hour < 22{
		return "День"
	} else {
		return "Ночь"
	}
}

func nextFriday() int{
	now := TimeNow()
	today := now.Weekday()
	daysUntil := int(time.Friday - today)
	if daysUntil < 0{
		daysUntil += 7
	}
	return daysUntil
}

func CheckCurrentDayOfTheWeek(answer string) bool{
	now := currentDayOfTheWeek()
	return strings.EqualFold(answer, now)
}

func CheckNowDayOrNight(answer string) (bool, error){
	now := dayOrNight()
	if utf8.RuneCountInString(answer) != 4{
		return false, errors.New("исправь свой ответ, а лучше ложись поспать")
	}
	return strings.EqualFold(now, answer), nil
}