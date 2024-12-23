/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"sudhagar/glad/pkg/common"
	"sudhagar/glad/usecase/course"

	"sudhagar/glad/api/presenter"

	"sudhagar/glad/entity"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func listCourses(service course.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading courses"
		var data []*entity.Course
		var err error
		tenant := r.Header.Get(common.HttpHeaderTenantID)
		search := r.URL.Query().Get(httpParamQuery)
		page, _ := strconv.Atoi(r.URL.Query().Get(httpParamPage))
		limit, _ := strconv.Atoi(r.URL.Query().Get(httpParamLimit))
		tenantID, err := entity.StringToID(tenant)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte("Unable to parse tenant id"))
			return
		}

		switch {
		case search == "":
			data, err = service.ListCourses(tenantID, page, limit)
		default:
			// TODO: search need to be reworked; need to add a count
			// for search; also need to see how the caller generates
			// the search query request
			data, err = service.SearchCourses(tenantID, search, page, limit)
		}
		w.Header().Set("Content-Type", "application/json")
		if err != nil && err != entity.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage + ":" + err.Error()))
			return
		}

		total := service.GetCount(tenantID)
		w.Header().Set(httpHeaderTotalCount, strconv.Itoa(total))

		if data == nil {
			w.WriteHeader(http.StatusNotFound)
			_, _ = w.Write([]byte(errorMessage))
			return
		}
		var toJ []*presenter.Course
		for _, d := range data {
			pc := &presenter.Course{
				ID:           d.ID,
				Name:         &d.Name,
				Mode:         &d.Mode,
				CenterID:     &d.CenterID,
				Notes:        &d.Notes,
				Timezone:     &d.Timezone,
				Status:       &d.Status,
				MaxAttendees: &d.MaxAttendees,
				NumAttendees: &d.NumAttendees,
			}
			pc.Address = &presenter.Address{}
			pc.Address.CopyFrom(d.Address)

			toJ = append(toJ, pc)
		}
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.Header().Set(common.HttpHeaderTenantID, tenant)
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte("Unable to encode course"))
		}
	})
}

func createCourse(service course.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error adding course"
		var input struct {
			ExtID        *string
			Name         string                  `json:"name"`
			CenterID     entity.ID               `json:"centerId"`
			ProductID    entity.ID               `json:"productId"`
			Organizer    []entity.ID             `json:"organizer"`
			Contact      []entity.ID             `json:"contact"`
			Teacher      []entity.ID             `json:"teacher"`
			Notes        string                  `json:"notes"`
			Status       entity.CourseStatus     `json:"status"`
			MaxAttendees int32                   `json:"maxAttendees"`
			Dates        []entity.CourseDateTime `json:"dates"`
			Timezone     string                  `json:"timezone"`
			Address      entity.CourseAddress    `json:"address"`
			Mode         entity.CourseMode       `json:"mode"`
			Notify       []entity.ID             `json:"notify"`
		}

		tenant := r.Header.Get(common.HttpHeaderTenantID)
		tenantID, err := entity.StringToID(tenant)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte("Missing tenant ID"))
			return
		}

		err = json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte("Unable to decode the data. " + err.Error()))
			return
		}

		id, err := service.CreateCourse(
			tenantID,
			input.ExtID,
			input.CenterID,
			input.ProductID,
			input.Name,
			input.Notes,
			input.Timezone,
			input.Address,
			input.Status,
			input.Mode,
			input.MaxAttendees,
			// numAttendees
			0,
		)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage + ":" + err.Error()))
			return
		}
		toJ := &presenter.Course{
			ID: id,
		}

		w.Header().Set(common.HttpHeaderTenantID, tenant)
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage))
			return
		}
	})
}

func getCourse(service course.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading course"
		vars := mux.Vars(r)
		id, err := entity.StringToID(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(err.Error()))
			return
		}
		data, err := service.GetCourse(id)
		if err != nil && err != entity.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage + ":" + err.Error()))
			return
		}

		if data == nil {
			w.WriteHeader(http.StatusNotFound)
			_, _ = w.Write([]byte("Empty data returned"))
			return
		}

		toJ := &presenter.Course{
			ID:           data.ID,
			Name:         &data.Name,
			Mode:         &data.Mode,
			CenterID:     &data.CenterID,
			Notes:        &data.Notes,
			Timezone:     &data.Timezone,
			Status:       &data.Status,
			MaxAttendees: &data.MaxAttendees,
			NumAttendees: &data.NumAttendees,
		}

		toJ.Address = &presenter.Address{}
		toJ.Address.CopyFrom(data.Address)

		w.Header().Set(common.HttpHeaderTenantID, data.TenantID.String())
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte("Unable to encode course"))
		}
	})
}

func deleteCourse(service course.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error removing course"
		vars := mux.Vars(r)
		id, err := entity.StringToID(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage))
			return
		}
		err = service.DeleteCourse(id)
		switch err {
		case nil:
			w.WriteHeader(http.StatusOK)
			return
		case entity.ErrNotFound:
			w.WriteHeader(http.StatusNotFound)
			_, _ = w.Write([]byte("Course doesn't exist"))
			return
		default:
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage))
			return
		}
	})
}

func updateCourse(service course.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error updating course"

		vars := mux.Vars(r)
		id, err := entity.StringToID(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(err.Error()))
			return
		}

		var input entity.Course
		tenant := r.Header.Get(common.HttpHeaderTenantID)
		tenantID, err := entity.StringToID(tenant)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte("Missing tenant ID"))
			return
		}

		err = json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte("Unable to decode the data. " + err.Error()))
			return
		}

		input.ID = id
		input.TenantID = tenantID
		err = service.UpdateCourse(&input)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage + ":" + err.Error()))
			return
		}

		toJ := &presenter.Course{
			ID: input.ID,
		}

		w.Header().Set(common.HttpHeaderTenantID, tenant)
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage))
			return
		}
	})
}

// MakeCourseHandlers make url handlers
func MakeCourseHandlers(r *mux.Router, n negroni.Negroni, service course.UseCase) {
	r.Handle("/v1/courses", n.With(
		negroni.Wrap(listCourses(service)),
	)).Methods("GET", "OPTIONS").Name("listCourses")

	r.Handle("/v1/courses", n.With(
		negroni.Wrap(createCourse(service)),
	)).Methods("POST", "OPTIONS").Name("createCourse")

	r.Handle("/v1/courses/{id}", n.With(
		negroni.Wrap(getCourse(service)),
	)).Methods("GET", "OPTIONS").Name("getCourse")

	r.Handle("/v1/courses/{id}", n.With(
		negroni.Wrap(deleteCourse(service)),
	)).Methods("DELETE", "OPTIONS").Name("deleteCourse")

	r.Handle("/v1/courses/{id}", n.With(
		negroni.Wrap(updateCourse(service)),
	)).Methods("PUT", "OPTIONS").Name("updateCourse")
}
