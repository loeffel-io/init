package init

import (
	"github.com/getsentry/raven-go"
	"github.com/google/gops/agent"
	"github.com/joho/godotenv"
	"github.com/loeffel-io/helper"
	"log"
	"os"
)

// Load env file
func Dotenv() {
	if err := godotenv.Load(helper.LoadFile(".env")); err != nil {
		log.Fatal(err.Error())
	}
}

// Setup Sentry DNS
func Sentry() {
	if err := raven.SetDSN(os.Getenv("SENTRY")); err != nil {
		raven.CaptureError(err, nil)
		log.Fatal(err.Error())
	}
}

// Setup agent
func Agent() {
	if os.Getenv("ENV") != "local" {
		return
	}

	if err := agent.Listen(agent.Options{
		ShutdownCleanup: true,
	}); err != nil {
		raven.CaptureError(err, nil)
		log.Fatal(err)
	}
}
