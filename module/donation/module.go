package donationmodule

import (
	"context"

	donationdomain "github.com/twin-te/twinte-back/module/donation/domain"
	"github.com/twin-te/twinte-back/module/shared/domain/idtype"
)

// UseCase represents application specific business rules.
//
// The following error codes are not stated explicitly in the each method, but may be returned.
//   - shared.InvalidArgument
//   - shared.Unauthenticated
//   - shared.Unauthorized
type UseCase interface {
	// CreateOneTimeCheckoutSession
	//
	// [Authentication] optional
	CreateOneTimeCheckoutSession(ctx context.Context, amount int) (idtype.CheckoutSessionID, error)

	// CreateSubscriptionCheckoutSession
	//
	// [Authentication] required
	CreateSubscriptionCheckoutSession(ctx context.Context, subscriptionPlanID idtype.SubscriptionPlanID) (idtype.CheckoutSessionID, error)

	// GetOrCreatePaymentUser
	//
	// [Authentication] required
	GetOrCreatePaymentUser(ctx context.Context) (*donationdomain.PaymentUser, error)

	// UpdateOrCreatePaymentUser
	//
	// [Authentication] required
	UpdateOrCreatePaymentUser(ctx context.Context, in UpdateOrCreatePaymentUserIn) (*donationdomain.PaymentUser, error)

	// GetPaymentHistories
	//
	// [Authentication] required
	GetPaymentHistories(ctx context.Context) ([]*donationdomain.PaymentHistory, error)

	// GetSubscriptions
	//
	// [Authentication] required
	GetSubscriptions(ctx context.Context) ([]*donationdomain.Subscription, error)

	// Unsubscribe
	//
	// [Authentication] required
	//
	// [Error Code]
	//   - donation.SubscriptionNotFound
	Unsubscribe(ctx context.Context, id idtype.SubscriptionID) error

	// GetTotalAmount
	//
	// [Authentication] not required
	GetTotalAmount(ctx context.Context) (int, error)

	// GetContributors
	//
	// [Authentication] not required
	GetContributors(ctx context.Context) ([]*donationdomain.PaymentUser, error)
}

type UpdateOrCreatePaymentUserIn struct {
	DisplayName *string
	Link        *string
}
