package authdomain

import (
	"fmt"

	"github.com/twin-te/twinte-back/base"
	shareddomain "github.com/twin-te/twinte-back/module/shared/domain"
)

//go:generate go run golang.org/x/tools/cmd/stringer -type=Provider -trimprefix=Provider -output=provider_string.gen.go
type Provider int

const (
	ProviderGoogle Provider = iota + 1
	ProviderApple
	ProviderTwitter
)

var AllProviders = []Provider{
	ProviderGoogle,
	ProviderApple,
	ProviderTwitter,
}

func ParseProvider(s string) (Provider, error) {
	ret, ok := base.FindByString(AllProviders, s)
	if ok {
		return ret, nil
	}
	return 0, fmt.Errorf("failed to parse provider %#v", s)
}

// SocialID represents provider's user id
type SocialID struct {
	shareddomain.RequiredString
}

func ParseSocialID(s string) (SocialID, error) {
	rs, err := shareddomain.NewRequiredStringParser("social id")(s)
	return SocialID{rs}, err
}

type UserAuthentication struct {
	Provider Provider
	SocialID SocialID
}

func NewUserAuthentication(provider Provider, socialID SocialID) UserAuthentication {
	return UserAuthentication{Provider: provider, SocialID: socialID}
}
