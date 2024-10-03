package main

import (
	"context"
	"fmt"

	"github.com/turistikrota/osm"
)

func main() {
	ctx := context.Background()
	result, err := osm.Lookup(ctx, "way", 123456789) // Example OSM type and ID
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Lookup Result:", result)
}
