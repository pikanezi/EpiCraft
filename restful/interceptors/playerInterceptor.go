package interceptors

import (
	"github.com/gocarina/pi"
	"EpiCraft/restful/db"
	"time"
)

// ExtendsSessionInterceptor extends the duration of the session of the player.
func ExtendsSessionInterceptor(c *pi.RequestContext) error {
	if err := db.ExtendsPlayerAccountTTLByAccessToken(c.GetHeader(apiTokenHeader), time.Now().Add(30 * time.Minute)); err != nil {
		return pi.NewError(500, err)
	}
	return nil
}
