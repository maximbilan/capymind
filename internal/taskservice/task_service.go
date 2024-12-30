//coverage:ignore file

package taskservice

import "context"

type TaskService interface {
	Connect(ctx *context.Context)
	Schedule()
	Close()
}
