package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact Page</h1><p>To get in touch, email me at <a href=\"mailto:neilbriancohen@gmail.com\">neilbriancohen@gmail.com</a>")
}

//func pathHandler(w http.ResponseWriter, r *http.Request) {
//	switch r.URL.Path {
//	case "/":
//		homeHandler(w, r)
//	case "/contact":
//		contactHandler(w, r)
//	default:
//		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
//	}
//}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	//	fmt.Fprint(w, "<h2><b>Q: How do I contact support?</b></h2><p>A: Email me at neilbriancohen@gmail.com</p>")
	fmt.Fprint(w, `<h1>FAQ Page</h1>
	<ul>
		<li><b> Is there a free version?</b> Yes! we offer a 30 day free trial</li>
		<li><b> How do I contact support?</b> Email me at <a href="mailto:neilbriancohen@gmail.com">neilbriancohen@gmail.com</a>"</li>
	</ul>`)
}

type Router struct{}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	case "/faq":
		faqHandler(w, r)
	default:
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}
}

func main() {
	var router Router
	//	http.HandleFunc("/", homeHandler)
	//	http.HandleFunc("/contact", contactHandler)
	//	http.HandleFunc("/", pathHandler)
	fmt.Println("Starting on port 4000")
	http.ListenAndServe(":4000", router)
}
