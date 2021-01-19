// +build !binary_log

package diode_test

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/diode"
)

func ExampleNewWriter() {
	// The default global level is Trace
	globalLevel := zerolog.GlobalLevel()
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	defer zerolog.SetGlobalLevel(globalLevel)

	w := diode.NewWriter(os.Stdout, 1000, 0, func(missed int) {
		fmt.Printf("Dropped %d messages\n", missed)
	})
	log := zerolog.New(w)
	log.Print("test")

	w.Close()

	// Output: {"level":"debug","message":"test"}
}
