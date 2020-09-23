// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strings"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/scylladb/scylla-operator/pkg/scyllaclient/internal/scylla/models"
)

// SnitchDatacenterGetReader is a Reader for the SnitchDatacenterGet structure.
type SnitchDatacenterGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SnitchDatacenterGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSnitchDatacenterGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSnitchDatacenterGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSnitchDatacenterGetOK creates a SnitchDatacenterGetOK with default headers values
func NewSnitchDatacenterGetOK() *SnitchDatacenterGetOK {
	return &SnitchDatacenterGetOK{}
}

/*SnitchDatacenterGetOK handles this case with default header values.

SnitchDatacenterGetOK snitch datacenter get o k
*/
type SnitchDatacenterGetOK struct {
	Payload string
}

func (o *SnitchDatacenterGetOK) GetPayload() string {
	return o.Payload
}

func (o *SnitchDatacenterGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSnitchDatacenterGetDefault creates a SnitchDatacenterGetDefault with default headers values
func NewSnitchDatacenterGetDefault(code int) *SnitchDatacenterGetDefault {
	return &SnitchDatacenterGetDefault{
		_statusCode: code,
	}
}

/*SnitchDatacenterGetDefault handles this case with default header values.

internal server error
*/
type SnitchDatacenterGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the snitch datacenter get default response
func (o *SnitchDatacenterGetDefault) Code() int {
	return o._statusCode
}

func (o *SnitchDatacenterGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *SnitchDatacenterGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *SnitchDatacenterGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}