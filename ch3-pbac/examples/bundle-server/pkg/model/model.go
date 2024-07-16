package model

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"sync"
)

const (
	ContentTypeGzip   string = "application/gzip"
	ContentTypeJson   string = "application/json"
	HeaderContentType string = "Content-Type"
	PageNotFound      string = "404 page not found"
	Forbidden         string = "403 forbidden"
	Unauthorized      string = "401 unauthorized"
)

type Logic struct {
	m sync.Mutex
}

type Controller struct {
	L *Logic
}

var serviceId = uuid.New()

func GetServiceId() string {
	return serviceId.String()
}

type info struct {
	NAME string `json:"service-name"`
	ID   string `json:"service-id"`
}

func (i info) String() string {
	out, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}

	return string(out)
}

var (
	L  Logic
	IC InfoController
	C  Controller
)

var ServiceInfo = info{}

type InfoController struct {
	ServiceInfo info
}

func (ic InfoController) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprintln(w, "OK")
}

func (ic InfoController) GetServiceInfo(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprintln(w, ServiceInfo.String())
}

func (c Controller) GetBundles(w http.ResponseWriter, r *http.Request) {
	j, e := RegBundles.Json()
	if e != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprintln(w, "")
	} else {
		w.Header().Set(HeaderContentType, ContentTypeJson)
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprintln(w, string(j))
	}
}

func (c Controller) GetServerConfig(w http.ResponseWriter, r *http.Request) {
	j, e := SC.Json()
	if e != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprintln(w, "")
	} else {
		w.Header().Set(HeaderContentType, ContentTypeJson)
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprintln(w, string(j))
	}
}
