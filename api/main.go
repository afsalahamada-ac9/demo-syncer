/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"sudhagar/glad/repository"

	"sudhagar/glad/usecase/account"
	"sudhagar/glad/usecase/contact"
	"sudhagar/glad/usecase/label"
	"sudhagar/glad/usecase/labeler"
	"sudhagar/glad/usecase/template"
	"sudhagar/glad/usecase/tenant"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"sudhagar/glad/api/handler"
	"sudhagar/glad/api/middleware"
	"sudhagar/glad/config"
	"sudhagar/glad/pkg/metric"

	"github.com/codegangsta/negroni"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

func main() {

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true",
		config.DB_USER,
		config.DB_PASSWORD,
		config.DB_HOST,
		config.DB_DATABASE)
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	templateRepo := repository.NewTemplateMySQL(db)
	templateService := template.NewService(templateRepo)

	tenantRepo := repository.NewTenantMySQL(db)
	tenantService := tenant.NewService(tenantRepo)

	accountRepo := repository.NewAccountMySQL(db)
	accountService := account.NewService(accountRepo, nil)

	contactRepo := repository.NewContactMySQL(db)
	contactService := contact.NewService(contactRepo)

	labelRepo := repository.NewLabelMySQL(db)
	labelService := label.NewService(labelRepo)

	labelerUseCase := labeler.NewService(contactService, labelService)

	metricService, err := metric.NewPrometheusService()
	if err != nil {
		log.Fatal(err.Error())
	}
	r := mux.NewRouter()
	// handlers
	n := negroni.New(
		negroni.HandlerFunc(middleware.Cors),
		negroni.HandlerFunc(middleware.Metrics(metricService)),
		negroni.NewLogger(),
	)
	// template
	handler.MakeTemplateHandlers(r, *n, templateService)

	// tenant
	handler.MakeTenantHandlers(r, *n, tenantService)

	// account
	handler.MakeAccountHandlers(r, *n, accountService)

	// contact
	handler.MakeContactHandlers(r, *n, contactService)

	// label
	handler.MakeLabelHandlers(r, *n, labelService)

	// labeler
	handler.MakeLabelerHandlers(r, *n, labelerUseCase)

	http.Handle("/", r)
	http.Handle("/metrics", promhttp.Handler())
	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	logger := log.New(os.Stderr, "logger: ", log.Lshortfile)
	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         ":" + strconv.Itoa(config.API_PORT),
		Handler:      context.ClearHandler(http.DefaultServeMux),
		ErrorLog:     logger,
	}
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}
}
