package grpc

import (
	"crypto/tls"

	gc "xinhari.com/xinhari/client/grpc"
	gs "xinhari.com/xinhari/server/grpc"
	"xinhari.com/xinhari/service"
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
