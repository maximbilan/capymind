//coverage:ignore file

package database

import "context"

type Database interface {
	Connect(ctx *context.Context)
	Close()
}
