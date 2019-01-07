package service

import (
	"github.com/godcong/role-manager-server/model"
	"github.com/godcong/role-manager-server/util"
	"time"
)

// Token ...
type Token struct {
	OID           string `json:"oid"`
	Name          string `json:"name"`
	Pwd           string `json:"pwd"`
	EffectiveTime int64  `json:"effectiveTime"`
}

// ToToken ...
func ToToken(u *model.User) (string, error) {
	t := Token{}
	t.Name = u.Name
	t.OID = u.ID.String()
	//t.Pwd = u.Password //TODO
	t.EffectiveTime = time.Now().Unix() + 3600*24*7

	sub, err := util.MarshalJSON(t)
	if err != nil {
		return "", err
	}

	token, err := EncryptJWT([]byte(globalKey), sub)

	return token, err
}
