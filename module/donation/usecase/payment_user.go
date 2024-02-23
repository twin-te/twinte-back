package donationusecase

import (
	"context"
	"errors"

	"github.com/samber/lo"
	"github.com/twin-te/twinte-back/base"
	donationmodule "github.com/twin-te/twinte-back/module/donation"
	donationdomain "github.com/twin-te/twinte-back/module/donation/domain"
	donationport "github.com/twin-te/twinte-back/module/donation/port"
	sharedport "github.com/twin-te/twinte-back/module/shared/port"
	"golang.org/x/sync/errgroup"
)

func (uc *impl) GetOrCreatePaymentUser(ctx context.Context) (*donationdomain.PaymentUser, error) {
	userID, err := uc.a.Authenticate(ctx)
	if err != nil {
		return nil, err
	}

	paymentUser, err := uc.r.FindPaymentUser(ctx, donationport.FindPaymentUserConds{
		UserID: userID,
	}, sharedport.LockNone)

	if err != nil && !errors.Is(err, sharedport.ErrNotFound) {
		return nil, err
	}

	if err == nil {
		return paymentUser, nil
	}

	paymentUser, err = uc.f.NewPaymentUser(userID, nil, nil)
	if err != nil {
		return nil, err
	}

	return paymentUser, uc.r.CreatePaymentUsers(ctx, paymentUser)
}

func (uc *impl) UpdateOrCreatePaymentUser(ctx context.Context, in donationmodule.UpdateOrCreatePaymentUserIn) (paymentUser *donationdomain.PaymentUser, err error) {
	userID, err := uc.a.Authenticate(ctx)
	if err != nil {
		return nil, err
	}

	err = uc.r.Transaction(ctx, func(rtx donationport.Repository) (err error) {
		paymentUser, err = rtx.FindPaymentUser(ctx, donationport.FindPaymentUserConds{UserID: userID}, sharedport.LockExclusive)
		if err != nil {
			return err
		}

		paymentUser.BeforeUpdateHook()
		paymentUser.Update(donationdomain.PaymentUserDataToUpdate{
			DisplayName: in.DisplayName,
			Link:        in.Link,
		})

		return rtx.UpdatePaymentUser(ctx, paymentUser)
	})

	if err != nil && !errors.Is(err, sharedport.ErrNotFound) {
		return
	}

	if err == nil {
		return
	}

	paymentUser, err = uc.f.NewPaymentUser(userID, in.DisplayName, in.Link)
	if err != nil {
		return
	}

	return paymentUser, uc.r.CreatePaymentUsers(ctx, paymentUser)
}

func (uc *impl) GetContributors(ctx context.Context) ([]*donationdomain.PaymentUser, error) {
	uc.contributorsCache.mu.RLock()
	defer uc.contributorsCache.mu.RUnlock()

	return base.MapByClone(uc.contributorsCache.v), nil
}

func (uc *impl) updateContributorsCache(ctx context.Context) error {
	paymentUsers, err := uc.r.ListPaymentUsers(ctx, donationport.ListPaymentUsersConds{
		RequireDisplayName: true,
	}, sharedport.LockNone)
	if err != nil {
		return err
	}

	mask := make([]bool, len(paymentUsers))

	eg, ctx := errgroup.WithContext(ctx)
	for i, paymentUser := range paymentUsers {
		i := i
		paymentUser := paymentUser
		eg.Go(func() error {
			select {
			case <-ctx.Done():
				return nil
			default:
				paymentHistories, err := uc.g.ListPaymentHistories(ctx, &paymentUser.ID)
				if err != nil {
					return err
				}

				isContributor := lo.SomeBy(paymentHistories, func(paymentHistory *donationdomain.PaymentHistory) bool {
					return paymentHistory.Status == donationdomain.PaymentStatusSucceeded
				})
				mask[i] = isContributor
				return nil
			}
		})
	}

	if err := eg.Wait(); err != nil {
		return err
	}

	contributors := lo.Filter(paymentUsers, func(_ *donationdomain.PaymentUser, index int) bool {
		return mask[index]
	})

	uc.contributorsCache.mu.Lock()
	uc.contributorsCache.v = contributors
	uc.contributorsCache.mu.Unlock()

	return nil
}
