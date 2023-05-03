package service

import (
	"errors"
	"github.com/go-resty/resty/v2"
	"net/http"
)

const sdnListURL = "https://www.treasury.gov/ofac/downloads/sdn.xml"

func (s *service) GetSdnList() (resp []byte, err error) {
	var response *resty.Response
	if response, err = s.httpLib.R().SetError(&err).Post(sdnListURL); err != nil {
		return
	}

	if response.StatusCode() != http.StatusOK {
		err = errors.New(string(response.Body()))
		return
	}
	return response.Body(), nil
}
