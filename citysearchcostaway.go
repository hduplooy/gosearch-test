// citysearchbreadth.go
// Author: Hannes du Plooy
// Revision Date: 3 Sep 2016
// Implements BestCostAwaySearch of hduplooy/gosearch to search for a road trip from one city to another
// It is similar to citysearchcost.go except we keep track of how far we travelled and how far away we are from the goal
package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	src "github.com/hduplooy/gosearch"
)

// City just keeps the city information in this implementation
// Name is the name of the city
// Neighbours are all the cities that can be reached directly
// Distances are the distances in km from this city to its neighbours
// Key Just to not use a string to search for we make the key the number of when it was created (so it is unique)
// Latitude and Longitude is the actual geo coordinates of the cities
// Away is the direct distance from this city to the goal
type City struct {
	Name       string
	Neighbours []*City
	Distances  []float64
	Latitude   float64
	Longitude  float64
	Key        int
	Away       float64
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
	return city.Name == destination
}

// Cost is the total distance travelled so far
func (city *CitySE) Cost() float64 { return city.TotCost }

// The distance to the destination city is returned (see Distance later in the code)
func (city *CitySE) Away() float64 {
	return city.Distance(cities[destination])
}

// Return the unique HistKey for the state
func (city *CitySE) Key() string {
	return city.HistKey
}

// Stringer func for CitySE
func (city *CitySE) String() string {
	return city.Name + " " + fmt.Sprintf("%.2f", city.TotCost) + "km"
}

// City database
var cities = make(map[string]*City)

// AddRoad will add cities if not in database and to each others neighbours as well as distance to each other
func AddRoad(city1, city2 string, dist float64) {
	c1, ok := cities[city1]
	if !ok {
		c1 = &City{Name: city1, Key: len(cities)}
		cities[city1] = c1
	}
	c2, ok := cities[city2]
	if !ok {
		c2 = &City{Name: city2, Key: len(cities)}
		cities[city2] = c2
	}
	c1.Neighbours = append(c1.Neighbours, c2)
	c1.Distances = append(c1.Distances, dist)
	c2.Neighbours = append(c2.Neighbours, c1)
	c2.Distances = append(c2.Distances, dist)
}

// Set the geo coordinates for a city
func SetCoords(city string, lat, long float64) {
	c, ok := cities[city]
	if ok {
		c.Latitude = lat
		c.Longitude = long
	}
}

// toRad converts degree values to radians
func toRad(val float64) float64 {
	return val * math.Pi / 180.0
}

// Determine the distance between cities based on their latitude and longitudes
func (city *CitySE) Distance(city2 *City) float64 {
	dlat := toRad(city.Latitude - city2.Latitude)
	dlon := toRad(city.Longitude - city2.Longitude)
	lat1 := toRad(city.Latitude)
	lat2 := toRad(city2.Latitude)
	a1 := math.Sin(dlat / 2.0)
	a2 := math.Sin(dlon / 2.0)
	a := a1*a1 + a2*a2*math.Cos(lat1)*math.Cos(lat2)
	return 6371.0 * 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
}

// Our distination city
var destination string

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
	SetCoords("Pretoria", -25.7313, 28.2184)
	SetCoords("Midrand", -25.98953, 28.12843)
	SetCoords("Bloemfontein", -29.1183, 26.2249)
	SetCoords("Cape Town", -33.9249, 18.4241)
	SetCoords("Johannesburg", -26.2041, 28.0473)
	SetCoords("Kempton", -26.1, 28.233334)
	SetCoords("Klerksdorp", -26.859823, 26.631750)
	SetCoords("Potchefstroom", -26.71667, 27.1)
	SetCoords("Kimberley", -28.741943, 24.771944)
	SetCoords("Vanderbijl", -26.703421, 27.807695)
	SetCoords("Vereeniging", -26.673611, 27.931944)
	SetCoords("Sasolburg", -26.810190, 27.827724)
	SetCoords("Kroonstad", -27.644606, 27.250900)
	SetCoords("Ventersburg", -28.08561, 27.13814)
	SetCoords("Beaufort West", -32.35671, 22.58295)
	SetCoords("Worcester", -33.64651, 19.44852)
	SetCoords("George", -33.963, 22.46173)
}

// searchRoute will search for a route from fromcity to tocity
func searchRoute(fromcity, tocity string) {
	destination = tocity
	city := cities[fromcity]
	// Call our search func
	cnt, ans, hist := src.BestCostAwaySearch(&CitySE{city, strconv.Itoa(city.Key), 0}, true)
	fmt.Printf("Done in %d steps\n", cnt)
	for i := len(hist) - 1; i >= 0; i-- {
		fmt.Printf("%v\n", hist[i])
	}
	fmt.Printf("%v\n", ans)
}

func main() {
	searchRoute("Pretoria", "Cape Town")
}
