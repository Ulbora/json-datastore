package datastore

//MockDataStore Datastore
type MockDataStore struct {
	Path              string
	MockSuccess       bool
	MockDeleteSuccess bool
	MockData          []byte
	MockDataList      [][]byte
}

//Save Save
func (d *MockDataStore) Save(name string, data interface{}) bool {
	return d.MockSuccess
}

//Read Read
func (d *MockDataStore) Read(name string) *[]byte {
	return &d.MockData
}

//ReadAll ReadAll
func (d *MockDataStore) ReadAll() *[][]byte {
	return &d.MockDataList
}

//Delete Delete
func (d *MockDataStore) Delete(name string) bool {
	return d.MockDeleteSuccess
}

//GetNew GetNew
func (d *MockDataStore) GetNew() JSONDatastore {
	// Should call in main of application and then
	// should us dependency injection to inject JSONDataStore
	var jd JSONDatastore
	jd = d
	return jd
}
