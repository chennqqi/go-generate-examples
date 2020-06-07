package assets

import (
	"net/http"
)

// Assets contains project assets.
var ImageAssets http.FileSystem = http.Dir("assets/images")
