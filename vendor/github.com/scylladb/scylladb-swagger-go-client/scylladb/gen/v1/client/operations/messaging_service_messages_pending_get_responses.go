// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strings"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/scylladb/scylladb-swagger-go-client/scylladb/gen/v1/models"
)

// MessagingServiceMessagesPendingGetReader is a Reader for the MessagingServiceMessagesPendingGet structure.
type MessagingServiceMessagesPendingGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MessagingServiceMessagesPendingGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMessagingServiceMessagesPendingGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewMessagingServiceMessagesPendingGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewMessagingServiceMessagesPendingGetOK creates a MessagingServiceMessagesPendingGetOK with default headers values
func NewMessagingServiceMessagesPendingGetOK() *MessagingServiceMessagesPendingGetOK {
	return &MessagingServiceMessagesPendingGetOK{}
}

/*
MessagingServiceMessagesPendingGetOK handles this case with default header values.

MessagingServiceMessagesPendingGetOK messaging service messages pending get o k
*/
type MessagingServiceMessagesPendingGetOK struct {
	Payload []*models.MessageCounter
}

func (o *MessagingServiceMessagesPendingGetOK) GetPayload() []*models.MessageCounter {
	return o.Payload
}

func (o *MessagingServiceMessagesPendingGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMessagingServiceMessagesPendingGetDefault creates a MessagingServiceMessagesPendingGetDefault with default headers values
func NewMessagingServiceMessagesPendingGetDefault(code int) *MessagingServiceMessagesPendingGetDefault {
	return &MessagingServiceMessagesPendingGetDefault{
		_statusCode: code,
	}
}

/*
MessagingServiceMessagesPendingGetDefault handles this case with default header values.

internal server error
*/
type MessagingServiceMessagesPendingGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the messaging service messages pending get default response
func (o *MessagingServiceMessagesPendingGetDefault) Code() int {
	return o._statusCode
}

func (o *MessagingServiceMessagesPendingGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *MessagingServiceMessagesPendingGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *MessagingServiceMessagesPendingGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
