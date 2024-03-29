package klook

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/asaskevich/govalidator"
)

type Gateway struct {
	Client Client
}

func (gateway *Gateway) SetHeader() []Header {
	var headers []Header

	if len(gateway.Client.APIKey) == 0 {
		token, err := gateway.GetToken()
		if err != nil {
			fmt.Println(err.Error())
			return headers
		}

		headers = []Header{
			Header{
				Key:   "Authorization",
				Value: "Bearer " + token,
			},
			Header{
				Key:   "Content-Type",
				Value: "application/json",
			},
			Header{
				Key:   "Accepts-Language",
				Value: "id_ID",
			},
		}
	} else {
		headers = []Header{
			Header{
				Key:   "X-API-KEY",
				Value: gateway.Client.APIKey,
			},
			Header{
				Key:   "Content-Type",
				Value: "application/json",
			},
			Header{
				Key:   "Accepts-Language",
				Value: "id_ID",
			},
		}
	}

	return headers
}

// Get Token
func (gateway *Gateway) GetToken() (string, error) {
	var resp LoginResponse
	var url = gateway.Client.APILoginEnvType.String() + "/oauth/v3/token"

	data := new(Token)
	data.GrantType = "client_credentials"
	data.ClientId = gateway.Client.ClientID
	data.ClientSecret = gateway.Client.ClientSecret
	data.Scope = "order"

	jsonReq, _ := json.Marshal(data)

	headers := []Header{
		Header{
			Key:   "Content-Type",
			Value: "application/json",
		},
	}

	_ = gateway.Client.Call("POST", url, headers, bytes.NewBuffer(jsonReq), &resp)

	if !govalidator.IsNull(resp.Error) {
		return "", errors.New(resp.ErrorDescription)
	}

	return resp.AccessToken, nil
}

// Get Country
func (gateway *Gateway) GetCountry() Countries {
	var resp Countries
	var url = gateway.Client.APIEnvType.String() + "/v2/countries"

	headers := gateway.SetHeader()

	_ = gateway.Client.Call("GET", url, headers, nil, &resp)

	return resp
}

// Get List Activities
func (gateway *Gateway) GetActivities(req ActivityRequest) Activities {
	var resp Activities
	var url = gateway.Client.APIEnvType.String() + "/v2/activities"

	params := fmt.Sprintf("?page=%d&page_size=%d&country_id=%d&city_id=%d", req.Page, req.PageSize, req.CountryID, req.CityID)

	headers := gateway.SetHeader()

	_ = gateway.Client.Call("GET", url+params, headers, nil, &resp)

	return resp
}

// Get Detail Activity
func (gateway *Gateway) GetDetailActivity(activity_id string) (DetailActivityResponse, error) {
	var resp DetailActivityResponse
	var url = gateway.Client.APIEnvType.String() + "/v2/activities/" + activity_id
	headers := gateway.SetHeader()

	_ = gateway.Client.Call("GET", url, headers, nil, &resp)

	if !resp.Success {
		msg := resp.Error.Message
		if govalidator.IsNull(resp.Error.Message) {
			msg = "Unknown error. Could not get any response"
		}
		return resp, errors.New(msg)
	}

	return resp, nil
}

// Create Order
func (gateway *Gateway) CreateOrder(r OrderRequest) (Order, error) {
	var resp Order
	var url = gateway.Client.APIEnvType.String() + "/v2/orders"

	headers := gateway.SetHeader()

	jsonReq, _ := json.Marshal(r)

	_ = gateway.Client.Call("POST", url, headers, bytes.NewBuffer(jsonReq), &resp)
	if !govalidator.IsNull(resp.Error.Code) {
		return resp, errors.New(resp.Error.Message)
	}

	return resp, nil
}

// Cek Availability
func (gateway *Gateway) CekAvailability(r AvailabilityRequest) (bool, error) {
	var resp Availablilty
	var url = gateway.Client.APIEnvType.String() + "/v2/orders/request"

	headers := gateway.SetHeader()

	jsonReq, _ := json.Marshal(r)

	_ = gateway.Client.Call("POST", url, headers, bytes.NewBuffer(jsonReq), &resp)

	if !resp.Success {
		return false, errors.New(resp.Error.Message)
	}

	return true, nil
}

// Detail Order
func (gateway *Gateway) GetDetailOrder(hash string) (OrderDetail, error) {
	var resp OrderDetail
	var url = gateway.Client.APIEnvType.String() + "/v2/orders/" + hash

	headers := gateway.SetHeader()

	_ = gateway.Client.Call("GET", url, headers, nil, &resp)
	if !resp.Success {
		return resp, errors.New(resp.Error.Message)
	}

	return resp, nil
}

// Resend Voucher
func (gateway *Gateway) ResendVoucher(hash string) (bool, error) {
	var resp Availablilty
	var url = gateway.Client.APIEnvType.String() + "/v2/orders/" + hash + "/resend_voucher"

	headers := gateway.SetHeader()

	_ = gateway.Client.Call("GET", url, headers, nil, &resp)
	if !resp.Success {
		fmt.Println("Error Resend")
		return false, errors.New(resp.Error.Message)
	}

	return true, nil
}

// Cancellation
func (gateway *Gateway) CancelOrder(hash string) (Cancellation, error) {
	var resp Cancellation
	var url = gateway.Client.APIEnvType.String() + "/v2/orders/" + hash + "/cancel"

	headers := gateway.SetHeader()

	_ = gateway.Client.Call("POST", url, headers, nil, &resp)
	if !resp.Success {
		return resp, errors.New(resp.Error.Message)
	}

	return resp, nil
}

// Get Schedule
func (gateway *Gateway) GetSchedule(req ScheduleRequest) (Schedules, error) {
	var resp Schedules
	var url = gateway.Client.APIEnvType.String() + "/v2/products/" + req.ProductID + "/schedules"

	params := "?start_date=" + req.StartDate + "&end_date=" + req.EndDate

	headers := gateway.SetHeader()

	_ = gateway.Client.Call("GET", url+params, headers, nil, &resp)
	if !govalidator.IsNull(resp.Error.Code) {
		return resp, errors.New(resp.Error.Message)
	}

	return resp, nil
}

func (gateway *Gateway) GetBalance() (Balance, error) {
	var resp Balance
	var url = gateway.Client.APIEnvType.String() + "/v2/balance"

	headers := gateway.SetHeader()

	_ = gateway.Client.Call("GET", url, headers, nil, &resp)
	if !govalidator.IsNull(resp.Error.Code) {
		return resp, errors.New(resp.Error.Message)
	}

	return resp, nil
}
