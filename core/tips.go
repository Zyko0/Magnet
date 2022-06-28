package core

import (
	"math/rand"
	"time"
)

var (
	tips = []string{
		"Tip: Don't forget to dash in order to fight magnetism!",
		"Tip: Look deeper to anticipate future collisions!",
		"Tip: You are way too good for tips!",
		"Tip: If you need to be accurate, slide when you can!",
		"Tip: While sliding, you can still dash to slide faster!",
		"Tip: Your hitbox is actually a circle, and stays the same",
		"Tip: Do not listen to tips too much, trust yourself!",
	}
	tipIndex int
)

func init() {
	rand.Seed(time.Now().UnixNano())
	// Note: -1 to avoid showing last tip first :D
	tipIndex = rand.Intn(len(tips) - 1)
}

func increaseTipIndex() {
	tipIndex = (tipIndex + 1) % len(tips)
}

func GetTip() string {
	return tips[tipIndex]
}
