package dump

import (
	"github.com/Danial-Movahed/nuga-lib/device"
	"github.com/Danial-Movahed/nuga-lib/features/keys"
	"github.com/Danial-Movahed/nuga-lib/features/light"
)

// State represents device state. It contains data of all supported features
type State struct {
	Name     device.Model `json:"name"`
	Firmware string       `json:"firmware"`
	Lights   light.State  `json:"lights"`
	Keys     keys.State   `json:"keys"`
}
