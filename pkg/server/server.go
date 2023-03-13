package server

import (
	"time"

	"github.com/rs/zerolog"
	"github.com/sirjager/go_emails/cfg"
	"github.com/sirjager/go_emails/pkg/mail"

	rpc "github.com/sirjager/rpcs/emails/go"
)

type RPCEmailsServer struct {
	rpc.UnimplementedEmailsServer
	startTime time.Time
	config    cfg.Config
	logger    zerolog.Logger
	smtp      mail.SMTP
}

func NewEmailsServer(startTime time.Time, logger zerolog.Logger, config cfg.Config) (*RPCEmailsServer, error) {
	mailSmtp, err := mail.NewSMTP(config)
	if err != nil {
		return nil, err
	}

	srvic := &RPCEmailsServer{logger: logger, config: config, smtp: *mailSmtp, startTime: startTime}
	return srvic, nil
}
