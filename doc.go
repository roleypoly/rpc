// Package rpc should never be imported. Use a subpackage.
package rpc

//go:generate go run hack/gen.go

func init() {
	panic("Do not import this.")
}
