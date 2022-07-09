package erring

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_DomainErr_Err(t *testing.T) {
	t.Parallel()

	type test struct {
		name     string
		err      DomainErr
		expected string
	}
	tests := []test{
		{
			name: "when with value return with full error message",
			err: DomainErr{
				name:        "error 1",
				description: "a fake error for tests",
				err:         errors.New("external error"),
				fields:      map[string]interface{}{"value": 123},
				calle:       "func1",
			},
			expected: "{\n name: error 1\n description: a fake error for tests\n  calle: func1\n err: external error\n}",
		},
		{
			name:     "when with no value return with nils",
			err:      DomainErr{},
			expected: "{\n}",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			actual := tt.err.Error()

			assert.Equal(t, tt.expected, actual)
		})
	}
}

func Test_DomainErr_Is(t *testing.T) {
	t.Parallel()

	extErr := errors.New("123")

	type test struct {
		name     string
		err      DomainErr
		target   error
		expected bool
	}
	tests := []test{
		{
			name:     "when errors are domain errors and has diffrent name return false",
			err:      DomainErr{name: "123"},
			target:   DomainErr{name: "456"},
			expected: false,
		},
		{
			name:     "when errors are domain errors and has same name return true",
			err:      DomainErr{name: "123"},
			target:   DomainErr{name: "123"},
			expected: true,
		},
		{
			name:     "when external error don't match target return true",
			err:      DomainErr{err: extErr},
			target:   errors.New("456"),
			expected: false,
		},
		{
			name:     "when external error match target return true",
			err:      DomainErr{err: extErr},
			target:   extErr,
			expected: true,
		},
		{
			name:     "when wrapped with match return true",
			err:      DomainErr{err: DomainErr{err: extErr}},
			target:   extErr,
			expected: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			actual := errors.Is(tt.err, tt.target)

			assert.Equal(t, tt.expected, actual)
		})
	}
}
