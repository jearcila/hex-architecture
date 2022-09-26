package genovaService

import (
	"context"
	"encoding/json"

	model "github.com/jearcila/hex-architecture/core/model"

	transactions_context "github.com/mercadolibre/fury_gateway-kit/pkg/g2/framework/transactions/context"
	"github.com/mercadolibre/fury_go-core/pkg/telemetry/tracing"

	"github.com/mercadolibre/fury_go-core/pkg/rusty"
	"github.com/mercadolibre/fury_go-core/pkg/telemetry"
)

const (
	_targetIDBase = "g2.integration.genova."
)

type Client struct {
	Service *service
}

func (c *Client) Authorize(tctx transactions_context.Context, authorization model.FirstOperationRequest) (model.FirstOperationResponse, error) {
	ctx := getCtx(tctx, "authorize")
	params := []rusty.RequestOption{
		rusty.WithBody(authorization),
	}

	saveRequestInJournal(tctx, authorization)
	response, err := c.Service.endpointAuthorize.Post(ctx, params...)
	if err != nil {
		saveErrorInJournal(tctx, err)
		return model.FirstOperationResponse{}, err
	}
	saveResponseInJournal(tctx, response)

	var apiResponse model.FirstOperationResponse
	if err := json.Unmarshal(response.Body, &apiResponse); err != nil {
		return model.FirstOperationResponse{}, err
	}

	return apiResponse, nil
}

func (c *Client) Capture(tctx transactions_context.Context, capture model.CaptureRequest) (model.CaptureResponse, error) {
	ctx := getCtx(tctx, "capture")
	params := []rusty.RequestOption{
		rusty.WithBody(capture),
	}

	saveRequestInJournal(tctx, capture)
	response, err := c.Service.endpointCapture.Post(ctx, params...)
	if err != nil {
		saveErrorInJournal(tctx, err)
		return model.CaptureResponse{}, err
	}
	saveResponseInJournal(tctx, response)

	var apiResponse model.CaptureResponse
	if err := json.Unmarshal(response.Body, &apiResponse); err != nil {
		return model.CaptureResponse{}, err
	}

	return apiResponse, nil
}

func (c *Client) Cancel(tctx transactions_context.Context, cancel model.CancelRequest) (model.CancelResponse, error) {
	ctx := getCtx(tctx, "cancel")
	params := []rusty.RequestOption{
		rusty.WithBody(cancel),
	}

	saveRequestInJournal(tctx, cancel)
	response, err := c.Service.endpointCancel.Post(ctx, params...)
	if err != nil {
		saveErrorInJournal(tctx, err)
		return model.CancelResponse{}, err
	}
	saveResponseInJournal(tctx, response)

	var apiResponse model.CancelResponse
	if err := json.Unmarshal(response.Body, &apiResponse); err != nil {
		return model.CancelResponse{}, err
	}

	return apiResponse, nil
}

func (c *Client) Purchase(tctx transactions_context.Context, purchase model.FirstOperationRequest) (model.FirstOperationResponse, error) {
	ctx := getCtx(tctx, "purchase")
	params := []rusty.RequestOption{
		rusty.WithBody(purchase),
	}

	saveRequestInJournal(tctx, purchase)
	response, err := c.Service.endpointPurchase.Post(ctx, params...)
	if err != nil {
		saveErrorInJournal(tctx, err)
		return model.FirstOperationResponse{}, err
	}
	saveResponseInJournal(tctx, response)

	var apiResponse model.FirstOperationResponse
	if err := json.Unmarshal(response.Body, &apiResponse); err != nil {
		return model.FirstOperationResponse{}, err
	}

	return apiResponse, nil
}

func getCtx(tctx transactions_context.Context, operation string) context.Context {
	ctx := tctx.GetUncancellableRequestContext()
	ctx, sp := telemetry.StartSpan(ctx, "g2-i-genova.genova.client."+operation)
	defer sp.Finish()
	targetID := _targetIDBase + operation
	ctx = tracing.WithEndpointTemplate(ctx, targetID)
	ctx = tracing.WithTargetID(ctx, targetID)
	return ctx
}
