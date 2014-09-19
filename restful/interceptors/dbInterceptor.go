package interceptors

import (
	"EpiCraft/restful/db"
	"github.com/gocarina/pi"
)

const (
	httpSeed = 0xdeadbeef
)

// CreateMongoSession creates a new session for the request context.
func CreateMongoSession(c *pi.RequestContext) error {
	db.Copy(httpSeed)
	return nil
}

// CloseMongoSession closes the MongoDB session created for the request context.
func CloseMongoSession(c *pi.RequestContext) error {
	db.CloseCopy(httpSeed)
	return nil
}
