package erring

import (
	"errors"
	"fmt"
	"strings"
)

type DomainErr struct {
	calle       string
	fields      map[string]interface{}
	name        string
	description string
	err         error
}

func (e DomainErr) Error() string {
	acc := "{"
	if e.name != "" {
		acc += fmt.Sprintf("\n name: %s", e.name)
	}

	if e.description != "" {
		acc += fmt.Sprintf("\n description: %s", e.description)
	}

	if e.calle != "" {
		acc += fmt.Sprintf("\n  calle: %s", e.calle)
	}

	if e.err != nil {
		acc += fmt.Sprintf("\n err: %s", e.err.Error())
	}

	return acc + "\n}"
}

func (e DomainErr) Unwrap() error {
	return e.err
}

func (e DomainErr) Is(target error) bool {
	if targetAsDomain, ok := target.(DomainErr); ok && targetAsDomain.name == e.name {
		return true
	}

	if e.err != nil {
		return errors.Is(target, e.err)
	}

	return false
}

func NewWrapper(calle string) *DomainErr {
	return &DomainErr{calle: calle}
}

func NewErr(name, description string) DomainErr {
	return DomainErr{
		name:        strings.ReplaceAll(strings.ToLower(name), " ", "_"),
		description: description,
	}
}

func (e DomainErr) Field(key string, value interface{}) DomainErr {
	if e.fields == nil {
		e.fields = map[string]interface{}{}
	}

	e.fields[key] = value

	return e
}

func (e DomainErr) Fields(fields map[string]interface{}) DomainErr {
	if e.fields == nil {
		e.fields = map[string]interface{}{}
	}

	for k, v := range fields {
		e.fields[k] = v
	}

	return e
}

func (e DomainErr) Wrap(err error) DomainErr {
	e.err = err

	return e
}

func (e DomainErr) NameAndDescribe() (name, description string) {
	if e.name != "" && e.description != "" {
		return e.name, e.description
	}

	var dErr DomainErr
	if ok := errors.As(e.err, &dErr); ok {
		return dErr.NameAndDescribe()
	}

	return "internal server error", "the server has encountered a situation it does not know how to handle."
}

func (e DomainErr) Calls() string {
	acc := e.calle

	var dErr DomainErr
	if ok := errors.As(e.err, &dErr); ok && dErr.calle != "" {
		acc = fmt.Sprintf("%s > %s", acc, dErr.calle)
	}

	return acc
}

func (e DomainErr) GetFields() map[string]interface{} {
	var dErr DomainErr
	ok := errors.As(e.err, &dErr)
	if !ok {
		return e.fields
	}

	acc := dErr.GetFields()
	for k, v := range e.fields {
		acc[k] = v
	}

	return acc
}

func (e DomainErr) InternalErr() error {
	err := e.err
	var dErr DomainErr
	for err != nil && errors.As(err, &dErr) {
		err = dErr.err
	}

	return err
}

func (e DomainErr) Err() error {
	return e
}
