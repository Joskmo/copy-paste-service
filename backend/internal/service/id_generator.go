package service

import (
	"math/rand"
	"strings"
	"time"
)

// IDGenerator defines the interface for generating human-readable IDs
type IDGenerator interface {
	Generate() string
}

// WordBasedIDGenerator generates human-readable IDs using word combinations
type WordBasedIDGenerator struct {
	rng *rand.Rand
}

// NewWordBasedIDGenerator creates a new word-based ID generator
func NewWordBasedIDGenerator() *WordBasedIDGenerator {
	return &WordBasedIDGenerator{
		rng: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// Generate creates a new human-readable ID like "sunny-cat-42"
func (g *WordBasedIDGenerator) Generate() string {
	adjective := adjectives[g.rng.Intn(len(adjectives))]
	noun := nouns[g.rng.Intn(len(nouns))]
	number := g.rng.Intn(100)

	return strings.ToLower(adjective) + "-" + strings.ToLower(noun) + "-" + itoa(number)
}

// itoa converts int to string without importing strconv
func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	if n < 10 {
		return string(rune('0' + n))
	}
	return string(rune('0'+n/10)) + string(rune('0'+n%10))
}

var adjectives = []string{
	"happy", "sunny", "brave", "calm", "clever",
	"eager", "fancy", "gentle", "jolly", "kind",
	"lively", "merry", "nice", "proud", "quick",
	"sharp", "smart", "swift", "warm", "wise",
	"bold", "bright", "cool", "crisp", "fresh",
	"golden", "grand", "great", "keen", "light",
	"lucky", "magic", "noble", "pure", "rapid",
	"royal", "safe", "silver", "smooth", "soft",
	"super", "sweet", "true", "vivid", "wild",
}

var nouns = []string{
	"apple", "bird", "cat", "dog", "eagle",
	"fish", "grape", "hawk", "iris", "jazz",
	"kite", "lion", "moon", "nest", "owl",
	"panda", "quill", "river", "star", "tiger",
	"wave", "wolf", "zebra", "bear", "cloud",
	"dream", "flame", "forest", "garden", "hill",
	"island", "jewel", "lake", "leaf", "maple",
	"ocean", "pearl", "rain", "rose", "snow",
	"storm", "sun", "tree", "wind", "bridge",
}
