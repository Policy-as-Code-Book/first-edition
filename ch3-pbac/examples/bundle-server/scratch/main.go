package main

import (
	"encoding/json"
	"fmt"
)

type MyJsonName struct {
	Auth struct {
		Audience string `json:"audience"`
		Enable   bool   `json:"enable"`
		Secret   string `json:"secret"`
	} `json:"auth"`
	Bundles struct {
		Bundle_out_dir string `json:"bundle-out-dir"`
		Bundle_uri     string `json:"bundle-uri"`
		Bundles        []struct {
			Build              bool     `json:"build"`
			Bundle_file_name   string   `json:"bundle-file-name"`
			Bundle_in_dir      string   `json:"bundle-in-dir"`
			Bundle_out_dir     string   `json:"bundle-out-dir"`
			Bundle_revision    string   `json:"bundle-revision"`
			Bundle_roots       []string `json:"bundle-roots"`
			Bundle_signing_key string   `json:"bundle-signing-key"`
		} `json:"bundles"`
	} `json:"bundles"`
	Init struct {
		Allow_dir_list         bool   `json:"allow-dir-list"`
		Enable_etag            bool   `json:"enable-etag"`
		Enable_request_logging bool   `json:"enable-request-logging"`
		Log_level              string `json:"log-level"`
	} `json:"init"`
	Metadata struct {
		Name string `json:"name"`
		Tags struct {
			Billing string `json:"billing"`
			Env     string `json:"env"`
			Owner   string `json:"owner"`
		} `json:"tags"`
	} `json:"metadata"`
	Network struct {
		Server_address string `json:"server-address"`
		Server_port    string `json:"server-port"`
		TLS            struct {
			Cert   string `json:"cert"`
			Enable bool   `json:"enable"`
			Key    string `json:"key"`
		} `json:"tls"`
		Uris struct {
			Health string `json:"health"`
			Info   string `json:"info"`
		} `json:"uris"`
	} `json:"network"`
}

var config MyJsonName

func main() {
	/* First: declared map of string with empty interface
	which will hold the value of the parsed json. */
	//var result map[string]interface{}
	/* Second: Unmarshal the json string string by converting
	   it to byte into map */
	//json.Unmarshal([]byte(configJson), &result)
	json.Unmarshal([]byte(configJson), &config)
	/* Third: Read the Value by its key */
	//fmt.Printf("%+v", result)
	fmt.Printf("%+v", config)
}

const configJson string = `
{
  "metadata": {
    "name": "opa-bundle-server",
    "tags": {
      "owner": "jimmy",
      "env": "dev",
      "billing": "lob-cc"
    }
  },
  "network": {
    "server-address": "10.0.2.2",
    "server-port": "8443",
    "tls": {
      "enable": true,
      "cert": "tls/localhost.crt",
      "key": "tls/localhost.key"
    },
    "uris": {
      "health": "/healthz",
      "info": "/info"
    }
  },
  "init": {
    "allow-dir-list": false,
    "enable-request-logging": true,
    "enable-etag": true,
    "log-level": "debug"
  },
  "auth": {
    "enable": true,
    "secret": "a3c1ea92-6474-11ed-81ce-0242ac120002",
    "audience": "opa-bundle-api"
  },
  "bundles": {
    "bundle-uri": "/v1/bundles",
    "bundle-out-dir": "bundles",
    "bundles": [
      {
        "build": true,
        "bundle-file-name": "bundle.tar.gz",
        "bundle-in-dir": "bundle-material",
        "bundle-out-dir": "bundles",
        "bundle-revision": "v0.1.0",
        "bundle-roots": [
          "pacbook"
        ],
        "bundle-signing-key": "keys/key.pem"
      }
    ]
  }
}
`
