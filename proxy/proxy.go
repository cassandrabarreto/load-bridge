package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

type Proxy struct {}

func New() *Proxy{
	return &Proxy{}
}
/* */
func (*Proxy) ForwardRequest(writer http.ResponseWriter, request *http.Request, target string){
	url , err := url.Parse(target)

	if err != nil {
		http.Error(writer, "Invalid backend URL", http.StatusInternalServerError)
		return
	}

	reverseProxy := &httputil.ReverseProxy{
		Rewrite: func(pr *httputil.ProxyRequest) {
			pr.SetURL(url)
			pr.SetXForwarded()
		},
	}
	reverseProxy.ServeHTTP(writer, request)
}