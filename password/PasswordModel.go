package password

import (
	"encoding/json"
	"izbasar.link/web/account"
	"net/http"
	"strconv"
)

type Password struct {
	Content   string `json:"content"`
	AccountId uint64 `json:"account_id"`
}

func PasswordHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		err = PasswordGet(w, r)
		break
	case "POST":
		break
	case "PUT":
		break
	case "DELETE":
		break
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
func PasswordGet(w http.ResponseWriter, r *http.Request) (err error) {
	r.ParseForm()
	value := r.Form["id"][0]
	id, _ := strconv.ParseUint(value, 0, strconv.IntSize)
	if account.IdMax < id {
		w.WriteHeader(400)
		return err
	} else {
		account := account.AccountById[id]
		if account == nil {
			w.WriteHeader(400)
			return err
		} else {
			password := account.Password
			passwordFinal := Password{Content: password, AccountId: id}
			output, err := json.MarshalIndent(&passwordFinal, "", "\t\t")
			if err != nil {
				return err
			} else {
				w.Header().Set("Content-Type", "application/json")
				w.Write(output)
				w.WriteHeader(200)
			}
		}
	}
	return err
}
