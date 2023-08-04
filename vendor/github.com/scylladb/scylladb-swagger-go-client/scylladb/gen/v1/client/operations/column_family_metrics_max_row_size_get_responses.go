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

// ColumnFamilyMetricsMaxRowSizeGetReader is a Reader for the ColumnFamilyMetricsMaxRowSizeGet structure.
type ColumnFamilyMetricsMaxRowSizeGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ColumnFamilyMetricsMaxRowSizeGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewColumnFamilyMetricsMaxRowSizeGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewColumnFamilyMetricsMaxRowSizeGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewColumnFamilyMetricsMaxRowSizeGetOK creates a ColumnFamilyMetricsMaxRowSizeGetOK with default headers values
func NewColumnFamilyMetricsMaxRowSizeGetOK() *ColumnFamilyMetricsMaxRowSizeGetOK {
	return &ColumnFamilyMetricsMaxRowSizeGetOK{}
}

/*
ColumnFamilyMetricsMaxRowSizeGetOK handles this case with default header values.

ColumnFamilyMetricsMaxRowSizeGetOK column family metrics max row size get o k
*/
type ColumnFamilyMetricsMaxRowSizeGetOK struct {
	Payload interface{}
}

func (o *ColumnFamilyMetricsMaxRowSizeGetOK) GetPayload() interface{} {
	return o.Payload
}

func (o *ColumnFamilyMetricsMaxRowSizeGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewColumnFamilyMetricsMaxRowSizeGetDefault creates a ColumnFamilyMetricsMaxRowSizeGetDefault with default headers values
func NewColumnFamilyMetricsMaxRowSizeGetDefault(code int) *ColumnFamilyMetricsMaxRowSizeGetDefault {
	return &ColumnFamilyMetricsMaxRowSizeGetDefault{
		_statusCode: code,
	}
}

/*
ColumnFamilyMetricsMaxRowSizeGetDefault handles this case with default header values.

internal server error
*/
type ColumnFamilyMetricsMaxRowSizeGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the column family metrics max row size get default response
func (o *ColumnFamilyMetricsMaxRowSizeGetDefault) Code() int {
	return o._statusCode
}

func (o *ColumnFamilyMetricsMaxRowSizeGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *ColumnFamilyMetricsMaxRowSizeGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *ColumnFamilyMetricsMaxRowSizeGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
