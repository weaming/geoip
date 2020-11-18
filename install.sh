#!/bin/bash

wget https://raw.githubusercontent.com/wp-statistics/GeoLite2-City/master/GeoLite2-City.mmdb.gz
gzip -d GeoLite2-City.mmdb.gz && mv GeoLite2-City.mmdb /etc/
go install .
