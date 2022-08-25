package dto

import "xjtu.teach/ginEssential/model"

type Userdto struct {
	Name      string `json:"name"`
	Telephone string `json:"telephone"`
}

func ToUserDto(user model.User) *Userdto {
	return &Userdto{
		Name:      user.Name,
		Telephone: user.Telephone,
	}

}
