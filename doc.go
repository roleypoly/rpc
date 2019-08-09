// Package rpc should never be imported. Use a subpackage.
package rpc

//go:generate ./hack/gen.sh

func init() {
	panic("Do not import this.")
}
