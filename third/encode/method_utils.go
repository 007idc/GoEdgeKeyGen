package encode

import (
	"errors"
	"reflect"
)

type MethodInterface interface {
	Init(key []byte, iv []byte) error
	Encrypt(src []byte) (dst []byte, err error)
	Decrypt(dst []byte) (src []byte, err error)
}

var methods = map[string]reflect.Type{
	"aes-256-cfb": reflect.TypeOf(new(AES256CFBMethod)).Elem(),
}

func NewMethodInstance(method string, key string, iv string) (MethodInterface, error) {
	valueType, ok := methods[method]
	if !ok {
		return nil, errors.New("method '" + method + "' not found")
	}
	instance, ok := reflect.New(valueType).Interface().(MethodInterface)
	if !ok {
		return nil, errors.New("method '" + method + "' must implement MethodInterface")
	}
	err := instance.Init([]byte(key), []byte(iv))
	return instance, err
}

func RecoverMethodPanic(err interface{}) error {
	if err != nil {
		s, ok := err.(string)
		if ok {
			return errors.New(s)
		}

		e, ok := err.(error)
		if ok {
			return e
		}

		return errors.New("unknown error")
	}
	return nil
}
