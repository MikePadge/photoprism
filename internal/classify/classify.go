/*
This package encapsulates image classification using TensorFlow

Additional information can be found in our Developer Guide:

https://github.com/mikepadge/photoprism/wiki
*/
package classify

import (
	"github.com/mikepadge/photoprism/internal/event"
)

//go:generate go run gen.go

var log = event.Log
