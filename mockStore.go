package datastore

//MockDataStore Datastore
type MockDataStore struct {
	Path        string
	MockSuccess bool
	MockData    []byte
}

//Save Save
func (d *MockDataStore) Save(name string, data interface{}) bool {
	return d.MockSuccess
}

func (d *MockDataStore) Read(name string) []byte {
	return d.MockData
}

//GetNew GetNew
func (d *MockDataStore) GetNew() JSONDatastore {
	// Should call in main of application and then
	// should us dependency injection to inject JSONDataStore
	var jd JSONDatastore
	jd = d
	return jd
}
