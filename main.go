package main

import (
	"log"
	"os"

	"github.com/chrisself/switchboard/internal/library"
	"github.com/chrisself/switchboard/internal/logic"
)

// The directory where the patch articulation data is stored.
var libraryRoot = os.DirFS("library")

// The directory where the generated articulation sets are written.
const buildDir = "generated"

func main() {
	patches, err := library.LoadPatches(libraryRoot)
	if err != nil {
		log.Fatalf("switchboard: failed to load patches: %v", err)
	}

	if err := logic.Generate(patches, buildDir); err != nil {
		log.Fatalf("switchboard: failed to generate articulation sets: %v", err)
	}
}
