package gen

import (
	"fmt"
	"math/rand/v2"
	"strings"
)

// Product represents a product listing.
type Product struct {
	Name        string  `json:"name"`
	Category    string  `json:"category"`
	Price       float64 `json:"price"`
	Discount    float64 `json:"discount_percent"`
	Description string  `json:"description"`
	Stock       int     `json:"stock"`
}

var (
	productPrefixes = []string{
		"Ultra", "Pro", "Smart", "Eco", "Max", "Mini", "Turbo", "Flex",
		"Aero", "Neo", "Prime", "Essential", "Advanced", "Classic", "Modern",
		"Diamond", "Platinum", "Quantum", "Stellar", "Vertex", "Zen",
	}
	productNames = []string{
		"Headphones", "Keyboard", "Monitor", "Mouse", "Laptop Stand", "Webcam",
		"Charger", "Speaker", "Microphone", "SSD Drive", "RAM Module", "GPU Cooler",
		"Router", "Switch", "Tablet", "Watch", "Fitness Tracker", "Smart Bulb",
		"Printer", "Scanner", "Projector", "Drone", "Camera Lens", "Tripod",
		"Backpack", "Sneakers", "Blender", "Coffee Maker", "Air Purifier",
		"Robot Vacuum", "Electric Toothbrush", "Shower Head", "Desk Lamp",
	}
	categories = []string{
		"Electronics", "Clothing", "Home & Kitchen", "Sports", "Books",
		"Toys", "Health", "Beauty", "Automotive", "Garden", "Music", "Office",
	}
	adjectives = []string{
		"premium", "ergonomic", "wireless", "compact", "durable", "lightweight",
		"advanced", "sleek", "portable", "intelligent", "versatile", "high-performance",
		"revolutionary", "innovative", "sustainable", "efficient", "reliable", "powerful",
	}
)

// GenerateProduct creates a single random product.
func GenerateProduct(rng *rand.Rand) Product {
	prefix := productPrefixes[rng.IntN(len(productPrefixes))]
	name := productNames[rng.IntN(len(productNames))]
	fullName := prefix + " " + name
	category := categories[rng.IntN(len(categories))]
	price := roundTo(4.99+rng.Float64()*995.01, 2)
	discount := roundTo(0+rng.Float64()*25, 2)
	adj1 := adjectives[rng.IntN(len(adjectives))]
	adj2 := adjectives[rng.IntN(len(adjectives))]
	for adj2 == adj1 {
		adj2 = adjectives[rng.IntN(len(adjectives))]
	}
	description := fmt.Sprintf("A %s %s %s designed for quality and performance. Features %s build with premium materials.",
		adj1, adj2, strings.ToLower(name), adj1)
	stock := rng.IntN(500)

	return Product{
		Name:        fullName,
		Category:    category,
		Price:       price,
		Discount:    discount,
		Description: description,
		Stock:       stock,
	}
}
