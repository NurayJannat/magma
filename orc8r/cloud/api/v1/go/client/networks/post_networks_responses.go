// Code generated by go-swagger; DO NOT EDIT.

package networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "magma/orc8r/cloud/api/v1/go/models"
)

// PostNetworksReader is a Reader for the PostNetworks structure.
type PostNetworksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostNetworksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostNetworksCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostNetworksDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostNetworksCreated creates a PostNetworksCreated with default headers values
func NewPostNetworksCreated() *PostNetworksCreated {
	return &PostNetworksCreated{}
}

/*PostNetworksCreated handles this case with default header values.

Success
*/
type PostNetworksCreated struct {
}

func (o *PostNetworksCreated) Error() string {
	return fmt.Sprintf("[POST /networks][%d] postNetworksCreated ", 201)
}

func (o *PostNetworksCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostNetworksDefault creates a PostNetworksDefault with default headers values
func NewPostNetworksDefault(code int) *PostNetworksDefault {
	return &PostNetworksDefault{
		_statusCode: code,
	}
}

/*PostNetworksDefault handles this case with default header values.

Unexpected Error
*/
type PostNetworksDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post networks default response
func (o *PostNetworksDefault) Code() int {
	return o._statusCode
}

func (o *PostNetworksDefault) Error() string {
	return fmt.Sprintf("[POST /networks][%d] PostNetworks default  %+v", o._statusCode, o.Payload)
}

func (o *PostNetworksDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostNetworksDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}