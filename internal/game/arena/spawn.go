package arena

import (
	"online_shooter/internal/game/geometry"
)

// generateSpawns generates unique spawn points for each player on the edges of the arena.
func (a *Arena) generateSpawns() {
	totalSides := 4

	// a spawns amount is equal to the players amount
	spawnsAmount := a.SquaresAmount

	// create spawn points in the corners
	spawnsAmount -= totalSides
	a.Spawns = append(a.Spawns, geometry.Point{
		X: 0,
		Y: 0,
	})
	a.Spawns = append(a.Spawns, geometry.Point{
		X: a.Width,
		Y: 0,
	})
	a.Spawns = append(a.Spawns, geometry.Point{
		X: 0,
		Y: a.Height,
	})
	a.Spawns = append(a.Spawns, geometry.Point{
		X: a.Width,
		Y: a.Height,
	})

	// count spawn points per side
	spawnsPerSide := (spawnsAmount + totalSides - 1) / totalSides

	// generate spawns on the top side
	for i := 1; i < spawnsPerSide+1 && i <= spawnsAmount; i++ {
		spawn := geometry.Point{
			X: float32(i) * (a.Width / float32(spawnsPerSide+1)),
			Y: 0,
		}
		a.Spawns = append(a.Spawns, spawn)
	}

	// generate spawns on the right side
	for i := 1; i < spawnsPerSide+1 && (i+spawnsPerSide) <= spawnsAmount; i++ {
		spawn := geometry.Point{
			X: a.Width,
			Y: float32(i) * (a.Height / float32(spawnsPerSide+1)),
		}
		a.Spawns = append(a.Spawns, spawn)
	}

	// generate spawns on the bottom side
	for i := 1; i < spawnsPerSide+1 && (i+2*spawnsPerSide) <= spawnsAmount; i++ {
		spawn := geometry.Point{
			X: float32(i) * (a.Width / float32(spawnsPerSide+1)),
			Y: a.Height,
		}
		a.Spawns = append(a.Spawns, spawn)
	}

	// generate spawns on the left side
	for i := 1; i < spawnsPerSide+1 && (i+3*spawnsPerSide) <= spawnsAmount; i++ {
		spawn := geometry.Point{
			X: 0,
			Y: float32(i) * (a.Height / float32(spawnsPerSide+1)),
		}
		a.Spawns = append(a.Spawns, spawn)
	}
}
