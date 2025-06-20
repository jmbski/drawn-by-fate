package main

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

//go:embed assets/images/**
var embeddedImages embed.FS

//go:embed assets/**
var embeddedAssets embed.FS

var embedPath = "/home/joseph/coding_base/drawn-by-fate/embed_test/embedded/assets"

const (
	playerPng = "/home/joseph/coding_base/drawn-by-fate/embed_test/assets/images/sprites/player.png"
	imagesDir = "/home/joseph/coding_base/drawn-by-fate/embed_test/assets/images"
	assetsDir = "/home/joseph/coding_base/drawn-by-fate/embed_test/assets"
)

type Utils struct{}

// PathExists is a helper function that checks if a given path exists
func (u *Utils) PathExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func (u *Utils) IsDir(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}

func (u *Utils) IsValidDir(path string) error {
	if !utils.PathExists(path) {
		return fmt.Errorf("Error, path: '%s' does not exist", path)
	}

	if !utils.IsDir(path) {
		return fmt.Errorf("Error, path: '%s' is not a directory", path)
	}
	return nil
}

var utils *Utils

type PPaths struct{}

func (p *PPaths) GetUserDataDir() (string, error) {
	return "/home/joseph/coding_base/drawn-by-fate/embed_test/embedded", nil
}

var ppaths *PPaths

type DBFFileSystem interface {
	ReadDir(string) ([]fs.DirEntry, error)
	ReadFile(string) ([]byte, error)
}

type DBFDirectory struct {
	Path      string
	EmbedPath string
	Dirs      []DBFDirectory
	Files     []string
	Embedded  bool
}

func (d *DBFDirectory) GetFilePaths() []string {
	fullPaths := make([]string, 0, len(d.Files))
	for _, file := range d.Files {
		fullPaths = append(fullPaths, filepath.Join(d.Path, file))
	}
	return fullPaths
}

func (d *DBFDirectory) GetEmbedPaths() []string {
	embedPaths := make([]string, 0, len(d.Files))
	for _, file := range d.Files {
		embedPaths = append(embedPaths, filepath.Join(d.EmbedPath, file))
	}
	return embedPaths
}

func (d *DBFDirectory) PrintPaths(indent ...string) {
	indentStr := ""
	if len(indent) > 0 {
		indentStr = indent[0]
	}
	fmt.Println(indentStr + d.EmbedPath)
	for _, subDir := range d.Dirs {
		subDir.PrintPaths(indentStr + "  ")
	}
	for _, file := range d.Files {
		fmt.Println(indentStr + "  " + file)
	}
}

func (d *DBFDirectory) Embed() error {
	if !d.Embedded {
		return fmt.Errorf("directory '%s' is not embeddable", d.Path)
	}

	if !utils.PathExists(d.EmbedPath) {
		if err := os.MkdirAll(d.EmbedPath, 0755); err != nil {
			return fmt.Errorf("Failed to create directory %s: %w", d.EmbedPath, err)
		}
	}
	internalPaths := d.GetFilePaths()

	for i, path := range d.GetEmbedPaths() {
		if !utils.PathExists(path) {
			iPath := internalPaths[i]
			data, err := embeddedAssets.ReadFile(iPath)
			if err != nil {
				fmt.Printf("failed to read embedded file %s: %s\n", iPath, err)
				continue
			}
			if err := os.WriteFile(path, data, 0644); err != nil {
				fmt.Printf("failed to write embedded file %s: %s\n", iPath, err)
			}
		}
	}

	for _, subdir := range d.Dirs {
		subdir.Embed()
	}

	return nil
}

func processDirEntries(dir *DBFDirectory) (err error) {
	var entries []fs.DirEntry
	if !dir.Embedded {
		entries, err = os.ReadDir(dir.Path)
		if err != nil {
			return err
		}
	} else {
		entries, err = embeddedAssets.ReadDir(dir.Path)
		if err != nil {
			return err
		}
	}

	for _, entry := range entries {
		if entry.IsDir() {
			subDir, err := newDbfDir(filepath.Join(dir.Path, entry.Name()), dir.Embedded)
			if err != nil {
				return fmt.Errorf("failed to read subdirectory '%s': %w", entry.Name(), err)
			}
			dir.Dirs = append(dir.Dirs, *subDir)
		} else {
			dir.Files = append(dir.Files, entry.Name())
		}
	}

	return nil
}

func newDbfDir(path string, embedded bool) (*DBFDirectory, error) {

	if err := utils.IsValidDir(path); err != nil {
		return nil, err
	}

	dir := &DBFDirectory{
		Path:     path,
		Embedded: embedded,
	}

	if embedded {
		home, err := ppaths.GetUserDataDir()
		if err != nil {
			return nil, err
		}
		dir.EmbedPath = filepath.Join(home, "core", path)
	}

	if err := processDirEntries(dir); err != nil {
		return nil, err
	}
	return dir, nil
}

func NewEmbeddedDir(path string) (*DBFDirectory, error) {

	return newDbfDir(path, true)
}

func NewDBFDirectory(path string) (*DBFDirectory, error) {

	return newDbfDir(path, false)
}

func EmbedDir(dirPath, relPath string, embeddedData embed.FS) error {
	entries, err := embeddedData.ReadDir(relPath)
	if err != nil {
		return fmt.Errorf("failed to read embedded directory: %w", err)
	}

	for _, entry := range entries {
		fmt.Println("entry:", entry.Name())
	}
	return nil
}

func main() {
	//path := "/home/joseph/coding_base/_ici_temp/dbox_4-11/cws3_upload/cws3_upload"

	test := []int{}
	for _, i := range test {
		fmt.Println(i)
	}
	fmt.Println("Done")
	/* dir, err := NewEmbeddedDir("assets")
	if err != nil {
		panic(err)
	}

	dir.Embed() */
	/* entries, err := embeddedAssets.ReadDir("assets")
	if err != nil {
		panic(err)
	}
	for _, entry := range entries {
		fmt.Println(entry.Name())
	} */
	/* err = EmbedDir(filepath.Join(embedPath, "images"), "assets/images", embeddedImages)
	if err != nil {
		fmt.Println(err)
	} */
}
