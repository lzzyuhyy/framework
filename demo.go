package main

import (
	"encoding/json"
	"fmt"
	"github.com/lzzyuhyy/framework/es/esgo"
	"github.com/lzzyuhyy/framework/es/httpes"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/get/list", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode("404 not found")
			return
		}
		q := r.URL.Query()
		index := q.Get("index")
		id := q.Get("id")

		info, err := httpes.GetInfo(index, id, nil)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode("data can't get")
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"data": info,
		})
	})

	http.HandleFunc("/put/list", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "PUT" {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode("404 not found")
			return
		}

		data := struct {
			Index string
			Id    string
			Data  httpes.R
		}{}

		// 读取请求体
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Unable to read request body", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		// 解析 JSON 数据
		err = json.Unmarshal(body, &data)
		if err != nil {
			http.Error(w, "Invalid JSON"+err.Error(), http.StatusBadRequest)
			return
		}

		info, err := httpes.PutInfo(data.Index, data.Id, data.Data)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode("data can't get")
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"data": info,
		})
	})

	http.HandleFunc("/create/index", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode("404 not found")
			return
		}

		data := struct {
			Index string
		}{}

		// 读取请求体
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Unable to read request body", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		// 解析 JSON 数据
		err = json.Unmarshal(body, &data)
		if err != nil {
			http.Error(w, "Invalid JSON"+err.Error(), http.StatusBadRequest)
			return
		}

		rep, err := esgo.CreateIndex(data.Index)
		if err != nil {
			http.Error(w, "create index err"+err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprintln(w, rep)

	})

	http.HandleFunc("/put/data", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "PUT" {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode("404 not found")
			return
		}

		data := struct {
			Index string
			Info  map[string]interface{}
		}{}

		// 读取请求体
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Unable to read request body", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		// 解析 JSON 数据
		err = json.Unmarshal(body, &data)
		if err != nil {
			http.Error(w, "Invalid JSON"+err.Error(), http.StatusBadRequest)
			return
		}

		info, err := esgo.PutInfo(data.Index, data.Info)
		if err != nil {
			return
		}
		if err != nil {
			http.Error(w, "create index err"+err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprintln(w, info)

	})

	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		panic(err)
	}

}
