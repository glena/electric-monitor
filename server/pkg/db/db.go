package db

import (
	"fmt"
	"time"

	"github.com/influxdata/influxdb/client/v2"
)

type DB struct {
	client      client.Client
	database    string
	batchPoints client.BatchPoints
}

func Init(host string, database string, username string, password string) (db DB, err error) {

	// Create a new HTTPClient
	dbClient, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     host,
		Username: username,
		Password: password,
	})

	if err != nil {
		return DB{}, err
	}

	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  database,
		Precision: "s",
	})
	if err != nil {
		return DB{}, err
	}

	obj := DB{
		database:    database,
		client:      dbClient,
		batchPoints: bp,
	}

	return obj, nil
}

func (me DB) Disconnect() {
	me.client.Close()
}

func (me DB) Query(cmd string) (res []client.Result, err error) {
	q := client.Query{
		Command:  cmd,
		Database: me.database,
	}
	if response, err := me.client.Query(q); err == nil {
		if response.Error() != nil {
			return res, response.Error()
		}
		res = response.Results
	} else {
		return res, err
	}
	return res, nil
}

func (me DB) InsertPoint(name string, tags map[string]string, fields map[string]interface{}) error {
	pt, err := client.NewPoint(name, tags, fields, time.Now())
	if err != nil {
		fmt.Println("NewPoint")
		return err
	}

	me.batchPoints.AddPoint(pt)

	err = me.client.Write(me.batchPoints)
	// Write the batch
	if err != nil {
		fmt.Println("Write the batch")
		return err
	}

	return nil
}
