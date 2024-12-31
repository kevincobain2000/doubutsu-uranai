package main

import (
	"testing"
)

func TestParseCSV(t *testing.T) {
	mockCSV := `F,1800-01-05,希望キャラ：狼
F,1800-01-05,意思決定キャラ：ひつじ
M,1800-01-06,本質キャラ：ゾウ`

	expected := [][]string{
		{"F", "1800-01-05", "希望キャラ：狼"},
		{"F", "1800-01-05", "意思決定キャラ：ひつじ"},
		{"M", "1800-01-06", "本質キャラ：ゾウ"},
	}

	result := parseCSV(mockCSV)

	if len(result) != len(expected) {
		t.Fatalf("Expected %d records, got %d", len(expected), len(result))
	}

	for i, record := range result {
		for j, field := range record {
			if field != expected[i][j] {
				t.Errorf("Expected field %s, got %s", expected[i][j], field)
			}
		}
	}
}

func TestFilterRecords(t *testing.T) {
	records := [][]string{
		{"F", "1800-01-05", "希望キャラ：狼"},
		{"F", "1800-01-05", "意思決定キャラ：ひつじ"},
		{"M", "1800-01-06", "本質キャラ：ゾウ"},
	}

	tests := []struct {
		gender   string
		dob      string
		expected int
	}{
		{"F", "1800-01-05", 2},
		{"M", "1800-01-06", 1},
		{"F", "1800-01-06", 0},
	}

	for _, tt := range tests {
		result := filterRecords(records, tt.gender, tt.dob)
		if len(result) != tt.expected {
			t.Errorf("For gender %s and DOB %s, expected %d records, got %d", tt.gender, tt.dob, tt.expected, len(result))
		}
	}
}

func TestAnimalDescriptions(t *testing.T) {
	animalDescriptions := map[string]string{
		"狼":   "Wolf: Independent and confident, a natural leader.",
		"ひつじ": "Sheep: Gentle, nurturing, and values harmony in groups.",
		"ゾウ":  "Elephant: Wise, dependable, and emotionally strong.",
	}

	tests := []struct {
		animal    string
		expected  string
		shouldErr bool
	}{
		{"狼", "Wolf: Independent and confident, a natural leader.", false},
		{"ひつじ", "Sheep: Gentle, nurturing, and values harmony in groups.", false},
		{"未知", "", true}, // "未知" means "unknown"
	}

	for _, tt := range tests {
		desc, ok := animalDescriptions[tt.animal]
		if !ok && !tt.shouldErr {
			t.Errorf("Expected description for animal %s, got none", tt.animal)
		} else if desc != tt.expected && !tt.shouldErr {
			t.Errorf("For animal %s, expected '%s', got '%s'", tt.animal, tt.expected, desc)
		}
	}
}
