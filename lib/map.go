package lib

import "fmt"

func Map() {

	type Vertex struct {
		Lat, Long float64
	}

	m := make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	m["Google"] = Vertex{
		37.42202, -122.08408,
	}

	mx := map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}

	mx["Apple"] = Vertex{10.68433, -74.39967}

	fmt.Println(len(m))

	fmt.Println(m)
	fmt.Println("mx", mx)

}
