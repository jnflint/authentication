// Code generated by goa v3.0.6, DO NOT EDIT.
//
// password-auth HTTP client encoders and decoders
//
// Command:
// $ goa gen github.com/anshap1719/authentication/design

package client

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"

	passwordauth "github.com/anshap1719/authentication/controllers/gen/password_auth"
	passwordauthviews "github.com/anshap1719/authentication/controllers/gen/password_auth/views"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// BuildRegisterRequest instantiates a HTTP request object with method and path
// set to call the "password-auth" service "register" endpoint
func (c *Client) BuildRegisterRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: RegisterPasswordAuthPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("password-auth", "register", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeRegisterRequest returns an encoder for requests sent to the
// password-auth register server.
func EncodeRegisterRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*passwordauth.RegisterParams)
		if !ok {
			return goahttp.ErrInvalidType("password-auth", "register", "*passwordauth.RegisterParams", v)
		}
		if p.Authorization != nil {
			req.Header.Set("Authorization", *p.Authorization)
		}
		if p.XSession != nil {
			req.Header.Set("X-Session", *p.XSession)
		}
		if p.APIKey != nil {
			req.Header.Set("API-Key", *p.APIKey)
		}
		body := NewRegisterRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("password-auth", "register", err)
		}
		return nil
	}
}

// DecodeRegisterResponse returns a decoder for responses returned by the
// password-auth register endpoint. restoreBody controls whether the response
// body should be restored after having been read.
func DecodeRegisterResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body RegisterResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("password-auth", "register", err)
			}
			var (
				authorization string
				xSession      string
			)
			authorizationRaw := resp.Header.Get("Authorization")
			if authorizationRaw == "" {
				err = goa.MergeErrors(err, goa.MissingFieldError("Authorization", "header"))
			}
			authorization = authorizationRaw
			xSessionRaw := resp.Header.Get("X-Session")
			if xSessionRaw == "" {
				err = goa.MergeErrors(err, goa.MissingFieldError("X-Session", "header"))
			}
			xSession = xSessionRaw
			if err != nil {
				return nil, goahttp.ErrValidationError("password-auth", "register", err)
			}
			p := NewRegisterUserMediaOK(&body, authorization, xSession)
			view := "default"
			vres := &passwordauthviews.UserMedia{p, view}
			if err = passwordauthviews.ValidateUserMedia(vres); err != nil {
				return nil, goahttp.ErrValidationError("password-auth", "register", err)
			}
			res := passwordauth.NewUserMedia(vres)
			return res, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("password-auth", "register", resp.StatusCode, string(body))
		}
	}
}

// BuildLoginRequest instantiates a HTTP request object with method and path
// set to call the "password-auth" service "login" endpoint
func (c *Client) BuildLoginRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: LoginPasswordAuthPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("password-auth", "login", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeLoginRequest returns an encoder for requests sent to the password-auth
// login server.
func EncodeLoginRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*passwordauth.LoginParams)
		if !ok {
			return goahttp.ErrInvalidType("password-auth", "login", "*passwordauth.LoginParams", v)
		}
		if p.APIKey != nil {
			req.Header.Set("API-Key", *p.APIKey)
		}
		values := req.URL.Query()
		if p.Token != nil {
			values.Add("token", *p.Token)
		}
		req.URL.RawQuery = values.Encode()
		body := NewLoginRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("password-auth", "login", err)
		}
		return nil
	}
}

// DecodeLoginResponse returns a decoder for responses returned by the
// password-auth login endpoint. restoreBody controls whether the response body
// should be restored after having been read.
func DecodeLoginResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body LoginResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("password-auth", "login", err)
			}
			var (
				authorization string
				xSession      string
			)
			authorizationRaw := resp.Header.Get("Authorization")
			if authorizationRaw == "" {
				err = goa.MergeErrors(err, goa.MissingFieldError("Authorization", "header"))
			}
			authorization = authorizationRaw
			xSessionRaw := resp.Header.Get("X-Session")
			if xSessionRaw == "" {
				err = goa.MergeErrors(err, goa.MissingFieldError("X-Session", "header"))
			}
			xSession = xSessionRaw
			if err != nil {
				return nil, goahttp.ErrValidationError("password-auth", "login", err)
			}
			p := NewLoginUserMediaOK(&body, authorization, xSession)
			view := "default"
			vres := &passwordauthviews.UserMedia{p, view}
			if err = passwordauthviews.ValidateUserMedia(vres); err != nil {
				return nil, goahttp.ErrValidationError("password-auth", "login", err)
			}
			res := passwordauth.NewUserMedia(vres)
			return res, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("password-auth", "login", resp.StatusCode, string(body))
		}
	}
}

// BuildRemoveRequest instantiates a HTTP request object with method and path
// set to call the "password-auth" service "remove" endpoint
func (c *Client) BuildRemoveRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: RemovePasswordAuthPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("password-auth", "remove", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeRemoveRequest returns an encoder for requests sent to the
// password-auth remove server.
func EncodeRemoveRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*passwordauth.RemovePayload)
		if !ok {
			return goahttp.ErrInvalidType("password-auth", "remove", "*passwordauth.RemovePayload", v)
		}
		if p.Authorization != nil {
			req.Header.Set("Authorization", *p.Authorization)
		}
		if p.XSession != nil {
			req.Header.Set("X-Session", *p.XSession)
		}
		if p.APIKey != nil {
			req.Header.Set("API-Key", *p.APIKey)
		}
		return nil
	}
}

// DecodeRemoveResponse returns a decoder for responses returned by the
// password-auth remove endpoint. restoreBody controls whether the response
// body should be restored after having been read.
func DecodeRemoveResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			return nil, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("password-auth", "remove", resp.StatusCode, string(body))
		}
	}
}

// BuildChangePasswordRequest instantiates a HTTP request object with method
// and path set to call the "password-auth" service "change-password" endpoint
func (c *Client) BuildChangePasswordRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ChangePasswordPasswordAuthPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("password-auth", "change-password", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeChangePasswordRequest returns an encoder for requests sent to the
// password-auth change-password server.
func EncodeChangePasswordRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*passwordauth.ChangePasswordParams)
		if !ok {
			return goahttp.ErrInvalidType("password-auth", "change-password", "*passwordauth.ChangePasswordParams", v)
		}
		req.Header.Set("Authorization", p.Authorization)
		req.Header.Set("X-Session", p.XSession)
		if p.APIKey != nil {
			req.Header.Set("API-Key", *p.APIKey)
		}
		body := NewChangePasswordRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("password-auth", "change-password", err)
		}
		return nil
	}
}

// DecodeChangePasswordResponse returns a decoder for responses returned by the
// password-auth change-password endpoint. restoreBody controls whether the
// response body should be restored after having been read.
func DecodeChangePasswordResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			return nil, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("password-auth", "change-password", resp.StatusCode, string(body))
		}
	}
}

// BuildResetRequest instantiates a HTTP request object with method and path
// set to call the "password-auth" service "reset" endpoint
func (c *Client) BuildResetRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ResetPasswordAuthPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("password-auth", "reset", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeResetRequest returns an encoder for requests sent to the password-auth
// reset server.
func EncodeResetRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*passwordauth.ResetPayload)
		if !ok {
			return goahttp.ErrInvalidType("password-auth", "reset", "*passwordauth.ResetPayload", v)
		}
		if p.APIKey != nil {
			req.Header.Set("API-Key", *p.APIKey)
		}
		values := req.URL.Query()
		if p.Email != nil {
			values.Add("email", *p.Email)
		}
		req.URL.RawQuery = values.Encode()
		return nil
	}
}

// DecodeResetResponse returns a decoder for responses returned by the
// password-auth reset endpoint. restoreBody controls whether the response body
// should be restored after having been read.
func DecodeResetResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			return nil, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("password-auth", "reset", resp.StatusCode, string(body))
		}
	}
}

// BuildConfirmResetRequest instantiates a HTTP request object with method and
// path set to call the "password-auth" service "confirm-reset" endpoint
func (c *Client) BuildConfirmResetRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ConfirmResetPasswordAuthPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("password-auth", "confirm-reset", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeConfirmResetRequest returns an encoder for requests sent to the
// password-auth confirm-reset server.
func EncodeConfirmResetRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*passwordauth.ResetPasswordParams)
		if !ok {
			return goahttp.ErrInvalidType("password-auth", "confirm-reset", "*passwordauth.ResetPasswordParams", v)
		}
		if p.APIKey != nil {
			req.Header.Set("API-Key", *p.APIKey)
		}
		body := NewConfirmResetRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("password-auth", "confirm-reset", err)
		}
		return nil
	}
}

// DecodeConfirmResetResponse returns a decoder for responses returned by the
// password-auth confirm-reset endpoint. restoreBody controls whether the
// response body should be restored after having been read.
func DecodeConfirmResetResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			return nil, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("password-auth", "confirm-reset", resp.StatusCode, string(body))
		}
	}
}

// BuildCheckEmailAvailableRequest instantiates a HTTP request object with
// method and path set to call the "password-auth" service
// "check-email-available" endpoint
func (c *Client) BuildCheckEmailAvailableRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: CheckEmailAvailablePasswordAuthPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("password-auth", "check-email-available", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeCheckEmailAvailableRequest returns an encoder for requests sent to the
// password-auth check-email-available server.
func EncodeCheckEmailAvailableRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*passwordauth.CheckEmailAvailablePayload)
		if !ok {
			return goahttp.ErrInvalidType("password-auth", "check-email-available", "*passwordauth.CheckEmailAvailablePayload", v)
		}
		if p.APIKey != nil {
			req.Header.Set("API-Key", *p.APIKey)
		}
		values := req.URL.Query()
		if p.Email != nil {
			values.Add("email", *p.Email)
		}
		req.URL.RawQuery = values.Encode()
		return nil
	}
}

// DecodeCheckEmailAvailableResponse returns a decoder for responses returned
// by the password-auth check-email-available endpoint. restoreBody controls
// whether the response body should be restored after having been read.
func DecodeCheckEmailAvailableResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body bool
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("password-auth", "check-email-available", err)
			}
			return body, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("password-auth", "check-email-available", resp.StatusCode, string(body))
		}
	}
}

// BuildCheckPhoneAvailableRequest instantiates a HTTP request object with
// method and path set to call the "password-auth" service
// "check-phone-available" endpoint
func (c *Client) BuildCheckPhoneAvailableRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: CheckPhoneAvailablePasswordAuthPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("password-auth", "check-phone-available", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeCheckPhoneAvailableRequest returns an encoder for requests sent to the
// password-auth check-phone-available server.
func EncodeCheckPhoneAvailableRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*passwordauth.CheckPhoneAvailablePayload)
		if !ok {
			return goahttp.ErrInvalidType("password-auth", "check-phone-available", "*passwordauth.CheckPhoneAvailablePayload", v)
		}
		if p.APIKey != nil {
			req.Header.Set("API-Key", *p.APIKey)
		}
		values := req.URL.Query()
		if p.Phone != nil {
			values.Add("phone", *p.Phone)
		}
		req.URL.RawQuery = values.Encode()
		return nil
	}
}

// DecodeCheckPhoneAvailableResponse returns a decoder for responses returned
// by the password-auth check-phone-available endpoint. restoreBody controls
// whether the response body should be restored after having been read.
func DecodeCheckPhoneAvailableResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body bool
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("password-auth", "check-phone-available", err)
			}
			return body, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("password-auth", "check-phone-available", resp.StatusCode, string(body))
		}
	}
}
