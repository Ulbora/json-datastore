package datastore

import (
	"encoding/json"
	"fmt"
	"testing"
)

type testData struct {
	Name    string
	Address string
	Other   []string
}

var ds DataStore

func TestDataStore_Save(t *testing.T) {
	//var ds DataStore
	ds.Path = "./testFiles"
	d := ds.GetNew()

	var td testData
	td.Address = "some address"
	td.Name = "tester"
	td.Other = []string{"att1", "att2", "att3"}

	suc := d.Save("test", td)
	//fmt.Println("found cache data", ds.cache[td.Name])
	if !suc || ds.cache == nil {
		t.Fail()
	}
}

func TestDataStore_Read(t *testing.T) {
	ds.Path = "./testFiles"
	d := ds.GetNew()

	pg := d.Read("test")
	//fmt.Println("read cache data", pg)
	var td testData
	err := json.Unmarshal(pg, &td)
	if err != nil {
		fmt.Println("read err", err)
	}
	//fmt.Println("read testData", td)
	if pg == nil || td.Address != "some address" {
		t.Fail()
	}
}

func TestDataStore_ReadERR(t *testing.T) {
	ds.Path = "./testFiles"
	d := ds.GetNew()

	pg := d.Read("test2")
	//fmt.Println("read cache data", pg)
	if pg != nil {
		t.Fail()
	}
}

func TestDataStore_Read3(t *testing.T) {
	ds.Path = "./testFiles"
	d := ds.GetNew()

	pg := d.Read("test3")
	//fmt.Println("read cache data", pg)

	var td testData
	err := json.Unmarshal(pg, &td)
	if err != nil {
		fmt.Println("read err", err)
	}
	//fmt.Println("read testData", td)

	if pg == nil || td.Address != "some address" {
		t.Fail()
	}
}
