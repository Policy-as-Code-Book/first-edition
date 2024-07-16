package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	Log "github.com/sirupsen/logrus"
	"io"
	"jimmyray.io/opa-bundle-api/pkg/model"
	"jimmyray.io/opa-bundle-api/pkg/utils"
	"net/http"
	"net/http/httptest"
	"os"
	"runtime"
	"strings"
	"testing"
)

const (
	employeePayload       string = `{"id":"218000","fname":"Indrajit","lname":"Raney","sex":"M","dob":"1964-08-04T00:00:00Z","hireDate":"1989-08-31T00:00:00Z","position":"Senior Engineer","salary":64633,"dept":{"id":"d006","name":"Quality Management","mgrId":"110854"},"address":{"street":"4200 Old Us Highway395n N","city":"Washoe Valley","county":"Washoe","state":"NV","zipcode":"89704"}}`
	employeePayloadCreate string = `{"id":"500000","fname":"Jimmy","lname":"Ray","sex":"M","dob":"1964-08-04T00:00:00Z","hireDate":"1989-08-31T00:00:00Z","position":"Developer Advocate","salary":1000000,"dept":{"id":"d006","name":"Quality Management","mgrId":"110854"},"address":{"street":"4200 Old Us Highway395n N","city":"Washoe Valley","county":"Washoe","state":"NV","zipcode":"89704"}}`
	employeePayloadPatch  string = `{"id":"218000","fname":"Indrajit","lname":"Raney","sex":"M","dob":"1964-08-04T00:00:00Z","hireDate":"1989-08-31T00:00:00Z","position":"Engineer","salary":64633,"dept":{"id":"d006","name":"Quality Management","mgrId":"110854"},"address":{"street":"4200 Old Us Highway395n N","city":"Washoe Valley","county":"Washoe","state":"NV","zipcode":"89704"}}`
	employeeId            string = "218000"
	expectedCount         int    = 161
)

var employeeCount int

func getCurrentFuncName() string {
	pc, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(pc).Name()
}

func cleanData() {
	//Clean-up, reset service data
	model.L.ServiceData = make(map[string]model.Employee)
	err := model.LoadMockData()
	if err != nil {
		fmt.Println("Fatal: Could not load mock data for testing.")
	}

	fmt.Println("Finished running ", getCurrentFuncName())

	//os.Exit(1)
}

func TestMain(m *testing.M) {
	fmt.Println("Test Setup...")
	model.InitValidator()
	utils.InitLogs(nil, Log.InfoLevel)

	model.C = model.Controller{
		L:        &model.L,
		Validate: model.Validate,
	}

	model.IC = model.InfoController{
		ServiceInfo: model.ServiceInfo,
	}

	cleanData()

	employeeCount = len(model.L.ServiceData)

	fmt.Println("Finished running ", getCurrentFuncName())

	exitVal := m.Run()
	os.Exit(exitVal)
}

func TestCount(t *testing.T) {
	//Test employee count
	if employeeCount != expectedCount {
		t.Errorf("Expected employee count %d, received %d", expectedCount, employeeCount)
	}

	fmt.Println("Finished running ", getCurrentFuncName())
}

func TestGetAllData(t *testing.T) {
	u := "/data"
	r := httptest.NewRequest(http.MethodGet, u, nil)
	w := httptest.NewRecorder()

	model.C.GetAllData(w, r)
	res := w.Result()
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)

	if err != nil {
		t.Errorf("expected error to be nil, received %v", err)
	}

	e := model.Employees{}
	err = json.Unmarshal([]byte(data), &e)

	if err != nil {
		t.Error("Error parsing returned data byte array into Employees struct")
	}

	if len(e) != expectedCount {
		t.Errorf("expected count was %d, received %d", expectedCount, len(e))
	}

	fmt.Println("Finished running ", getCurrentFuncName())
}

func TestGetData(t *testing.T) {
	u := fmt.Sprintf("/data/%s", employeeId)
	r := httptest.NewRequest(http.MethodPatch, u, nil)
	w := httptest.NewRecorder()

	vars := map[string]string{
		"id": employeeId,
	}
	r = mux.SetURLVars(r, vars)

	model.C.GetData(w, r)
	res := w.Result()
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)

	if err != nil {
		t.Errorf("expected error to be nil, received %v", err)
	}

	if strings.Trim(string(data), "\n") != strings.Trim(employeePayload, "\n") {
		t.Errorf("Expected %v, received %v", employeePayload, string(data))
	}

	fmt.Println("Finished running ", getCurrentFuncName())
}

func TestPatchData(t *testing.T) {
	// Test data conflict
	u := "/data"
	r := httptest.NewRequest(http.MethodPatch, u, bytes.NewReader([]byte(employeePayload)))
	w := httptest.NewRecorder()

	model.C.PatchData(w, r)
	res := w.Result()
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	//t.Logf("Return = %s", string(data))

	if err != nil {
		t.Errorf("expected error to be nil, received %v", err)
	}

	if string(data) != model.DataConflictErr {
		t.Errorf("expected %v, received %v", model.DataConflictErr, string(data))
	}

	// Test patch by comparing strings
	u = "/data"
	r = httptest.NewRequest(http.MethodPatch, u, bytes.NewReader([]byte(employeePayloadPatch)))
	w = httptest.NewRecorder()

	model.C.PatchData(w, r)
	res = w.Result()
	defer res.Body.Close()

	data, err = io.ReadAll(res.Body)
	//t.Logf("Return = %s", string(data))

	if err != nil {
		fmt.Println(err.Error())
		t.Errorf("Expected error to be nil, received %v", err)
	}

	if strings.Trim(string(data), "\n") != strings.Trim(employeePayloadPatch, "\n") {
		t.Errorf("Expected %v, received %v", employeePayloadPatch, string(data))
	}

	//Compare structs
	var e1 model.Employee
	err = json.Unmarshal(data, &e1)
	if err != nil {
		t.Error("Error parsing returned data byte array into Employee struct")
	}

	var e2 model.Employee
	err = json.Unmarshal([]byte(employeePayloadPatch), &e2)
	if err != nil {
		t.Error("Error parsing employeePayload byte array into Employee struct")
	}

	if e1 != e2 {
		t.Errorf("Employee structs did not match. Expected %v, received %v", e2, e1)
	}

	cleanData()

	fmt.Println("Finished running ", getCurrentFuncName())
}

func TestDeleteData(t *testing.T) {
	u := fmt.Sprintf("/data/%s", employeeId)
	r := httptest.NewRequest(http.MethodDelete, u, nil)
	w := httptest.NewRecorder()

	vars := map[string]string{
		"id": employeeId,
	}
	r = mux.SetURLVars(r, vars)

	model.C.DeleteData(w, r)
	res := w.Result()
	defer res.Body.Close()

	_, err := io.ReadAll(res.Body)

	if err != nil {
		t.Errorf("expected error to be nil, received %v", err)
	}

	if len(model.L.ServiceData) != expectedCount-1 {
		t.Errorf("Expected count: %v, received count: %v", expectedCount-1, len(model.L.ServiceData))
	}

	cleanData()

	fmt.Println("Finished running ", getCurrentFuncName())
}

func TestCreateData(t *testing.T) {
	u := "/data"
	r := httptest.NewRequest(http.MethodPut, u, bytes.NewReader([]byte(employeePayloadCreate)))
	w := httptest.NewRecorder()

	model.C.CreateData(w, r)
	res := w.Result()
	defer res.Body.Close()

	_, err := io.ReadAll(res.Body)
	//t.Logf("Return = %s", string(data))

	if err != nil {
		t.Errorf("expected error to be nil, received %v", err)
	}

	if w.Code != http.StatusCreated {
		t.Errorf("Expected code: %v, received code: %v", http.StatusCreated, w.Code)
	}

	if len(model.L.ServiceData) != expectedCount+1 {
		t.Errorf("Expected count: %v, received count: %v", expectedCount+1, len(model.L.ServiceData))
	}

	cleanData()

	fmt.Println("Finished running ", getCurrentFuncName())
}
