package erring

import (
	"errors"
	"fmt"
	"strings"
)

type DomainErr struct {
	calle  string
	fields map[string]interface{}
	msg    string
	err    error
}

func (e DomainErr) Error() string {
	acc := "{"
	if e.msg != "" {
		acc += fmt.Sprintf("\n msg: %s", e.msg)
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
	if targetAsDomain, ok := target.(DomainErr); ok {
		return targetAsDomain.msg == e.msg
	}

	if e.err != nil {
		return errors.Is(target, e.err)
	}

	return false
}

func New(calle string) *DomainErr {
	return &DomainErr{calle: calle}
}

func (e *DomainErr) Field(key string, value interface{}) *DomainErr {
	e.fields[key] = value

	return e
}

func (e *DomainErr) Fields(fields map[string]interface{}) *DomainErr {
	for k, v := range fields {
		e.fields[k] = v
	}

	return e
}

func (e *DomainErr) Wrap(err error) *DomainErr {
	e.err = err

	return e
}

func (e *DomainErr) Err(msg ...string) error {
	e.msg = strings.Join(msg, ", ")

	return e
}
