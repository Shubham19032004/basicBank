package gapi

import "context"

type Metadata struct {
	UserAgent string
	ClientIp  string
}

func (server *Server) extractMetadata(ctx context.Context) *Metadata {
	mtdt:=&Metadata{}
	return mtdt
}
