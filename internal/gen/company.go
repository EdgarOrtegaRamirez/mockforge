package gen

import (
	"fmt"
	"math/rand/v2"
	"strings"
)

// Company represents a company profile.
type Company struct {
	Name       string `json:"name"`
	Industry   string `json:"industry"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Website    string `json:"website"`
	Slogan     string `json:"slogan"`
	Founded    int    `json:"founded"`
	Employeees int    `json:"employees"`
}

var (
	prefixes = []string{"Apex", "Nova", "Peak", "Summit", "Titan", "Atlas", "Prism",
		"Vertex", "Pulse", "Fusion", "Core", "Nexus", "Zenith", "Orbit", "Spark",
		"Vista", "Iron", "Silver", "Gold", "Blue", "Red", "Green", "Alpha", "Omega",
		"Quantum", "Hyper", "Meta", "Proto", "Dynamo", "Stellar", "Cosmic", "Solar",
		"Storm", "Blaze", "Frost", "Thunder", "Sapphire", "Cobalt", "Amber", "Crimson",
		"Echo", "Flux", "Grid", "Halo", "Ion", "Jade", "Knot", "Lumen", "Mist", "Noble",
	}
	suffixes = []string{
		"Technologies", "Solutions", "Systems", "Labs", "Global", "International",
		"Corp", "Group", "Industries", "Partners", "Ventures", "Dynamics", "Enterprises",
		"Capital", "Consulting", "Digital", "Software", "Networks", "Services",
		"Innovations", "Manufacturing", "Analytics", "Security", "Health", "Energy",
	}

	industries = []string{
		"Technology", "Healthcare", "Finance", "Education", "Manufacturing",
		"Retail", "Energy", "Real Estate", "Transportation", "Entertainment",
		"Telecommunications", "Construction", "Agriculture", "Aviation", "Chemicals",
	}

	slogans = []string{
		"Innovation you can trust", "Building tomorrow today", "Powering progress",
		"Excellence in every detail", "Where ideas become reality",
		"Leading the way forward", "Transforming industries",
		"Driven by excellence", "Empowering your success",
		"Redefining what's possible", "Your vision, our mission",
		"Quality that speaks for itself", "Innovation at scale",
	}
)

// GenerateCompany creates a single random company profile.
func GenerateCompany(rng *rand.Rand) Company {
	prefix := prefixes[rng.IntN(len(prefixes))]
	suffix := suffixes[rng.IntN(len(suffixes))]
	industry := industries[rng.IntN(len(industries))]
	name := prefix + " " + suffix
	slug := strings.ToLower(strings.ReplaceAll(name, " ", ""))
	email := fmt.Sprintf("info@%s.com", slug)
	phone := formatPhone(rng)
	website := fmt.Sprintf("www.%s.com", slug)
	slogan := slogans[rng.IntN(len(slogans))]
	founded := 1960 + rng.IntN(64)
	employees := 10 + rng.IntN(99990)

	return Company{
		Name:       name,
		Industry:   industry,
		Email:      email,
		Phone:      phone,
		Website:    website,
		Slogan:     slogan,
		Founded:    founded,
		Employeees: employees,
	}
}