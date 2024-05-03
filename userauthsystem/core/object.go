package core

import "time"

type AccountSignIn struct {
	Id          string
	LastLogIn   time.Time
	AccessToken string
}
