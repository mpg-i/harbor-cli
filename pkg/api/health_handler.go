package api

import (
	"context"
	"fmt"

	"github.com/goharbor/go-client/pkg/sdk/v2.0/client/health"
	"github.com/goharbor/harbor-cli/pkg/utils"
	"github.com/spf13/viper"
)

func GetHealth() (*health.GetHealthOK, error) {
	ctx, client, err := utils.ContextWithClient()
	if err != nil {
		return err
	}
	params := health.NewGetHealthParams().WithContext(ctx)

	response, err := client.Health.GetHealth(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("error getting health status: %w", err)
	}

	return response, nil
}

 