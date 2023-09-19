package apiservice

import (
	"fmt"

	apigen "github.com/twin-te/twinte-back/api/gen"
	authentity "github.com/twin-te/twinte-back/module/auth/entity"
)

func fromPBProvider(pbProvider apigen.Provider) (authentity.Provider, error) {
	switch pbProvider {
	case apigen.Provider_PROVIDER_GOOGLE:
		return authentity.ProviderGoogle, nil
	case apigen.Provider_PROVIDER_TWITTER:
		return authentity.ProviderTwitter, nil
	case apigen.Provider_PROVIDER_APPLE:
		return authentity.ProviderApple, nil
	}
	return 0, fmt.Errorf("unknown pb provider %d", pbProvider)
}

func toPBProvider(provider authentity.Provider) (apigen.Provider, error) {
	switch provider {
	case authentity.ProviderGoogle:
		return apigen.Provider_PROVIDER_GOOGLE, nil
	case authentity.ProviderTwitter:
		return apigen.Provider_PROVIDER_TWITTER, nil
	case authentity.ProviderApple:
		return apigen.Provider_PROVIDER_APPLE, nil
	}
	return 0, fmt.Errorf("unknown provider %d", provider)
}
