// Package main implements a gRPC server that handles Staffjoy emails.
package main

import (
	"os"
	"sync"

	"github.com/keighl/mandrill"
	"github.com/sirupsen/logrus"

	"v2.staffjoy.com/environments"
)

const (
	// ServiceName identifies this app in logs
	ServiceName         = "emailapi"
	fromName            = "Staffjoy"
	from                = "staffjoy@126.com"
	staffjoyEmailSuffix = "@126.com"
	mandrillTemplate    = "staffjoy-base"
)

var (
	logger *logrus.Entry
	config environments.Config
)

type emailServer struct {
	logger      *logrus.Entry
	errorClient environments.SentryClient
	client      *mandrill.Client
	clientMutex *sync.Mutex
	config      *environments.Config
}

// Setup environments, logger, etc
func init() {
	// Set the ENV environments variable to control dev/stage/prod behavior
	var err error
	config, err = environments.GetConfig(os.Getenv(environments.EnvVar))
	if err != nil {
		panic("Unable to determine emailserver configuration")
	}
	logger = config.GetLogger(ServiceName)
}

func main() {
	logger.Debugf("Boot")
}
