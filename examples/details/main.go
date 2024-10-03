package main

import (
	"context"
	"fmt"

	"github.com/turistikrota/osm"
)

func main() {
	ctx := context.Background()
	result, err := osm.DetailsWithPlaceID(ctx, 240109189) // PlaceID for the Empire State Building
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Details Result:", result)
}
