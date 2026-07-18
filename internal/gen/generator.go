package gen

// OutputFormat specifies the output format for generated data.
type OutputFormat string

const (
	OutputJSON OutputFormat = "json"
	OutputCSV  OutputFormat = "csv"
	OutputSQL  OutputFormat = "sql"
	OutputYAML OutputFormat = "yaml"
)

// MockGenerator is the main entry point for generating mock data.
type MockGenerator struct {
	Format OutputFormat
	Seed   int64
}

// New creates a new MockGenerator with the given format and optional seed.
func New(format OutputFormat, seed ...int64) *MockGenerator {
	var s int64
	if len(seed) > 0 {
		s = seed[0]
	}
	return &MockGenerator{Format: format, Seed: s}
}

// GeneratePersons generates n person profiles.
func (g *MockGenerator) GeneratePersons(n int) []Person {
	rng := randForSeed(g.Seed)
	people := make([]Person, n)
	for i := 0; i < n; i++ {
		people[i] = GeneratePerson(rng)
	}
	return people
}

// GenerateAddresses generates n addresses.
func (g *MockGenerator) GenerateAddresses(n int) []Address {
	rng := randForSeed(g.Seed)
	addrs := make([]Address, n)
	for i := 0; i < n; i++ {
		addrs[i] = GenerateAddress(rng)
	}
	return addrs
}

// GenerateCompanies generates n company profiles.
func (g *MockGenerator) GenerateCompanies(n int) []Company {
	rng := randForSeed(g.Seed)
	companies := make([]Company, n)
	for i := 0; i < n; i++ {
		companies[i] = GenerateCompany(rng)
	}
	return companies
}

// GenerateProducts generates n products.
func (g *MockGenerator) GenerateProducts(n int) []Product {
	rng := randForSeed(g.Seed)
	products := make([]Product, n)
	for i := 0; i < n; i++ {
		products[i] = GenerateProduct(rng)
	}
	return products
}

// GenerateMixed generates a mix of all data types.
func (g *MockGenerator) GenerateMixed(n int) []map[string]interface{} {
	rng := randForSeed(g.Seed)
	results := make([]map[string]interface{}, n)
	for i := 0; i < n; i++ {
		switch rng.IntN(4) {
		case 0:
			p := GeneratePerson(rng)
			results[i] = map[string]interface{}{
				"type":     "person",
				"first":    p.FirstName,
				"last":     p.LastName,
				"email":    p.Email,
				"phone":    p.Phone,
				"dob":      p.DOB,
				"gender":   p.Gender,
				"username": p.Username,
			}
		case 1:
			a := GenerateAddress(rng)
			results[i] = map[string]interface{}{
				"type":    "address",
				"street":  a.Street,
				"city":    a.City,
				"state":   a.State,
				"country": a.Country,
				"zip":     a.Zip,
				"lat":     a.Lat,
				"lng":     a.Lng,
			}
		case 2:
			c := GenerateCompany(rng)
			results[i] = map[string]interface{}{
				"type":      "company",
				"name":      c.Name,
				"industry":  c.Industry,
				"email":     c.Email,
				"phone":     c.Phone,
				"website":   c.Website,
				"slogan":    c.Slogan,
				"founded":   c.Founded,
				"employees": c.Employeees,
			}
		default:
			p := GenerateProduct(rng)
			results[i] = map[string]interface{}{
				"type":        "product",
				"name":        p.Name,
				"category":    p.Category,
				"price":       p.Price,
				"discount":    p.Discount,
				"description": p.Description,
				"stock":       p.Stock,
			}
		}
	}
	return results
}
