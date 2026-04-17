package logic

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	_ "embed"

	"github.com/chrisself/switchboard/internal/catalog"
	"github.com/chrisself/switchboard/internal/midi"
)

const plistFileExt = ".plist"

//go:embed plist.tmpl
var templateContent string

// The root data structure applied to the .plist template.
type articulationSet struct {
	// An articulation set .plist stores its own filename but Logic does not
	// appear to use it.
	Filename string

	Articulations []articulation
}

// An individual articulation. Some relevant observations regarding the workings
// of articulation sets in Logic:
//
//  1. ArticulationID is the value assigned to MIDI note events. When a user
//     creates a new articulation set the values are by default one-indexed,
//     but can be changed as desired.
//  2. ID does not appear in the application and is not user-assignable. When
//     a user creates a new articulation set the values are generated starting
//     from 1001.
//  3. Selecting an articulation set using the track inspector copies the set
//     into the project data. Any subsequent changes to the articulation data
//     are made only to the project unless the user also saves the updated set
//     back to disk via the track inspector.
//  4. Logic does not automatically reload articulation set data from disk. Any
//     changes made outside of the applications requires reselecting the same
//     set via the track inspector.
type articulation struct {
	// The identifier assigned to MIDI note events.
	ArticulationID int

	// A logic-internal identifier.
	ID int

	// The name of the articulation, appearing in the articulation set editor,
	// piano roll, note event, etc.
	Name string

	// The MIDI note number.
	Note int
}

func Generate(patches []catalog.Patch, buildDir string) error {
	tmpl, err := template.New("plist.tmpl").Parse(templateContent)
	if err != nil {
		return fmt.Errorf("failed to parse template: %v", err)
	}

	for _, patch := range patches {
		directory := filepath.Join(buildDir, patch.DirectoryPath)

		if err := os.MkdirAll(directory, 0700); err != nil {
			return fmt.Errorf("failed to create output directory: %v", err)
		}

		filename := patch.Name + plistFileExt

		articulationSet := articulationSet{
			Filename:      filename,
			Articulations: convertMappings(patch.Mappings),
		}

		file, err := os.Create(filepath.Join(directory, filename))
		if err != nil {
			return fmt.Errorf("failed to create file: %v", err)
		}

		if err := tmpl.Execute(file, articulationSet); err != nil {
			return fmt.Errorf("failed to execute template: %v", err)
		}
	}

	fmt.Printf("switchboard: generated %d articulation sets\n", len(patches))
	return nil
}

func convertMappings(mappings []catalog.Mapping) []articulation {
	articulations := make([]articulation, len(mappings))

	// For now, assign identifiers to articulations in a given patch using their
	// ordinal positions. This means articulations cannot be reordered without
	// changing the articulations existing MIDI note events in a project refer
	// to. Eventually the patches should be rewritten so that each articulation
	// mapping has a persistent identifier.
	for index, mapping := range mappings {
		articulationID := index + 1

		articulation := articulation{
			ArticulationID: articulationID,
			ID:             articulationID + 1000,
			Name:           mapping.Name,
			Note:           midi.NoteToNum(mapping.Keyswitch.Note, mapping.Keyswitch.Octave),
		}
		articulations[index] = articulation
	}

	return articulations
}
