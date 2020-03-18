package datastore

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

//JSONDatastore JSONDatastore
type JSONDatastore interface {
	Save(name string, data interface{}) bool
	Read(name string) *[]byte
	ReadAll() *[][]byte
	Delete(name string) bool
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

func (d *DataStore) Read(name string) *[]byte {
	d.mu.Lock()
	defer d.mu.Unlock()
	var rtn []byte
	if d.cache != nil && d.cache[name] != nil {
		rtn = d.cache[name]
	}
	return &rtn
}

//ReadAll ReadAll
func (d *DataStore) ReadAll() *[][]byte {
	d.mu.Lock()
	defer d.mu.Unlock()
	var rtn [][]byte
	if d.cache != nil {
		for _, v := range d.cache {
			rtn = append(rtn, v)
		}
	}
	return &rtn
}

//Delete Delete
func (d *DataStore) Delete(name string) bool {
	d.mu.Lock()
	defer d.mu.Unlock()
	var rtn bool
	if name != "" && d.Path != "" {
		delete(d.cache, name)
		var fileName = d.Path + string(filepath.Separator) + name + ".json"
		jerr := os.Remove(fileName)
		if jerr == nil {
			rtn = true
		}
	}
	return rtn
}

//GetNew GetNew
func (d *DataStore) GetNew() JSONDatastore {
	// Should call in main of application and then
	// should us dependency injection to inject JSONDataStore
	d.mu.Lock()
	defer d.mu.Unlock()
	var jd JSONDatastore
	if d.cache == nil {
		fmt.Println("initializing cache")
		d.cache = make(map[string][]byte)
		files, err := ioutil.ReadDir(d.Path)
		if err == nil {
			for _, f := range files {
				isDir := f.IsDir()
				if !isDir {
					var fileName = d.Path + string(filepath.Separator) + f.Name()
					file, err2 := ioutil.ReadFile(fileName)
					if err2 == nil {
						var name = f.Name()
						fmt.Println("name: ", name)
						name = name[0 : len(name)-5]
						fmt.Println("name2: ", name)
						d.cache[name] = file
					}
				}
			}
		}
	}
	jd = d
	return jd
}

//go mod init github.com/Ulbora/json-datastore
