package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println(time.Now())
	var month Month
	bytes, err := os.ReadFile("2020_OCTOBER.json")
	if err != nil {
		panic("Could not read file")
	}
	err = json.Unmarshal(bytes, &month)
	if err != nil {
		panic("Could not convert from JSON")
	}

	i := 0
	for _, _ = range month.TimelineObjects {
		// fmt.Printf("timelineObject: %v\n", timelineObject.PlaceVisit.Location.Address)
		i++
	}
	fmt.Println(i)
	fmt.Println(time.Now())
}
