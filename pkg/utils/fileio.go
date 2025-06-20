package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

// Generic function to read JSON from a file into a given Go struct
func ReadJSON[T any](path string, objs ...T) (*T, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read %s: %w", path, err)
	}

	var obj T
	if len(objs) > 0 {
		obj = objs[0]
	}

	if err := json.Unmarshal(data, &obj); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON from %s: %w", path, err)
	}

	return &obj, nil
}

// Generic function to write a Go struct as JSON to a file
func WriteJSON[T any](path string, obj T) error {
	data, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}

	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return fmt.Errorf("failed to create directories for %s: %w", path, err)
	}

	if err := os.WriteFile(path, data, 0644); err != nil {
		return fmt.Errorf("failed to write JSON to %s: %w", path, err)
	}

	return nil
}

func WriteToFile(data string, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(data)
	if err != nil {
		return err
	}
	return nil
}

func WriteBytesToFile(data []byte, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return err
	}
	return nil
}

// PathExists is a helper function that checks if a given path exists
func PathExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// IsDir is a helper function that checks if the given path is a directory
func IsDir(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}

func IsValidDir(path string) error {
	if !PathExists(path) {
		return fmt.Errorf("Error, path: '%s' does not exist", path)
	}

	if !IsDir(path) {
		return fmt.Errorf("Error, path: '%s' is not a directory", path)
	}
	return nil
}

// CheckIncludeExclude checks if a given value matches any regular expression in the includes slice
// and matches none in the excludes slice. Returns true if it should be included, false otherwise.
func CheckIncludeExclude(value string, includes, excludes []string) (bool, error) {

	// Check matches against exclude patterns first
	if matched, err := MatchStrPatterns(value, excludes...); err != nil {
		return false, fmt.Errorf("invalid exclude regex: %v", err)
	} else if matched {
		return false, nil
	}

	// Check matches against include patterns
	if matched, err := MatchStrPatterns(value, includes...); err != nil {
		return false, fmt.Errorf("invalid include regex: %v", err)
	} else if matched {
		return true, nil
	}

	return false, nil
}

func MatchStrPatterns(value string, patterns ...string) (bool, error) {
	for _, pattern := range patterns {
		matched, err := regexp.MatchString(pattern, value)
		if err != nil {
			return false, fmt.Errorf("invalid regex pattern: %v", err)
		}
		if matched {
			return true, nil
		}
	}
	return len(patterns) == 0, nil
}
