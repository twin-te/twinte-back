package authentity

import (
	_ "embed"
)

type OAuth2State string

func (state OAuth2State) String() string {
	return string(state)
}

type OAuth2Code string

func (code OAuth2Code) String() string {
	return string(code)
}

type OAuth2ConsentPageURL string

func (url OAuth2ConsentPageURL) String() string {
	return string(url)
}

func NewOAuth2StateFromString(s string) OAuth2State {
	return OAuth2State(s)
}

func NewOAuth2CodeFromString(s string) OAuth2Code {
	return OAuth2Code(s)
}

func NewOAuth2ConsentPageURLFromString(s string) OAuth2ConsentPageURL {
	return OAuth2ConsentPageURL(s)
}
