// Code generated by go-swagger; DO NOT EDIT.

package health_check

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ridhompra/pkg/models"
)

// GetHealthCheckReader is a Reader for the GetHealthCheck structure.
type GetHealthCheckReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetHealthCheckReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetHealthCheckOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetHealthCheckDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetHealthCheckOK creates a GetHealthCheckOK with default headers values
func NewGetHealthCheckOK() *GetHealthCheckOK {
	return &GetHealthCheckOK{}
}

/*
GetHealthCheckOK describes a response with status code 200, with default header values.

success
*/
type GetHealthCheckOK struct {
	Payload *models.BaseResponse
}

// IsSuccess returns true when this get health check o k response has a 2xx status code
func (o *GetHealthCheckOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get health check o k response has a 3xx status code
func (o *GetHealthCheckOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get health check o k response has a 4xx status code
func (o *GetHealthCheckOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get health check o k response has a 5xx status code
func (o *GetHealthCheckOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get health check o k response a status code equal to that given
func (o *GetHealthCheckOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetHealthCheckOK) Error() string {
	return fmt.Sprintf("[GET /health][%d] getHealthCheckOK  %+v", 200, o.Payload)
}

func (o *GetHealthCheckOK) String() string {
	return fmt.Sprintf("[GET /health][%d] getHealthCheckOK  %+v", 200, o.Payload)
}

func (o *GetHealthCheckOK) GetPayload() *models.BaseResponse {
	return o.Payload
}

func (o *GetHealthCheckOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BaseResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHealthCheckDefault creates a GetHealthCheckDefault with default headers values
func NewGetHealthCheckDefault(code int) *GetHealthCheckDefault {
	return &GetHealthCheckDefault{
		_statusCode: code,
	}
}

/*
GetHealthCheckDefault describes a response with status code -1, with default header values.

error
*/
type GetHealthCheckDefault struct {
	_statusCode int

	Payload *models.BaseResponse
}

// Code gets the status code for the get health check default response
func (o *GetHealthCheckDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get health check default response has a 2xx status code
func (o *GetHealthCheckDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get health check default response has a 3xx status code
func (o *GetHealthCheckDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get health check default response has a 4xx status code
func (o *GetHealthCheckDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get health check default response has a 5xx status code
func (o *GetHealthCheckDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get health check default response a status code equal to that given
func (o *GetHealthCheckDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetHealthCheckDefault) Error() string {
	return fmt.Sprintf("[GET /health][%d] getHealthCheck default  %+v", o._statusCode, o.Payload)
}

func (o *GetHealthCheckDefault) String() string {
	return fmt.Sprintf("[GET /health][%d] getHealthCheck default  %+v", o._statusCode, o.Payload)
}

func (o *GetHealthCheckDefault) GetPayload() *models.BaseResponse {
	return o.Payload
}

func (o *GetHealthCheckDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BaseResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
