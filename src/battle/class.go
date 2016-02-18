package class
import (
	"fmt"
	"os"
)
type moves struct {
	name1 string
	damage1 int
	accuracy2 int
	name2 string
	damage2 int
	accuracy2 int
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
	name2: "Freeze Ray"
	damage2: "50"
	
