# gosearch-test

## Some examples on how to use [hduplooy/gosearch](https://github.com/hduplooy/gosearch)

More examples will be added over time as more algorithms are added to hduplooy/gosearch

### 8queensdepth

This is the classical puzzle where 8 queens must be placed on a standard 8x8 chess board without any queen being able to capture any other queen. It is implemented making use of the Depth First Search algorithm.

### citysearchbreadth

This is an example of searching for a route from one city to another city. A number of South African cities/towns are provided and some of their neighbours. In this instance Breadth First Search is used. This will search for the smallest number of steps but not necessarily the shortest distance. This will take 3223 steps to get to the goal.

### citysearchcost

This is an example of searching for a route from one city to another city. A number of South African cities/towns are provided and some of their neighbours. In this instance Best Cost Search is used. This will search based on the shortest distance travelled so far. The route will be reasonably the shortest but it wouldn't necessarily be the fastest. This will take 125 steps to get to the goal.

### citysearchcostaway

This is an example of searching for a route from one city to another city. A number of South African cities/towns are provided and some of their neighbours. In this instance Best Cost Away Search is used. This will search based on the shortest distance travelled so far as well as the estimated distance to the goal. The route will be reasonably the shortest and the fastest. This will take 53 steps to get to the goal.

### webcitysearch

This is just a web implementation of citysearchcostaway at port 8080.


