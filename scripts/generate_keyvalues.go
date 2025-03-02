package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const (
	numberOfPairs  = 1000 // Total number of key-value pairs to generate
	duplicateRate  = 0.3  // 30% chance of generating a duplicate key
	maxKeyLength   = 10   // Maximum length of generated keys
	maxValueLength = 20   // Maximum length of generated values
	outputFile     = "data/data.txt"
)

// generateRandomString creates a random string of specified length
func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}

// generateKeyValuePair creates a random key-value pair
func generateKeyValuePair(existingKeys []string) (string, string) {
	// Decide whether to use an existing key (duplicate) or create a new one
	if len(existingKeys) > 0 && rand.Float64() < duplicateRate {
		return existingKeys[rand.Intn(len(existingKeys))], generateRandomString(rand.Intn(maxValueLength) + 1)
	}

	return generateRandomString(rand.Intn(maxKeyLength) + 1), generateRandomString(rand.Intn(maxValueLength) + 1)
}

func main() {
	// Initialize random seed
	rand.New(rand.NewSource(time.Now().UnixNano()))

	// Create or truncate the output file
	file, err := os.Create(outputFile)
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return
	}
	defer file.Close()

	var existingKeys []string
	var pairs []string

	// Generate key-value pairs
	for i := 0; i < numberOfPairs; i++ {
		key, value := generateKeyValuePair(existingKeys)
		pair := fmt.Sprintf("\"%s\"=\"%s\"", key, value)
		pairs = append(pairs, pair)

		// Only store new keys for potential duplication
		if !contains(existingKeys, key) {
			existingKeys = append(existingKeys, key)
		}
	}

	// Write to file
	content := strings.Join(pairs, "\n")
	_, err = file.WriteString(content)
	if err != nil {
		fmt.Printf("Error writing to file: %v\n", err)
		return
	}

	fmt.Printf("Successfully generated %d key-value pairs in %s\n", numberOfPairs, outputFile)

	// Print statistics
	uniqueKeys := len(existingKeys)
	duplicateCount := numberOfPairs - uniqueKeys
	fmt.Printf("Unique keys: %d\n", uniqueKeys)
	fmt.Printf("Duplicate keys: %d\n", duplicateCount)
	fmt.Printf("Duplication rate: %.2f%%\n", float64(duplicateCount)/float64(numberOfPairs)*100)
}

// contains checks if a string exists in a slice
func contains(slice []string, str string) bool {
	for _, v := range slice {
		if v == str {
			return true
		}
	}
	return false
}
