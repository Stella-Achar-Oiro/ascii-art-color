package main

import (
	colour "asci-art/color/functions"
	"os"
	"reflect"
	"testing"
)

// TestLoadBanner tests the LoadBanner function to ensure it correctly loads banner characters from a file.

// TestColorCheck tests the ColorCheck function for various input scenarios.
func TestColorCheck(t *testing.T) {
	tests := []struct {
		name           string
		colors         []string
		expectedBool   bool
		expectedColors []string
	}{
		{
			name:           "empty input",
			colors:         []string{},
			expectedBool:   false,
			expectedColors: []string{},
		},
		{
			name:           "single valid color",
			colors:         []string{"#FFFFFF"},
			expectedBool:   false,
			expectedColors: []string{"#FFFFFF"},
		},
		{
			name:           "multiple valid colors",
			colors:         []string{"#FFFFFF", "#000000"},
			expectedBool:   true,
			expectedColors: []string{"#FFFFFF", "#000000"},
		},
		{
			name:           "empty input and color included",
			colors:         []string{"#FFFFFF", "", "#000000"},
			expectedBool:   true,                           // This might need to be adjusted based on how you handle errors in RgbExtract.
			expectedColors: []string{"#FFFFFF", "#000000"}, // Assuming RgbExtract skips invalid colors.
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotBool, gotColors := colour.ColorCheck(tt.colors)
			if gotBool != tt.expectedBool {
				t.Errorf("ColorCheck() gotBool = %v, want %v", gotBool, tt.expectedBool)
			}
			if !equalSlices(gotColors, tt.expectedColors) {
				t.Errorf("ColorCheck() gotColors = %v, want %v", gotColors, tt.expectedColors)
			}
		})
	}
}

// equalSlices checks if two slices of strings are equal.
func equalSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

// mockArgs simulates command line arguments for testing.
func mockArgs(input string) {
	os.Args = []string{"cmd", "subcmd", input}
}

// TestColorAssign tests the ColorAssign function with predefined inputs.
func TestColorAssign(t *testing.T) {
	// Setup mock for os.Args to avoid dependency in tests.
	originalArgs := os.Args
	defer func() { os.Args = originalArgs }() // Ensure os.Args is restored after tests.

	tests := []struct {
		name    string
		colors  []string
		arg     string
		want    map[rune]string
		wantErr bool
	}{
		{
			name:   "single color",
			colors: []string{"#FFFFFF"},
			arg:    "test",
			want: map[rune]string{
				't': "#FFFFFF",
				'e': "#FFFFFF",
				's': "#FFFFFF",
			},
			wantErr: false,
		},
		{
			name:   "multiple colors",
			colors: []string{"#FFFFFF", "red", "blue"},
			arg:    "hey ray",
			want: map[rune]string{
				'h': "#FFFFFF",
				'e': "red",
				'y': "blue",
			},
			wantErr: false,
		},
		// Add more test cases as needed.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockArgs(tt.arg) // Mocking os.Args with test-specific input.
			got := colour.ColorAssighn(tt.colors)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ColorAssign() = %v, want %v", got, tt.want)
			}
		})
	}
}
