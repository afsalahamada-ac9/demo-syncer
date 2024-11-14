package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"sudhagar/glad/usecase/template"

	"sudhagar/glad/api/presenter"

	"sudhagar/glad/entity"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

const (
	httpHeaderTenantID   = "X-Messaging-TenantID"
	httpHeaderTotalCount = "X-Total-Count"
)

func listTemplates(service template.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading templates"
		var data []*entity.Template
		var err error
		tenant := r.Header.Get(httpHeaderTenantID)
		search := r.URL.Query().Get("search")

		tenantID, err := entity.StringToID(tenant)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Unable to parse tenant id"))
			return
		}

		switch {
		case search == "":
			data, err = service.ListTemplates(tenantID)
		default:
			// TODO: search need to be reworked; need to add a count
			// for search; also need to see how the react-admin generates
			// the search query request
			data, err = service.SearchTemplates(tenantID, search)
		}
		w.Header().Set("Content-Type", "application/json")
		if err != nil && err != entity.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage + ":" + err.Error()))
			return
		}

		total := service.GetCount(tenantID)
		w.Header().Set(httpHeaderTotalCount, strconv.Itoa(total))

		if data == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(errorMessage))
			return
		}
		var toJ []*presenter.Template
		for _, d := range data {
			toJ = append(toJ, &presenter.Template{
				ID:       d.ID,
				TenantID: d.TenantID,
				Name:     d.Name,
				Type:     d.Type,
				Content:  d.Content,
			})
		}
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.Header().Set(httpHeaderTenantID, tenant)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Unable to encode template"))
		}
	})
}

func createTemplate(service template.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error adding template"
		var input struct {
			Name    string              `json:"name"`
			Type    entity.TemplateType `json:"type"`
			Content string              `json:"content"`
		}

		tenant := r.Header.Get(httpHeaderTenantID)
		tenantID, err := entity.StringToID(tenant)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Missing tenant ID"))
			return
		}

		err = json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Unable to decode the data. " + err.Error()))
			return
		}

		id, err := service.CreateTemplate(
			tenantID,
			input.Name,
			input.Type,
			input.Content)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage + ":" + err.Error()))
			return
		}
		toJ := &presenter.Template{
			ID:       id,
			Name:     input.Name,
			Type:     entity.TemplateText,
			Content:  input.Content,
			TenantID: tenantID,
		}

		w.Header().Set(httpHeaderTenantID, tenant)
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
	})
}

func getTemplate(service template.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading template"
		vars := mux.Vars(r)
		id, err := entity.StringToID(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		data, err := service.GetTemplate(id)
		if err != nil && err != entity.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage + ":" + err.Error()))
			return
		}

		if data == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Empty data returned"))
			return
		}

		toJ := &presenter.Template{
			ID:      data.ID,
			Name:    data.Name,
			Type:    data.Type,
			Content: data.Content,
		}

		w.Header().Set(httpHeaderTenantID, data.TenantID.String())
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Unable to encode template"))
		}
	})
}

func deleteTemplate(service template.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error removing template"
		vars := mux.Vars(r)
		id, err := entity.StringToID(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		err = service.DeleteTemplate(id)
		switch err {
		case nil:
			w.WriteHeader(http.StatusOK)
			return
		case entity.ErrNotFound:
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Template doesn't exist"))
			return
		default:
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
	})
}

func updateTemplate(service template.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error updating template"

		vars := mux.Vars(r)
		id, err := entity.StringToID(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		var input entity.Template
		tenant := r.Header.Get(httpHeaderTenantID)
		tenantID, err := entity.StringToID(tenant)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Missing tenant ID"))
			return
		}

		err = json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Unable to decode the data. " + err.Error()))
			return
		}

		input.ID = id
		input.TenantID = tenantID
		err = service.UpdateTemplate(&input)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage + ":" + err.Error()))
			return
		}

		toJ := &presenter.Template{
			ID:       input.ID,
			TenantID: tenantID,
			Name:     input.Name,
			Type:     input.Type,
			Content:  input.Content,
		}

		w.Header().Set(httpHeaderTenantID, tenant)
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
	})
}

// MakeTemplateHandlers make url handlers
func MakeTemplateHandlers(r *mux.Router, n negroni.Negroni, service template.UseCase) {
	r.Handle("/v1/templates", n.With(
		negroni.Wrap(listTemplates(service)),
	)).Methods("GET", "OPTIONS").Name("listTemplates")

	r.Handle("/v1/templates", n.With(
		negroni.Wrap(createTemplate(service)),
	)).Methods("POST", "OPTIONS").Name("createTemplate")

	r.Handle("/v1/templates/{id}", n.With(
		negroni.Wrap(getTemplate(service)),
	)).Methods("GET", "OPTIONS").Name("getTemplate")

	r.Handle("/v1/templates/{id}", n.With(
		negroni.Wrap(deleteTemplate(service)),
	)).Methods("DELETE", "OPTIONS").Name("deleteTemplate")

	r.Handle("/v1/templates/{id}", n.With(
		negroni.Wrap(updateTemplate(service)),
	)).Methods("PUT", "OPTIONS").Name("updateTemplate")
}
