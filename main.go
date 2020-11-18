package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	"github.com/oschwald/geoip2-golang"
)

func main() {
	path := getDBPath()
	db, err := geoip2.Open(path)
	if err != nil {
		log.Fatal(err)

	}
	defer db.Close()

	if len(os.Args) < 2 {
		log.Fatal("missing IP")
	}

	ipStr := os.Args[1]
	if ipStr == "" {
		log.Fatal("missing IP")
	}

	ip := net.ParseIP(ipStr)
	if ip == nil {
		log.Fatal("invalid IP")
	}

	record, err := db.City(ip)
	if err != nil {
		log.Fatal(err)
	}

	bin, e := json.MarshalIndent(record, "", "  ")
	if e != nil {
		log.Fatal(e)
	}

	fmt.Println(string(bin))
}

func getDBPath() string {
	paths := []string{"./GeoLite2-City.mmdb", os.ExpandEnv("$HOME/.config/GeoLite2-City.mmdb"), "/etc/GeoLite2-City.mmdb"}
	for _, x := range paths {
		if _, e := os.Stat(x); e == nil {
			return x
		}
	}

	log.Fatal("missing geoip.mmdb:", strings.Join(paths, ", "))
	return ""
}
