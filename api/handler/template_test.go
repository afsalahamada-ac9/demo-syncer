/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"sudhagar/glad/api/presenter"
	"sudhagar/glad/entity"

	mock "sudhagar/glad/usecase/template/mock"

	"github.com/codegangsta/negroni"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

const (
	tenantAlice       entity.ID = 13790492210917015554
	tenantNonExistent entity.ID = 0
)

func Test_listTemplates(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeTemplateHandlers(r, *n, service)
	path, err := r.GetRoute("listTemplates").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/templates", path)
	tmpl := &entity.Template{
		ID:       entity.NewID(),
		TenantID: tenantAlice,
		Name:     "default-0",
	}
	service.EXPECT().GetCount(tmpl.TenantID).Return(1)
	service.EXPECT().
		ListTemplates(tmpl.TenantID).
		Return([]*entity.Template{tmpl}, nil)
	ts := httptest.NewServer(listTemplates(service))
	defer ts.Close()

	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodGet, ts.URL, nil)
	req.Header.Set(httpHeaderTenantID, tenantAlice.String())
	res, err := client.Do(req)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func Test_listTemplates_NotFound(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockUseCase(controller)
	ts := httptest.NewServer(listTemplates(service))
	defer ts.Close()
	tenantID := tenantAlice
	service.EXPECT().GetCount(tenantID).Return(0)
	service.EXPECT().
		SearchTemplates(tenantID, "non-existent").
		Return(nil, entity.ErrNotFound)

	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodGet,
		ts.URL+"?search=non-existent",
		nil)
	req.Header.Set(httpHeaderTenantID, tenantAlice.String())
	res, err := client.Do(req)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusNotFound, res.StatusCode)
}

func Test_listTemplates_Search(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockUseCase(controller)
	tmpl := &entity.Template{
		ID:       entity.NewID(),
		TenantID: tenantAlice,
		Name:     "default-0",
	}
	service.EXPECT().GetCount(tmpl.TenantID).Return(1)
	service.EXPECT().
		SearchTemplates(tmpl.TenantID, "default").
		Return([]*entity.Template{tmpl}, nil)
	ts := httptest.NewServer(listTemplates(service))
	defer ts.Close()

	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodGet,
		ts.URL+"?search=default",
		nil)
	req.Header.Set(httpHeaderTenantID, tenantAlice.String())
	res, err := client.Do(req)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func Test_createTemplate(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeTemplateHandlers(r, *n, service)
	path, err := r.GetRoute("createTemplate").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/templates", path)

	id := entity.NewID()
	service.EXPECT().
		CreateTemplate(gomock.Any(),
			gomock.Any(),
			gomock.Any(),
			gomock.Any()).
		Return(id, nil)
	h := createTemplate(service)

	ts := httptest.NewServer(h)
	defer ts.Close()

	payload := struct {
		TenantID entity.ID           `json:"tenant_id"`
		Name     string              `json:"name"`
		Type     entity.TemplateType `json:"type"`
		Content  string              `json:"content"`
	}{TenantID: tenantAlice,
		Name:    "default-0",
		Type:    (entity.TemplateText),
		Content: "This is a default message"}
	payloadBytes, err := json.Marshal(payload)
	assert.Nil(t, err)

	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodPost,
		ts.URL+"/v1/templates",
		bytes.NewReader(payloadBytes))
	req.Header.Set(httpHeaderTenantID, tenantAlice.String())
	req.Header.Set("Content-Type", "application/json")
	res, err := client.Do(req)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, res.StatusCode)

	var tmpl *presenter.Template
	json.NewDecoder(res.Body).Decode(&tmpl)
	assert.Equal(t, id, tmpl.ID)
	assert.Equal(t, payload.Content, tmpl.Content)
	assert.Equal(t, payload.Name, tmpl.Name)
	assert.Equal(t, payload.Type, tmpl.Type)
	assert.Equal(t, tenantAlice.String(), res.Header.Get(httpHeaderTenantID))
}

func Test_getTemplate(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeTemplateHandlers(r, *n, service)
	path, err := r.GetRoute("getTemplate").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/templates/{id}", path)
	tmpl := &entity.Template{
		ID:       entity.NewID(),
		TenantID: tenantAlice,
		Name:     "default-0",
		Type:     entity.TemplateText,
		Content:  "This is a default message",
	}
	service.EXPECT().
		GetTemplate(tmpl.ID).
		Return(tmpl, nil)
	handler := getTemplate(service)
	r.Handle("/v1/templates/{id}", handler)
	ts := httptest.NewServer(r)
	defer ts.Close()
	res, err := http.Get(ts.URL + "/v1/templates/" + tmpl.ID.String())
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)

	// presenter.Template is returned by the api (http) server
	var d *presenter.Template
	json.NewDecoder(res.Body).Decode(&d)
	assert.NotNil(t, d)

	assert.Equal(t, tmpl.ID, d.ID)
	assert.Equal(t, tmpl.Content, d.Content)
	assert.Equal(t, tmpl.Name, d.Name)
	assert.Equal(t, tmpl.Type, d.Type)
	assert.Equal(t, tenantAlice.String(), res.Header.Get(httpHeaderTenantID))
}

func Test_deleteTemplate(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeTemplateHandlers(r, *n, service)
	path, err := r.GetRoute("deleteTemplate").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/templates/{id}", path)

	id := entity.NewID()
	service.EXPECT().DeleteTemplate(id).Return(nil)
	handler := deleteTemplate(service)
	req, _ := http.NewRequest("DELETE", "/v1/templates/"+id.String(), nil)
	r.Handle("/v1/templates/{id}", handler).Methods("DELETE", "OPTIONS")
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}

func Test_deleteTemplateNonExistent(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeTemplateHandlers(r, *n, service)
	path, err := r.GetRoute("deleteTemplate").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/templates/{id}", path)

	id := entity.NewID()
	service.EXPECT().DeleteTemplate(id).Return(entity.ErrNotFound)
	handler := deleteTemplate(service)
	req, _ := http.NewRequest("DELETE", "/v1/templates/"+id.String(), nil)
	r.Handle("/v1/templates/{id}", handler).Methods("DELETE", "OPTIONS")
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusNotFound, rr.Code)
}

// TODO: Test case for updating template
