package partyrobot

import "fmt"

func Welcome(name string) string {
	return fmt.Sprintf("Welcome to my party, %s!", name)
}

func HappyBirthday(name string, age int) string {
	return fmt.Sprintf("Happy birthday %s! You are now %d years old!", name, age)
}

func AssignTable(name string, tableNumber int, seatmate string, direction string, distance float64) string {
	formattedTableNumber := fmt.Sprintf("%03d", tableNumber)
	return fmt.Sprintf(
		"Welcome to my party, %s!\n"+
			"You have been assigned to table %s. Your table is %s, exactly %.1f meters from here.\n"+
			"You will be sitting next to %s.",
		name, formattedTableNumber, direction, distance, seatmate,
	)
}
