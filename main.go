// GeoOrderTest .
//
// Used for customer clients, then can place order between different places, take order, and list orders. Use Google Maps API to get the distance for the order.
//
//     Schemes: http
//     Version: 0.1
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//
// swagger:meta
package main

import "github.com/jchprj/GeoOrderTest/cmd"

func main() {
	cmd.Execute()
}
