package main

import (
	"math"
	"fmt"
	"flag"
	"os"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"time"
)

const (
	//earthRadiusMi = 3958 // radius of the earth in miles.
	earthRaidusKm = 6371 // radius of the earth in kilometers.
)

// flag help menu
var usage = `Usage: distance [options...] 
Options:
  -d  distance in Km to be used for comparison
  -h  help menu
  
  `

var (
	v = flag.Bool("v", false, "verbose")
	d = flag.Float64("d", 7.5, "the distance used for comparison")
	h = flag.Bool("h", false, "help")
	results Generated
	rainfallresults RainfallGenerated
	db DB
	
)

func usageAndExit(msg string) {
	if msg != "" {
		fmt.Fprintf(os.Stderr, msg)
		fmt.Fprintf(os.Stderr, "\n\n")
	}
	flag.Usage()
	fmt.Fprintf(os.Stderr, "\n")
	os.Exit(1)
}

//Data structure for the retreive JSON data
type Generated struct {
	AreaMetadata []struct {
		Name          string `json:"name"`
		LabelLocation struct {
			Lat  float64 `json:"latitude"`
			Lon float64 `json:"longitude"`
		} `json:"label_location"`
	} `json:"area_metadata"`
}

type RainfallGenerated struct {
	Metadata struct {
		Stations []struct {
			ID       string `json:"id"`
			DeviceID string `json:"device_id"`
			Name     string `json:"name"`
			Location struct {
				Lat  float64 `json:"latitude"`
				Lon float64 `json:"longitude"`
			} `json:"location"`
		} `json:"stations"`
	} `json:"metadata"`
}


type DB struct {
	Dist 	float64 	`json:"dist"`
	Station []Stations `json:"stations"`
}

//Struct field must be Capital, cannot be lowercase.
type Stations struct {
		Name string `json:"name"`
		Id string 	`json:"id"`
		Tb string 	`json:"tb"`
}

// Function to retrieve Weather anf Rainfall data from GovTech API
func getWeatherRainfall() {
	t := time.Now()
	resp, err := http.Get("https://api.data.gov.sg/v1/environment/2-hour-weather-forecast?date_time=" + t.Format("2006-01-02T15:04:05"))
	if err != nil {
		fmt.Println(err)
	}
	data, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println(err)
	}
	
	if e := json.Unmarshal([]byte(data), &results); e != nil {
		fmt.Println(e.Error())
	}

	resp1, err := http.Get("https://api.data.gov.sg/v1/environment/rainfall?date_time=" + t.Format("2006-01-02T15:04:05"))
	if err != nil {
		fmt.Println(err)
	}
	data1, err := ioutil.ReadAll(resp1.Body)
	defer resp1.Body.Close()
	if err != nil {
		fmt.Println(err)
	}

	if e := json.Unmarshal([]byte(data1), &rainfallresults); e != nil {
		fmt.Println(e.Error())
	}
}

// degreesToRadians converts from degrees to radians.
func degreesToRadians(d float64) float64 {
	return d * math.Pi / 180
}

// Distance calculates the shortest path between two coordinates on the surface
// of the Earth. This function returns two units of measure, the first is the
// distance in miles, the second is the distance in kilometers.
func Distance(p, q, r, s float64) (km float64) {
	lat1 := degreesToRadians(p)
	lon1 := degreesToRadians(q)
	lat2 := degreesToRadians(r)
	lon2 := degreesToRadians(s)

	diffLat := lat2 - lat1
	diffLon := lon2 - lon1

	a := math.Pow(math.Sin(diffLat/2), 2) + math.Cos(lat1)*math.Cos(lat2)*
		math.Pow(math.Sin(diffLon/2), 2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	//mi = c * earthRadiusMi
	km = c * earthRaidusKm

	return km
}

func main() {
		

		flag.Usage = func() {
		fmt.Fprint(os.Stderr, fmt.Sprintf(usage))
		}

		flag.Parse()

		if *h == true {
			usageAndExit("*** Help Menu ***")
		}

		// Retrieve Weather Forecast and Rainfall data
		getWeatherRainfall()
		
		//Start of the comparison logic to list Rain Gauge Stations with "distance"
		for _, w := range results.AreaMetadata {
			//fmt.Println(w.Name)
			//fmt.Println(w.LabelLocation.Lat)
			//fmt.Println(w.LabelLocation.Lon)
			for _, r := range rainfallresults.Metadata.Stations {
				//fmt.Println(r.ID)
				//fmt.Println(r.Name)
				//fmt.Println(r.Location.Lat)
				//fmt.Println(r.Location.Lon)
				
				km := Distance(w.LabelLocation.Lat, w.LabelLocation.Lon, r.Location.Lat, r.Location.Lon)
				if km < *d {
					
					z := Stations{w.Name, r.ID, r.Name}
					db.Station = append(db.Station, z)
					
					//fmt.Println(db.stations)
					/*
					fmt.Println("name:", w.Name)
					fmt.Println("id:", r.ID)
					fmt.Println("tb:",r.Name)
					*/
					// fmt.Println("Kilometers:", km)
					
				}
			}
		}
		//fmt.Println(db)
		js, err := json.MarshalIndent(db, "", "  ") 
			if err != nil {
			fmt.Println(err)
			}
		fmt.Printf("%s\n", js)
        //fmt.Println("Kilometers:", km)
    }