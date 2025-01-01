package main

import (
	"bufio"
	_ "embed"
	"encoding/csv"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/fatih/color"
	"github.com/kevincobain2000/doubutsu-uranai/pkg"
)

//go:embed gold.csv
var goldData string

func main() {
	animalsEn := map[string]string{
		"狼":    "Wolf",
		"ひつじ":  "Sheep",
		"ゾウ":   "Elephant",
		"黒ひょう": "Black Panther",
		"こじか":  "Fawn",
		"たぬき":  "Raccoon Dog",
		"猿":    "Monkey",
		"ライオン": "Lion",
		"子守熊":  "Koala",
		"虎":    "Tiger",
		"ペガサス": "Pegasus",
		"チータ":  "Cheetah",
	}

	// Translation map for categories
	categoryTranslations := map[string]string{
		"希望キャラ":   "Aspiration Character",
		"意思決定キャラ": "Decision-Making Character",
		"本質キャラ":   "Core Character",
		"表面キャラ":   "Surface Character",
		"隠れキャラ":   "Hidden Character",
	}

	// Create softer pastel-like color functions
	softGreen := color.New(color.FgHiGreen).SprintFunc()
	softYellow := color.New(color.FgHiYellow).SprintFunc()
	softCyan := color.New(color.FgHiCyan).SprintFunc()
	softMagenta := color.New(color.FgHiMagenta).SprintFunc()
	softRed := color.New(color.FgHiRed).SprintFunc()

	// Prompt user input
	reader := bufio.NewReader(os.Stdin)

	// Validate gender
	var gender string
	for {
		fmt.Print(softCyan("Enter your gender (M/F): "))
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToUpper(input))

		if input == "M" || input == "F" {
			gender = input
			break
		}
		fmt.Println(softRed("Invalid input. Please enter 'M' for male or 'F' for female."))
	}

	// Validate date of birth
	var dob string
	// Compile the regular expression
	dateRegex := `^\d{4}-\d{2}-\d{2}$`
	re, err := regexp.Compile(dateRegex)
	if err != nil {
		fmt.Println("Error compiling regular expression:", err)
		return
	}

	// Start reading user input
	for {
		fmt.Print("Enter your date of birth (YYYY-MM-DD): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		// Use the compiled regular expression to match the input
		if re.MatchString(input) {
			dob = input
			break
		}
		fmt.Println("Invalid date format. Please enter the date in YYYY-MM-DD format.")
	}

	// Parse the CSV data
	records := parseCSV(goldData)

	// Find matching records
	matches := filterRecords(records, gender, dob)

	// Display results
	if len(matches) == 0 {
		fmt.Println(softMagenta("No data found for the given gender and date of birth."))
		return
	}

	fmt.Println(softCyan(fmt.Sprintf("Results for Gender: %s, Date of Birth: %s\n", gender, dob)))
	for _, match := range matches {
		parts := strings.Split(match[2], "：")
		if len(parts) != 2 {
			continue
		}
		categoryJP := parts[0]
		animalJP := parts[1]

		categoryEN := categoryTranslations[categoryJP]
		animalEN := pkg.AnimalDescriptions[animalJP]

		fmt.Println(softGreen(fmt.Sprintf("Your %s(%s): %s(%s)", categoryJP, categoryEN, animalJP, animalsEn[animalJP])))
		fmt.Println(softYellow("Description"))
		for _, line := range strings.Split(animalEN, "\n") {
			fmt.Println(line)
		}
		fmt.Println()
	}
}

// parseCSV reads and parses the embedded CSV data.
func parseCSV(data string) [][]string {
	reader := csv.NewReader(strings.NewReader(data))
	reader.FieldsPerRecord = -1 // Allow variable fields
	records, _ := reader.ReadAll()
	return records
}

// filterRecords filters records by gender and date of birth.
func filterRecords(records [][]string, gender, dob string) [][]string {
	var matches [][]string
	for _, record := range records {
		if len(record) < 3 {
			continue
		}
		if strings.EqualFold(record[0], gender) && record[1] == dob {
			matches = append(matches, record)
		}
	}
	return matches
}
