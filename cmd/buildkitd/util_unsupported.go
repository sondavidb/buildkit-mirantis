//go:build !linux
// +build !linux

package main

import (
	"github.com/moby/sys/user"
	"github.com/pkg/errors"
)

func parseIdentityMapping(str string) (*user.IdentityMapping, error) {
	if str == "" {
		return nil, nil
	}
	return nil, errors.New("user namespaces are only supported on linux")
}
