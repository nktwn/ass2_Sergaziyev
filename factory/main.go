package main

import (
	"Sergaziyev_ass2/factory/gamestore"
)

func main() {
	store := gamestore.NewGameStore()

	actionFactory := &gamestore.ActionFactory{}
	adventureFactory := &gamestore.AdventureFactory{}
	rpgFactory := &gamestore.RPGFactory{}

	store.PurchaseGame(actionFactory, "World of Warcraft")
	store.PurchaseGame(adventureFactory, "Dota 2")
	store.PurchaseGame(rpgFactory, "Dota")

	store.ListGames()
}
