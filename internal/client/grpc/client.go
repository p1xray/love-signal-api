package grpcclient

import (
	lsuserspb "github.com/p1xray/love-signal-protos/gen/go/users"
	qrcodepb "github.com/p1xray/pxr-qrcode/pkg/grpc/gen/go/qrcode"
	urlshortenerpb "github.com/p1xray/pxr-url-shortener/pkg/grpc/gen/go/urlshortener"
)

// GRPCClient provides gRPC clients.
type GRPCClient struct {
	Users        lsuserspb.UsersClient
	UrlShortener urlshortenerpb.UrlShortenerClient
	QRCode       qrcodepb.QrCodeClient
}

// New creates new gRPC client instance.
func New(
	users lsuserspb.UsersClient,
	urlShortener urlshortenerpb.UrlShortenerClient,
	qrCode qrcodepb.QrCodeClient,
) *GRPCClient {
	return &GRPCClient{
		Users:        users,
		UrlShortener: urlShortener,
		QRCode:       qrCode,
	}
}
