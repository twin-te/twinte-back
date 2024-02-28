package restv3

import (
	"context"
	"fmt"

	openapi_types "github.com/oapi-codegen/runtime/types"
	"github.com/samber/lo"
	"github.com/twin-te/twinte-back/base"
	"github.com/twin-te/twinte-back/handler/api/rest/v3/openapi"
	donationmodule "github.com/twin-te/twinte-back/module/donation"
	donationdomain "github.com/twin-te/twinte-back/module/donation/domain"
	"github.com/twin-te/twinte-back/module/shared/domain/idtype"
)

func toPaymentStatus(paymentStatus donationdomain.PaymentStatus) (openapi.PaymentStatus, error) {
	switch paymentStatus {
	case donationdomain.PaymentStatusPending:
		return openapi.PaymentStatusPending, nil
	case donationdomain.PaymentStatusCanceled:
		return openapi.PaymentStatusCanceled, nil
	case donationdomain.PaymentStatusSucceeded:
		return openapi.PaymentStatusSucceeded, nil
	}
	return "", fmt.Errorf("invalid %#v", paymentStatus)
}

func toPaymentType(paymentType donationdomain.PaymentType) (openapi.PaymentType, error) {
	switch paymentType {
	case donationdomain.PaymentTypeOneTime:
		return openapi.PaymentTypeOneTime, nil
	case donationdomain.PaymentTypeSubscription:
		return openapi.PaymentTypeSubscription, nil
	}
	return "", fmt.Errorf("invalid %#v", paymentType)
}

func toApiPayment(paymentHistory *donationdomain.PaymentHistory) (apiPayment openapi.Payment, err error) {
	apiPayment = openapi.Payment{
		Amount:  paymentHistory.Amount,
		Created: openapi_types.Date{Time: paymentHistory.CreatedAt},
		Id:      paymentHistory.ID.String(),
	}

	apiPayment.Status, err = toPaymentStatus(paymentHistory.Status)
	if err != nil {
		return
	}

	apiPayment.Type, err = toPaymentType(paymentHistory.Type)
	if err != nil {
		return
	}

	return
}

func toApiSubscription(subscription *donationdomain.Subscription) openapi.Subscription {
	apiSubscription := openapi.Subscription{
		Created: openapi_types.Date{Time: subscription.CreatedAt},
		Id:      subscription.ID.String(),
		Status:  lo.Ternary(subscription.IsActive, openapi.SubscriptionStatusActive, openapi.SubscriptionStatusCanceled),
	}

	apiSubscription.Plans = base.Map(subscription.Plans, func(plan *donationdomain.SubscriptionPlan) struct {
		Amount int    `json:"amount"`
		Id     string `json:"id"`
		Name   string `json:"name"`
	} {
		return struct {
			Amount int    `json:"amount"`
			Id     string `json:"id"`
			Name   string `json:"name"`
		}{
			Amount: plan.Amount,
			Id:     plan.ID.String(),
			Name:   plan.Name,
		}
	})

	return apiSubscription
}

func toApiPaymentUser(paymentUser *donationdomain.PaymentUser) openapi.PaymentUser {
	return openapi.PaymentUser{
		DisplayName:   paymentUser.DisplayName,
		Link:          paymentUser.Link,
		PaymentUserId: paymentUser.ID.String(),
		TwinteUserId:  toApiUUID(paymentUser.UserID),
	}
}

// 寄付総額を取得
// (GET /donation/aggregate/totalAmount)
func (h *impl) GetDonationAggregateTotalAmount(ctx context.Context, request openapi.GetDonationAggregateTotalAmountRequestObject) (res openapi.GetDonationAggregateTotalAmountResponseObject, err error) {
	totalAmount, err := h.donationUseCase.GetTotalAmount(ctx)
	if err != nil {
		return
	}

	res = openapi.GetDonationAggregateTotalAmount200JSONResponse{
		Total: totalAmount,
	}

	return
}

// 寄付してくれたユーザーで掲載OKのユーザー一覧を取得
// (GET /donation/aggregate/users)
func (h *impl) GetDonationAggregateUsers(ctx context.Context, request openapi.GetDonationAggregateUsersRequestObject) (res openapi.GetDonationAggregateUsersResponseObject, err error) {
	contributors, err := h.donationUseCase.GetContributors(ctx)
	if err != nil {
		return
	}

	res = openapi.GetDonationAggregateUsers200JSONResponse(
		base.Map(contributors, func(contributor *donationdomain.PaymentUser) struct {
			DisplayName *string `json:"displayName,omitempty"`
			Link        *string `json:"link,omitempty"`
		} {
			return struct {
				DisplayName *string `json:"displayName,omitempty"`
				Link        *string `json:"link,omitempty"`
			}{
				DisplayName: contributor.DisplayName,
				Link:        contributor.Link,
			}
		}),
	)

	return
}

// 請求一覧を取得
// (GET /donation/payment)
func (h *impl) GetDonationPayment(ctx context.Context, request openapi.GetDonationPaymentRequestObject) (res openapi.GetDonationPaymentResponseObject, err error) {
	paymentHistories, err := h.donationUseCase.GetPaymentHistories(ctx)
	if err != nil {
		return
	}

	apiPaymentHistories, err := base.MapWithErr(paymentHistories, toApiPayment)
	if err != nil {
		return
	}

	res = openapi.GetDonationPayment200JSONResponse(apiPaymentHistories)

	return
}

// 単発寄付のセッションを作成
// (POST /donation/session/onetime)
func (h *impl) PostDonationSessionOnetime(ctx context.Context, request openapi.PostDonationSessionOnetimeRequestObject) (res openapi.PostDonationSessionOnetimeResponseObject, err error) {
	id, err := h.donationUseCase.CreateOneTimeCheckoutSession(ctx, request.Body.Amount)
	if err != nil {
		return
	}

	res = openapi.PostDonationSessionOnetime200JSONResponse{
		SessionId: id.String(),
	}

	return
}

// サブスク寄付のセッションを作成
// (POST /donation/session/subscription)
func (h *impl) PostDonationSessionSubscription(ctx context.Context, request openapi.PostDonationSessionSubscriptionRequestObject) (res openapi.PostDonationSessionSubscriptionResponseObject, err error) {
	id, err := idtype.ParseSubscriptionPlanID(request.Body.PlanId)
	if err != nil {
		return
	}

	checkoutSessionID, err := h.donationUseCase.CreateSubscriptionCheckoutSession(ctx, id)
	if err != nil {
		return
	}

	res = openapi.PostDonationSessionSubscription200JSONResponse{
		SessionId: checkoutSessionID.String(),
	}

	return
}

// 契約中のサブスクを取得する
// (GET /donation/subscriptions)
func (h *impl) GetDonationSubscriptions(ctx context.Context, request openapi.GetDonationSubscriptionsRequestObject) (res openapi.GetDonationSubscriptionsResponseObject, err error) {
	subscriptions, err := h.donationUseCase.GetSubscriptions(ctx)
	if err != nil {
		return
	}

	apiSubscriptions := base.Map(subscriptions, toApiSubscription)

	res = openapi.GetDonationSubscriptions200JSONResponse(apiSubscriptions)

	return
}

// 指定したサブスクを解除
// (DELETE /donation/subscriptions/{id})
func (h *impl) DeleteDonationSubscriptionsId(ctx context.Context, request openapi.DeleteDonationSubscriptionsIdRequestObject) (res openapi.DeleteDonationSubscriptionsIdResponseObject, err error) {
	id, err := idtype.ParseSubscriptionID(request.Id)
	if err != nil {
		return
	}

	err = h.donationUseCase.Unsubscribe(ctx, id)

	return
}

// 支払いユーザー取得
// (GET /donation/users/me)
func (h *impl) GetDonationUsersMe(ctx context.Context, request openapi.GetDonationUsersMeRequestObject) (res openapi.GetDonationUsersMeResponseObject, err error) {
	paymentUser, err := h.donationUseCase.GetOrCreatePaymentUser(ctx)
	if err != nil {
		return
	}

	apiPaymentUser := toApiPaymentUser(paymentUser)

	res = openapi.GetDonationUsersMe200JSONResponse(apiPaymentUser)

	return
}

// 支払いユーザー情報更新
// (PATCH /donation/users/me)
func (h *impl) PatchDonationUsersMe(ctx context.Context, request openapi.PatchDonationUsersMeRequestObject) (res openapi.PatchDonationUsersMeResponseObject, err error) {
	paymentUser, err := h.donationUseCase.UpdateOrCreatePaymentUser(ctx, donationmodule.UpdateOrCreatePaymentUserIn{
		DisplayName: request.Body.DisplayName,
		Link:        request.Body.Link,
	})
	if err != nil {
		return
	}

	apiPaymentUser := toApiPaymentUser(paymentUser)

	res = openapi.PatchDonationUsersMe200JSONResponse(apiPaymentUser)

	return
}
