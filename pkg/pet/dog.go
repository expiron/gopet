package pet

import (
	log "github.com/sirupsen/logrus"
)

type Pet struct {
	Name string
}

func (p *Pet) Breathe() {
	log.Infof("the dog %s breathe", p.Name)
}

func (p *Pet) Eat(something string) {
	log.Infof("the dog %s eat %s", p.Name, something)
}

func NewPet(name string) Pet {
	return Pet{Name: name}
}
