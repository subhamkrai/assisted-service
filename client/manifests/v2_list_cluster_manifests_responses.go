// Code generated by go-swagger; DO NOT EDIT.

package manifests

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openshift/assisted-service/models"
)

// V2ListClusterManifestsReader is a Reader for the V2ListClusterManifests structure.
type V2ListClusterManifestsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V2ListClusterManifestsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV2ListClusterManifestsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewV2ListClusterManifestsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewV2ListClusterManifestsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewV2ListClusterManifestsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewV2ListClusterManifestsMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewV2ListClusterManifestsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewV2ListClusterManifestsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV2ListClusterManifestsOK creates a V2ListClusterManifestsOK with default headers values
func NewV2ListClusterManifestsOK() *V2ListClusterManifestsOK {
	return &V2ListClusterManifestsOK{}
}

/*V2ListClusterManifestsOK handles this case with default header values.

Success.
*/
type V2ListClusterManifestsOK struct {
	Payload models.ListManifests
}

func (o *V2ListClusterManifestsOK) Error() string {
	return fmt.Sprintf("[GET /v2/clusters/{cluster_id}/manifests][%d] v2ListClusterManifestsOK  %+v", 200, o.Payload)
}

func (o *V2ListClusterManifestsOK) GetPayload() models.ListManifests {
	return o.Payload
}

func (o *V2ListClusterManifestsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2ListClusterManifestsUnauthorized creates a V2ListClusterManifestsUnauthorized with default headers values
func NewV2ListClusterManifestsUnauthorized() *V2ListClusterManifestsUnauthorized {
	return &V2ListClusterManifestsUnauthorized{}
}

/*V2ListClusterManifestsUnauthorized handles this case with default header values.

Unauthorized.
*/
type V2ListClusterManifestsUnauthorized struct {
	Payload *models.InfraError
}

func (o *V2ListClusterManifestsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v2/clusters/{cluster_id}/manifests][%d] v2ListClusterManifestsUnauthorized  %+v", 401, o.Payload)
}

func (o *V2ListClusterManifestsUnauthorized) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *V2ListClusterManifestsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2ListClusterManifestsForbidden creates a V2ListClusterManifestsForbidden with default headers values
func NewV2ListClusterManifestsForbidden() *V2ListClusterManifestsForbidden {
	return &V2ListClusterManifestsForbidden{}
}

/*V2ListClusterManifestsForbidden handles this case with default header values.

Forbidden.
*/
type V2ListClusterManifestsForbidden struct {
	Payload *models.InfraError
}

func (o *V2ListClusterManifestsForbidden) Error() string {
	return fmt.Sprintf("[GET /v2/clusters/{cluster_id}/manifests][%d] v2ListClusterManifestsForbidden  %+v", 403, o.Payload)
}

func (o *V2ListClusterManifestsForbidden) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *V2ListClusterManifestsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2ListClusterManifestsNotFound creates a V2ListClusterManifestsNotFound with default headers values
func NewV2ListClusterManifestsNotFound() *V2ListClusterManifestsNotFound {
	return &V2ListClusterManifestsNotFound{}
}

/*V2ListClusterManifestsNotFound handles this case with default header values.

Error.
*/
type V2ListClusterManifestsNotFound struct {
	Payload *models.Error
}

func (o *V2ListClusterManifestsNotFound) Error() string {
	return fmt.Sprintf("[GET /v2/clusters/{cluster_id}/manifests][%d] v2ListClusterManifestsNotFound  %+v", 404, o.Payload)
}

func (o *V2ListClusterManifestsNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2ListClusterManifestsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2ListClusterManifestsMethodNotAllowed creates a V2ListClusterManifestsMethodNotAllowed with default headers values
func NewV2ListClusterManifestsMethodNotAllowed() *V2ListClusterManifestsMethodNotAllowed {
	return &V2ListClusterManifestsMethodNotAllowed{}
}

/*V2ListClusterManifestsMethodNotAllowed handles this case with default header values.

Method Not Allowed.
*/
type V2ListClusterManifestsMethodNotAllowed struct {
	Payload *models.Error
}

func (o *V2ListClusterManifestsMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /v2/clusters/{cluster_id}/manifests][%d] v2ListClusterManifestsMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *V2ListClusterManifestsMethodNotAllowed) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2ListClusterManifestsMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2ListClusterManifestsConflict creates a V2ListClusterManifestsConflict with default headers values
func NewV2ListClusterManifestsConflict() *V2ListClusterManifestsConflict {
	return &V2ListClusterManifestsConflict{}
}

/*V2ListClusterManifestsConflict handles this case with default header values.

Error.
*/
type V2ListClusterManifestsConflict struct {
	Payload *models.Error
}

func (o *V2ListClusterManifestsConflict) Error() string {
	return fmt.Sprintf("[GET /v2/clusters/{cluster_id}/manifests][%d] v2ListClusterManifestsConflict  %+v", 409, o.Payload)
}

func (o *V2ListClusterManifestsConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2ListClusterManifestsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2ListClusterManifestsInternalServerError creates a V2ListClusterManifestsInternalServerError with default headers values
func NewV2ListClusterManifestsInternalServerError() *V2ListClusterManifestsInternalServerError {
	return &V2ListClusterManifestsInternalServerError{}
}

/*V2ListClusterManifestsInternalServerError handles this case with default header values.

Error.
*/
type V2ListClusterManifestsInternalServerError struct {
	Payload *models.Error
}

func (o *V2ListClusterManifestsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v2/clusters/{cluster_id}/manifests][%d] v2ListClusterManifestsInternalServerError  %+v", 500, o.Payload)
}

func (o *V2ListClusterManifestsInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2ListClusterManifestsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
