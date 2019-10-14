package helpers

import (
	"fmt"
	"os"
	"os/signal"
	"rare/pkg/extractor"
	"rare/pkg/multiterm"
	"sync"
	"sync/atomic"
	"time"
)

// RunAggregationLoop is a helper that takes care of output sync
// And the main async loops for you, it has two inputs (in addition to the extractor)
//   matchProcessor - to process a match
//   writeOutput - triggered after a delay, only if there's an update
// The two functions are guaranteed to never happen at the same time
func RunAggregationLoop(ext *extractor.Extractor, matchProcessor func(*extractor.Match), writeOutput func()) {

	defer multiterm.ResetCursor()

	// Updater sync variables
	outputDone := make(chan bool)
	var outputMutex sync.Mutex
	var hasUpdates atomic.Value
	hasUpdates.Store(false)

	// Updating loop
	go func() {
		for {
			select {
			case <-outputDone:
				return
			case <-time.After(100 * time.Millisecond):
				if hasUpdates.Load().(bool) {
					hasUpdates.Store(false)
					outputMutex.Lock()
					writeOutput()
					outputMutex.Unlock()
				}
			}
		}
	}()

	// Processing data from extractor
	exitSignal := make(chan os.Signal)
	signal.Notify(exitSignal, os.Interrupt)
	reader := ext.ReadChan()
PROCESSING_LOOP:
	for {
		select {
		case <-exitSignal:
			break PROCESSING_LOOP
		case match, more := <-reader:
			if !more {
				break PROCESSING_LOOP
			}
			outputMutex.Lock()
			matchProcessor(match)
			outputMutex.Unlock()
			hasUpdates.Store(true)
		}
	}
	outputDone <- true

	writeOutput()
	fmt.Println()

	WriteExtractorSummary(ext)
}
