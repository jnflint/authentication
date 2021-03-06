// Code generated by goa v3.0.6, DO NOT EDIT.
//
// session HTTP client types
//
// Command:
// $ goa gen github.com/anshap1719/authentication/design

package client

import (
	session "github.com/anshap1719/authentication/controllers/gen/session"
	sessionviews "github.com/anshap1719/authentication/controllers/gen/session/views"
	goa "goa.design/goa/v3/pkg"
)

// RedeemTokenRequestBody is the type of the "session" service "redeemToken"
// endpoint HTTP request body.
type RedeemTokenRequestBody struct {
	// A merge token for merging into an account
	Token     string  `form:"token" json:"token" xml:"token"`
	UserAgent *string `form:"User-Agent,omitempty" json:"User-Agent,omitempty" xml:"User-Agent,omitempty"`
}

// GetSessionsResponseBody is the type of the "session" service "get-sessions"
// endpoint HTTP response body.
type GetSessionsResponseBody struct {
	CurrentSession *SessionResponseBody          `form:"currentSession,omitempty" json:"currentSession,omitempty" xml:"currentSession,omitempty"`
	OtherSessions  SessionCollectionResponseBody `form:"otherSessions,omitempty" json:"otherSessions,omitempty" xml:"otherSessions,omitempty"`
}

// SessionResponseBody is used to define fields on response body types.
type SessionResponseBody struct {
	// Unique unchanging session ID
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// ID of the user this session is for
	UserID *string `form:"userId,omitempty" json:"userId,omitempty" xml:"userId,omitempty"`
	// Time that this session was last used
	LastUsed *string `form:"lastUsed,omitempty" json:"lastUsed,omitempty" xml:"lastUsed,omitempty"`
	// The browser and browser version connected with this session
	Browser *string `form:"browser,omitempty" json:"browser,omitempty" xml:"browser,omitempty"`
	// The OS of the system where this session was used
	Os *string `form:"os,omitempty" json:"os,omitempty" xml:"os,omitempty"`
	// The last IP address where this session was used
	IP *string `form:"ip,omitempty" json:"ip,omitempty" xml:"ip,omitempty"`
	// A humanReadable string describing the last known location of the session
	Location *string `form:"location,omitempty" json:"location,omitempty" xml:"location,omitempty"`
	// The latitude of the last known location of the session
	Latitude *string `form:"latitude,omitempty" json:"latitude,omitempty" xml:"latitude,omitempty"`
	// The longitude of the last known location of the session
	Longitude *string `form:"longitude,omitempty" json:"longitude,omitempty" xml:"longitude,omitempty"`
	// Whether the session was from a mobile device
	IsMobile *bool `form:"isMobile,omitempty" json:"isMobile,omitempty" xml:"isMobile,omitempty"`
	// The URL of the Google map to show the location, suitable for using in an img
	// tag
	MapURL *string `form:"mapUrl,omitempty" json:"mapUrl,omitempty" xml:"mapUrl,omitempty"`
}

// SessionCollectionResponseBody is used to define fields on response body
// types.
type SessionCollectionResponseBody []*SessionResponseBody

// NewRedeemTokenRequestBody builds the HTTP request body from the payload of
// the "redeemToken" endpoint of the "session" service.
func NewRedeemTokenRequestBody(p *session.RedeemTokenPayload) *RedeemTokenRequestBody {
	body := &RedeemTokenRequestBody{
		Token:     p.Token,
		UserAgent: p.UserAgent,
	}
	return body
}

// NewRefreshResultOK builds a "session" service "refresh" endpoint result from
// a HTTP "OK" response.
func NewRefreshResultOK(authorization string, xSession string) *session.RefreshResult {
	return &session.RefreshResult{
		Authorization: &authorization,
		XSession:      &xSession,
	}
}

// NewGetSessionsAllSessionsOK builds a "session" service "get-sessions"
// endpoint result from a HTTP "OK" response.
func NewGetSessionsAllSessionsOK(body *GetSessionsResponseBody) *sessionviews.AllSessionsView {
	v := &sessionviews.AllSessionsView{}
	if body.CurrentSession != nil {
		v.CurrentSession = unmarshalSessionResponseBodyToSessionviewsSessionView(body.CurrentSession)
	}
	if body.OtherSessions != nil {
		v.OtherSessions = make([]*sessionviews.SessionView, len(body.OtherSessions))
		for i, val := range body.OtherSessions {
			v.OtherSessions[i] = unmarshalSessionResponseBodyToSessionviewsSessionView(val)
		}
	}
	return v
}

// NewRedeemTokenResultCreated builds a "session" service "redeemToken"
// endpoint result from a HTTP "Created" response.
func NewRedeemTokenResultCreated(authorization string, xSession string) *session.RedeemTokenResult {
	return &session.RedeemTokenResult{
		Authorization: &authorization,
		XSession:      &xSession,
	}
}

// ValidateSessionResponseBody runs the validations defined on
// SessionResponseBody
func ValidateSessionResponseBody(body *SessionResponseBody) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.UserID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("userId", "body"))
	}
	if body.LastUsed == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("lastUsed", "body"))
	}
	if body.Browser == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("browser", "body"))
	}
	if body.Os == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("os", "body"))
	}
	if body.IP == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("ip", "body"))
	}
	if body.Location == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("location", "body"))
	}
	if body.Latitude == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("latitude", "body"))
	}
	if body.Longitude == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("longitude", "body"))
	}
	if body.IsMobile == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("isMobile", "body"))
	}
	if body.MapURL == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("mapUrl", "body"))
	}
	if body.LastUsed != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.lastUsed", *body.LastUsed, goa.FormatDateTime))
	}
	return
}

// ValidateSessionCollectionResponseBody runs the validations defined on
// SessionCollectionResponseBody
func ValidateSessionCollectionResponseBody(body SessionCollectionResponseBody) (err error) {
	for _, e := range body {
		if e != nil {
			if err2 := ValidateSessionResponseBody(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}
