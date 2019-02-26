package handlers

import (
	"encoding/json"
	"github.com/saromanov/go-blog/internal/trace"
	"github.com/saromanov/go-blog/internal/platform"
	"github.com/saromanov/go-blog/internal/platform/db/postgresql"
)
// User defines handler for users 
type User struct {
	db *platform.Storage
}

// Create provides inserting user to db
func (u *User) Create(ctx context.Context, log *log.Logger, w http.ResponseWriter, r *http.Request, params map[string]string) error {
	ctx, span := trace.StartSpan(ctx, "handlers.User.Create")
	defer span.End()

	dbConn := u.MasterDB.Copy()
	defer dbConn.Close()

	v, ok := ctx.Value(web.KeyValues).(*web.Values)
	if !ok {
		return web.Shutdown("web value missing from context")
	}

	var newU user.User
	if err := json.Unmarshal(r.Body, &newU); err != nil {
		return errors.Wrap(err, "")
	}
	return nil
}
