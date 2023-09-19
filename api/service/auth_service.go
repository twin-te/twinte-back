package apiservice

import (
	"context"
	"errors"

	_ "embed"

	"github.com/bufbuild/connect-go"
	apigen "github.com/twin-te/twinte-back/api/gen"
	"github.com/twin-te/twinte-back/apperr"
	authmodule "github.com/twin-te/twinte-back/module/auth"
)

type AuthService struct {
	authUseCase authmodule.UseCase
}

func (s *AuthService) GetMe(ctx context.Context, req *connect.Request[apigen.GetMeRequest]) (*connect.Response[apigen.GetMeResponse], error) {
	user, err := s.authUseCase.AuthorizeAuthenticatedUser(ctx)
	if err != nil {
		if errors.Is(err, apperr.ErrUnauthenticated) {
			return nil, connect.NewError(connect.CodeUnauthenticated, err)
		}
		return nil, err
	}

	pbAuthentications := make([]*apigen.Authentication, 0, len(user.Authentications))
	for _, authentication := range user.Authentications {
		provider, err := toPBProvider(authentication.Provider)
		if err != nil {
			return nil, err
		}

		pbAuthentications = append(pbAuthentications, &apigen.Authentication{
			Provider: provider,
			SocialID: authentication.SocialID.String(),
		})
	}

	res := connect.NewResponse(&apigen.GetMeResponse{
		Id:              user.ID.String(),
		Authentications: pbAuthentications,
	})
	return res, nil
}

func (s *AuthService) DeleteAccount(ctx context.Context, req *connect.Request[apigen.DeleteAccountRequest]) (*connect.Response[apigen.DeleteAccountResponse], error) {
	user, err := s.authUseCase.AuthorizeAuthenticatedUser(ctx)
	if err != nil {
		if errors.Is(err, apperr.ErrUnauthenticated) {
			return nil, connect.NewError(connect.CodeUnauthenticated, err)
		}
		return nil, err
	}

	err = s.authUseCase.DeleteAccount(ctx, user)
	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(&apigen.DeleteAccountResponse{})
	return res, nil
}

func (s *AuthService) DeleteAuthentication(ctx context.Context, req *connect.Request[apigen.DeleteAuthenticationRequest]) (*connect.Response[apigen.DeleteAuthenticationResponse], error) {
	user, err := s.authUseCase.AuthorizeAuthenticatedUser(ctx)
	if err != nil {
		if errors.Is(err, apperr.ErrUnauthenticated) {
			return nil, connect.NewError(connect.CodeUnauthenticated, err)
		}
		return nil, err
	}

	provider, err := fromPBProvider(req.Msg.Provider)
	if err != nil {
		return nil, err
	}

	if err := s.authUseCase.DeleteUserAuthentication(ctx, user, provider); err != nil {
		return nil, err
	}

	res := connect.NewResponse(&apigen.DeleteAuthenticationResponse{})
	return res, nil
}

func NewAuthService(authUseCase authmodule.UseCase) *AuthService {
	return &AuthService{
		authUseCase: authUseCase,
	}
}
