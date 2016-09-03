// citysearchcost.go
// Author: Hannes du Plooy
// Revision Date: 3 Sep 2016
// Implements BestCostSearch of hduplooy/gosearch to search for a road trip from one city to another
// It is similar to citysearchbreadth.go except we keep track of how far we travelled
package main

import (
	"fmt"
	"strconv"
	"strings"

	src "github.com/hduplooy/gosearch"
)

// City just keeps the city information in this implementation
// Name is the name of the city
// Neighbours are all the cities that can be reached directly
// Distances are the distances in km from this city to its neighbours
// Key Just to not use a string to search for we make the key the number of when it was created (so it is unique)
type City struct {
	Name       string
	Neighbours []*City
	Distances  []float64
	Key        int
}

// This is the actual state (because we can actually go to the same city again, if we really want to)
// Includes City
// HistKey is a concatenation of the Key values of all the cities visited so far, so this is unique for each state
// TotCost is the total distance travelled so far
type CitySE struct {
	*City
	HistKey string
	TotCost float64
}

// Descendants get all the neighbours of a city
// A neighbour is only valid if it has not been visited before (it's Key does not appear in the HistKey)
func (city *CitySE) Descendants() []src.SearchF {
	tmp := make([]src.SearchF, 0, len(city.Neighbours))
	for i, val := range city.Neighbours {
		k := strconv.Itoa(val.Key)
		if strings.Index(city.HistKey, k) >= 0 {
			continue
		}
		tmp = append(tmp, &CitySE{val, city.HistKey + "-" + k, city.TotCost + city.Distances[i]})
	}
	return tmp
}

// Done is true once we reach Cape Town
func (city *CitySE) Done() bool {
	return city.Name == "Cape Town"
}

// Cost is the total distance travelled so far
func (city *CitySE) Cost() float64 { return city.TotCost }

// Away is not use
func (city *CitySE) Away() float64 { return 0.0 }

// Key is just the HistKey
func (city *CitySE) Key() string {
	return city.HistKey
}

// Stringer func for CitySE - gives name and distance
func (city *CitySE) String() string {
	return city.Name + " " + fmt.Sprintf("%.2f", city.TotCost) + "km"
}

// Our city database
var cities = make(map[string]*City)

// AddRoad add cities to database and add them to each others neighbours with distance between them
func AddRoad(city1, city2 string, dist float64) {
	c1, ok := cities[city1]
	// Add if not in database
	if !ok {
		c1 = &City{Name: city1, Key: len(cities)}
		cities[city1] = c1
	}
	c2, ok := cities[city2]
	// Add if not in database
	if !ok {
		c2 = &City{Name: city2, Key: len(cities)}
		cities[city2] = c2
	}
	// Add to neighbours and add distance
	c1.Neighbours = append(c1.Neighbours, c2)
	c1.Distances = append(c1.Distances, dist)
	c2.Neighbours = append(c2.Neighbours, c1)
	c2.Distances = append(c2.Distances, dist)
}

// init initializes our database
func init() {
	AddRoad("Pretoria", "Midrand", 28)
	AddRoad("Midrand", "Johannesburg", 25)
	AddRoad("Pretoria", "Kempton", 54)
	AddRoad("Johannesburg", "Kempton", 25)
	AddRoad("Johannesburg", "Klerksdorp", 172)
	AddRoad("Klerksdorp", "Potchefstroom", 47)
	AddRoad("Potchefstroom", "Kimberley", 358)
	AddRoad("Johannesburg", "Vanderbijl", 72)
	AddRoad("Vanderbijl", "Sasolburg", 17)
	AddRoad("Johannesburg", "Vereeniging", 63)
	AddRoad("Vereeniging", "Sasolburg", 29)
	AddRoad("Johannesburg", "Kroonstad", 190)
	AddRoad("Sasolburg", "Kroonstad", 124)
	AddRoad("Kroonstad", "Ventersburg", 52)
	AddRoad("Ventersburg", "Bloemfontein", 159)
	AddRoad("Bloemfontein", "Kimberley", 168)
	AddRoad("Bloemfontein", "Beaufort West", 570)
	AddRoad("Kimberley", "Beaufort West", 453)
	AddRoad("Beaufort West", "Worcester", 356)
	AddRoad("Worcester", "Cape Town", 111)
	AddRoad("Beaufort West", "George", 241)
	AddRoad("George", "Cape Town", 431)
}

func main() {
	// Start at Pretoria and go to Cape Town
	city := cities["Pretoria"]
	// Search for path and get history seeing that that is the cities we have to travel through
	cnt, ans, hist := src.BestCostSearch(&CitySE{city, strconv.Itoa(city.Key), 0}, true)
	fmt.Printf("Done in %d steps\n", cnt)
	for i := len(hist) - 1; i >= 0; i-- {
		fmt.Printf("%v\n", hist[i])
	}
	fmt.Printf("%v\n", ans)
}
