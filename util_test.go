package main

import "testing"

func TestBox_CheckCollision(t *testing.T) {
	//{244 346 8 2} //bullet
	//{251 384 64 64} //enemy
	bullet := &Box{
		X: 244,
		Y: 346,
		W: 8,
		H: 2,
	}

	enemy := &Box{
		X: 251,
		Y: 384,
		W: 64,
		H: 64,
	}

	dist := bullet.CheckCollision(AxisY, enemy)
	println(dist)
	dist = bullet.CheckCollision(AxisX, enemy)
	println(dist)

	println()

	dist = enemy.CheckCollision(AxisY, bullet)
	println(dist)
	dist = enemy.CheckCollision(AxisX, bullet)
	println(dist)
}
