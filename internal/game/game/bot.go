package game

import (
	"github.com/chewxy/math32"
	"math/rand"
	"online_shooter/internal/game/entity"
	"online_shooter/internal/game/geometry"
)

const (
	botBlindZoneLevel = 8
)

// CountMovingVector counts a vector which the bot moves with.
//
// Accepts a pointer to the bot, a pointer to the enemy which is an aim
// and a distance to the enemy.
//
// Returns a bot's moving vector.
func (g *Game) CountMovingVector(bot *entity.Square, enemy *entity.Square, distance float32) geometry.Vector {
	bot.RLock()
	defer bot.RUnlock()

	// if there are no enemies, or they are too far away
	if enemy == nil || distance > bot.Size*botBlindZoneLevel {
		// move towards a random vector
		vector := &geometry.Vector{
			X: rand.Float32() - 0.5*4 + g.Arena.Width/2 - bot.Position.X,
			Y: rand.Float32() - 0.5*4 + g.Arena.Height/2 - bot.Position.Y,
		}

		// normalize vector
		vector.Normalize()

		return *vector
	}

	enemy.RLock()
	defer enemy.RUnlock()

	// count the vector to the enemy
	vector := &geometry.Vector{
		X: enemy.Position.X - bot.Position.X,
		Y: enemy.Position.Y - bot.Position.Y,
	}

	// normalize vector
	vector.Normalize()

	// if the enemy's health is more than the bot's health
	if bot.Health < enemy.Health {
		// bot moves in the opposite direction from the enemy
		vector.X *= -1
		vector.Y *= -1
	}

	// navigate bot in the corners of the Arena
	g.cornerNavigation(bot, vector)

	// otherwise bot move in the same direction of the enemy
	return *vector
}

// FindEnemy finds the nearest enemy.
//
// Accepts a pointer to the bot that finds enemies.
//
// Returns a pointer to the enemy square and the distance to it.
func (g *Game) FindEnemy(bot *entity.Square) (*entity.Square, float32) {
	bot.RLock()
	defer bot.RUnlock()

	// create variables
	minDistance := math32.Hypot(g.Arena.Width, g.Arena.Height)
	var nearestEnemy *entity.Square

	for _, enemy := range g.Squares {
		// check if a potential enemy is not the finder itself
		if enemy == bot {
			continue
		}

		enemy.RLock()

		// get the distance
		distance := geometry.GetDistanceBetweenTwoPoints(bot.Position, enemy.Position)

		// if the distance is less than the min distance
		if distance < minDistance {
			// update variables
			minDistance = distance
			nearestEnemy = enemy
		}

		enemy.RUnlock()
	}

	return nearestEnemy, minDistance
}

// cornerNavigation navigates the bot in the corners of the Arena.
//
// Accepts a pointer to the bot and a vector of bot's moving direction.
func (g *Game) cornerNavigation(bot *entity.Square, vector *geometry.Vector) {
	if bot.Position.X > g.Arena.Width*90 && vector.X > 0 {
		if bot.Position.Y > g.Arena.Height*90 && vector.Y > 0 {
			vector.Y *= -1
		}
		if bot.Position.Y < g.Arena.Height*10 && vector.Y < 0 {
			vector.Y *= -1
		}
		vector.X *= -1
	}

	if bot.Position.X < g.Arena.Width*10 && vector.X < 0 {
		if bot.Position.Y > g.Arena.Height*90 && vector.Y > 0 {
			vector.Y *= -1
		}
		if bot.Position.Y < g.Arena.Height*10 && vector.Y < 0 {
			vector.Y *= -1
		}
		vector.X *= -1
	}

}

// CountShootingPoint counts a point which needs to be shot.
//
// Accepts a pointer to an enemy square and a distance to it.
//
// Returns a pointer to the shot point.
func (g *Game) CountShootingPoint(enemy *entity.Square, distance float32) *geometry.Point {
	// check if the enemy exists
	if enemy == nil {
		return nil
	}

	enemy.RLock()
	defer enemy.RUnlock()

	// check the bot's blind zone for the shooting
	if distance > enemy.Size*botBlindZoneLevel {
		return nil
	}

	// count the center of the aim
	enemyCenter := &geometry.Point{
		X: enemy.Position.X + enemy.Size/2,
		Y: enemy.Position.Y + enemy.Size/2,
	}

	return enemyCenter
}

// FindWeakestBot finds the weakest bot in the game.
//
// Returns a pointer to the weakest bot if it exists
// otherwise returns nil.
func (g *Game) FindWeakestBot() *entity.Square {
	var minRecord int16 = 32767
	var weakestBot *entity.Square

	for _, s := range g.Squares {
		if s.IsBot {
			botRecord := int16(s.Kills - s.Deaths)
			if minRecord > botRecord {
				minRecord = botRecord
				weakestBot = s
			}
		}
	}

	return weakestBot
}
