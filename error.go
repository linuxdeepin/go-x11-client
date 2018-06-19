package x

import (
	"fmt"
)

// Error is an interface that can contain any of the errors returned by
// the server. Use a type assertion switch to extract the Error structs.
//type Error interface {
//SequenceId() uint16
//BadId() uint32
//Error() string
//}

//// NewErrorFun is the type of function use to construct errors from raw bytes.
//// It should not be used. It is exported for use in the extension sub-packages.
//type NewErrorFun func(buf []byte) Error

//// NewErrorFuncs is a map from error numbers to functions that create
//// the corresponding error. It should not be used. It is exported for use in
//// the extension sub-packages.
//var NewErrorFuncs = make(map[int]NewErrorFun)

//// NewExtErrorFuncs is a temporary map that stores error constructor functions
//// for each extension. When an extension is initialized, each error for that
//// extension is added to the 'NewErrorFuncs' map. It should not be used. It is
//// exported for use in the extension sub-packages.
//var NewExtErrorFuncs = make(map[string]map[int]NewErrorFun)

// eventOrError corresponds to values that can be either an event or an
// error.

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
