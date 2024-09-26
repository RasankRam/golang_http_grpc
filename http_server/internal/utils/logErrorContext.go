package utils

import (
	"context"
	"github.com/mdobak/go-xerrors"
	"log/slog"
)

func LogErrorContext(logger *slog.Logger, ctx context.Context, error error) {
	err := xerrors.New(error)
	logger.ErrorContext(ctx, "", slog.Any("error", err))
}
