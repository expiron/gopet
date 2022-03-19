package keeper

import "app/pkg/pet"

func FeedAPet(name string) {
	pet := pet.NewPet(name)
	pet.Breathe()
	pet.Eat("hotdog")
}
