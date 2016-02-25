package class
import (
	"fmt"
	"os"
)
type moves struct {
	name1 string
	damage1 int
	accuracy1 int
	speed1 int
	name2 string
	damage2 int
	accuracy2 int
	speed2 int
}

type class struct {
	moves struct,
	health, agility, shield int
}

wizard := class{
	health: 300
	agility: 750
	shield: 200
	attacks: moves{
	name1: "Burning Fist"
	damage1: "100"
	accuracy: "50"
	speed: "50"
	name2: "Freeze Ray"
	damage2: "50"
	accuracy2: "80"
	speed2: "70"
	}
}
archer := class{
	health: 400
	agility: 700
	shield: 150
	attacks: moves{
	name1: "Sniper"
	damage1: "200"
	accuracy1: "15"
	speed1: "45"
	name2: "Point Blank"
	damage2: "75"
	accuracy2: "90"
	speed2: "80"
	}
}

warrior := class{
	health: 600
	agility: 200
	shield: 550
	attacks: moves{
	name1: "Cleave"
	damage1: "300"
	accuracy1: "30"
	speed1: "20"
	name2: "Thrust"
	damage2: "150"
	accuracy2: "90"
	speed2: "30"
	}
}

ranger := class{
	health: 500
	agility: 800
	shield: 150
	attacks: moves{
	name1: "Stab"
	damage1: "75"
	accuracy1: "90"
	speed1: "80"
	name2: "Club"
	damage2: "60"
	accuracy2: "70"
	speed2: "65"
	}
}
