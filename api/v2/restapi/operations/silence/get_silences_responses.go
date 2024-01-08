// Code generated by go-swagger; DO NOT EDIT.

// Copyright Prometheus Team
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package silence

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/prometheus/alertmanager/api/v2/models"
)

// GetSilencesOKCode is the HTTP code returned for type GetSilencesOK
const GetSilencesOKCode int = 200

/*
GetSilencesOK Get silences response

swagger:response getSilencesOK
*/
type GetSilencesOK struct {

	/*
	  In: Body
	*/
	Payload models.GettableSilences `json:"body,omitempty"`
}

// NewGetSilencesOK creates GetSilencesOK with default headers values
func NewGetSilencesOK() *GetSilencesOK {

	return &GetSilencesOK{}
}

// WithPayload adds the payload to the get silences o k response
func (o *GetSilencesOK) WithPayload(payload models.GettableSilences) *GetSilencesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get silences o k response
func (o *GetSilencesOK) SetPayload(payload models.GettableSilences) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSilencesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = models.GettableSilences{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetSilencesBadRequestCode is the HTTP code returned for type GetSilencesBadRequest
const GetSilencesBadRequestCode int = 400

/*
GetSilencesBadRequest Bad request

swagger:response getSilencesBadRequest
*/
type GetSilencesBadRequest struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewGetSilencesBadRequest creates GetSilencesBadRequest with default headers values
func NewGetSilencesBadRequest() *GetSilencesBadRequest {

	return &GetSilencesBadRequest{}
}

// WithPayload adds the payload to the get silences bad request response
func (o *GetSilencesBadRequest) WithPayload(payload string) *GetSilencesBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get silences bad request response
func (o *GetSilencesBadRequest) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSilencesBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetSilencesInternalServerErrorCode is the HTTP code returned for type GetSilencesInternalServerError
const GetSilencesInternalServerErrorCode int = 500

/*
GetSilencesInternalServerError Internal server error

swagger:response getSilencesInternalServerError
*/
type GetSilencesInternalServerError struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewGetSilencesInternalServerError creates GetSilencesInternalServerError with default headers values
func NewGetSilencesInternalServerError() *GetSilencesInternalServerError {

	return &GetSilencesInternalServerError{}
}

// WithPayload adds the payload to the get silences internal server error response
func (o *GetSilencesInternalServerError) WithPayload(payload string) *GetSilencesInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get silences internal server error response
func (o *GetSilencesInternalServerError) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSilencesInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}