package main

import (
	"fmt"
	"strconv"
)

type Item struct {
	Name   string
	Cost   int
	Damage int
	Armor  int
}

type Day21Player struct {
	HP    int
	ATK   int
	ARM   int
	Items []Item
}

func (p *Day21Player) RealATK() (out int) {
	out += p.ATK
	for _, i := range p.Items {
		out += i.Damage
	}
	return out
}

func (p *Day21Player) RealARM() (out int) {
	out += p.ARM
	for _, i := range p.Items {
		out += i.Armor
	}
	return out
}

type Day21Boss struct {
	HP  int
	ATK int
	ARM int
}

func Day21Fight(p Day21Player, b Day21Boss) bool {
	//fmt.Println(p)
	//fmt.Printf("player's real atk %d and real def %d, boss atk %d and def %d\n", p.RealATK(), p.RealARM(), b.ATK, b.ARM)
	for p.HP > 0 && b.HP > 0 {
		myDmg := p.RealATK() - b.ARM
		bossDmg := b.ATK - p.RealARM()
		if myDmg < 1 {
			myDmg = 1
		}
		if bossDmg < 1 {
			bossDmg = 1
		}

		b.HP = b.HP - myDmg
		if b.HP > 0 {
			p.HP = p.HP - bossDmg
		}
		//fmt.Printf("Player: %d HP Boss: %d HP\n", p.HP, b.HP)
	}
	return p.HP > 0
}

func day21sideA(lines []string) string {
	player := Day21Player{HP: 100, ATK: 0, ARM: 0}
	boss := Day21Boss{HP: 104, ATK: 8, ARM: 1}

	weaponOptions := []Item{}
	weaponOptions = append(weaponOptions, Item{Name: "Dagger", Cost: 8, Damage: 4, Armor: 0})
	weaponOptions = append(weaponOptions, Item{Name: "Shortsword", Cost: 10, Damage: 5, Armor: 0})
	weaponOptions = append(weaponOptions, Item{Name: "Warhammer", Cost: 25, Damage: 6, Armor: 0})
	weaponOptions = append(weaponOptions, Item{Name: "Longsword", Cost: 40, Damage: 7, Armor: 0})
	weaponOptions = append(weaponOptions, Item{Name: "Greataxe", Cost: 74, Damage: 8, Armor: 0})
	armorOptions := []Item{}
	armorOptions = append(armorOptions, Item{Name: "Bare Ass Naked", Cost: 0, Damage: 0, Armor: 0})
	armorOptions = append(armorOptions, Item{Name: "Leather", Cost: 13, Damage: 0, Armor: 1})
	armorOptions = append(armorOptions, Item{Name: "Chainmail", Cost: 31, Damage: 0, Armor: 2})
	armorOptions = append(armorOptions, Item{Name: "Splintmail", Cost: 53, Damage: 0, Armor: 3})
	armorOptions = append(armorOptions, Item{Name: "Bandedmail", Cost: 75, Damage: 0, Armor: 4})
	armorOptions = append(armorOptions, Item{Name: "Platemail", Cost: 102, Damage: 0, Armor: 5})
	ringOptions := []Item{}
	ringOptions = append(ringOptions, Item{Name: "No+No", Cost: 0, Damage: 0, Armor: 0})
	ringOptions = append(ringOptions, Item{Name: "No+A1", Cost: 25, Damage: 1, Armor: 0})
	ringOptions = append(ringOptions, Item{Name: "No+A2", Cost: 50, Damage: 2, Armor: 0})
	ringOptions = append(ringOptions, Item{Name: "No+A3", Cost: 100, Damage: 3, Armor: 0})
	ringOptions = append(ringOptions, Item{Name: "No+D1", Cost: 20, Damage: 0, Armor: 1})
	ringOptions = append(ringOptions, Item{Name: "No+D2", Cost: 40, Damage: 0, Armor: 2})
	ringOptions = append(ringOptions, Item{Name: "No+D3", Cost: 80, Damage: 0, Armor: 3})
	ringOptions = append(ringOptions, Item{Name: "A1+A2", Cost: 75, Damage: 3, Armor: 0})
	ringOptions = append(ringOptions, Item{Name: "A1+A3", Cost: 125, Damage: 4, Armor: 0})
	ringOptions = append(ringOptions, Item{Name: "A2+A3", Cost: 150, Damage: 5, Armor: 0})
	ringOptions = append(ringOptions, Item{Name: "D1+D2", Cost: 60, Damage: 0, Armor: 3})
	ringOptions = append(ringOptions, Item{Name: "D1+D3", Cost: 100, Damage: 0, Armor: 4})
	ringOptions = append(ringOptions, Item{Name: "D2+D3", Cost: 120, Damage: 0, Armor: 5})
	ringOptions = append(ringOptions, Item{Name: "A1+D1", Cost: 45, Damage: 1, Armor: 1})
	ringOptions = append(ringOptions, Item{Name: "A1+D2", Cost: 65, Damage: 1, Armor: 2})
	ringOptions = append(ringOptions, Item{Name: "A1+D3", Cost: 105, Damage: 1, Armor: 3})
	ringOptions = append(ringOptions, Item{Name: "A2+D1", Cost: 70, Damage: 2, Armor: 1})
	ringOptions = append(ringOptions, Item{Name: "A2+D2", Cost: 90, Damage: 2, Armor: 2})
	ringOptions = append(ringOptions, Item{Name: "A2+D3", Cost: 130, Damage: 2, Armor: 3})
	ringOptions = append(ringOptions, Item{Name: "A3+D1", Cost: 120, Damage: 3, Armor: 1})
	ringOptions = append(ringOptions, Item{Name: "A3+D2", Cost: 140, Damage: 3, Armor: 2})
	ringOptions = append(ringOptions, Item{Name: "A3+D3", Cost: 180, Damage: 3, Armor: 3})

	best := 9999999
	var bestItems []Item

	// 91 is too high - there's a better option my code is missing somehow

	for _, ring := range ringOptions {
		for _, weapon := range weaponOptions {
			for _, armor := range armorOptions {
				player.Items = []Item{}
				player.Items = append(player.Items, ring)
				player.Items = append(player.Items, weapon)
				player.Items = append(player.Items, armor)

				if Day21Fight(player, boss) {
					cost := ring.Cost + weapon.Cost + armor.Cost
					if cost < best {
						best = cost
						bestItems = player.Items
					}
				}
			}
		}
	}

	fmt.Println(bestItems)

	return strconv.Itoa(best)
}

func day21sideB(lines []string) string {
	player := Day21Player{HP: 100, ATK: 0, ARM: 0}
	boss := Day21Boss{HP: 104, ATK: 8, ARM: 1}

	weaponOptions := []Item{}
	weaponOptions = append(weaponOptions, Item{Name: "Dagger", Cost: 8, Damage: 4, Armor: 0})
	weaponOptions = append(weaponOptions, Item{Name: "Shortsword", Cost: 10, Damage: 5, Armor: 0})
	weaponOptions = append(weaponOptions, Item{Name: "Warhammer", Cost: 25, Damage: 6, Armor: 0})
	weaponOptions = append(weaponOptions, Item{Name: "Longsword", Cost: 40, Damage: 7, Armor: 0})
	weaponOptions = append(weaponOptions, Item{Name: "Greataxe", Cost: 74, Damage: 8, Armor: 0})
	armorOptions := []Item{}
	armorOptions = append(armorOptions, Item{Name: "Bare Ass Naked", Cost: 0, Damage: 0, Armor: 0})
	armorOptions = append(armorOptions, Item{Name: "Leather", Cost: 13, Damage: 0, Armor: 1})
	armorOptions = append(armorOptions, Item{Name: "Chainmail", Cost: 31, Damage: 0, Armor: 2})
	armorOptions = append(armorOptions, Item{Name: "Splintmail", Cost: 53, Damage: 0, Armor: 3})
	armorOptions = append(armorOptions, Item{Name: "Bandedmail", Cost: 75, Damage: 0, Armor: 4})
	armorOptions = append(armorOptions, Item{Name: "Platemail", Cost: 102, Damage: 0, Armor: 5})
	ringOptions := []Item{}
	ringOptions = append(ringOptions, Item{Name: "No+No", Cost: 0, Damage: 0, Armor: 0})
	ringOptions = append(ringOptions, Item{Name: "No+A1", Cost: 25, Damage: 1, Armor: 0})
	ringOptions = append(ringOptions, Item{Name: "No+A2", Cost: 50, Damage: 2, Armor: 0})
	ringOptions = append(ringOptions, Item{Name: "No+A3", Cost: 100, Damage: 3, Armor: 0})
	ringOptions = append(ringOptions, Item{Name: "No+D1", Cost: 20, Damage: 0, Armor: 1})
	ringOptions = append(ringOptions, Item{Name: "No+D2", Cost: 40, Damage: 0, Armor: 2})
	ringOptions = append(ringOptions, Item{Name: "No+D3", Cost: 80, Damage: 0, Armor: 3})
	ringOptions = append(ringOptions, Item{Name: "A1+A2", Cost: 75, Damage: 3, Armor: 0})
	ringOptions = append(ringOptions, Item{Name: "A1+A3", Cost: 125, Damage: 4, Armor: 0})
	ringOptions = append(ringOptions, Item{Name: "A2+A3", Cost: 150, Damage: 5, Armor: 0})
	ringOptions = append(ringOptions, Item{Name: "D1+D2", Cost: 60, Damage: 0, Armor: 3})
	ringOptions = append(ringOptions, Item{Name: "D1+D3", Cost: 100, Damage: 0, Armor: 4})
	ringOptions = append(ringOptions, Item{Name: "D2+D3", Cost: 120, Damage: 0, Armor: 5})
	ringOptions = append(ringOptions, Item{Name: "A1+D1", Cost: 45, Damage: 1, Armor: 1})
	ringOptions = append(ringOptions, Item{Name: "A1+D2", Cost: 65, Damage: 1, Armor: 2})
	ringOptions = append(ringOptions, Item{Name: "A1+D3", Cost: 105, Damage: 1, Armor: 3})
	ringOptions = append(ringOptions, Item{Name: "A2+D1", Cost: 70, Damage: 2, Armor: 1})
	ringOptions = append(ringOptions, Item{Name: "A2+D2", Cost: 90, Damage: 2, Armor: 2})
	ringOptions = append(ringOptions, Item{Name: "A2+D3", Cost: 130, Damage: 2, Armor: 3})
	ringOptions = append(ringOptions, Item{Name: "A3+D1", Cost: 120, Damage: 3, Armor: 1})
	ringOptions = append(ringOptions, Item{Name: "A3+D2", Cost: 140, Damage: 3, Armor: 2})
	ringOptions = append(ringOptions, Item{Name: "A3+D3", Cost: 180, Damage: 3, Armor: 3})

	worst := 0
	var bestItems []Item

	// 91 is too high - there's a better option my code is missing somehow

	for _, ring := range ringOptions {
		for _, weapon := range weaponOptions {
			for _, armor := range armorOptions {
				player.Items = []Item{}
				player.Items = append(player.Items, ring)
				player.Items = append(player.Items, weapon)
				player.Items = append(player.Items, armor)

				if !Day21Fight(player, boss) {
					cost := ring.Cost + weapon.Cost + armor.Cost
					if cost > worst {
						worst = cost
						bestItems = player.Items
					}
				}
			}
		}
	}

	fmt.Println(bestItems)

	return strconv.Itoa(worst)
}
