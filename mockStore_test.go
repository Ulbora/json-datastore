package datastore

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMockDataStore_Save(t *testing.T) {
	var mds MockDataStore
	mds.Path = "./testFiles"
	mds.MockSuccess = true
	d := mds.GetNew()

	var td testData
	td.Address = "some address"
	td.Name = "tester"
	td.Other = []string{"att1", "att2", "att3"}

	suc := d.Save("test", td)
	//fmt.Println("found cache data", ds.cache[td.Name])
	if !suc {
		t.Fail()
	}
}

func TestMockDataStore_Read(t *testing.T) {
	var mds MockDataStore
	mds.Path = "./testFiles"

	var td testData
	td.Address = "some address"
	td.Name = "tester"
	td.Other = []string{"att1", "att2", "att3"}

	aJSON, err := json.Marshal(td)
	if err != nil {
		fmt.Println("err: ", err)
	}
	mds.MockData = aJSON
	d := mds.GetNew()

	pg := d.Read("test")
	var td2 testData

	err2 := json.Unmarshal(pg, &td2)
	if err2 != nil {
		fmt.Println("read err", err2)
	}
	//fmt.Println("read testData", td)
	if pg == nil || td.Address != "some address" {
		t.Fail()
	}
}
