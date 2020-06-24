package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func homepageHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello you"))
}

func fileGetHandler(w http.ResponseWriter, r *http.Request) {
	var reqFile requestedFile
	if err := json.NewDecoder(r.Body).Decode(&reqFile); err != nil {
		msg := httpResponse{
			Status:  "error",
			Code:	101,
			Message: "Received invalid json",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&msg)
		return
	}

	if !fileExists(reqFile.FileName) {
		msg := httpResponse{
			Status:  "error",
			Code: 	 102,
			Message: "File not found",
		}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(&msg)
		return
	}

	cont, err := ioutil.ReadFile(reqFile.FileName)
	if err != nil {
		msg := httpResponse{
			Status:  "error",
			Code: 	 103,
			Message: "File could not be read",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&msg)
		return
	}

	returnData := httpResponse{
		Status:  "success",
		Code:    100,
		Message: "",
		Data:    requestedFile{
			FileName:    reqFile.FileName,
			FileContent: string(cont),
		},
	}

	//w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&returnData)
}

func fileSaveHandler(w http.ResponseWriter, r *http.Request) {
	var reqFile requestedFile
	if err := json.NewDecoder(r.Body).Decode(&reqFile); err != nil {
		msg := httpResponse{
			Status:  "error",
			Code:	101,
			Message: "Received invalid json",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&msg)
		return
	}

	if !fileExists(reqFile.FileName) {
		msg := httpResponse{
			Status:  "error",
			Code: 	 102,
			Message: "File not found",
		}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(&msg)
		return
	}

	err := ioutil.WriteFile(reqFile.FileName, []byte(reqFile.FileContent), 0664)
	if err != nil {
		msg := httpResponse{
			Status:  "error",
			Code:	 104,
			Message: "Could not write file",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&msg)
		return
	}

	returnData := httpResponse{
		Status:  "success",
		Code:    100,
	}

	json.NewEncoder(w).Encode(&returnData)
}