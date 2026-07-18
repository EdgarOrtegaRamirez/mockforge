package gen

import (
	"encoding/json"
	"fmt"
	"strings"
)

// FormatOutput formats the generated data into the specified output format.
func FormatOutput(data interface{}, format OutputFormat) (string, error) {
	switch format {
	case OutputJSON:
		return formatJSON(data)
	case OutputCSV:
		return formatCSV(data)
	case OutputSQL:
		return formatSQL(data)
	case OutputYAML:
		return formatYAML(data)
	default:
		return "", fmt.Errorf("unsupported format: %s", format)
	}
}

func formatJSON(data interface{}) (string, error) {
	b, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return "", fmt.Errorf("json marshal error: %w", err)
	}
	return string(b) + "\n", nil
}

func formatCSV(data interface{}) (string, error) {
	var headers []string
	var rows [][]string

	switch v := data.(type) {
	case []Person:
		headers = []string{"first_name", "last_name", "email", "phone", "date_of_birth", "gender", "username"}
		rows = make([][]string, len(v))
		for i, p := range v {
			rows[i] = []string{p.FirstName, p.LastName, p.Email, p.Phone, p.DOB, p.Gender, p.Username}
		}
	case []Address:
		headers = []string{"street", "city", "state", "country", "zip", "lat", "lng"}
		rows = make([][]string, len(v))
		for i, a := range v {
			rows[i] = []string{a.Street, a.City, a.State, a.Country, a.Zip, fmt.Sprintf("%.4f", a.Lat), fmt.Sprintf("%.4f", a.Lng)}
		}
	case []Company:
		headers = []string{"name", "industry", "email", "phone", "website", "slogan", "founded", "employees"}
		rows = make([][]string, len(v))
		for i, c := range v {
			rows[i] = []string{c.Name, c.Industry, c.Email, c.Phone, c.Website, c.Slogan, fmt.Sprintf("%d", c.Founded), fmt.Sprintf("%d", c.Employeees)}
		}
	case []Product:
		headers = []string{"name", "category", "price", "discount_percent", "description", "stock"}
		rows = make([][]string, len(v))
		for i, p := range v {
			rows[i] = []string{p.Name, p.Category, fmt.Sprintf("%.2f", p.Price), fmt.Sprintf("%.2f", p.Discount), p.Description, fmt.Sprintf("%d", p.Stock)}
		}
	case []map[string]interface{}:
		headers = []string{"type", "data_summary"}
		rows = make([][]string, len(v))
		for i, m := range v {
			tt := m["type"].(string)
			var summary string
			switch tt {
			case "person":
				summary = fmt.Sprintf("%s %s", m["first"], m["last"])
			case "address":
				summary = fmt.Sprintf("%s, %s", m["city"], m["state"])
			case "company":
				summary = m["name"].(string)
			case "product":
				summary = fmt.Sprintf("%s ($%.2f)", m["name"], m["price"])
			default:
				summary = fmt.Sprintf("%v", m)
			}
			rows[i] = []string{tt, summary}
		}
	default:
		return "", fmt.Errorf("unsupported type for CSV: %T", data)
	}

	var sb strings.Builder
	sb.WriteString(strings.Join(headers, ",") + "\n")
	for _, row := range rows {
		sb.WriteString(strings.Join(row, ",") + "\n")
	}
	return sb.String(), nil
}

func formatSQL(data interface{}) (string, error) {
	var sb strings.Builder

	switch v := data.(type) {
	case []Person:
		sb.WriteString("-- Mock Persons\n")
		sb.WriteString("INSERT INTO persons (first_name, last_name, email, phone, date_of_birth, gender, username) VALUES\n")
		for i, p := range v {
			term := ","
			if i == len(v)-1 {
				term = ";"
			}
			sb.WriteString(fmt.Sprintf("  ('%s', '%s', '%s', '%s', '%s', '%s', '%s')%s\n",
				p.FirstName, p.LastName, p.Email, p.Phone, p.DOB, p.Gender, p.Username, term))
		}
	case []Address:
		sb.WriteString("-- Mock Addresses\n")
		sb.WriteString("INSERT INTO addresses (street, city, state, country, zip, lat, lng) VALUES\n")
		for i, a := range v {
			term := ","
			if i == len(v)-1 {
				term = ";"
			}
			sb.WriteString(fmt.Sprintf("  ('%s', '%s', '%s', '%s', '%s', %.4f, %.4f)%s\n",
				a.Street, a.City, a.State, a.Country, a.Zip, a.Lat, a.Lng, term))
		}
	case []Company:
		sb.WriteString("-- Mock Companies\n")
		sb.WriteString("INSERT INTO companies (name, industry, email, phone, website, slogan, founded, employees) VALUES\n")
		for i, c := range v {
			term := ","
			if i == len(v)-1 {
				term = ";"
			}
			sb.WriteString(fmt.Sprintf("  ('%s', '%s', '%s', '%s', '%s', '%s', %d, %d)%s\n",
				c.Name, c.Industry, c.Email, c.Phone, c.Website, c.Slogan, c.Founded, c.Employeees, term))
		}
	case []Product:
		sb.WriteString("-- Mock Products\n")
		sb.WriteString("INSERT INTO products (name, category, price, discount_percent, description, stock) VALUES\n")
		for i, p := range v {
			term := ","
			if i == len(v)-1 {
				term = ";"
			}
			sb.WriteString(fmt.Sprintf("  ('%s', '%s', %.2f, %.2f, '%s', %d)%s\n",
				p.Name, p.Category, p.Price, p.Discount, p.Description, p.Stock, term))
		}
	case []map[string]interface{}:
		sb.WriteString("-- Mock Mixed Data\n")
		sb.WriteString("INSERT INTO mixed_data (type, name, details) VALUES\n")
		for i, m := range v {
			tt := m["type"].(string)
			var name, details string
			switch tt {
			case "person":
				name = fmt.Sprintf("%s %s", m["first"], m["last"])
				details = m["email"].(string)
			case "address":
				name = m["city"].(string)
				details = fmt.Sprintf("%s %s", m["state"], m["zip"])
			case "company":
				name = m["name"].(string)
				details = m["industry"].(string)
			case "product":
				name = m["name"].(string)
				details = fmt.Sprintf("$%.2f", m["price"])
			}
			term := ","
			if i == len(v)-1 {
				term = ";"
			}
			sb.WriteString(fmt.Sprintf("  ('%s', '%s', '%s')%s\n", tt, name, details, term))
		}
	default:
		return "", fmt.Errorf("unsupported type for SQL: %T", data)
	}

	return sb.String(), nil
}

func formatYAML(data interface{}) (string, error) {
	var sb strings.Builder
	switch v := data.(type) {
	case []Person:
		sb.WriteString("# Mock Persons\n")
		for i, p := range v {
			sb.WriteString(fmt.Sprintf("- first_name: %s\n", p.FirstName))
			sb.WriteString(fmt.Sprintf("  last_name: %s\n", p.LastName))
			sb.WriteString(fmt.Sprintf("  email: %s\n", p.Email))
			sb.WriteString(fmt.Sprintf("  phone: %s\n", p.Phone))
			sb.WriteString(fmt.Sprintf("  date_of_birth: %s\n", p.DOB))
			sb.WriteString(fmt.Sprintf("  gender: %s\n", p.Gender))
			sb.WriteString(fmt.Sprintf("  username: %s\n", p.Username))
			if i < len(v)-1 {
				sb.WriteString("---\n")
			}
		}
	case []Address:
		sb.WriteString("# Mock Addresses\n")
		for i, a := range v {
			sb.WriteString(fmt.Sprintf("- street: %s\n", a.Street))
			sb.WriteString(fmt.Sprintf("  city: %s\n", a.City))
			sb.WriteString(fmt.Sprintf("  state: %s\n", a.State))
			sb.WriteString(fmt.Sprintf("  country: %s\n", a.Country))
			sb.WriteString(fmt.Sprintf("  zip: %s\n", a.Zip))
			sb.WriteString(fmt.Sprintf("  lat: %.4f\n", a.Lat))
			sb.WriteString(fmt.Sprintf("  lng: %.4f\n", a.Lng))
			if i < len(v)-1 {
				sb.WriteString("---\n")
			}
		}
	case []Company:
		sb.WriteString("# Mock Companies\n")
		for i, c := range v {
			sb.WriteString(fmt.Sprintf("- name: %s\n", c.Name))
			sb.WriteString(fmt.Sprintf("  industry: %s\n", c.Industry))
			sb.WriteString(fmt.Sprintf("  email: %s\n", c.Email))
			sb.WriteString(fmt.Sprintf("  phone: %s\n", c.Phone))
			sb.WriteString(fmt.Sprintf("  website: %s\n", c.Website))
			sb.WriteString(fmt.Sprintf("  slogan: %s\n", c.Slogan))
			sb.WriteString(fmt.Sprintf("  founded: %d\n", c.Founded))
			sb.WriteString(fmt.Sprintf("  employees: %d\n", c.Employeees))
			if i < len(v)-1 {
				sb.WriteString("---\n")
			}
		}
	case []Product:
		sb.WriteString("# Mock Products\n")
		for i, p := range v {
			sb.WriteString(fmt.Sprintf("- name: %s\n", p.Name))
			sb.WriteString(fmt.Sprintf("  category: %s\n", p.Category))
			sb.WriteString(fmt.Sprintf("  price: %.2f\n", p.Price))
			sb.WriteString(fmt.Sprintf("  discount_percent: %.2f\n", p.Discount))
			sb.WriteString(fmt.Sprintf("  description: %s\n", p.Description))
			sb.WriteString(fmt.Sprintf("  stock: %d\n", p.Stock))
			if i < len(v)-1 {
				sb.WriteString("---\n")
			}
		}
	case []map[string]interface{}:
		sb.WriteString("# Mock Mixed Data\n")
		for i, m := range v {
			tt := m["type"].(string)
			sb.WriteString(fmt.Sprintf("- type: %s\n", tt))
			switch tt {
			case "person":
				sb.WriteString(fmt.Sprintf("  name: %s %s\n", m["first"], m["last"]))
				sb.WriteString(fmt.Sprintf("  email: %s\n", m["email"]))
			case "address":
				sb.WriteString(fmt.Sprintf("  city: %s\n", m["city"]))
				sb.WriteString(fmt.Sprintf("  state: %s\n", m["state"]))
			case "company":
				sb.WriteString(fmt.Sprintf("  name: %s\n", m["name"]))
				sb.WriteString(fmt.Sprintf("  industry: %s\n", m["industry"]))
			case "product":
				sb.WriteString(fmt.Sprintf("  name: %s\n", m["name"]))
				sb.WriteString(fmt.Sprintf("  price: %.2f\n", m["price"]))
			}
			if i < len(v)-1 {
				sb.WriteString("---\n")
			}
		}
	default:
		return "", fmt.Errorf("unsupported type for YAML: %T", data)
	}
	return sb.String(), nil
}
