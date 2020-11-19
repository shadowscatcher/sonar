package api

import (
	"net/http"

	"github.com/bi-zone/sonar/internal/actions"
)

func (api *API) DNSRecordsCreate(w http.ResponseWriter, r *http.Request) {

	var params actions.DNSRecordsCreateParams

	if err := fromJSON(r, &params); err != nil {
		api.handleError(w, r, err)
		return
	}

	res, err := api.actions.DNSRecordsCreate(r.Context(), params)
	if err != nil {
		api.handleError(w, r, err)
		return
	}

	responseJSON(w, res, http.StatusCreated)
}

func (api *API) DNSRecordsDelete(w http.ResponseWriter, r *http.Request) {

	var params actions.DNSRecordsDeleteParams

	if err := fromPath(r, &params); err != nil {
		api.handleError(w, r, err)
		return
	}

	res, err := api.actions.DNSRecordsDelete(r.Context(), params)
	if err != nil {
		api.handleError(w, r, err)
		return
	}

	responseJSON(w, res, http.StatusOK)
}

func (api *API) DNSRecordsList(w http.ResponseWriter, r *http.Request) {

	var params actions.DNSRecordsListParams

	if err := fromPath(r, &params); err != nil {
		api.handleError(w, r, err)
		return
	}

	res, err := api.actions.DNSRecordsList(r.Context(), params)
	if err != nil {
		api.handleError(w, r, err)
		return
	}

	responseJSON(w, res, http.StatusOK)
}

func (api *API) PayloadsCreate(w http.ResponseWriter, r *http.Request) {

	var params actions.PayloadsCreateParams

	if err := fromJSON(r, &params); err != nil {
		api.handleError(w, r, err)
		return
	}

	res, err := api.actions.PayloadsCreate(r.Context(), params)
	if err != nil {
		api.handleError(w, r, err)
		return
	}

	responseJSON(w, res, http.StatusCreated)
}

func (api *API) PayloadsDelete(w http.ResponseWriter, r *http.Request) {

	var params actions.PayloadsDeleteParams

	if err := fromPath(r, &params); err != nil {
		api.handleError(w, r, err)
		return
	}

	res, err := api.actions.PayloadsDelete(r.Context(), params)
	if err != nil {
		api.handleError(w, r, err)
		return
	}

	responseJSON(w, res, http.StatusOK)
}

func (api *API) PayloadsList(w http.ResponseWriter, r *http.Request) {

	var params actions.PayloadsListParams

	if err := fromQuery(r, &params); err != nil {
		api.handleError(w, r, err)
		return
	}

	res, err := api.actions.PayloadsList(r.Context(), params)
	if err != nil {
		api.handleError(w, r, err)
		return
	}

	responseJSON(w, res, http.StatusOK)
}

func (api *API) PayloadsUpdate(w http.ResponseWriter, r *http.Request) {

	var params actions.PayloadsUpdateParams

	if err := fromJSON(r, &params); err != nil {
		api.handleError(w, r, err)
		return
	}

	if err := fromPath(r, &params); err != nil {
		api.handleError(w, r, err)
		return
	}

	res, err := api.actions.PayloadsUpdate(r.Context(), params)
	if err != nil {
		api.handleError(w, r, err)
		return
	}

	responseJSON(w, res, http.StatusOK)
}

func (api *API) UserCurrent(w http.ResponseWriter, r *http.Request) {

	res, err := api.actions.UserCurrent(r.Context())
	if err != nil {
		api.handleError(w, r, err)
		return
	}

	responseJSON(w, res, http.StatusOK)
}

func (api *API) UsersCreate(w http.ResponseWriter, r *http.Request) {

	var params actions.UsersCreateParams

	if err := fromJSON(r, &params); err != nil {
		api.handleError(w, r, err)
		return
	}

	res, err := api.actions.UsersCreate(r.Context(), params)
	if err != nil {
		api.handleError(w, r, err)
		return
	}

	responseJSON(w, res, http.StatusCreated)
}

func (api *API) UsersDelete(w http.ResponseWriter, r *http.Request) {

	var params actions.UsersDeleteParams

	if err := fromPath(r, &params); err != nil {
		api.handleError(w, r, err)
		return
	}

	res, err := api.actions.UsersDelete(r.Context(), params)
	if err != nil {
		api.handleError(w, r, err)
		return
	}

	responseJSON(w, res, http.StatusOK)
}