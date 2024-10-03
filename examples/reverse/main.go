package main

import (
	"context"
	"fmt"

	"github.com/turistikrota/osm"
)

func main() {
	ctx := context.Background()
	result, err := osm.Reverse(ctx, 40.748817, -73.985428) // Latitude and Longitude for the Empire State Building
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Reverse Geocoding Result:", result)
}
