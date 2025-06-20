package config

import (
	"drawn-by-fate/pkg/logging"
	"drawn-by-fate/pkg/ppaths"
	"drawn-by-fate/pkg/utils"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

type PackageDir struct {
	Path        string
	EmbedPath   string
	Dirs        []PackageDir
	Files       []string
	Embedded    bool
	CommonPaths CommonPaths
}

func (p *PackageDir) GetFilePaths() []string {
	fullPaths := make([]string, 0, len(p.Files))
	for _, file := range p.Files {
		fullPaths = append(fullPaths, filepath.Join(p.Path, file))
	}
	return fullPaths
}

func (p *PackageDir) GetEmbedPaths() []string {
	embedPaths := make([]string, 0, len(p.Files))
	for _, file := range p.Files {
		embedPaths = append(embedPaths, filepath.Join(p.EmbedPath, file))
	}
	return embedPaths
}

func (p *PackageDir) PrintPaths(indent ...string) {
	indentStr := ""
	if len(indent) > 0 {
		indentStr = indent[0]
	}
	fmt.Println(indentStr + p.EmbedPath)
	for _, subDir := range p.Dirs {
		subDir.PrintPaths(indentStr + "  ")
	}
	for _, file := range p.Files {
		fmt.Println(indentStr + "  " + file)
	}
}

func (p *PackageDir) Embed(override ...bool) error {
	overWrite := len(override) > 0 && override[0]
	if !p.Embedded {
		return fmt.Errorf("directory '%s' is not embeddable", p.Path)
	}

	if !utils.PathExists(p.Path) {
		if err := os.MkdirAll(p.Path, 0755); err != nil {
			return fmt.Errorf("Failed to create directory %s: %w", p.Path, err)
		}
	}
	embedPaths := p.GetEmbedPaths()

	for i, path := range p.GetFilePaths() {
		if !utils.PathExists(path) || overWrite {
			ePath := embedPaths[i]
			data, err := embeddedPackage.ReadFile(ePath)
			if err != nil {
				fmt.Printf("failed to read embedded file %s: %s\n", ePath, err)
				continue
			}

			if err := os.WriteFile(path, data, 0644); err != nil {
				fmt.Printf("failed to write embedded file %s: %s\n", ePath, err)
			}
		}
	}

	for _, subdir := range p.Dirs {
		if err := subdir.Embed(overWrite); err != nil {
			return err
		}
	}

	return nil
}

func (p *PackageDir) CreateSubdir(dirName string) {
	subdir := PackageDir{
		Path:      filepath.Join(p.Path, dirName),
		EmbedPath: filepath.Join(p.EmbedPath, dirName),
		Embedded:  p.Embedded,
	}
	if err := processDirEntries(&subdir); err != nil {
		logging.LogWarning("Dir not processed", dirName)
	} else {
		p.Dirs = append(p.Dirs, subdir)
	}
}

func processDirEntries(dir *PackageDir, ignored ...string) (err error) {
	var entries []fs.DirEntry
	if !dir.Embedded {
		entries, err = os.ReadDir(dir.Path)
		if err != nil {
			return err
		}
	} else {
		entries, err = embeddedPackage.ReadDir(dir.EmbedPath)
		if err != nil {
			return err
		}
	}

	for _, entry := range entries {
		if matched, err := utils.CheckIncludeExclude(entry.Name(), []string{}, ignored); err != nil {
			logging.LogWarning(err)
		} else if matched {
			logging.LogDebug("ignored", entry.Name())
			continue
		}

		if entry.IsDir() {
			/* subDir, err := NewPackageDir(filepath.Join(dir.Path, entry.Name()), ignored...)
			if err != nil {
				return fmt.Errorf("failed to read subdirectory '%s': %w", entry.Name(), err)
			}
			dir.Dirs = append(dir.Dirs, *subDir) */
			dir.CreateSubdir(entry.Name())
		} else {
			dir.Files = append(dir.Files, entry.Name())
		}
	}

	return nil
}

func NewPackageDir(name string, ignored ...string) (*PackageDir, error) {

	path, err := ppaths.GetModuleRoot(name)
	if err != nil {
		return nil, err
	}

	embedded := name == "core"

	cmnPaths, err := NewCommonPaths(path, !embedded)
	if err != nil {
		return nil, err
	}

	dir := &PackageDir{
		Path:        path,
		Embedded:    embedded,
		CommonPaths: cmnPaths,
	}

	if embedded {
		dir.EmbedPath = "package"
	}

	logging.LogDebug("package path:", path, "embed path", dir.EmbedPath)

	if err := processDirEntries(dir, ignored...); err != nil {
		return nil, err
	}
	return dir, nil
}
