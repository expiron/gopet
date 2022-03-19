package keeper

import "github.com/expiron/gopet/pkg/pet"

func FeedAPet(name string) {
	pet := pet.NewPet(name)
	pet.Breathe()
	pet.Eat("hotdog")
}
