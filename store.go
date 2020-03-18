package datastore

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"sync"
)

//JSONDatastore JSONDatastore
type JSONDatastore interface {
	Save(name string, data interface{}) bool
	Read(name string) []byte
}

//DataStore Datastore
type DataStore struct {
	Path  string
	cache map[string][]byte
	mu    sync.Mutex
}

//Save Save
func (d *DataStore) Save(name string, data interface{}) bool {
	d.mu.Lock()
	defer d.mu.Unlock()
	var rtn bool
	aJSON, err := json.Marshal(data)
	if err == nil && name != "" && d.Path != "" {
		d.cache[name] = aJSON
		var fileName = d.Path + string(filepath.Separator) + name + ".json"
		jerr := ioutil.WriteFile(fileName, aJSON, 0644)
		if jerr == nil {
			rtn = true
		}
	}
	return rtn
}

func (d *DataStore) Read(name string) []byte {
	d.mu.Lock()
	defer d.mu.Unlock()
	var rtn []byte
	if d.cache != nil && d.cache[name] != nil {
		rtn = d.cache[name]
	} else {
		var fileName = d.Path + string(filepath.Separator) + name + ".json"
		file, err := ioutil.ReadFile(fileName)
		if err == nil {
			rtn = file
		} else {
			log.Println("Reading Json file err: ", err)
		}
	}
	return rtn
}

//GetNew GetNew
func (d *DataStore) GetNew() JSONDatastore {
	// Should call in main of application and then
	// should us dependency injection to inject JSONDataStore
	var jd JSONDatastore
	if d.cache == nil {
		fmt.Println("initializing cache")
		d.cache = make(map[string][]byte)
	}
	jd = d
	return jd
}

//go mod init github.com/Ulbora/json-datastore
