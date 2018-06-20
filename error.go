package x

import (
	"fmt"
)

// generic error
type GenericError struct {
	// ResponseType uint8  // type of the response === 0
	ErrorCode  uint8  // ErrorCode
	Sequence   uint16 // Sequence number
	ResourceID uint32 // Resource ID for requests with side effects only
	MinorCode  uint16 // Minor opcode of the failed request
	MajorCode  uint8  // Major opcode of the faield request
	Data       []byte
}

func (err *GenericError) GetCode() uint8 {
	return err.ErrorCode
}

func (err *GenericError) GetSequenceNumber() uint16 {
	return err.Sequence
}

func (err *GenericError) GetMajorOpcode() uint8 {
	return err.MajorCode
}

func (err *GenericError) GetMinorOpcode() uint16 {
	return err.MinorCode
}

func (err *GenericError) Error() string {
	return fmt.Sprintf("GenericError{ ErrorCode: %d, Sequence: %d, ResourceID: %d, MajorCode: %d, MinorCode: %d }",
		err.ErrorCode, err.Sequence, err.ResourceID, err.MajorCode, err.MinorCode)
}

func NewGenericError(data []byte) *GenericError {
	var v GenericError

	responseType := data[0]
	if responseType != ResponseTypeError {
		panic("not error")
	}
	b := 1

	v.ErrorCode = data[b]
	b += 1

	v.Sequence = Get16(data[b:])
	b += 2

	v.ResourceID = Get32(data[b:])
	b += 4

	v.MinorCode = Get16(data[b:])
	b += 2

	v.MajorCode = data[b]
	//b += 1

	v.Data = data

	return &v
}

func NewError(data []byte) Error {
	genericErr := NewGenericError(data)
	// TODO:
	readErrFunc, ok := readErrorFuncMap[genericErr.ErrorCode]
	if ok {
		r := NewReaderFromData(data)
		return readErrFunc(r)
	} else {
		return genericErr
	}
}

func (c *Conn) NewError(data []byte) Error {
	genericErr := NewGenericError(data)
	errCode := genericErr.ErrorCode

	if 1 <= errCode && errCode <= 127 {
		// core error code in range [1,127]
		fn, ok := readErrorFuncMap[errCode]
		if ok {
			r := NewReaderFromData(data)
			return fn(r)
		}
	} else {
		// ext error code in range [128,255]
		err := c.ext.newError(errCode, data)
		if err != nil {
			return err
		}
	}
	return genericErr
}
