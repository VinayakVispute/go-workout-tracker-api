package main

import "fmt"

type Item struct {
	Name string
	Type string
}

type Player struct {
	Name      string
	Inventory []Item
}

func main() {
	player := Player{Name: "Vinayak"}

	// Example items
	shield := Item{Name: "Shield", Type: "Armor"}
	potion := Item{Name: "Potion", Type: "HealthItem"}
	bow := Item{Name: "Bow", Type: "Weapon"}

	// Picking up items
	player.pickUpItem(potion)
	player.pickUpItem(potion)
	player.pickUpItem(bow)
	player.pickUpItem(shield)
	player.pickUpItem(potion)

	fmt.Println("Player Inventory after pickups:", player.Inventory)

	// Dropping an item
	player.dropItem("Bow")
	fmt.Println("Player Inventory after dropping Bow:", player.Inventory)

	// Using items
	player.useItem("Bow")    // Should show "Item not in inventory"
	player.useItem("Potion") // Should restore health
	player.useItem("Shield") // Should use the Shield
}

func (p *Player) pickUpItem(item Item) {
	p.Inventory = append(p.Inventory, item)
	fmt.Printf("Picked up: %s (%s)\n", item.Name, item.Type)
}

func (p *Player) dropItem(itemName string) {
	indexOfItem := -1

	for index, value := range p.Inventory {
		if value.Name == itemName {
			indexOfItem = index
			break
		}
	}

	if indexOfItem != -1 {
		p.Inventory = append(p.Inventory[:indexOfItem], p.Inventory[indexOfItem+1:]...)
		fmt.Printf("Dropped: %s\n", itemName)
	} else {
		fmt.Printf("Cannot drop %s: Not in inventory\n", itemName)
	}
}

func (p *Player) useItem(itemName string) {
	switch itemName {
	case "Bow":
		fmt.Println("Shot an arrow with the Bow!")
	case "Potion":
		fmt.Println("You drank a Potion. Health restored!")
	case "Shield":
		fmt.Println("You raised your Shield for defense!")
	default:
		fmt.Printf("%s is not in inventory.\n", itemName)
	}

	// Attempt to drop the used item
	p.dropItem(itemName)
}
