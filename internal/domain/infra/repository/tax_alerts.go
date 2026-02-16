package repository

import (
	"context"
	"gopher-finance-engine/internal/domain/entity"
)

type TaxAlertsRepositoryI interface {
	CreateTaxAlerts(ctx context.Context, alerts entity.TaxAlert) error
}
