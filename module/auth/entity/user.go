package authentity

import (
	"fmt"

	"github.com/twin-te/twinte-back/idtype"
)

type Provider int

const (
	ProviderGoogle Provider = iota + 1
	ProviderApple
	ProviderTwitter
)

type SocialID string

func (sid SocialID) String() string {
	return string(sid)
}

type UserAuthentication struct {
	Provider Provider
	SocialID SocialID
}

type User struct {
	ID              idtype.UserID
	Authentications []UserAuthentication
}

func (u *User) AddAuthentication(newAuthentication UserAuthentication) error {
	for _, authentication := range u.Authentications {
		if newAuthentication.Provider == authentication.Provider {
			return fmt.Errorf("the provider %d already registered", newAuthentication.Provider)
		}
	}
	u.Authentications = append(u.Authentications, newAuthentication)
	return nil
}

func (u *User) DeleteAuthentication(provider Provider) error {
	if len(u.Authentications) == 1 {
		return fmt.Errorf("the user associated with only one authentication can't delete user authentication")
	}

	for i, authentication := range u.Authentications {
		if provider == authentication.Provider {
			u.Authentications = append(u.Authentications[:i], u.Authentications[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("the provider %d is not registered", provider)
}

func NewSocialIDFromString(s string) SocialID {
	return SocialID(s)
}
