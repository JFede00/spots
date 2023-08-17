package main

import (
	"database/sql"
	"fmt"
	"log"
	"math"
)

// Define the Point struct 
type Point struct {
	Latitude  float64
	Longitude float64
	Radious   float64
}

// create type_ and area variables
var type_ string
var area float64

// DB credentials
const (
	host     = "localhost"
	port     = 0000
	user     = "postgres"
	password = "your-password"
	dbname   = "spots"
)

func main() {

	// choise if square or circle
	fmt.Println("Choice: square or circle")
	fmt.Scan(&type_)
	switch(type_){
	    case "circle":
	        fmt.Println("your choice is: ", type_)
	    break;
	    
	    case "square":
	        fmt.Println("your choice is: ", type_)
	    break;
	    
	    default:
	        fmt.Println("Please enter circle or square :)")
	    return
	}
	

	// set area dimension
	fmt.Println("Set area")
	fmt.Scan(&area)
	if(area > 0) {
    	fmt.Println("The area is:", area*100)
	}else{
	    fmt.Println("Please enter value different form 0 :)")
	    return
	}
	

	// set the point coordinates
	point_1 := Point{1000.4, 1200.5, 11}

	point_2 := Point{1432.3, 11123.5, 6}
	var distance float64
	if !(point_1.Latitude > area*100 || point_1.Longitude > area*100 || point_2.Latitude > area*100 || point_2.Longitude > area*100 || point_1.Radious > area*100 || point_2.Radious > area*100) {
	fmt.Println("First Point coordinates : ", point_1)
	fmt.Println("Second Point coordinates : ", point_2)
	
	// calculate distance from point
	distance := math.Sqrt((point_2.Latitude - point_1.Latitude) + (point_2.Longitude - point_1.Longitude))
	fmt.Println( "The distance between: ", distance * 2)
	} else {
		fmt.Println("Please enter coordinates that can be into the square or circle area.")
	}

	

	/***

	  SQL QUERY

	 ***/

	// create connection with DB
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, password, host, port, dbname)

	// Open the connection
	db, err := sql.Open("mysql", dsn)

	// Check connection
	if err != nil {
		log.Fatalf("impossible to create the connection: %s", err)
	}
	defer db.Close()

	// Create INSERT query
	sqlStatement := `
        INSERT INTO spots (point_1, point_2, distance)
        VALUES ($1, $2, $3)
        RETURNING id`

	id := 0
	err = db.QueryRow(sqlStatement, 30, point_1, point_2, distance).Scan(&id)
	if err != nil {
		panic(err)
	}
	fmt.Println("New record ID is:", id)

	// GROUP BY distance query
	sqlQuery := "SELECT * from spots WHERE distance > 50 GROUP BY distance "
	err = db.QueryRow(sqlQuery).Scan(&distance)
	}
