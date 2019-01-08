package model

// Organization ...
type Organization struct {
	Model       `bson:",inline"`
	Name        string `bson:"name"`
	Verify      string `json:"verify"`
	Description string `bson:"description"`
}

// CreateIfNotExist ...
func (o *Organization) CreateIfNotExist() error {
	return CreateIfNotExist(o)
}

// Create ...
func (o *Organization) Create() error {
	return InsertOne(o)
}

// Update ...
func (o *Organization) Update() error {
	return UpdateOne(o)
}

// Delete ...
func (o *Organization) Delete() error {
	return DeleteByID(o)
}

// Find ...
func (o *Organization) Find() error {
	return FindByID(o)
}

func (o *Organization) _Name() string {
	return "organization"
}
