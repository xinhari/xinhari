// Package os runs processes locally
package os

import (
	"xinhari.com/xinhari/runtime/local/process"
)

type Process struct{}

func NewProcess(opts ...process.Option) process.Process {
	return &Process{}
}
