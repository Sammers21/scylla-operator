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

// StorageServiceRPCServerDeleteReader is a Reader for the StorageServiceRPCServerDelete structure.
type StorageServiceRPCServerDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageServiceRPCServerDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageServiceRPCServerDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStorageServiceRPCServerDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStorageServiceRPCServerDeleteOK creates a StorageServiceRPCServerDeleteOK with default headers values
func NewStorageServiceRPCServerDeleteOK() *StorageServiceRPCServerDeleteOK {
	return &StorageServiceRPCServerDeleteOK{}
}

/*
StorageServiceRPCServerDeleteOK handles this case with default header values.

StorageServiceRPCServerDeleteOK storage service Rpc server delete o k
*/
type StorageServiceRPCServerDeleteOK struct {
}

func (o *StorageServiceRPCServerDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStorageServiceRPCServerDeleteDefault creates a StorageServiceRPCServerDeleteDefault with default headers values
func NewStorageServiceRPCServerDeleteDefault(code int) *StorageServiceRPCServerDeleteDefault {
	return &StorageServiceRPCServerDeleteDefault{
		_statusCode: code,
	}
}

/*
StorageServiceRPCServerDeleteDefault handles this case with default header values.

internal server error
*/
type StorageServiceRPCServerDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the storage service Rpc server delete default response
func (o *StorageServiceRPCServerDeleteDefault) Code() int {
	return o._statusCode
}

func (o *StorageServiceRPCServerDeleteDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *StorageServiceRPCServerDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *StorageServiceRPCServerDeleteDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
