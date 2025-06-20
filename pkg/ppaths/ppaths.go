package ppaths

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

// GameID is used to name your config/save folders
const GameID = "drawn_by_fate"

// GetUserGameDir returns the OS-appropriate directory for user-specific data
func GetUserGameDir() (string, error) {
	switch runtime.GOOS {
	case "windows":
		if dir := os.Getenv("LOCALAPPDATA"); dir != "" {
			return filepath.Join(dir, GameID), nil
		}
		if dir := os.Getenv("APPDATA"); dir != "" {
			return filepath.Join(dir, GameID), nil
		}
		return "", errors.New("unable to determine Windows data directory")
	case "darwin":
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		return filepath.Join(home, "Library", "Application Support", GameID), nil
	default:
		// Linux and other Unix-like systems
		if dir := os.Getenv("XDG_DATA_HOME"); dir != "" {
			return filepath.Join(dir, GameID), nil
		}
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		return filepath.Join(home, ".local", "share", GameID), nil
	}
}

func GetModuleRoot(name string) (string, error) {
	path, err := GetUserGameDir()
	if err != nil {
		return "", err
	}

	if name == "core" {
		path = filepath.Join(path, name)
	} else {
		path = filepath.Join(path, "mods", name)
	}

	return path, nil
}

// JoinUserDirPath returns a full file path under the user data directory
func JoinUserDirPath(pathParts ...string) (string, error) {
	dir, err := GetUserGameDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(append([]string{dir}, pathParts...)...), nil
}

// EnsureDataDir makes sure the user data directory exists
func EnsureDataDir() (string, error) {
	dir, err := GetUserGameDir()
	if err != nil {
		return "", err
	}
	dirs := []string{
		dir,
		filepath.Join(dir, "core"),
		filepath.Join(dir, "mods"),
	}
	for _, path := range dirs {
		err = os.MkdirAll(path, 0755)
		if err != nil {
			return "", err
		}
	}
	return dir, nil
}

// AssetExists checks if an asset exists at the given relative path from base assets dir
func AssetExists(baseDir, relPath string) bool {
	fullPath := filepath.Join(baseDir, relPath)
	if _, err := os.Stat(fullPath); err == nil {
		return true
	}
	return false
}

// IsWritable tests whether a file or directory is writable
func IsWritable(path string) bool {
	file, err := os.CreateTemp(path, ".writetest")
	if err != nil {
		return false
	}
	file.Close()
	os.Remove(file.Name())
	return true
}

// DebugPrintPaths is a helper for logging important paths
func DebugPrintPaths() {
	userData, err := GetUserGameDir()
	if err != nil {
		fmt.Println("[Error] Getting data dir:", err)
	} else {
		fmt.Println("User Data Dir:", userData)
	}
	cwd, _ := os.Getwd()
	fmt.Println("Working Dir:", cwd)
	fmt.Println("Runtime OS:", runtime.GOOS)
}
