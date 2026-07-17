# MockForge

Realistic test data generation CLI for developers. Generate mock person profiles, addresses, companies, products, or mixed data in JSON, CSV, SQL, or YAML format.

## Features

- **5 data types**: Persons, addresses, companies, products, and mixed
- **4 output formats**: JSON, CSV, SQL INSERT statements, YAML
- **Reproducible generation**: Use a seed for consistent, repeatable results
- **Pure Go**: Zero external dependencies at runtime (stdlib only)

## Installation

```bash
go install github.com/EdgarOrtegaRamirez/mockforge/cmd/mockforge@latest
```

Or build from source:

```bash
git clone https://github.com/EdgarOrtegaRamirez/mockforge
cd mockforge
go build -o mockforge ./cmd/mockforge/
```

## Usage

```bash
# Generate 5 person profiles as JSON
mockforge person -n 5 -f json

# Generate 3 companies as CSV
mockforge company -n 3 -f csv

# Generate 10 products as SQL
mockforge product -n 10 -f sql

# Generate mixed data with a fixed seed
mockforge mixed -n 20 -f json --seed 42

# Generate addresses as YAML
mockforge address -n 5 -f yaml
```

## Commands

| Command   | Description                  |
|-----------|------------------------------|
| `person`  | Generate mock person profiles |
| `address` | Generate mock addresses      |
| `company` | Generate mock company profiles |
| `product` | Generate mock products       |
| `mixed`   | Generate a mix of all types  |

## Options

| Option         | Short | Description                          | Default |
|----------------|-------|--------------------------------------|---------|
| `--count`      | `-n`  | Number of records to generate        | 10      |
| `--format`     | `-f`  | Output format: json, csv, sql, yaml  | json    |
| `--seed`       |       | Seed for reproducible generation     | random  |

## Output Formats

### JSON (default)
```json
[
  {
    "first_name": "Joseph",
    "last_name": "Flores",
    "email": "jFlores@example.com",
    "phone": "(503) 570-1342",
    "date_of_birth": "1952-10-28",
    "gender": "male",
    "username": "asmith363"
  }
]
```

### CSV
```csv
first_name,last_name,email,phone,date_of_birth,gender,username
Joseph,Flores,jFlores@example.com,(503) 570-1342,1952-10-28,male,asmith363
```

### SQL
```sql
-- Mock Persons
INSERT INTO persons (first_name, last_name, email, phone, date_of_birth, gender, username) VALUES
  ('Joseph', 'Flores', 'jFlores@example.com', '(503) 570-1342', '1952-10-28', 'male', 'asmith363');
```

### YAML
```yaml
# Mock Persons
- first_name: Joseph
  last_name: Flores
  email: jFlores@example.com
  phone: (503) 570-1342
  date_of_birth: 1952-10-28
  gender: male
  username: asmith363
```

## Library Usage

```go
import "github.com/EdgarOrtegaRamirez/mockforge/internal/gen"

g := gen.New(gen.OutputJSON, 42)
people := g.GeneratePersons(10)
```

## Reproducible Generation

Use `--seed` for consistent results:

```bash
# These produce identical output:
mockforge person -n 3 --seed 42
mockforge person -n 3 --seed 42

# These produce different output:
mockforge person -n 3 --seed 42
mockforge person -n 3 --seed 99
```

## License

MIT