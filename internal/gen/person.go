package gen

import (
	"fmt"
	"math/rand/v2"
	"strings"
	"time"
)

// Person represents a person profile with realistic mock data.
type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	DOB       string `json:"date_of_birth"`
	Gender    string `json:"gender"`
	Username  string `json:"username"`
}

var (
	firstNames = []string{
		"James", "Mary", "John", "Patricia", "Robert", "Jennifer", "Michael", "Linda",
		"William", "Elizabeth", "David", "Barbara", "Richard", "Susan", "Joseph", "Jessica",
		"Thomas", "Sarah", "Charles", "Karen", "Christopher", "Lisa", "Daniel", "Nancy",
		"Matthew", "Betty", "Anthony", "Margaret", "Mark", "Sandra", "Donald", "Ashley",
		"Steven", "Kimberly", "Paul", "Emily", "Andrew", "Donna", "Joshua", "Michelle",
		"Kenneth", "Dorothy", "Kevin", "Carol", "Brian", "Amanda", "George", "Melissa",
		"Timothy", "Deborah", "Ronald", "Stephanie", "Edward", "Rebecca", "Jason", "Sharon",
		"Jeffrey", "Laura", "Ryan", "Cynthia", "Jacob", "Kathleen", "Gary", "Amy",
		"Nicholas", "Angela", "Eric", "Shirley", "Jonathan", "Anna", "Stephen", "Brenda",
		"Larry", "Pamela", "Justin", "Emma", "Scott", "Nicole", "Brandon", "Helen",
		"Benjamin", "Samantha", "Samuel", "Katherine", "Raymond", "Christine", "Gregory", "Debra",
		"Frank", "Rachel", "Alexander", "Carolyn", "Patrick", "Janet", "Jack", "Catherine",
		"Daniel", "Maria", "Henry", "Heather", "Arthur", "Diane", "Peter", "Julie",
		"Juan", "Joyce", "Jack", "Virginia", "Dennis", "Valerie",
	}

	lastNames = []string{
		"Smith", "Johnson", "Williams", "Brown", "Jones", "Garcia", "Miller", "Davis",
		"Rodriguez", "Martinez", "Hernandez", "Lopez", "Gonzalez", "Wilson", "Anderson",
		"Thomas", "Taylor", "Moore", "Jackson", "Martin", "Lee", "Perez", "Thompson",
		"White", "Harris", "Sanchez", "Clark", "Ramirez", "Lewis", "Robinson", "Walker",
		"Young", "Allen", "King", "Wright", "Scott", "Torres", "Nguyen", "Hill", "Flores",
		"Green", "Adams", "Nelson", "Baker", "Hall", "Rivera", "Campbell", "Mitchell",
		"Carter", "Roberts", "Gomez", "Phillips", "Evans", "Turner", "Diaz", "Parker",
		"Cruz", "Edwards", "Collins", "Reyes", "Stewart", "Morris", "Morales", "Murphy",
		"Cook", "Rogers", "Gutierrez", "Ortiz", "Morgan", "Cooper", "Peterson", "Bailey",
		"Reed", "Kelly", "Howard", "Ramos", "Kim", "Cox", "Ward", "Richardson",
		"Watson", "Brooks", "Chavez", "Wood", "James", "Bennett", "Gray", "Mendoza",
		"Ruiz", "Hughes", "Price", "Alvarez", "Castillo", "Sanders", "Patel", "Myers",
		"Long", "Ross", "Foster", "Jimenez", "Powell", "Jenkins", "Perry", "Russell",
		"Sullivan", "Bell", "Coleman", "Butler", "Henderson", "Barnes", "Gonzales", "Fisher",
		"Vasquez", "Simmons", "Romero", "Jordan", "Patterson", "Alexander", "Hamilton", "Graham",
		"Reynolds", "Griffin", "Wallace", "Moreno", "West", "Cole", "Hayes", "Barnett",
		"Graves", "Mendez", "Castro", "Sutton", "Gregory", "McKinney", "Lucas", "Miles",
	}

	domains = []string{
		"example.com", "testmail.io", "mailtest.org", "devops.net", "quickmail.com",
		"fastmail.net", "inbox.io", "postbox.com", "emailpro.org", "sendit.net",
		"contactme.io", "getmail.com", "reachout.org", "message.me", "connect.io",
	}

	usernames = []string{
		"jdoe", "asmith", "bwilson", "cjohnson", "dlee", "emartinez", "fgarcia",
		"ghernandez", "hlopez", "igonzalez", "jwilson", "kanderson", "lthomas",
		"mtaylor", "nmoore", "oturner", "pjackson", "qwhite", "rharris", "smartin",
		"tclark", "ulewis", "vrobinson", "wwalker", "xyoung", "zallen",
	}
)

// GeneratePerson creates a single random person profile.
func GeneratePerson(rng *rand.Rand) Person {
	first := firstNames[rng.IntN(len(firstNames))]
	last := lastNames[rng.IntN(len(lastNames))]
	email := strings.ToLower(first[:1]) + last + "@" + domains[rng.IntN(len(domains))]
	phone := formatPhone(rng)
	dob := randomDOB(rng)
	gender := pickGender(rng, first)
	username := usernames[rng.IntN(len(usernames))] + fmt.Sprintf("%d", rng.IntN(999))

	return Person{
		FirstName: first,
		LastName:  last,
		Email:     email,
		Phone:     phone,
		DOB:       dob,
		Gender:    gender,
		Username:  username,
	}
}

func formatPhone(rng *rand.Rand) string {
	area := 200 + rng.IntN(800)
	prefix := 200 + rng.IntN(800)
	line := 1000 + rng.IntN(9000)
	return fmt.Sprintf("(%d) %d-%d", area, prefix, line)
}

func randomDOB(rng *rand.Rand) string {
	start := time.Date(1950, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2005, 1, 1, 0, 0, 0, 0, time.UTC)
	d := start.Add(time.Duration(rng.Int64N(int64(end.Sub(start)))))
	return d.Format("2006-01-02")
}

func pickGender(rng *rand.Rand, firstName string) string {
	femaleNames := map[string]bool{
		"Mary": true, "Patricia": true, "Jennifer": true, "Linda": true,
		"Elizabeth": true, "Susan": true, "Jessica": true, "Sarah": true,
		"Karen": true, "Lisa": true, "Betty": true, "Margaret": true,
		"Sandra": true, "Ashley": true, "Kimberly": true, "Emily": true,
		"Michelle": true, "Dorothy": true, "Carol": true, "Amanda": true,
		"Melissa": true, "Deborah": true, "Stephanie": true, "Rebecca": true,
		"Sharon": true, "Laura": true, "Cynthia": true, "Kathleen": true,
		"Amy": true, "Angela": true, "Shirley": true, "Brenda": true,
		"Pamela": true, "Emma": true, "Nicole": true, "Helen": true,
		"Samantha": true, "Katherine": true, "Christine": true, "Debra": true,
		"Rachel": true, "Carolyn": true, "Janet": true, "Catherine": true,
		"Maria": true, "Heather": true, "Diane": true, "Julie": true,
		"Joyce": true, "Virginia": true, "Valerie": true, "Anna": true,
	}
	if femaleNames[firstName] {
		return "female"
	}
	if rng.Float64() < 0.48 {
		return "female"
	}
	return "male"
}