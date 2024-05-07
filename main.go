package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Item struct {
	Name string
}

type Location struct {
	Name     string
	State    string
	Items    []Item
	Entities []Entity
}

type Entity interface {
	Describe() string
	Interact(player *Player) string
}

type Player struct {
	Location  string
	Inventory []Item
}

type GameState struct {
	Locations map[string]Location
	Player    Player
}

// HandleCommand processes player commands
func (gameState *GameState) handleCommand(command string, args []string) {
	switch command {
	case "осмотреться":
		gameState.DescribeLocation()
	case "идти":
		if len(args) < 1 {
			fmt.Println("Укажите, куда идти.")
			return
		}
		gameState.ChangeLocation(args[0])
	case "взять":
		if len(args) < 1 {
			fmt.Println("Укажите, что взять.")
			return
		}
		gameState.TakeItem(args[0])
	default:
		fmt.Println("Неизвестная команда")
	}
}

// DescribeLocation displays description of current location
func (gameState *GameState) DescribeLocation() {
	location, exists := gameState.Locations[gameState.Player.Location]
	if !exists {
		fmt.Println("Локация не найдена.")
		return
	}
	fmt.Println("Вы находитесь в", location.Name)
	fmt.Println("Описание:", location.State)
	fmt.Println("Предметы в локации:")
	for _, item := range location.Items {
		fmt.Println("-", item.Name)
	}
}

// ChangeLocation moves player to a new location
func (gameState *GameState) ChangeLocation(newLocation string) {
	_, exists := gameState.Locations[newLocation]
	if !exists {
		fmt.Println("Локация не найдена.")
		return
	}
	gameState.Player.Location = newLocation
	fmt.Println("Перемещение в", newLocation)
}

// TakeItem allows player to pick up an item from the location
func (gameState *GameState) TakeItem(itemName string) {
	location, exists := gameState.Locations[gameState.Player.Location]
	if !exists {
		fmt.Println("Локация не найдена.")
		return
	}
	for i, item := range location.Items {
		if item.Name == itemName {
			gameState.Player.Inventory = append(gameState.Player.Inventory, item)
			location.Items = append(location.Items[:i], location.Items[i+1:]...)
			fmt.Println("Вы взяли", itemName)
			return
		}
	}
	fmt.Println("Предмет не найден.")
}

func main() {
	locations := map[string]Location{
		"кухня":   {Name: "кухня", State: "описание кухни", Items: []Item{{Name: "ключи"}}},
		"коридор": {Name: "коридор", State: "описание коридора", Items: []Item{}},
		"комната": {Name: "комната", State: "описание комнаты", Items: []Item{{Name: "рюкзак"}, {Name: "конспекты"}}},
		"улица":   {Name: "улица", State: "описание улицы", Items: []Item{}},
	}

	gameState := &GameState{
		Locations: locations,
		Player:    Player{Location: "кухня", Inventory: []Item{}},
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Доступные команды: осмотреться, идти <название локации>, взять <предмет>")

	for {
		fmt.Print("Введите команду: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSuffix(input, "\n")
		parts := strings.Split(input, " ")
		command := parts[0]
		args := parts[1:]

		gameState.handleCommand(command, args)
	}
}
