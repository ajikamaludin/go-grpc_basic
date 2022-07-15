package jsonplaceholder

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/ajikamaludin/go-grpc_basic/pkg/v1/requestapi"
	"github.com/ajikamaludin/go-grpc_basic/pkg/v1/utils/constants"
)

type Response []struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Address  struct {
		Street  string `json:"street"`
		Suite   string `json:"suite"`
		City    string `json:"city"`
		Zipcode string `json:"zipcode"`
		Geo     struct {
			Lat string `json:"lat"`
			Lng string `json:"lng"`
		} `json:"geo"`
	} `json:"address"`
	Phone   string `json:"phone"`
	Website string `json:"website"`
	Company struct {
		Name        string `json:"name"`
		CatchPhrase string `json:"catchPhrase"`
		Bs          string `json:"bs"`
	} `json:"company"`
}

func GetListUser() (*Response, error) {
	res, err := requestapi.Invoke(&requestapi.ReqInfo{
		URL:    fmt.Sprintf("%v/users", constants.Host_Reqres),
		Method: constants.MethodGET,
	}, 60*time.Second, nil)
	if err != nil {
		return nil, err
	}

	var resp Response

	err = json.Unmarshal(res.Body, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
