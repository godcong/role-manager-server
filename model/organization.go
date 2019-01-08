package model

// Organization ...
type Organization struct {
	Model       `bson:",inline"`
	IsDefault   bool   `bson:"is_default"`
	Verify      string `bson:"verify"`   //验证状态
	Name        string `bson:"name"`     //商户名称
	Code        string `bson:"code"`     //社会统一信用代码
	Contact     string `bson:"contact"`  //商户联系人
	Position    string `bson:"position"` //联系人职位
	Phone       string `bson:"phone"`    //联系人手机号
	Mailbox     string `bson:"mailbox"`  //联系人邮箱
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
