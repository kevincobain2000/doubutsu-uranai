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
)

//go:embed gold.csv
var goldData string

func main() {
	// Animal descriptions
	animalDescriptions := map[string]string{
		"狼":    "Wolf: Independent and confident, a natural leader. People born under the sign of the Wolf are known for their strong sense of individuality and leadership qualities. They tend to be self-reliant and prefer to work alone but are highly effective when they take charge. They are intuitive and perceptive, often able to understand situations and people at a deeper level. However, their need for independence can make them seem distant at times, and they may struggle with relying on others.",
		"ひつじ":  "Sheep: Gentle, nurturing, and values harmony in groups. Sheep people are compassionate, kind-hearted, and always look for ways to help those around them. They are sensitive to the needs of others and excel in creating harmonious environments. While they are caring and selfless, they may sometimes struggle with asserting themselves, as they dislike conflict and confrontation. Their empathetic nature makes them excellent listeners and trusted friends.",
		"ゾウ":   "Elephant: Wise, dependable, and emotionally strong. Elephants are deeply loyal and steadfast, always there for those they care about. They possess a calm and grounded presence, providing guidance to those in need. Their wisdom comes from experience, and they are known to approach challenges with patience and thoughtfulness. However, they may be seen as slow to act, as they prefer to carefully consider all options before making decisions. They are nurturing and protective, making them great caregivers.",
		"黒ひょう": "Black Panther: Charismatic and stylish, thrives on individuality. People with the Black Panther personality are confident, stylish, and love to stand out. They have a magnetic presence that draws others to them, but they value their independence and may not always conform to societal expectations. Their charm and mystery often captivate those around them, though they can sometimes be perceived as aloof or secretive. They are highly creative and unafraid to take risks to pursue their goals.",
		"こじか":  "Fawn: Innocent and pure, a seeker of peace and affection. Those born under the sign of the Fawn are gentle and sensitive individuals who seek emotional connection and harmony in their relationships. They are soft-spoken and have a pure, innocent outlook on life, often preferring peaceful surroundings. Fawns are deeply affectionate and care about the well-being of others, but their trusting nature can sometimes make them vulnerable to being hurt. They thrive in nurturing and loving environments.",
		"たぬき":  "Raccoon Dog: Adaptable, resourceful, and full of surprises. Raccoon Dog people are incredibly flexible and quick to adapt to any situation. They are resourceful, often finding creative solutions to problems that others might miss. They enjoy novelty and variety in their lives, which keeps them on their toes. While they are clever and versatile, they may sometimes appear unpredictable, as their curiosity often leads them down unexpected paths. Their playful spirit and quick wit make them a joy to be around.",
		"猿":    "Monkey: Playful, energetic, and highly intelligent. People born under the Monkey sign are curious, quick-witted, and often the life of the party. They thrive on intellectual challenges and are always looking for new ways to solve problems. Monkeys are sociable and love to interact with others, using their humor and intelligence to make a lasting impression. While they are highly adaptable, they can sometimes be seen as restless or unpredictable, as they seek constant stimulation and excitement.",
		"ライオン": "Lion: Bold, courageous, and naturally commands respect. Lions are natural-born leaders, full of energy, confidence, and determination. They are charismatic and have the ability to inspire and motivate others with their vision. People with the Lion personality tend to be ambitious and always strive for success. While they are generous and protective of their loved ones, they can sometimes be perceived as domineering or stubborn due to their strong will and desire for control.",
		"子守熊":  "Koala: Easy-going, nurturing, and values comfort and security. Koala people are calm, caring, and laid-back. They appreciate a comfortable, secure lifestyle and thrive when surrounded by tranquility. They are deeply nurturing, offering support to friends and family without expecting much in return. However, their peaceful nature can sometimes make them resistant to change or challenges, as they prefer stability and familiarity. Koalas are sensitive and compassionate, often putting others' needs before their own.",
		"虎":    "Tiger: Fierce, determined, and unafraid to take risks. Tiger people are bold, confident, and have an unstoppable drive. They are fiercely independent and not afraid to take risks in order to achieve their goals. Tigers are natural leaders who inspire others with their energy and passion. While their courage is admirable, they can sometimes come off as too intense or aggressive, as they are always striving to reach the top. They are competitive and thrive in situations where they can challenge themselves.",
		"ペガサス": "Pegasus: Dreamy, creative, and unbound by limitations. Pegasus individuals are imaginative and visionary, often thinking outside the box. They are idealistic and have a strong desire to explore new possibilities, unencumbered by conventional limits. Their creativity knows no bounds, and they are drawn to artistic or unconventional pursuits. However, their dreamy nature can sometimes make them seem detached or impractical, as they are more focused on their visions than the everyday realities of life.",
		"チータ":  "Cheetah: Fast, focused, and thrives under pressure. People with the Cheetah personality are quick thinkers and action-oriented individuals who excel in fast-paced environments. They are highly focused on their goals and able to maintain clarity under pressure, making them excellent problem solvers. Cheetahs are determined and driven, always ready to seize opportunities and take action. While their speed and focus are impressive, they can sometimes be impatient and struggle with situations that require slow, steady progress.",
	}
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
		animalEN := animalDescriptions[animalJP]

		fmt.Println(softGreen(fmt.Sprintf("Your %s(%s): %s(%s)", categoryJP, categoryEN, animalJP, animalsEn[animalJP])))
		fmt.Println(softYellow("Description"), fmt.Sprint(animalEN))
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
