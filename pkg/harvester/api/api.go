package api

import (
	"context"
	"errors"

	"github.com/HewlettPackard/galadriel/pkg/common"
	"github.com/HewlettPackard/galadriel/pkg/common/telemetry"
	"github.com/HewlettPackard/galadriel/pkg/harvester/controller"
)

type API interface {
	common.RunnablePlugin
	GetTrustBundle(context.Context, string) (string, error)
	AddTrustBundle(context.Context, string) error
	// ApproveFederationRelationship(context.Context, string) (common.FederationRelationship, error)
	// DenyFederationRelationship(context.Context, string) (common.FederationRelationship, error)
	ManageFederationRelationship(context.Context, string) (common.FederationRelationship, error)
	GetFederationRelationshipsByStatus(context.Context, string) ([]common.FederationRelationship, error)
}

type HTTPApi struct {
	controller controller.HarvesterController
	logger     common.Logger
}

func NewHTTPApi(controller controller.HarvesterController) API {
	// TODO: Add listen address and port
	return &HTTPApi{
		controller: controller,
		logger:     *common.NewLogger(telemetry.HTTPApi),
	}
}

func (a *HTTPApi) Run(ctx context.Context) error {
	a.logger.Info("Starting HTTP API")
	// TODO: implement

	<-ctx.Done()
	return nil
}

func (a *HTTPApi) GetTrustBundle(ctx context.Context, spiffeID string) (string, error) {
	telemetry.Count(ctx, telemetry.HTTPApi, telemetry.TrustBundle, telemetry.Get)
	return "", errors.New("not implemented")
}

func (a *HTTPApi) AddTrustBundle(ctx context.Context, spiffeID string) error {
	telemetry.Count(ctx, telemetry.HTTPApi, telemetry.TrustBundle, telemetry.Add)
	return errors.New("not implemented")
}

// POST: federation-relationship/{relationshipId} {action: approve/deny}
func (a *HTTPApi) ManageFederationRelationship(ctx context.Context, id string) (common.FederationRelationship, error) {

	var fr common.FederationRelationship

	// if body.action == telemetry.Approve {
	// 	telemetry.Count(ctx, telemetry.HTTPApi, telemetry.TrustBundle, telemetry.Approve)
	// 	fr, err := a.controller.ApproveFederationRelationship(ctx, id)
	//  fr.
	// }

	// if body.action == telemetry.Deny {
	// 	telemetry.Count(ctx, telemetry.HTTPApi, telemetry.TrustBundle, telemetry.Approve)
	//   fr, err := a.controller.DenyFederationRelationship(ctx, id)
	// }

	return fr, errors.New("operation not implemented")
}

func (a *HTTPApi) GetFederationRelationshipsByStatus(ctx context.Context, status string) ([]common.FederationRelationship, error) {
	telemetry.Count(ctx, telemetry.HTTPApi, telemetry.FederationRelationship, telemetry.Get)

	var fr []common.FederationRelationship

	fr, err := a.controller.GetFederationRelationshipsByStatus(ctx, status)
	return fr, err
}
