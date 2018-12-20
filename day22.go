package main

import (
	"fmt"
	"math/rand"
	"time"
)

type player struct {
	HP     int
	Mana   int
	Spells []spell
	ARM    int
}

func (p *player) validSpells(fx map[string]*effect) (validSpells []spell, bail bool) {
	validSpells = []spell{}
	for _, spell := range p.Spells {
		if spell.Cost < p.Mana && !fx[spell.Effect].Active {
			validSpells = append(validSpells, spell)
		}
	}
	return validSpells, len(validSpells) == 0
}

type boss struct {
	HP  int
	ATK int
}

type spell struct {
	Name   string
	Cost   int
	Damage int
	Heal   int
	Effect string
}

type effect struct {
	Duration  int
	Remaining int
	Active    bool
	Proc      func(p *player, b *boss)
}

func newEffect(duration int, proc func(p *player, b *boss)) *effect {
	ef := effect{duration, 0, false, proc}
	return &ef
}

func (e *effect) string() string {
	if e.Active {
		return fmt.Sprintf("ACTIVE at %d/%d", e.Remaining, e.Duration)
	}
	return fmt.Sprintf("INACTIVE")
}

func (e *effect) apply(p *player, b *boss) {
	e.Proc(p, b)
	e.Remaining--
	if e.Remaining == 0 {
		e.Active = false
	}
}

func applyEffects(p *player, b *boss, fx map[string]*effect) {
	for _, e := range fx {
		if e.Active {
			e.apply(p, b)
		}
	}
}

func applySpell(s spell, p *player, b *boss, fx map[string]*effect) int {
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

func round(p *player, b *boss, fx map[string]*effect) (continueFight bool, bail bool, manaSpent int) {
	// player's turn
	p.HP--
	if p.HP == 0 {
		return false, false, manaSpent
	}

	applyEffects(p, b, fx)

	validSpells, bail := p.validSpells(fx)
	if bail {
		return false, true, manaSpent
	}

	choice := validSpells[rand.Intn(len(validSpells))]

	manaSpent = applySpell(choice, p, b, fx)

	p.ARM = 0

	// boss's turn
	applyEffects(p, b, fx)
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
	}
	return true, false, manaSpent
}

func fight(p player, b boss) (playerWon bool, manaCost int) {
	fx := map[string]*effect{
		"None":     newEffect(0, func(p *player, b *boss) {}),
		"Poison":   newEffect(6, func(p *player, b *boss) { b.HP = b.HP - 3 }),
		"Shield":   newEffect(6, func(p *player, b *boss) { p.ARM = 7 }),
		"Recharge": newEffect(5, func(p *player, b *boss) { p.Mana = p.Mana + 101 }),
	}

	continueBattle := true
	manaSpentThatTurn := 0
	bail := false

	for continueBattle && !bail {
		continueBattle, bail, manaSpentThatTurn = round(&p, &b, fx)
		manaCost = manaCost + manaSpentThatTurn
	}
	return p.HP > 0 && !bail, manaCost
}

func day22sideA(lines []string) string {
	spells := []spell{
		spell{"Magic Missile", 53, 4, 0, "None"},
		spell{"Drain", 73, 2, 2, "None"},
		spell{"Shield", 113, 0, 0, "Shield"},
		spell{"Poison", 173, 0, 0, "Poison"},
		spell{"Recharge", 229, 0, 0, "Recharge"},
	}
	seed := time.Now().UnixNano()
	rand.Seed(seed)
	p := player{HP: 50, ARM: 0, Mana: 500, Spells: spells}
	b := boss{HP: 51, ATK: 9}
	bestCost := 999999999
	for {
		won, cost := fight(p, b)
		if won && (cost < bestCost) {
			bestCost = cost
			fmt.Println(cost)
		}
	}
}

func day22sideB(lines []string) string {
	return "n/i"
}
