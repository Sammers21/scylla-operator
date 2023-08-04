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

// CacheServiceMetricsRowRequestsMovingAvrageGetReader is a Reader for the CacheServiceMetricsRowRequestsMovingAvrageGet structure.
type CacheServiceMetricsRowRequestsMovingAvrageGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CacheServiceMetricsRowRequestsMovingAvrageGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCacheServiceMetricsRowRequestsMovingAvrageGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCacheServiceMetricsRowRequestsMovingAvrageGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCacheServiceMetricsRowRequestsMovingAvrageGetOK creates a CacheServiceMetricsRowRequestsMovingAvrageGetOK with default headers values
func NewCacheServiceMetricsRowRequestsMovingAvrageGetOK() *CacheServiceMetricsRowRequestsMovingAvrageGetOK {
	return &CacheServiceMetricsRowRequestsMovingAvrageGetOK{}
}

/*
CacheServiceMetricsRowRequestsMovingAvrageGetOK handles this case with default header values.

CacheServiceMetricsRowRequestsMovingAvrageGetOK cache service metrics row requests moving avrage get o k
*/
type CacheServiceMetricsRowRequestsMovingAvrageGetOK struct {
	Payload *models.RateMovingAverage
}

func (o *CacheServiceMetricsRowRequestsMovingAvrageGetOK) GetPayload() *models.RateMovingAverage {
	return o.Payload
}

func (o *CacheServiceMetricsRowRequestsMovingAvrageGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RateMovingAverage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCacheServiceMetricsRowRequestsMovingAvrageGetDefault creates a CacheServiceMetricsRowRequestsMovingAvrageGetDefault with default headers values
func NewCacheServiceMetricsRowRequestsMovingAvrageGetDefault(code int) *CacheServiceMetricsRowRequestsMovingAvrageGetDefault {
	return &CacheServiceMetricsRowRequestsMovingAvrageGetDefault{
		_statusCode: code,
	}
}

/*
CacheServiceMetricsRowRequestsMovingAvrageGetDefault handles this case with default header values.

internal server error
*/
type CacheServiceMetricsRowRequestsMovingAvrageGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the cache service metrics row requests moving avrage get default response
func (o *CacheServiceMetricsRowRequestsMovingAvrageGetDefault) Code() int {
	return o._statusCode
}

func (o *CacheServiceMetricsRowRequestsMovingAvrageGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CacheServiceMetricsRowRequestsMovingAvrageGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *CacheServiceMetricsRowRequestsMovingAvrageGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
