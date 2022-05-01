package entity

import "github.com/o-t-k-t/ent-exercise/ent"

type User ent.User

func (u User) LocalName() string {
	return u.Name + "さん"
}
