package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"sadam.com/m/account"
	"sadam.com/m/password"
	"sadam.com/m/video"
	"time"
)

func headerHandlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "有人访问了该接口", r.Header)
	fmt.Println(rand.Intn(1000000))
	var testAccount = account.Account{Name: "gitee", Username: "sadam98", Password: "1a2b3c4d5@S", Url: "https://gitee.com", Tel: "18810720138", Email: "1903249375@qq.com"}
	testAccount.Save()
}

func cookieHandlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "有人访问了该接口", r.Header.Get("cookie"))
	account.Load()
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

func main() {
	server := http.Server{
		Addr: "0.0.0.0:7005",
	}

	http.Handle("/header", log(headerHandlerFunc))
	http.Handle("/cookie", log(cookieHandlerFunc))
	http.Handle("/account", log(account.AccountHandler))
	http.Handle("/password", log(password.PasswordHandler))
	http.Handle("/video", log(video.VideoHandler))
	account.Load()
	server.ListenAndServe()

}
