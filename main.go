package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

func headerHandlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "有人访问了该接口", r.Header)
	fmt.Println(rand.Intn(1000000))
	var testAccount = Account{Name: "gitee", Username: "sadam98", Password: "1a2b3c4d5@S", Url: "https://gitee.com", Tel: "18810720138", Email: "1903249375@qq.com"}
	testAccount.save()
}

func cookieHandlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "有人访问了该接口", r.Header.Get("cookie"))
	load()
}

func log(h http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		err := request.ParseForm()
		if err == nil {

		} else {
			//fmt.Printf("IP:%s Path:%s,Method:%s,Header:%s\n",request.RemoteAddr,request.URL,request.Method,request.Header)
		}
		fmt.Printf("Time:[%s], IP:[%s] Path:[%s],Method:[%s],Form:[%s],Header:[%s]\n", time.Now(), request.RemoteAddr, request.URL, request.Method, request.Form, request.Header)
		//fmt.Println("This is Form",request.Form)
		//name:=runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
		//fmt.Println("This is Host",request.Host)
		//fmt.Println("This is Response",request.Response)  <nil>
		//fmt.Println("This is Body",request.Body)  {}
		//fmt.Println("This is Cancel",request.Cancel)  <nil>
		//fmt.Println("This is ContentLength",request.ContentLength)  8
		//fmt.Println("This is Close",request.Close)  false
		//fmt.Println("This is GetBody",request.GetBody)  <nil>
		//fmt.Println("This is Proto:",request.Proto)  HTTP /1.1
		//fmt.Println("This is PostForm:",request.PostForm) map[]
		//fmt.Println("This is MultipartForm",request.MultipartForm)  <nil>
		//fmt.Println("This is ProtoMajor:",request.ProtoMajor)  1
		//fmt.Println("This is ProtoMinor:",request.ProtoMinor)  1
		//fmt.Println("This is RequestURI:",request.RequestURI)
		h(writer, request)
	}
}
func MyPrint(request *http.Request) {

}

//func account(w http.ResponseWriter,r *http.Request)  {
//}

type Account struct {
	Id       uint64
	Name     string
	Username string
	Password string
	Url      string
	Tel      string
	Email    string
}
type Password struct {
	Content   string `json:"content"`
	AccountId uint64 `json:"account_id"`
}

var CsvDataFileName string = "accounts.csv"
var IdMax uint64 = 0
var AccountById map[uint64]*Account = make(map[uint64]*Account)
var AccountByName map[string]*Account = make(map[string]*Account)

func (account *Account) store() {
	AccountById[account.Id] = account
	AccountByName[account.Name] = account
}
func (account *Account) save() {
	IdMax++
	account.Id = IdMax
	account.store()
	csvDataFile, err := os.Create(CsvDataFileName)
	if err != nil {
		panic(err)
	}
	defer csvDataFile.Close()
	writer := csv.NewWriter(csvDataFile)
	for _, value := range AccountById {
		line := []string{strconv.Itoa(int(IdMax)), strconv.Itoa(int(value.Id)), value.Name, value.Username, value.Password, value.Tel, value.Email, value.Url}
		err := writer.Write(line)
		if err != nil {
			panic(err)
		}
	}
	writer.Flush()
	csvDataFile.Close()
}

func load() {
	file, err := os.Open(CsvDataFileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	for true {
		row, err := reader.Read()
		if err == nil {
			IdMax, err = strconv.ParseUint(row[0], 0, 64)
			if err != nil {
				panic(err)
			} else {
				id, err := strconv.ParseUint(row[1], 0, 64)
				if err != nil {
					panic(err)
				} else {
					account := Account{Id: id, Name: row[2], Username: row[3], Password: row[4], Tel: row[5], Email: row[6], Url: row[7]}
					account.store()
				}
				fmt.Println(row)

			}
		} else if err == io.EOF {
			//	说明已经读到尾了，读完了
			break
		} else {
			panic(err)
		}
	}
}

func main() {
	server := http.Server{
		Addr: "0.0.0.0:7005",
	}

	http.Handle("/header", log(headerHandlerFunc))
	http.Handle("/cookie", log(cookieHandlerFunc))
	http.Handle("/account", log(AccountHandler))
	http.Handle("/password", log(PasswordHandler))
	load()
	server.ListenAndServe()

}
func AccountHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		err = AccountGET(w, r)
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
	if IdMax < id {
		w.WriteHeader(400)
		return err
	} else {
		account := AccountById[id]
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
func AccountGET(w http.ResponseWriter, r *http.Request) (err error) {
	var accounts []*Account
	for _, account := range AccountById {
		accounts = append(accounts, account)
	}
	t, _ := template.ParseFiles("tmp.html")
	t.Execute(w, accounts)
	return nil
}
