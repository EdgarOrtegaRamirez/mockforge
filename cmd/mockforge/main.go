package main

import (
	"fmt"
	"os"

	"github.com/EdgarOrtegaRamirez/mockforge/internal/gen"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "person":
		runPerson()
	case "address":
		runAddress()
	case "company":
		runCompany()
	case "product":
		runProduct()
	case "mixed":
		runMixed()
	case "help", "--help", "-h":
		printUsage()
	default:
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n\n", command)
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("MockForge — Realistic Test Data Generation CLI")
	fmt.Println()
	fmt.Println("Usage: mockforge <command> [options]")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("  person    Generate mock person profiles")
	fmt.Println("  address   Generate mock addresses")
	fmt.Println("  company   Generate mock company profiles")
	fmt.Println("  product   Generate mock products")
	fmt.Println("  mixed     Generate a mix of all data types")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("  --count, -n int       Number of records to generate (default: 10)")
	fmt.Println("  --format, -f string   Output format: json, csv, sql, yaml (default: json)")
	fmt.Println("  --seed int            Seed for reproducible generation (default: random)")
	fmt.Println("  --help, -h            Show this help message")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  mockforge person -n 5 -f json")
	fmt.Println("  mockforge company -n 3 -f csv")
	fmt.Println("  mockforge mixed -n 20 -f json --seed 42")
	fmt.Println("  mockforge product -n 10 -f sql")
}

func runPerson() {
	count := 10
	format := gen.OutputJSON
	var seed int64

	for i := 2; i < len(os.Args); i++ {
		switch os.Args[i] {
		case "--count", "-n":
			if i+1 < len(os.Args) {
				fmt.Sscanf(os.Args[i+1], "%d", &count)
				i++
			}
		case "--format", "-f":
			if i+1 < len(os.Args) {
				format = gen.OutputFormat(os.Args[i+1])
				i++
			}
		case "--seed":
			if i+1 < len(os.Args) {
				fmt.Sscanf(os.Args[i+1], "%d", &seed)
				i++
			}
		}
	}

	g := gen.New(format, seed)
	data := g.GeneratePersons(count)
	output, err := gen.FormatOutput(data, format)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Print(output)
}

func runAddress() {
	count := 10
	format := gen.OutputJSON
	var seed int64

	for i := 2; i < len(os.Args); i++ {
		switch os.Args[i] {
		case "--count", "-n":
			if i+1 < len(os.Args) {
				fmt.Sscanf(os.Args[i+1], "%d", &count)
				i++
			}
		case "--format", "-f":
			if i+1 < len(os.Args) {
				format = gen.OutputFormat(os.Args[i+1])
				i++
			}
		case "--seed":
			if i+1 < len(os.Args) {
				fmt.Sscanf(os.Args[i+1], "%d", &seed)
				i++
			}
		}
	}

	g := gen.New(format, seed)
	data := g.GenerateAddresses(count)
	output, err := gen.FormatOutput(data, format)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Print(output)
}

func runCompany() {
	count := 10
	format := gen.OutputJSON
	var seed int64

	for i := 2; i < len(os.Args); i++ {
		switch os.Args[i] {
		case "--count", "-n":
			if i+1 < len(os.Args) {
				fmt.Sscanf(os.Args[i+1], "%d", &count)
				i++
			}
		case "--format", "-f":
			if i+1 < len(os.Args) {
				format = gen.OutputFormat(os.Args[i+1])
				i++
			}
		case "--seed":
			if i+1 < len(os.Args) {
				fmt.Sscanf(os.Args[i+1], "%d", &seed)
				i++
			}
		}
	}

	g := gen.New(format, seed)
	data := g.GenerateCompanies(count)
	output, err := gen.FormatOutput(data, format)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Print(output)
}

func runProduct() {
	count := 10
	format := gen.OutputJSON
	var seed int64

	for i := 2; i < len(os.Args); i++ {
		switch os.Args[i] {
		case "--count", "-n":
			if i+1 < len(os.Args) {
				fmt.Sscanf(os.Args[i+1], "%d", &count)
				i++
			}
		case "--format", "-f":
			if i+1 < len(os.Args) {
				format = gen.OutputFormat(os.Args[i+1])
				i++
			}
		case "--seed":
			if i+1 < len(os.Args) {
				fmt.Sscanf(os.Args[i+1], "%d", &seed)
				i++
			}
		}
	}

	g := gen.New(format, seed)
	data := g.GenerateProducts(count)
	output, err := gen.FormatOutput(data, format)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Print(output)
}

func runMixed() {
	count := 10
	format := gen.OutputJSON
	var seed int64

	for i := 2; i < len(os.Args); i++ {
		switch os.Args[i] {
		case "--count", "-n":
			if i+1 < len(os.Args) {
				fmt.Sscanf(os.Args[i+1], "%d", &count)
				i++
			}
		case "--format", "-f":
			if i+1 < len(os.Args) {
				format = gen.OutputFormat(os.Args[i+1])
				i++
			}
		case "--seed":
			if i+1 < len(os.Args) {
				fmt.Sscanf(os.Args[i+1], "%d", &seed)
				i++
			}
		}
	}

	g := gen.New(format, seed)
	data := g.GenerateMixed(count)
	output, err := gen.FormatOutput(data, format)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Print(output)
}
