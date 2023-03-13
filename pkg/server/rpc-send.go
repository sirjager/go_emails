package server

import (
	"context"

	"github.com/sirjager/go_emails/pkg/mail"
	rpc "github.com/sirjager/rpcs/emails/go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (srvr *RPCEmailsServer) EmailsSendMail(ctx context.Context, req *rpc.EmailsSendMailRequest) (*rpc.EmailsSendMailResponse, error) {
	// This only validates if the request is authorized or not
	if err := srvr.authorize(ctx); err != nil {
		return nil, unAuthenticatedError(err)
	}
	mail := mail.Mail{To: req.Receiver, From: req.Sender, Subject: req.Subject, Body: req.Body}
	err := srvr.smtp.SendMail(mail)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &rpc.EmailsSendMailResponse{Message: "email sent successfully"}, nil
}
