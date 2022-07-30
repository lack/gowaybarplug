package gowaybarplug

import (
	"encoding/json"
	"fmt"
)

// Status represents the json structure expected from waybar custom plugins.
// See waybar-custom(5) for more details.
type Status struct {
	// Text is usually the label, represented by {} in a waybar format
	Text string `json:"text"`
	// Tooltip appears when you mouse-hover over the custom entry
	Tooltip string `json:"tooltip,omitempty"`
	// Class is a list of css classes that will be added to the waybat entry
	Class []string `json:"class,omitempty"`
	// Percentage can be added to format strings via the {percent} format string, but can also affect which icon is set in {icon} if the config specifies format-icons as an array.
	Percentage *int `json:"percentage,omitempty"`
	// Alt os the key used to look up the {icon} of format-icons is specified as a map.
	Alt string `json:"alt,omitempty"`
}

// String renders the Status as a json string
func (s *Status) String() string {
	b, err := json.Marshal(s)
	if err != nil {
		// Fake json error reporting
		return fmt.Sprintf(`"text": "Marshal error", "tooltip": "Marshal error: %s"}`, err.Error())
	}
	return string(b)
}
