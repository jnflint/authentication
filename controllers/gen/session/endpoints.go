// Code generated by goa v3.0.6, DO NOT EDIT.
//
// session endpoints
//
// Command:
// $ goa gen github.com/anshap1719/go-authentication/design

package session

import (
	"context"

	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
)

// Endpoints wraps the "session" service endpoints.
type Endpoints struct {
	Refresh         goa.Endpoint
	Logout          goa.Endpoint
	LogoutOther     goa.Endpoint
	LogoutSpecific  goa.Endpoint
	GetSessions     goa.Endpoint
	RedeemToken     goa.Endpoint
	CleanSessions   goa.Endpoint
	CleanLoginToken goa.Endpoint
	CleanMergeToken goa.Endpoint
}

// NewEndpoints wraps the methods of the "session" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	// Casting service to Auther interface
	a := s.(Auther)
	return &Endpoints{
		Refresh:         NewRefreshEndpoint(s, a.APIKeyAuth),
		Logout:          NewLogoutEndpoint(s, a.JWTAuth, a.APIKeyAuth),
		LogoutOther:     NewLogoutOtherEndpoint(s, a.JWTAuth, a.APIKeyAuth),
		LogoutSpecific:  NewLogoutSpecificEndpoint(s, a.JWTAuth, a.APIKeyAuth),
		GetSessions:     NewGetSessionsEndpoint(s, a.JWTAuth, a.APIKeyAuth),
		RedeemToken:     NewRedeemTokenEndpoint(s, a.APIKeyAuth),
		CleanSessions:   NewCleanSessionsEndpoint(s, a.JWTAuth, a.APIKeyAuth),
		CleanLoginToken: NewCleanLoginTokenEndpoint(s, a.JWTAuth, a.APIKeyAuth),
		CleanMergeToken: NewCleanMergeTokenEndpoint(s, a.JWTAuth, a.APIKeyAuth),
	}
}

// Use applies the given middleware to all the "session" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.Refresh = m(e.Refresh)
	e.Logout = m(e.Logout)
	e.LogoutOther = m(e.LogoutOther)
	e.LogoutSpecific = m(e.LogoutSpecific)
	e.GetSessions = m(e.GetSessions)
	e.RedeemToken = m(e.RedeemToken)
	e.CleanSessions = m(e.CleanSessions)
	e.CleanLoginToken = m(e.CleanLoginToken)
	e.CleanMergeToken = m(e.CleanMergeToken)
}

// NewRefreshEndpoint returns an endpoint function that calls the method
// "refresh" of service "session".
func NewRefreshEndpoint(s Service, authAPIKeyFn security.AuthAPIKeyFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*RefreshPayload)
		var err error
		sc := security.APIKeyScheme{
			Name:           "api_key",
			Scopes:         []string{},
			RequiredScopes: []string{},
		}
		var key string
		if p.APIKey != nil {
			key = *p.APIKey
		}
		ctx, err = authAPIKeyFn(ctx, key, &sc)
		if err != nil {
			return nil, err
		}
		return s.Refresh(ctx, p)
	}
}

// NewLogoutEndpoint returns an endpoint function that calls the method
// "logout" of service "session".
func NewLogoutEndpoint(s Service, authJWTFn security.AuthJWTFunc, authAPIKeyFn security.AuthAPIKeyFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*LogoutPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{},
			RequiredScopes: []string{},
		}
		var token string
		if p.Authorization != nil {
			token = *p.Authorization
		}
		ctx, err = authJWTFn(ctx, token, &sc)
		if err == nil {
			sc := security.APIKeyScheme{
				Name:           "api_key",
				Scopes:         []string{},
				RequiredScopes: []string{},
			}
			var key string
			if p.APIKey != nil {
				key = *p.APIKey
			}
			ctx, err = authAPIKeyFn(ctx, key, &sc)
		}
		if err != nil {
			return nil, err
		}
		return nil, s.Logout(ctx, p)
	}
}

// NewLogoutOtherEndpoint returns an endpoint function that calls the method
// "logout-other" of service "session".
func NewLogoutOtherEndpoint(s Service, authJWTFn security.AuthJWTFunc, authAPIKeyFn security.AuthAPIKeyFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*LogoutOtherPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{},
			RequiredScopes: []string{},
		}
		var token string
		if p.Authorization != nil {
			token = *p.Authorization
		}
		ctx, err = authJWTFn(ctx, token, &sc)
		if err == nil {
			sc := security.APIKeyScheme{
				Name:           "api_key",
				Scopes:         []string{},
				RequiredScopes: []string{},
			}
			var key string
			if p.APIKey != nil {
				key = *p.APIKey
			}
			ctx, err = authAPIKeyFn(ctx, key, &sc)
		}
		if err != nil {
			return nil, err
		}
		return nil, s.LogoutOther(ctx, p)
	}
}

// NewLogoutSpecificEndpoint returns an endpoint function that calls the method
// "logout-specific" of service "session".
func NewLogoutSpecificEndpoint(s Service, authJWTFn security.AuthJWTFunc, authAPIKeyFn security.AuthAPIKeyFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*LogoutSpecificPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{},
			RequiredScopes: []string{},
		}
		var token string
		if p.Authorization != nil {
			token = *p.Authorization
		}
		ctx, err = authJWTFn(ctx, token, &sc)
		if err == nil {
			sc := security.APIKeyScheme{
				Name:           "api_key",
				Scopes:         []string{},
				RequiredScopes: []string{},
			}
			var key string
			if p.APIKey != nil {
				key = *p.APIKey
			}
			ctx, err = authAPIKeyFn(ctx, key, &sc)
		}
		if err != nil {
			return nil, err
		}
		return nil, s.LogoutSpecific(ctx, p)
	}
}

// NewGetSessionsEndpoint returns an endpoint function that calls the method
// "get-sessions" of service "session".
func NewGetSessionsEndpoint(s Service, authJWTFn security.AuthJWTFunc, authAPIKeyFn security.AuthAPIKeyFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*GetSessionsPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{},
			RequiredScopes: []string{},
		}
		var token string
		if p.Authorization != nil {
			token = *p.Authorization
		}
		ctx, err = authJWTFn(ctx, token, &sc)
		if err == nil {
			sc := security.APIKeyScheme{
				Name:           "api_key",
				Scopes:         []string{},
				RequiredScopes: []string{},
			}
			var key string
			if p.APIKey != nil {
				key = *p.APIKey
			}
			ctx, err = authAPIKeyFn(ctx, key, &sc)
		}
		if err != nil {
			return nil, err
		}
		res, err := s.GetSessions(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedAllSessions(res, "default")
		return vres, nil
	}
}

// NewRedeemTokenEndpoint returns an endpoint function that calls the method
// "redeemToken" of service "session".
func NewRedeemTokenEndpoint(s Service, authAPIKeyFn security.AuthAPIKeyFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*RedeemTokenPayload)
		var err error
		sc := security.APIKeyScheme{
			Name:           "api_key",
			Scopes:         []string{},
			RequiredScopes: []string{},
		}
		var key string
		if p.APIKey != nil {
			key = *p.APIKey
		}
		ctx, err = authAPIKeyFn(ctx, key, &sc)
		if err != nil {
			return nil, err
		}
		return s.RedeemToken(ctx, p)
	}
}

// NewCleanSessionsEndpoint returns an endpoint function that calls the method
// "clean-sessions" of service "session".
func NewCleanSessionsEndpoint(s Service, authJWTFn security.AuthJWTFunc, authAPIKeyFn security.AuthAPIKeyFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*CleanSessionsPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{},
			RequiredScopes: []string{},
		}
		var token string
		if p.Authorization != nil {
			token = *p.Authorization
		}
		ctx, err = authJWTFn(ctx, token, &sc)
		if err == nil {
			sc := security.APIKeyScheme{
				Name:           "api_key",
				Scopes:         []string{},
				RequiredScopes: []string{},
			}
			var key string
			if p.APIKey != nil {
				key = *p.APIKey
			}
			ctx, err = authAPIKeyFn(ctx, key, &sc)
		}
		if err != nil {
			return nil, err
		}
		return nil, s.CleanSessions(ctx, p)
	}
}

// NewCleanLoginTokenEndpoint returns an endpoint function that calls the
// method "clean-login-token" of service "session".
func NewCleanLoginTokenEndpoint(s Service, authJWTFn security.AuthJWTFunc, authAPIKeyFn security.AuthAPIKeyFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*CleanLoginTokenPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{},
			RequiredScopes: []string{},
		}
		var token string
		if p.Authorization != nil {
			token = *p.Authorization
		}
		ctx, err = authJWTFn(ctx, token, &sc)
		if err == nil {
			sc := security.APIKeyScheme{
				Name:           "api_key",
				Scopes:         []string{},
				RequiredScopes: []string{},
			}
			var key string
			if p.APIKey != nil {
				key = *p.APIKey
			}
			ctx, err = authAPIKeyFn(ctx, key, &sc)
		}
		if err != nil {
			return nil, err
		}
		return nil, s.CleanLoginToken(ctx, p)
	}
}

// NewCleanMergeTokenEndpoint returns an endpoint function that calls the
// method "clean-merge-token" of service "session".
func NewCleanMergeTokenEndpoint(s Service, authJWTFn security.AuthJWTFunc, authAPIKeyFn security.AuthAPIKeyFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*CleanMergeTokenPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{},
			RequiredScopes: []string{},
		}
		var token string
		if p.Authorization != nil {
			token = *p.Authorization
		}
		ctx, err = authJWTFn(ctx, token, &sc)
		if err == nil {
			sc := security.APIKeyScheme{
				Name:           "api_key",
				Scopes:         []string{},
				RequiredScopes: []string{},
			}
			var key string
			if p.APIKey != nil {
				key = *p.APIKey
			}
			ctx, err = authAPIKeyFn(ctx, key, &sc)
		}
		if err != nil {
			return nil, err
		}
		return nil, s.CleanMergeToken(ctx, p)
	}
}