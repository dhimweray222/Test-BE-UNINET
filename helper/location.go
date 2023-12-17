package helper

import (
	"log"
	"net"
	"os"
	"strconv"
	"strings"

	"github.com/ipinfo/go/v2/ipinfo"
)

type Location struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lon"`
}

func GetLocation(ip_address string) (Location, error) {
	token := os.Getenv("TOKEN_MAPS")

	client := ipinfo.NewClient(nil, nil, token)

	info, err := client.GetIPInfo(net.ParseIP(ip_address))

	if err != nil {
		log.Println(err)
		return Location{}, err
	}

	// Split the string using the comma as the delimiter
	parts := strings.Split(info.Location, ",")

	latitude := parts[0]
	longitude := parts[1]
	fLatitude, err := strconv.ParseFloat(latitude, 64)
	fLongitude, err := strconv.ParseFloat(longitude, 64)
	data := Location{
		Latitude:  fLatitude,
		Longitude: fLongitude,
	}

	return data, nil

}
