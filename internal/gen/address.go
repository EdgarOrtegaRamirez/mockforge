package gen

import (
	"fmt"
	"math/rand/v2"
)

// Address represents a physical address.
type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	Country string `json:"country"`
	Zip     string `json:"zip"`
	Lat     float64 `json:"lat"`
	Lng     float64 `json:"lng"`
}

var (
	streetPrefixes = []string{"Main", "Oak", "Elm", "Maple", "Cedar", "Pine", "Washington",
		"Lake", "Hill", "Park", "Sunset", "Spring", "Forest", "River", "Mill",
		"Church", "School", "Highland", "Valley", "Meadow", "Brook", "Shore",
		"Ridge", "Cliff", "Glen", "Stone", "Willow", "Magnolia", "Cherry", "Poplar"}
	streetSuffixes = []string{"Street", "Avenue", "Road", "Drive", "Lane", "Boulevard",
		"Court", "Place", "Way", "Trail", "Circle", "Terrace", "Way", "Path"}

	cities = []string{
		"New York", "Los Angeles", "Chicago", "Houston", "Phoenix", "Philadelphia",
		"San Antonio", "San Diego", "Dallas", "San Jose", "Austin", "Jacksonville",
		"Fort Worth", "Columbus", "Charlotte", "San Francisco", "Indianapolis",
		"Seattle", "Denver", "Nashville", "Portland", "Las Vegas", "Memphis",
		"Louisville", "Baltimore", "Milwaukee", "Albuquerque", "Tucson", "Fresno",
		"Sacramento", "Kansas City", "Mesa", "Atlanta", "Omaha", "Colorado Springs",
		"Raleigh", "Miami", "Virginia Beach", "Oakland", "Minneapolis", "Tampa",
		"Tulsa", "Arlington", "New Orleans", "Wichita", "Cleveland", "Bakersfield",
		"Aurora", "Anaheim", "Honolulu", "Santa Ana", "Riverside", "Corpus Christi",
		"Lexington", "Henderson", "Stockton", "Saint Paul", "St. Louis", "Cincinnati",
		"Pittsburgh", "Greensboro", "Anchorage", "Plano", "Lincoln", "Orlando",
		"Irvine", "Newark", "Toledo", "Durham", "Chula Vista", "Fort Wayne",
		"St. Petersburg", "Laredo", "Jersey City", "Chandler", "Madison",
	}

	states = map[string]string{
		"NY": "New York", "CA": "California", "IL": "Illinois", "TX": "Texas",
		"AZ": "Arizona", "PA": "Pennsylvania", "FL": "Florida", "OH": "Ohio",
		"GA": "Georgia", "NC": "North Carolina", "MI": "Michigan", "WA": "Washington",
		"CO": "Colorado", "TN": "Tennessee", "NV": "Nevada", "MD": "Maryland",
		"MA": "Massachusetts", "NM": "New Mexico", "VA": "Virginia", "OR": "Oregon",
		"MN": "Minnesota", "SC": "South Carolina", "AL": "Alabama", "LA": "Louisiana",
		"KY": "Kentucky", "UT": "Utah", "OK": "Oklahoma", "CT": "Connecticut",
		"IA": "Iowa", "MS": "Mississippi", "AR": "Arkansas", "KS": "Kansas",
		"DE": "Delaware", "HI": "Hawaii", "NH": "New Hampshire", "MT": "Montana",
		"RI": "Rhode Island", "WV": "West Virginia", "ID": "Idaho", "NE": "Nebraska",
		"SD": "South Dakota", "ND": "North Dakota", "AK": "Alaska",
		"VT": "Vermont", "WY": "Wyoming",
	}

	countries = []string{"US", "CA", "MX", "GB", "DE", "FR", "AU", "JP", "BR", "IN"}
)

// GenerateAddress creates a single random address.
func GenerateAddress(rng *rand.Rand) Address {
	streetNum := 100 + rng.IntN(9900)
	prefix := streetPrefixes[rng.IntN(len(streetPrefixes))]
	suffix := streetSuffixes[rng.IntN(len(streetSuffixes))]
	street := fmt.Sprintf("%d %s %s", streetNum, prefix, suffix)

	city := cities[rng.IntN(len(cities))]
	state := stateCode(rng)
	country := countries[rng.IntN(len(countries))]
	zip := fmt.Sprintf("%05d", 10000+rng.IntN(90000))
	lat := roundTo(40.7128 + (rng.Float64()-0.5)*40, 4)
	lng := roundTo(-74.0060 + (rng.Float64()-0.5)*60, 4)

	return Address{
		Street:  street,
		City:    city,
		State:   state,
		Country: country,
		Zip:     zip,
		Lat:     lat,
		Lng:     lng,
	}
}

func stateCode(rng *rand.Rand) string {
	keys := make([]string, 0, len(states))
	for k := range states {
		keys = append(keys, k)
	}
	return keys[rng.IntN(len(keys))]
}

func roundTo(val float64, places int) float64 {
	pow := 1.0
	for i := 0; i < places; i++ {
		pow *= 10
	}
	return float64(int(val*pow+0.5)) / pow
}