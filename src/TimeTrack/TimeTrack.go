package TimeTrack

import (
	"fmt"
	"time"
)

func TimeTrack(start time.Time, name string) {
	finish := time.Since(start)
	fmt.Printf("%s took %s", name, finish)
}
