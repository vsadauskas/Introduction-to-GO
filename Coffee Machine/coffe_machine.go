package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/kyokomi/emoji/v2"
)

// global variables
var money int = 550
var water int = 400
var milk int = 540
var beans int = 120
var cups int = 9

// new comment
var emojiCoffee string = emoji.Sprint(":coffee:")
var emojiWater string = emoji.Sprint(":droplet:")
var emojiBeans string = emoji.Sprint(":beans:")
var emojiCups string = emoji.Sprint(":pouring_liquid:")
var emojiMilk string = emoji.Sprint(":glass_of_milk:")

// function checks if we have enough stuff to make coffee
func check_enough(name string) {
	if name == "espresso" {
		// For one espresso, the coffee machine needs
		// 250 ml of water and
		// 16 g of coffee beans.
		// 1 cup
		// It costs $4.
		if (water >= 250) && (beans >= 16) && (cups >= 1) {
			color.Green("I have enough resources, making you a coffee! " + emojiCoffee)
			water -= 250
			beans -= 16
			cups--
			money += 4
		} else if water < 250 {
			color.Red("Sorry, not enough water! " + emojiWater)
		} else if beans < 6 {
			color.Red("Sorry, not enough beans! " + emojiBeans)
		} else if cups < 1 {
			color.Red("Sorry, not enough cups! " + emojiCups)
		}

	} else if name == "latte" {
		// For a latte, the coffee machine needs
		// 350 ml of water,
		// 75 ml of milk,
		// and 20 g of coffee beans.
		// 1 cup
		// It costs $7.
		if (water >= 350) && (milk >= 75) && (beans >= 20) && (cups >= 1) {
			color.Green("I have enough resources, making you a coffee! " + emojiCoffee)
			water -= 350
			milk -= 75
			beans -= 20
			cups--
			money += 7
		} else if water < 350 {
			color.Red("Sorry, not enough water! " + emojiWater)
		} else if milk < 75 {
			color.Red("Sorry, not enough milk! " + emojiMilk)
		} else if beans < 20 {
			color.Red("Sorry, not enough beans! " + emojiBeans)
		} else if cups < 1 {
			color.Red("Sorry, not enough cups! " + emojiCups)
		}

	} else if name == "cappuccino" {
		// And for a cappuccino, the coffee machine needs
		// 200 ml of water,
		// 100 ml of milk,
		// and 12 g of coffee beans.
		// 1 cup
		// It costs $6.
		if (water >= 200) && (milk >= 100) && (beans >= 12) && (cups >= 1) {
			color.Green("I have enough resources, making you a coffee! " + emojiCoffee)
			water -= 200
			milk -= 100
			beans -= 12
			cups--
			money += 6
		} else if water < 200 {
			color.Red("Sorry, not enough water! " + emojiWater)
		} else if milk < 100 {
			color.Red("Sorry, not enough milk! " + emojiMilk)
		} else if beans < 12 {
			color.Red("Sorry, not enough beans! " + emojiBeans)
		} else if cups < 1 {
			color.Red("Sorry, not enough cups! " + emojiCups)
		}
	}
}

func main() {
	var action string

	// main program loop
	for action != "exit" {

		color.Blue("Write action (buy, fill, take, remaining, exit):")

		fmt.Scan(&action)

		// *** action buy
		if action == "buy" {
			var kind string
			color.Blue("What do you want to buy? 1 - espresso, 2 - latte, 3 - cappuccino, back - to main menu:")
			fmt.Scan(&kind)
			if kind == "1" {
				check_enough("espresso")
			} else if kind == "2" {
				check_enough("latte")
			} else if kind == "3" {
				check_enough("cappuccino")
			} else if kind == "back" {

			} else {
				fmt.Println("")
			}
			// *** action fill
		} else if action == "fill" {
			var add_water, add_milk, add_beans, add_cups int
			color.Cyan("Write how many ml of water you want to add:")
			fmt.Scan(&add_water)
			color.Cyan("Write how many ml of milk you want to add:")
			fmt.Scan(&add_milk)
			color.Cyan("Write how many grams of coffee beans you want to add:")
			fmt.Scan(&add_beans)
			color.Cyan("Write how many disposable cups you want to add:")
			fmt.Scan(&add_cups)
			// mathematics
			water += add_water
			milk += add_milk
			beans += add_beans
			cups += add_cups
			// *** action take
		} else if action == "take" {
			fmt.Printf("I gave you $%v\n", money)
			money = 0
			// *** action remaining
		} else if action == "remaining" {
			fmt.Println("The coffee machine has:")
			fmt.Printf("%v ml of water\n", water)
			fmt.Printf("%v ml of milk\n", milk)
			fmt.Printf("%v g of coffee beans\n", beans)
			fmt.Printf("%v disposable cups\n", cups)
			fmt.Printf("%v of money\n", money)
			// *** action exit
		} else if action == "exit" {

		} else {
			fmt.Println("WTF?")
		}
	}
}
