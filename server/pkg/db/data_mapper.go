package db

import (
	"encoding/json"

	"github.com/influxdata/influxdb/client/v2"
)

type DataPoint struct {
	Timestamp string
	P         float64
	V         float64
	I         float64
}

func MapDataPoints(res []client.Result) []DataPoint {
	pointList := []DataPoint{}

	for _, d := range res[0].Series[0].Values {
		if d[1] != nil && d[2] != nil && d[3] != nil {
			P, err1 := d[1].(json.Number).Float64()
			V, err2 := d[2].(json.Number).Float64()
			I, err3 := d[3].(json.Number).Float64()

			if err1 == nil && err2 == nil && err3 == nil {
				o := DataPoint{
					Timestamp: d[0].(string),
					P:         P,
					V:         V,
					I:         I,
				}

				pointList = append(pointList, o)
			}
		}
	}

	return pointList
}
