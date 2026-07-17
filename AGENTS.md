# AGENTS.md — MockForge

## Project Overview

MockForge is a Go CLI tool for generating realistic mock/fake data for testing. It generates person profiles, addresses, companies, products, and mixed data in multiple output formats.

## Architecture

```
mockforge/
├── cmd/mockforge/main.go       — CLI entry point
├── internal/gen/               — Core data generation package
│   ├── generator.go            — MockGenerator struct and high-level methods
│   ├── person.go               — Person profile generation
│   ├── address.go              — Address generation
│   ├── company.go              — Company profile generation
│   ├── product.go              — Product generation
│   ├── output.go               — Output formatting (JSON, CSV, SQL, YAML)
│   ├── helpers.go              — RNG and utility helpers
│   └── generator_test.go       — Tests
└── go.mod                      — Module definition
```

## Building

```bash
go build -o mockforge ./cmd/mockforge/
```

## Testing

```bash
go test ./... -v -race
```

## Data Types

### Person
- First name, last name (US common names)
- Email (derived from name + domain)
- Phone (US format with area code)
- Date of birth (random between 1950-2005)
- Gender (based on name heuristic)
- Username (random prefix + number)

### Address
- Street number + name + suffix (US style)
- City (major US cities)
- State code (all 50 US states + territories)
- Country code (US, CA, MX, GB, DE, FR, AU, JP, BR, IN)
- ZIP code
- Latitude/longitude (roughly US-aligned)

### Company
- Name (prefix + suffix combination)
- Industry (15 categories)
- Email (info@company.com format)
- Phone (US format)
- Website (www.company.com)
- Slogan
- Founded year (1960-2024)
- Employee count (10-99,999)

### Product
- Name (adjective + product noun)
- Category (12 categories)
- Price ($4.99 - $999.99)
- Discount percent (0-25%)
- Description (template-based)
- Stock count (0-499)

## Adding a New Data Type

1. Create a new file in `internal/gen/` with the struct definition
2. Add generation function following existing patterns (use `*rand.Rand` parameter)
3. Add case in `Generator.GenerateMixed()` for mixed output
4. Add CSV headers and row conversion in `output.go`
5. Add SQL template in `output.go`
6. Add YAML template in `output.go`
7. Add tests in `generator_test.go`
8. Add CLI command in `main.go`

## Output Format Support

- **JSON**: Uses Go's `encoding/json` for proper escaping
- **CSV**: Simple comma-separated with headers
- **SQL**: INSERT INTO statements with proper quoting
- **YAML**: Manual formatting (no external dependency)

## Key Design Decisions

1. **No external dependencies** — Only stdlib + go.mod dependencies (cobra for CLI)
2. **Seed-based RNG** — Uses PCG for reproducible, deterministic output
3. **Simple CLI** — Command + flag pattern, no cobra subcommands needed
4. **Multiple formats** — Same data, multiple output styles