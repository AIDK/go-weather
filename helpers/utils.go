package helpers

import "os"

// Location returns the location from the command line arguments (if present)
func Location() (location string) {

	// default location
	location = "Cape%20Town"

	// we check if the length of the os.Args slice is greater than or equal to 2
	// if it is, we assign the second element of the slice to the location variable
	// the first element of the slice is the name of the program
	if len(os.Args) >= 2 {
		location = os.Args[1]
	}

	return location
}
