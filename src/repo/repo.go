package repo

import (
	"log"
	"sync"
)

var initDone = false
var initLock sync.Mutex

func EnsureResourcesLoaded(dataDir string) {
	if initDone {
		return
	}

	log.Print("Acquiring resource mutex")
	initLock.Lock()
	defer initLock.Unlock()

	if initDone {
		log.Print("Resources were loaded by another request")
		return
	}

	log.Print("Loading resources")
	ParseHskFiles(dataDir)
	// parseMandarinCompanionFiles(dataDir)
	ParseCcCeDict(dataDir)
	ParseWordFrequencyFile(dataDir)
	ParseCharFrequencyFile(dataDir)
	ParseCharacterCompositionFile(dataDir)
	initDone = true
	log.Print("Resources loaded")
}
