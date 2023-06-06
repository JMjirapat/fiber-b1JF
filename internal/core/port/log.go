package port

import "fiber/internal/core/model"

type LogRepo interface {
	Create(body *model.UsageLog) error
}
