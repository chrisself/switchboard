package catalog

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"
)

const PatchFileExtension = ".json"

type Patch struct {
	// The name of the patch.
	//
	// Examples: "Aluphone", "Celli", "Clarinets a3"
	Name string

	// The path of the directory in which the patch is located, excluding the
	// root directory itself.
	//
	// Example: "Spitfire Audio/BBC Symphony Orchestra/Brass/"
	DirectoryPath string

	// The patch's keyswitch mappings.
	Mappings []Mapping
}

// Maps a named articulation to a keyswitch.
type Mapping struct {
	// The articulation name.
	Name string `json:"name"`

	// The keyswitch.
	Keyswitch Keyswitch
}

type Keyswitch struct {
	// The name of the note.
	//
	// Examples: "C", "F#", "Ab"
	Note string `json:"note"`

	// The note's octave. See package midi for middle C convention details.
	Octave int `json:"octave"`
}

// LoadPatches searches the provided filesystem for patch files and loads any
// contained articulation mappings.
func LoadPatches(filesystem fs.FS) ([]Patch, error) {
	var patches []Patch
	var mappingsCount = 0

	fn := func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if entry.IsDir() {
			return nil
		}

		if filepath.Ext(path) != PatchFileExtension {
			return nil
		}

		file, err := filesystem.Open(path)
		if err != nil {
			return err
		}

		var mappings []Mapping
		if err := json.NewDecoder(file).Decode(&mappings); err != nil {
			return err
		}
		mappingsCount += len(mappings)

		patches = append(patches, Patch{
			Name:          getPatchName(path),
			DirectoryPath: getDirectoryPath(path),
			Mappings:      mappings,
		})

		return nil
	}

	if err := fs.WalkDir(filesystem, ".", fn); err != nil {
		return []Patch{}, err
	}

	fmt.Printf("switchboard: loaded %d patches containing %d mappings\n",
		len(patches), mappingsCount)

	return patches, nil
}

func getPatchName(path string) string {
	filename := filepath.Base(path)
	extension := filepath.Ext(path)

	return strings.TrimSuffix(filename, extension)
}

func getDirectoryPath(path string) string {
	filename := filepath.Base(path)
	return strings.TrimSuffix(path, filename)
}
