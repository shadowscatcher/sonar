package apiclient

import (
	"context"

	"github.com/bi-zone/sonar/internal/actions"
	"github.com/bi-zone/sonar/internal/utils/errors"
)

func (c *Client) DNSRecordsCreate(ctx context.Context, params actions.DNSRecordsCreateParams) (actions.DNSRecordsCreateResult, errors.Error) {
	var res actions.DNSRecordsCreateResult

	err := handle(c.client.R().
		SetBody(params).
		SetError(&APIError{}).
		SetResult(&res).
		SetContext(ctx).
		Post("/dnsrecords"))

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Client) DNSRecordsDelete(ctx context.Context, params actions.DNSRecordsDeleteParams) (actions.DNSRecordsDeleteResult, errors.Error) {
	var res actions.DNSRecordsDeleteResult

	err := handle(c.client.R().
		SetPathParams(toPath(params)).
		SetError(&APIError{}).
		SetResult(&res).
		SetContext(ctx).
		Delete("/dnsrecords/{payloadName}/{name}/{type}"))

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Client) DNSRecordsList(ctx context.Context, params actions.DNSRecordsListParams) (actions.DNSRecordsListResult, errors.Error) {
	var res actions.DNSRecordsListResult

	err := handle(c.client.R().
		SetPathParams(toPath(params)).
		SetError(&APIError{}).
		SetResult(&res).
		SetContext(ctx).
		Get("/dnsrecords/{payloadName}"))

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Client) PayloadsCreate(ctx context.Context, params actions.PayloadsCreateParams) (actions.PayloadsCreateResult, errors.Error) {
	var res actions.PayloadsCreateResult

	err := handle(c.client.R().
		SetBody(params).
		SetError(&APIError{}).
		SetResult(&res).
		SetContext(ctx).
		Post("/payloads"))

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Client) PayloadsDelete(ctx context.Context, params actions.PayloadsDeleteParams) (actions.PayloadsDeleteResult, errors.Error) {
	var res actions.PayloadsDeleteResult

	err := handle(c.client.R().
		SetPathParams(toPath(params)).
		SetError(&APIError{}).
		SetResult(&res).
		SetContext(ctx).
		Delete("/payloads/{name}"))

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Client) PayloadsList(ctx context.Context, params actions.PayloadsListParams) (actions.PayloadsListResult, errors.Error) {
	var res actions.PayloadsListResult

	err := handle(c.client.R().
		SetQueryParamsFromValues(toQuery(params)).
		SetError(&APIError{}).
		SetResult(&res).
		SetContext(ctx).
		Get("/payloads"))

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Client) PayloadsUpdate(ctx context.Context, params actions.PayloadsUpdateParams) (actions.PayloadsUpdateResult, errors.Error) {
	var res actions.PayloadsUpdateResult

	err := handle(c.client.R().
		SetBody(params).
		SetPathParams(toPath(params)).
		SetError(&APIError{}).
		SetResult(&res).
		SetContext(ctx).
		Put("/payloads/{name}"))

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Client) UserCurrent(ctx context.Context) (actions.UserCurrentResult, errors.Error) {
	var res actions.UserCurrentResult

	err := handle(c.client.R().SetError(&APIError{}).
		SetResult(&res).
		SetContext(ctx).
		Get("/user"))

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Client) UsersCreate(ctx context.Context, params actions.UsersCreateParams) (actions.UsersCreateResult, errors.Error) {
	var res actions.UsersCreateResult

	err := handle(c.client.R().
		SetBody(params).
		SetError(&APIError{}).
		SetResult(&res).
		SetContext(ctx).
		Post("/users"))

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Client) UsersDelete(ctx context.Context, params actions.UsersDeleteParams) (actions.UsersDeleteResult, errors.Error) {
	var res actions.UsersDeleteResult

	err := handle(c.client.R().
		SetPathParams(toPath(params)).
		SetError(&APIError{}).
		SetResult(&res).
		SetContext(ctx).
		Delete("/users/{name}"))

	if err != nil {
		return nil, err
	}

	return res, nil
}