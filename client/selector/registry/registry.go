// Package registry uses the go-micro registry for selection
package registry

import (
	"xinhari.com/client/selector"
)

// NewSelector returns a new registry selector
func NewSelector(opts ...selector.Option) selector.Selector {
	return selector.NewSelector(opts...)
}
