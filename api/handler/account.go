package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"sudhagar/glad/usecase/account"

	"sudhagar/glad/api/presenter"

	"sudhagar/glad/entity"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func listAccounts(service account.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading accounts"
		var data []*entity.Account
		var err error
		tenant := r.Header.Get(httpHeaderTenantID)

		tenantID, err := entity.StringToID(tenant)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Unable to parse tenant id"))
			return
		}

		data, err = service.ListAccounts(tenantID)
		w.Header().Set("Content-Type", "application/json")
		if err != nil && err != entity.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage + ":" + err.Error()))
			return
		}

		total := service.GetCount(tenantID)
		w.Header().Set(httpHeaderTotalCount, strconv.Itoa(total))
		w.Header().Set(httpHeaderTenantID, tenant)

		var toJ []*presenter.Account
		for _, d := range data {
			toJ = append(toJ, &presenter.Account{
				ID:       d.ID,
				TenantID: d.TenantID,
				Username: d.Username,
				Type:     d.Type,
			})
		}
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Unable to encode account"))
		}
	})
}

func getAccountQR(service account.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error adding account"

		tenant := r.Header.Get(httpHeaderTenantID)
		tenantID, err := entity.StringToID(tenant)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Missing tenant ID"))
			return
		}

		// var input struct {
		// 	Type entity.AccountType `json:"type"`
		// }
		// err = json.NewDecoder(r.Body).Decode(&input)
		// if err != nil {
		// 	log.Println(err.Error())
		// 	w.WriteHeader(http.StatusBadRequest)
		// 	w.Write([]byte("Unable to decode the data. " + err.Error()))
		// 	return
		// }

		// account type hardcoded to WhatsApp
		username, data, err := service.GetQR(
			tenantID,
			entity.AccountWhatsApp)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage + ":" + err.Error()))
			return
		}
		toJ := &presenter.Account{
			Username: username,
			Data:     data,
			TenantID: tenantID,
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

func getAccountStatus(service account.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error adding account"
		vars := mux.Vars(r)
		username := vars["username"]

		tenant := r.Header.Get(httpHeaderTenantID)
		tenantID, err := entity.StringToID(tenant)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Missing tenant ID"))
			return
		}

		status, err := service.GetStatus(
			username,
			tenantID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage + ":" + err.Error()))
			return
		}
		toJ := &presenter.Account{
			Username: username,
			Status:   status,
			TenantID: tenantID,
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

// func createAccount(service account.UseCase) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		errorMessage := "Error adding account"
// 		var input struct {
// 			Name    string             `json:"name"`
// 			Type    entity.AccountType `json:"type"`
// 			Content string             `json:"content"`
// 		}

// 		tenant := r.Header.Get(httpHeaderTenantID)
// 		tenantID, err := entity.StringToID(tenant)
// 		if err != nil {
// 			w.WriteHeader(http.StatusBadRequest)
// 			w.Write([]byte("Missing tenant ID"))
// 			return
// 		}

// 		err = json.NewDecoder(r.Body).Decode(&input)
// 		if err != nil {
// 			log.Println(err.Error())
// 			w.WriteHeader(http.StatusBadRequest)
// 			w.Write([]byte("Unable to decode the data. " + err.Error()))
// 			return
// 		}

// 		id, err := service.CreateAccount(
// 			tenantID,
// 			input.Name,
// 			input.Type,
// 			input.Content)
// 		if err != nil {
// 			w.WriteHeader(http.StatusInternalServerError)
// 			w.Write([]byte(errorMessage + ":" + err.Error()))
// 			return
// 		}
// 		toJ := &presenter.Account{
// 			ID:       id,
// 			Name:     input.Name,
// 			Type:     entity.AccountText,
// 			Content:  input.Content,
// 			TenantID: tenantID,
// 		}

// 		w.Header().Set(httpHeaderTenantID, tenant)
// 		w.WriteHeader(http.StatusCreated)
// 		if err := json.NewEncoder(w).Encode(toJ); err != nil {
// 			log.Println(err.Error())
// 			w.WriteHeader(http.StatusInternalServerError)
// 			w.Write([]byte(errorMessage))
// 			return
// 		}
// 	})
// }

func getAccount(service account.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading account"
		vars := mux.Vars(r)
		username := vars["username"]
		data, err := service.GetAccount(username)
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

		toJ := &presenter.Account{
			ID:       data.ID,
			Username: data.Username,
			Type:     data.Type,
		}

		w.Header().Set(httpHeaderTenantID, data.TenantID.String())
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Unable to encode account"))
		}
	})
}

func deleteAccount(service account.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error removing account"
		vars := mux.Vars(r)
		username := vars["username"]
		err := service.DeleteAccount(username)
		switch err {
		case nil:
			w.WriteHeader(http.StatusOK)
			return
		case entity.ErrNotFound:
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Account doesn't exist"))
			return
		default:
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
	})
}

func updateAccount(service account.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error updating account"

		vars := mux.Vars(r)
		username := vars["username"]

		var input entity.Account
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

		input.Username = username
		input.TenantID = tenantID
		err = service.UpdateAccount(&input)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage + ":" + err.Error()))
			return
		}

		toJ := &presenter.Account{
			ID:       input.ID,
			TenantID: tenantID,
			Username: input.Username,
			Type:     input.Type,
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

// MakeAccountHandlers make url handlers
func MakeAccountHandlers(r *mux.Router, n negroni.Negroni, service account.UseCase) {
	r.Handle("/v1/accounts", n.With(
		negroni.Wrap(listAccounts(service)),
	)).Methods("GET", "OPTIONS").Name("listAccounts")

	r.Handle("/v1/accounts/qr", n.With(
		negroni.Wrap(getAccountQR(service)),
	)).Methods("GET", "OPTIONS").Name("getAccountQR")

	r.Handle("/v1/accounts/{username}/status", n.With(
		negroni.Wrap(getAccountStatus(service)),
	)).Methods("GET", "OPTIONS").Name("getAccountStatus")

	// r.Handle("/v1/accounts", n.With(
	// 	negroni.Wrap(createAccount(service)),
	// )).Methods("POST", "OPTIONS").Name("createAccount")

	r.Handle("/v1/accounts/{username}", n.With(
		negroni.Wrap(getAccount(service)),
	)).Methods("GET", "OPTIONS").Name("getAccount")

	r.Handle("/v1/accounts/{username}", n.With(
		negroni.Wrap(deleteAccount(service)),
	)).Methods("DELETE", "OPTIONS").Name("deleteAccount")

	r.Handle("/v1/accounts/{username}", n.With(
		negroni.Wrap(updateAccount(service)),
	)).Methods("PUT", "OPTIONS").Name("updateAccount")
}
