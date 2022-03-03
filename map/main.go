package main

import "fmt"
import "sync"

var rwm sync.RWMutex

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	colors["yellow"] = "laksjdf"
	delete(colors, "yellow")

	printMap(colors)

	// map not thread safe so use Mutex
	set(colors, "gold", "######")
	c := get(colors, "gold")
	fmt.Printf("c: %s", c)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}

func get(m map[string]string, key string) string {
	rwm.RLock()
	defer rwm.RUnlock()
	return m[key]
}

func set(m map[string]string, key string, value string) {
	rwm.Lock()
	defer rwm.Unlock()
	m[key] = value
}
