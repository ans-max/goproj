package main

import (
	"flag"
	"fmt"
	"strings"
	"time"
)

func DayDate(day string) time.Time {
	today := time.Now()
	if day == "yesterday" {
		return today.AddDate(0, 0, -1)
	} else if day == "tomorrow" {
		return today.AddDate(0, 0, 1)
	} else {
		return today
	}
}

func main() {
	day := flag.String("day", "today", "Enter day as today, tomorrow, yesterday")
	sign := flag.String("sign", "aquarius", "Enter the Zodiac sign in Lower case")
	flag.Parse()
	stringSign := *sign
	stringDay := *day
	date := DayDate(stringDay)
	fmt.Println("\n", strings.Title(stringSign))
	fmt.Println("\n", date.Format("January 2, 2006"))
	ProphecyAstro := FromAstroCom(stringSign, stringDay)
	fmt.Println("\n#From Astrology.com for:", stringSign)
	fmt.Println("\n", ProphecyAstro)
	ProphecyBejan := FromGaneshSpeaks(stringSign, stringDay)
	fmt.Println("\n#From Ganeshaspeaks.com for:", stringSign)
	fmt.Println("\n", ProphecyBejan)
}
