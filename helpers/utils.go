package helpers

import "os"

// Location returns the location from the command line arguments (if present)
func Location() (location string) {

	location = "Johannesburg"

	if len(os.Args) >= 2 {
		location = os.Args[1]
	}

	return location
}
