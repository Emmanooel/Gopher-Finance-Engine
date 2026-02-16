package repository

import (
	"context"
	"gopher-finance-engine/internal/domain/entity"
	"gopher-finance-engine/internal/domain/infra/repository"
)

type TaxAlertsRepository struct{}

func NewTaxAlertsRepository() repository.TaxAlertsRepositoryI {
	return &TaxAlertsRepository{}
}

func (t *TaxAlertsRepository) CreateTaxAlerts(ctx context.Context, alerts entity.TaxAlert) error {
	return nil
}
