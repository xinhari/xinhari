package grpc

import (
	"crypto/tls"

	gc "xinhari.com/client/grpc"
	gs "xinhari.com/server/grpc"
	"xinhari.com/service"
)

// WithTLS sets the TLS config for the service
func WithTLS(t *tls.Config) service.Option {
	return func(o *service.Options) {
		o.Client.Init(
			gc.AuthTLS(t),
		)
		o.Server.Init(
			gs.AuthTLS(t),
		)
	}
}
