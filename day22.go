package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Player struct {
	HP     int
	Mana   int
	Spells []Spell
	ARM    int
}

func (p *Player) ValidSpells(fx map[string]*Effect) (validSpells []Spell, bail bool) {
	validSpells = []Spell{}
	for _, spell := range p.Spells {
		if spell.Cost < p.Mana && !fx[spell.Effect].Active {
			validSpells = append(validSpells, spell)
		}
	}
	return validSpells, len(validSpells) == 0
}

type Boss struct {
	HP  int
	ATK int
}

type Spell struct {
	Name   string
	Cost   int
	Damage int
	Heal   int
	Effect string
}

type Effect struct {
	Duration  int
	Remaining int
	Active    bool
	Proc      func(p *Player, b *Boss)
}

func NewEffect(duration int, proc func(p *Player, b *Boss)) *Effect {
	ef := Effect{duration, 0, false, proc}
	return &ef
}

func (e *Effect) String() string {
	if e.Active {
		return fmt.Sprintf("ACTIVE at %d/%d", e.Remaining, e.Duration)
	} else {
		return fmt.Sprintf("INACTIVE")
	}
}

func (e *Effect) Apply(p *Player, b *Boss) {
	e.Proc(p, b)
	e.Remaining -= 1
	if e.Remaining == 0 {
		e.Active = false
	}
}

func ApplyEffects(p *Player, b *Boss, fx map[string]*Effect) {
	for _, e := range fx {
		if e.Active {
			e.Apply(p, b)
		}
	}
}

func ApplySpell(s Spell, p *Player, b *Boss, fx map[string]*Effect) int {
	p.Mana -= s.Cost
	p.HP += s.Heal
	b.HP -= s.Damage
	if s.Effect != "None" {
		if fx[s.Effect].Active == true {
			panic("trying to activate an already-active status; something has gone awry")
		}
		fx[s.Effect].Active = true
		fx[s.Effect].Remaining = fx[s.Effect].Duration
	}
	return s.Cost
}

func Round(p *Player, b *Boss, fx map[string]*Effect) (continueFight bool, bail bool, manaSpent int) {
	// player's turn
	p.HP -= 1
	if p.HP == 0 {
		return false, false, manaSpent
	}

	ApplyEffects(p, b, fx)

	validSpells, bail := p.ValidSpells(fx)
	if bail {
		return false, true, manaSpent
	}

	choice := validSpells[rand.Intn(len(validSpells))]

	manaSpent = ApplySpell(choice, p, b, fx)

	p.ARM = 0

	// boss's turn
	ApplyEffects(p, b, fx)
	if b.HP > 0 {
		dmg := b.ATK - p.ARM
		if dmg < 1 {
			dmg = 1
		}
		p.HP -= dmg
	}

	p.ARM = 0

	if b.HP <= 0 || p.HP <= 0 {
		return false, false, manaSpent
	} else {
		return true, false, manaSpent
	}
}

func Fight(p Player, b Boss) (playerWon bool, manaCost int) {
	fx := map[string]*Effect{
		"None":     NewEffect(0, func(p *Player, b *Boss) {}),
		"Poison":   NewEffect(6, func(p *Player, b *Boss) { b.HP = b.HP - 3 }),
		"Shield":   NewEffect(6, func(p *Player, b *Boss) { p.ARM = 7 }),
		"Recharge": NewEffect(5, func(p *Player, b *Boss) { p.Mana = p.Mana + 101 }),
	}

	continueBattle := true
	manaSpentThatTurn := 0
	bail := false

	for continueBattle && !bail {
		continueBattle, bail, manaSpentThatTurn = Round(&p, &b, fx)
		manaCost = manaCost + manaSpentThatTurn
	}
	return p.HP > 0 && !bail, manaCost
}

func day22sideA(lines []string) string {
	spells := []Spell{
		Spell{"Magic Missile", 53, 4, 0, "None"},
		Spell{"Drain", 73, 2, 2, "None"},
		Spell{"Shield", 113, 0, 0, "Shield"},
		Spell{"Poison", 173, 0, 0, "Poison"},
		Spell{"Recharge", 229, 0, 0, "Recharge"},
	}
	seed := time.Now().UnixNano()
	rand.Seed(seed)
	player := Player{HP: 50, ARM: 0, Mana: 500, Spells: spells}
	boss := Boss{HP: 51, ATK: 9}
	//boss := Boss{HP: 58, ATK: 9} // jp
	bestCost := 999999999
	for {
		won, cost := Fight(player, boss)
		if won && (cost < bestCost) {
			bestCost = cost
			fmt.Println(cost)
		}
	}
	return "n/i"
}

func day22sideB(lines []string) string {
	return "n/i"
}
