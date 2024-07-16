package model

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"

	"github.com/open-policy-agent/frameworks/constraint/pkg/externaldata"
	"k8s.io/klog/v2"
)

type Logic struct {
	m sync.Mutex
}

type Controller struct {
	L *Logic
}

var (
	L Logic
	C Controller
)

func (c Controller) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprintln(w, "OK")
}

func (c Controller) Answer(w http.ResponseWriter, r *http.Request) {
	j, e := getAnswerJson()
	if e != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = fmt.Fprintln(w, "could not get answers")
	}

	fmt.Println("Answer: ", j)

	w.Header().Set(HeaderContentType, ContentTypeJson)
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprintln(w, j)
}

func (c Controller) Provide(w http.ResponseWriter, r *http.Request) {
	providerResponse := externaldata.ProviderResponse{
		APIVersion: apiVersion,
		Kind:       kind,
		Response: externaldata.Response{
			Idempotent: true,
		},
	}

	// read
	rb, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		providerResponse.Response.SystemError = err.Error()
		_, _ = fmt.Fprintln(w, providerResponse)
	}

	fmt.Println("Request body: ", string(rb))

	// parse
	var providerRequest externaldata.ProviderRequest
	err = json.Unmarshal(rb, &providerRequest)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		providerResponse.Response.SystemError = err.Error()
		_, _ = fmt.Fprintln(w, providerResponse)
	}

	a := AC.getAnswer()
	if a.Answer == "" || a.Code == "" {
		w.WriteHeader(http.StatusInternalServerError)
		providerResponse.Response.SystemError = err.Error()
		_, _ = fmt.Fprintln(w, providerResponse)
	}

	results := make([]externaldata.Item, 0)
	for _, key := range providerRequest.Request.Keys {
		klog.Info("key: ", key)
		// switch {
		// case key == "answer":
		// 	results = append(results, externaldata.Item{
		// 		Key:   key,
		// 		Value: fmt.Sprint(a.Answer),
		// 	})
		// case key == "code":
		// 	results = append(results, externaldata.Item{
		// 		Key:   key,
		// 		Value: fmt.Sprint(a.Code),
		// 	})
		// default:
		// 	results = append(results, externaldata.Item{
		// 		Key:   key,
		// 		Error: "unknown key",
		// 	})

		// 	providerResponse.Response.SystemError = "unknown keys"
		// }

		switch {
		case a.Code == "Y":
			results = append(results, externaldata.Item{
				Key:   key,
				Value: fmt.Sprint(a.Answer, ":", a.Code),
			})
		default:
			results = append(results, externaldata.Item{
				Key:   key,
				Value: fmt.Sprint(a.Answer, ":", a.Code),
				Error: fmt.Sprint(a.Answer, ":", a.Code, ":DENIED"),
			})
		}
	}

	providerResponse.Response.Items = results
	b, e := json.Marshal(providerResponse)
	if e != nil {
		klog.ErrorS(err, "unable to marshall response json")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, JsonProviderErrorMessage)
	}

	klog.Info("Response Body: ", string(b))
	w.Header().Set(HeaderContentType, ContentTypeJson)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, string(b))

}
