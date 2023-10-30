package fmsghosts

import (
	"context"
	"net"
)

func LookupHosts(domain *string) ([]string, error) {
	r := net.Resolver{}
	txts, err := r.LookupTXT(context.Background(), *domain)
	if err != nil {
		return nil, err
	}
	return txts, nil
}