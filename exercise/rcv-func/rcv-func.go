//--Summary:
//  Implement receiver functions to create stat modifications
//  for a video game character.
//
//--Requirements:
//* Implement a player having the following statistics:
//  - Health, Max Health
//  - Energy, Max Energy
//  - Name
//* Implement receiver functions to modify the `Health` and `Energy`
//  statistics of the player.
//  - Print out the statistic change within each function
//  - Execute each function at least once

package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Player struct {
	name              string
	health, maxHealth uint8
	energy, maxEnergy uint8
	dead              bool
}

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

var damageReasons = []string{
	"accidentally deployed Claude to production on Friday evening 💀",
	"read 900 AI doom tweets in one sitting 📉",
	"survived another Claude outage ⚠️",
	"got blackmailed by Claude during safety testing 🤖",
	"opened leaked Anthropic repo by mistake 🧠",
	"explained GPUs to finance department 🔥",
	"argued with Pentagon about AI safety 🪖",
	"investors asked 'when IPO?' for 7 hours 📊",
	"AGI prediction: '6 months' (again) 📆",
	"joined alignment meeting that solved nothing 🧓",
}

var healReasons = []string{
	"received fresh NVIDIA GPUs from cloud gods 🖥️",
	"successfully fixed production outage 🚑",
	"drank Silicon Valley emergency coffee ☕",
	"survived Hacker News debate 🧓",
	"closed 47 GitHub tabs and touched grass 🌱",
	"DNS issue fixed itself mysteriously 🌐",
	"demo didn’t crash 🎉",
	"restarted everything and it worked 🔌",
}

var energyLossReasons = []string{
	"debugged Claude Code until 4 AM 🌙",
	"handled 500 errors all day 💥",
	"6-hour AI alignment Zoom call 😭",
	"Kubernetes self-destruct sequence ☸️",
	"doomscrolled AI Twitter for 9 hours 🧠",
	"explained transformers to executives 📊",
}

var energyGainReasons = []string{
	"servers restarted successfully 🔌",
	"deployment finally said SUCCESS 🎉",
	"more GPUs approved ⚡",
	"copied Stack Overflow code perfectly 🧠",
	"bug was just a missing comma 🤡",
	"disabled one service and fixed everything 🔥",
}

func wasted(text string) {
	fmt.Println("\n======================================")
	fmt.Println("            W A S T E D")
	fmt.Println("======================================")
	fmt.Println(text)
	fmt.Println("======================================")
}

func random(arr []string) string {
	return arr[r.Intn(len(arr))]
}

func (p *Player) isDead() bool {
	return p.dead
}

func (p *Player) checkDeath() {
	if p.health == 0 || p.energy == 0 {
		p.dead = true
	}
}

func (p *Player) removeHealth(h uint8) {
	fmt.Printf("\n💥 %s %s\n", p.name, random(damageReasons))

	if p.isDead() {
		return
	}

	if h >= p.health {
		p.health = 0
	} else {
		p.health -= h
	}

	fmt.Printf("☠️  Lost %d health\n", h)
	fmt.Printf("❤️ Current Health: %d\n", p.health)

	p.checkDeath()

	if p.health == 0 {
		wasted(p.name + " got replaced by Claude Ultra Instinct Enterprise Edition.")
	}
}

func (p *Player) addHealth(h uint8) {
	if p.isDead() {
		return
	}

	fmt.Printf("\n🩹 %s %s\n", p.name, random(healReasons))

	if p.health+h >= p.maxHealth {
		p.health = p.maxHealth
	} else {
		p.health += h
	}

	fmt.Printf("✨ Gained %d health\n", h)
	fmt.Printf("❤️ Current Health: %d\n", p.health)
}

func (p *Player) removeEnergy(e uint8) {
	fmt.Printf("\n🥵 %s %s\n", p.name, random(energyLossReasons))

	if p.isDead() {
		return
	}

	if e >= p.energy {
		p.energy = 0
	} else {
		p.energy -= e
	}

	fmt.Printf("🔋 Lost %d energy\n", e)
	fmt.Printf("⚡ Current Energy: %d\n", p.energy)

	p.checkDeath()

	if p.energy == 0 {
		wasted(p.name + " got trapped in infinite Zoom meeting about AI safety.")
	}
}

func (p *Player) addEnergy(e uint8) {
	if p.isDead() {
		return
	}

	fmt.Printf("\n⚡ %s %s\n", p.name, random(energyGainReasons))

	if p.energy+e >= p.maxEnergy {
		p.energy = p.maxEnergy
	} else {
		p.energy += e
	}

	fmt.Printf("🚀 Gained %d energy\n", e)
	fmt.Printf("⚡ Current Energy: %d\n", p.energy)
}

func (p *Player) securityStrike() {
	fmt.Println("\n🛡️ FINAL STRIKE TEST")

	fmt.Println("💣 Attempting illegal over-decrease...")

	beforeH := p.health
	beforeE := p.energy

	p.removeHealth(255)
	p.removeEnergy(255)

	fmt.Println("\n🧠 SECURITY RESULT:")
	fmt.Println("✔ Health before:", beforeH, "after:", p.health)
	fmt.Println("✔ Energy before:", beforeE, "after:", p.energy)
	fmt.Println("✔ System prevented underflow + repeated death spam")
}

func (p *Player) restore50() {
	fmt.Println("\n🧬 SYSTEM RECOVERY INITIATED")

	p.dead = false
	p.health = p.maxHealth / 2
	p.energy = p.maxEnergy / 2

	fmt.Println("✨ Restored to 50% capacity")
}

func main() {
	player := Player{
		name:      "Dario The Destroyer",
		health:    100,
		maxHealth: 100,
		energy:    25,
		maxEnergy: 25,
	}

	fmt.Println("🎮 GAME STARTED")
	fmt.Println("🌍 Welcome to Anthropic Simulator 2026")
	fmt.Println("📢 Objective: survive AGI predictions and production outages.")
	fmt.Println("🏢 Difficulty: NIGHTMARE MODE")

	fmt.Println("\n📰 BREAKING NEWS:")
	fmt.Println("Dario says developers will be replaced in 6 months.")
	fmt.Println("Developers: 'bro the 6 months are recursive 💀'")

	fmt.Println()

	player.removeHealth(20)
	player.removeEnergy(5)

	player.removeHealth(15)
	player.removeEnergy(8)

	player.addHealth(10)
	player.addEnergy(5)

	player.removeHealth(35)
	player.removeEnergy(10)

	player.addHealth(20)
	player.addEnergy(15)

	player.removeHealth(50)
	player.removeEnergy(25)

	player.addHealth(100)
	player.addEnergy(100)

	player.securityStrike()

	fmt.Println("\n🏆 ACHIEVEMENTS UNLOCKED")
	fmt.Println("✔️ Survived AGI prediction cycle")
	fmt.Println("✔️ Pretended to understand Kubernetes")
	fmt.Println("✔️ Fixed production by restarting everything")
	fmt.Println("✔️ Attended useless AI alignment meeting")
	fmt.Println("✔️ Read AI Twitter without losing sanity")

	fmt.Println("\n💬 NPC Developer:")
	fmt.Println(`   "Every 6 months we're supposedly finished."`)
	fmt.Println(`   "At this point the timeline has recursion."`)

	fmt.Println("\n💬 Senior Engineer:")
	fmt.Println(`   "AI won't replace developers."`)
	fmt.Println(`   "But meetings definitely replaced productivity."`)

	fmt.Println("\n💬 Random Investor:")
	fmt.Println(`   "Can Claude make line go up?" 📈`)
}
