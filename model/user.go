package model

type User struct {
	*Model
	Name          string
	Username      string
	Email         string
	MobilePhone   string
	IDCardFacade  string
	IDCardObverse string
	Association   string
	Password      string
	Token         string
}

func (u *User) Update() error {
	panic("implement me")
}

func (u *User) Delete() error {
	panic("implement me")
}

func (u *User) Find() error {
	panic("implement me")
}

func (u *User) CollectionName() string {
	return "user"
}

func (u *User) Create() {
	mgo.Client.Database("")
}
