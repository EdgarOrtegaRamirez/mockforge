package gen

import (
	"math/rand/v2"
	"strings"
	"testing"
)

var testRng = rand.New(rand.NewPCG(12345, 67890))

func TestGeneratePerson(t *testing.T) {
	p := GeneratePerson(testRng)
	if p.FirstName == "" {
		t.Error("FirstName should not be empty")
	}
	if p.LastName == "" {
		t.Error("LastName should not be empty")
	}
	if !strings.Contains(p.Email, "@") {
		t.Errorf("Email should contain @: %s", p.Email)
	}
	if !strings.Contains(p.Phone, "(") {
		t.Errorf("Phone should contain area code: %s", p.Phone)
	}
	if p.Gender != "male" && p.Gender != "female" {
		t.Errorf("Gender should be male or female, got: %s", p.Gender)
	}
	if p.Username == "" {
		t.Error("Username should not be empty")
	}
}

func TestGeneratePersonSeedReproducible(t *testing.T) {
	rng1 := randForSeed(42)
	rng2 := randForSeed(42)
	p1 := GeneratePerson(rng1)
	p2 := GeneratePerson(rng2)
	if p1.FirstName != p2.FirstName || p1.LastName != p2.LastName {
		t.Error("Same seed should produce same results")
	}
}

func TestGenerateAddress(t *testing.T) {
	a := GenerateAddress(testRng)
	if a.Street == "" {
		t.Error("Street should not be empty")
	}
	if a.City == "" {
		t.Error("City should not be empty")
	}
	if a.State == "" {
		t.Error("State should not be empty")
	}
	if a.Zip == "" {
		t.Error("Zip should not be empty")
	}
	if a.Lat < -90 || a.Lat > 90 {
		t.Errorf("Lat should be between -90 and 90, got: %f", a.Lat)
	}
}

func TestGenerateCompany(t *testing.T) {
	c := GenerateCompany(testRng)
	if c.Name == "" {
		t.Error("Name should not be empty")
	}
	if c.Industry == "" {
		t.Error("Industry should not be empty")
	}
	if !strings.Contains(c.Email, "@") {
		t.Errorf("Email should contain @: %s", c.Email)
	}
	if c.Founded < 1900 || c.Founded > 2026 {
		t.Errorf("Founded should be reasonable, got: %d", c.Founded)
	}
	if c.Employeees < 1 {
		t.Error("Employeees should be positive")
	}
}

func TestGenerateProduct(t *testing.T) {
	p := GenerateProduct(testRng)
	if p.Name == "" {
		t.Error("Name should not be empty")
	}
	if p.Category == "" {
		t.Error("Category should not be empty")
	}
	if p.Price <= 0 {
		t.Errorf("Price should be positive, got: %f", p.Price)
	}
	if p.Discount < 0 || p.Discount > 100 {
		t.Errorf("Discount should be 0-100, got: %f", p.Discount)
	}
	if p.Stock < 0 {
		t.Error("Stock should not be negative")
	}
}

func TestGeneratePersons(t *testing.T) {
	g := New(OutputJSON, 42)
	people := g.GeneratePersons(5)
	if len(people) != 5 {
		t.Errorf("Expected 5 persons, got %d", len(people))
	}
}

func TestGenerateCompanies(t *testing.T) {
	g := New(OutputJSON, 42)
	companies := g.GenerateCompanies(3)
	if len(companies) != 3 {
		t.Errorf("Expected 3 companies, got %d", len(companies))
	}
}

func TestGenerateProducts(t *testing.T) {
	g := New(OutputJSON, 42)
	products := g.GenerateProducts(4)
	if len(products) != 4 {
		t.Errorf("Expected 4 products, got %d", len(products))
	}
}

func TestGenerateMixed(t *testing.T) {
	g := New(OutputJSON, 42)
	data := g.GenerateMixed(20)
	if len(data) != 20 {
		t.Errorf("Expected 20 mixed records, got %d", len(data))
	}
	for i, d := range data {
		tt := d["type"]
		if tt != "person" && tt != "address" && tt != "company" && tt != "product" {
			t.Errorf("Record %d has invalid type: %v", i, tt)
		}
	}
}

func TestGenerateAddresses(t *testing.T) {
	g := New(OutputJSON, 42)
	addrs := g.GenerateAddresses(5)
	if len(addrs) != 5 {
		t.Errorf("Expected 5 addresses, got %d", len(addrs))
	}
}
