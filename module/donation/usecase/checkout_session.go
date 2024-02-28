package donationusecase

import (
	"context"
	"errors"

	donationport "github.com/twin-te/twinte-back/module/donation/port"
	"github.com/twin-te/twinte-back/module/shared/domain/idtype"
	sharedport "github.com/twin-te/twinte-back/module/shared/port"
)

func (uc *impl) CreateOneTimeCheckoutSession(ctx context.Context, amount int) (idtype.CheckoutSessionID, error) {
	var paymentUserID *idtype.PaymentUserID

	userID, err := uc.a.Authenticate(ctx)
	if err == nil {
		paymentUser, err := uc.r.FindPaymentUser(ctx, donationport.FindPaymentUserConds{
			UserID: userID,
		}, sharedport.LockNone)

		switch {
		case err == nil:
			paymentUserID = &paymentUser.ID
		case !errors.Is(err, sharedport.ErrNotFound):
			return "", err
		}
	}

	return uc.g.CreateOneTimeCheckoutSession(ctx, paymentUserID, amount)
}

func (uc *impl) CreateSubscriptionCheckoutSession(ctx context.Context, subscriptionPlanID idtype.SubscriptionPlanID) (idtype.CheckoutSessionID, error) {
	_, err := uc.a.Authenticate(ctx)
	if err != nil {
		return "", err
	}

	paymentUser, err := uc.GetOrCreatePaymentUser(ctx)
	if err != nil {
		return "", nil
	}

	return uc.g.CreateSubscriptionCheckoutSession(ctx, paymentUser.ID, subscriptionPlanID)
}
