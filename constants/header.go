package constants

import "net/textproto"

var (
	// AuthHeader is the HTTP header used for authentication tokens
	XServiceName  = textproto.CanonicalMIMEHeaderKey("X-Service-Name")
	XApiKey       = textproto.CanonicalMIMEHeaderKey("X-Api-Key")
	XRequestAt    = textproto.CanonicalMIMEHeaderKey("X-Request-At")
	Authorization = textproto.CanonicalMIMEHeaderKey("Authorization")
)
