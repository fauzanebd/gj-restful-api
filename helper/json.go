package helper

import (
	"encoding/json"
	"net/http"

	"github.com/fauzanebd/gj-restful-api/model/web"
)

func ReadFromRequestBodyJSON(request *http.Request, v interface{}) error {
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(v)
	if err != nil {
		return err
	}
	return nil
}

func WriteToResponseBodyJSON(writer http.ResponseWriter, data interface{}) error {
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(data.(web.WebResponse).Code)

	encoder := json.NewEncoder(writer)
	err := encoder.Encode(data)
	if err != nil {
		return err
	}
	return nil
}
