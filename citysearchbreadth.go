// citysearchbreadth.go
// Author: Hannes du Plooy
// Revision Date: 3 Sep 2016
// Implements BreadthFirstSearch of hduplooy/gosearch to search for a road trip from one city to another
// BreadthFirst will search for the least steps but not necessarily the shortest real distance
package main

import (
	"fmt"

	src "github.com/hduplooy/gosearch"
)

// City represents our state
// Name is the name of the city
// Neighbours are all the cities that can be reached directly
type City struct {
	Name       string
	Neighbours []*City
}

// Descendants really just return all the neighbours
func (city *City) Descendants() []src.SearchF {
	tmp := make([]src.SearchF, len(city.Neighbours))
	for i, val := range city.Neighbours {
		tmp[i] = val
	}
	return tmp
}

// Done returns when the current state is "Cape Town" which is our goal state in this instance
func (city *City) Done() bool {
	return city.Name == "Cape Town"
}

// Cost is not used
func (city *City) Cost() float64 { return 0.0 }

// Away is not used
func (city *City) Away() float64 { return 0.0 }

// Key is just the name of the city
func (city *City) Key() string {
	return city.Name
}

// Stringer func for City
func (city *City) String() string {
	return city.Name
}

// A map to the cities that are in our database
var cities = make(map[string]*City)

// AddRoad will link two cities by making sure that they are in the database
// and then add them to the other's neighbours
func AddRoad(city1, city2 string) {
	c1, ok := cities[city1]
	// If not already in database then put it there
	if !ok {
		c1 = &City{Name: city1}
		cities[city1] = c1
	}
	c2, ok := cities[city2]
	// If not already in database then put it there
	if !ok {
		c2 = &City{Name: city2}
		cities[city2] = c2
	}
	c1.Neighbours = append(c1.Neighbours, c2)
	c2.Neighbours = append(c2.Neighbours, c1)
}

// Add all the direct links between cities
func init() {
	AddRoad("Pretoria", "Johannesburg")
	AddRoad("Pretoria", "Midrand")
	AddRoad("Midrand", "Johannesburg")
	AddRoad("Pretoria", "Kempton")
	AddRoad("Johannesburg", "Kempton")
	AddRoad("Johannesburg", "Klerksdorp")
	AddRoad("Klerksdorp", "Potchefstroom")
	AddRoad("Potchefstroom", "Kimberley")
	AddRoad("Johannesburg", "Vanderbijl")
	AddRoad("Vanderbijl", "Sasolburg")
	AddRoad("Johannesburg", "Vereeniging")
	AddRoad("Vereeniging", "Sasolburg")
	AddRoad("Johannesburg", "Kroonstad")
	AddRoad("Sasolburg", "Kroonstad")
	AddRoad("Kroonstad", "Ventersburg")
	AddRoad("Ventersburg", "Bloemfontein")
	AddRoad("Bloemfontein", "Kimberley")
	AddRoad("Bloemfontein", "Beaufort West")
	AddRoad("Kimberley", "Beaufort West")
	AddRoad("Beaufort West", "Worcester")
	AddRoad("Worcester", "Cape Town")
	AddRoad("Beaufort West", "George")
	AddRoad("George", "Cape Town")
}

func main() {
	// We are going to look for a path from Pretoria to Cape Town
	city := cities["Pretoria"]
	cnt, ans, hist := src.BreadthFirstSearch(city, true)
	fmt.Printf("Done in %d steps\n", cnt)
	for i := len(hist) - 1; i >= 0; i-- {
		fmt.Printf("%v\n", hist[i])
	}
	fmt.Printf("%v\n", ans)
}
