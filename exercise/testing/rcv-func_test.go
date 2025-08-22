//--Summary:
//  Copy your rcv-func solution to this directory and write unit tests.
//
//--Requirements:
//* Write unit tests that ensure:
//  - Health & energy can not go above their maximums
//  - Health & energy can not go below 0
//* If any of your  tests fail, make the necessary corrections
//  in the copy of your rcv-func solution file.
//
//--Notes:
//* Use `go test -v ./exercise/testing` to run these specific tests
//

package main

import "testing"

func TestIsDead(t *testing.T) {
	player := Player{
		name:      "Test",
		health:    100,
		maxHealth: 100,
		energy:    50,
		maxEnergy: 50,
		dead:      false,
	}

	if player.isDead() {
		t.Errorf("expected player to be alive")
	}

	player.dead = true

	if !player.isDead() {
		t.Errorf("expected player to be dead")
	}
}

func TestCheckDeathByHealth(t *testing.T) {
	player := Player{
		health: 0,
		energy: 10,
	}

	player.checkDeath()

	if !player.dead {
		t.Errorf("expected player to be dead when health is 0")
	}
}

func TestCheckDeathByEnergy(t *testing.T) {
	player := Player{
		health: 10,
		energy: 0,
	}

	player.checkDeath()

	if !player.dead {
		t.Errorf("expected player to be dead when energy is 0")
	}
}

func TestRemoveHealth(t *testing.T) {
	player := Player{
		name:      "Test",
		health:    100,
		maxHealth: 100,
		energy:    50,
		maxEnergy: 50,
	}

	player.removeHealth(25)

	if player.health != 75 {
		t.Errorf("expected health 75, got %d", player.health)
	}

	if player.dead {
		t.Errorf("expected player to still be alive")
	}
}

func TestRemoveHealthKillsPlayer(t *testing.T) {
	player := Player{
		name:      "Test",
		health:    20,
		maxHealth: 100,
		energy:    50,
		maxEnergy: 50,
	}

	player.removeHealth(25)

	if player.health != 0 {
		t.Errorf("expected health 0, got %d", player.health)
	}

	if !player.dead {
		t.Errorf("expected player to be dead")
	}
}

func TestAddHealth(t *testing.T) {
	player := Player{
		name:      "Test",
		health:    50,
		maxHealth: 100,
		energy:    50,
		maxEnergy: 50,
	}

	player.addHealth(20)

	if player.health != 70 {
		t.Errorf("expected health 70, got %d", player.health)
	}
}

func TestAddHealthCapsAtMax(t *testing.T) {
	player := Player{
		name:      "Test",
		health:    90,
		maxHealth: 100,
		energy:    50,
		maxEnergy: 50,
	}

	player.addHealth(50)

	if player.health != 100 {
		t.Errorf("expected health capped at 100, got %d", player.health)
	}
}

func TestRemoveEnergy(t *testing.T) {
	player := Player{
		name:      "Test",
		health:    100,
		maxHealth: 100,
		energy:    50,
		maxEnergy: 50,
	}

	player.removeEnergy(20)

	if player.energy != 30 {
		t.Errorf("expected energy 30, got %d", player.energy)
	}

	if player.dead {
		t.Errorf("expected player to still be alive")
	}
}

func TestRemoveEnergyKillsPlayer(t *testing.T) {
	player := Player{
		name:      "Test",
		health:    100,
		maxHealth: 100,
		energy:    10,
		maxEnergy: 50,
	}

	player.removeEnergy(20)

	if player.energy != 0 {
		t.Errorf("expected energy 0, got %d", player.energy)
	}

	if !player.dead {
		t.Errorf("expected player to be dead")
	}
}

func TestAddEnergy(t *testing.T) {
	player := Player{
		name:      "Test",
		health:    100,
		maxHealth: 100,
		energy:    20,
		maxEnergy: 50,
	}

	player.addEnergy(15)

	if player.energy != 35 {
		t.Errorf("expected energy 35, got %d", player.energy)
	}
}

func TestAddEnergyCapsAtMax(t *testing.T) {
	player := Player{
		name:      "Test",
		health:    100,
		maxHealth: 100,
		energy:    45,
		maxEnergy: 50,
	}

	player.addEnergy(20)

	if player.energy != 50 {
		t.Errorf("expected energy capped at 50, got %d", player.energy)
	}
}

func TestDeadPlayerCannotHeal(t *testing.T) {
	player := Player{
		name:      "Test",
		health:    0,
		maxHealth: 100,
		energy:    0,
		maxEnergy: 50,
		dead:      true,
	}

	player.addHealth(50)

	if player.health != 0 {
		t.Errorf("expected dead player health unchanged, got %d", player.health)
	}
}

func TestDeadPlayerCannotGainEnergy(t *testing.T) {
	player := Player{
		name:      "Test",
		health:    0,
		maxHealth: 100,
		energy:    0,
		maxEnergy: 50,
		dead:      true,
	}

	player.addEnergy(50)

	if player.energy != 0 {
		t.Errorf("expected dead player energy unchanged, got %d", player.energy)
	}
}

func TestRestore50(t *testing.T) {
	player := Player{
		name:      "Test",
		health:    0,
		maxHealth: 100,
		energy:    0,
		maxEnergy: 80,
		dead:      true,
	}

	player.restore50()

	if player.dead {
		t.Errorf("expected player to be revived")
	}

	if player.health != 50 {
		t.Errorf("expected health 50, got %d", player.health)
	}

	if player.energy != 40 {
		t.Errorf("expected energy 40, got %d", player.energy)
	}
}

func TestSecurityStrikePreventsUnderflow(t *testing.T) {
	player := Player{
		name:      "Test",
		health:    10,
		maxHealth: 100,
		energy:    10,
		maxEnergy: 50,
	}

	player.securityStrike()

	if player.health != 0 {
		t.Errorf("expected health 0 after security strike, got %d", player.health)
	}

	if player.energy != 0 {
		t.Errorf("expected energy 0 after security strike, got %d", player.energy)
	}

	if !player.dead {
		t.Errorf("expected player to be dead")
	}
}

func TestRandomReturnsValue(t *testing.T) {
	values := []string{"a", "b", "c"}

	result := random(values)

	found := false
	for _, v := range values {
		if result == v {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("random returned unexpected value: %s", result)
	}
}
