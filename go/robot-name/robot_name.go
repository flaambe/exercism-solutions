// Package robotname implements simple robots.
package robotname

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

const maxNamesPool = 26 * 26 * 10 * 10 * 10

var random = rand.New(rand.NewSource(time.Now().UnixNano()))
var usedNames = map[string]bool{}

// Robot represents a robot with a name.
type Robot struct {
	name string
}

// Name returns the robot's name or error if namespace is exhausted.
func (r *Robot) Name() (string, error) {
	if r.name != "" {
		return r.name, nil
	}

	if len(usedNames) >= maxNamesPool {
		return "", errors.New("namespace is exhausted")
	}

	r.name = generateRandomName()

	for usedNames[r.name] {
		r.name = generateRandomName()
	}

	usedNames[r.name] = true

	return r.name, nil
}

// Reset erases the robot's name.
func (r *Robot) Reset() {
	r.name = ""
}

func generateRandomName() string {
	letter1 := random.Intn(26) + 'A'
	letter2 := random.Intn(26) + 'A'
	digit := random.Intn(1000)

	return fmt.Sprintf("%c%c%03d", letter1, letter2, digit)
}
