package errors

import (
	"errors"
	"testing"
)

func TestErrors(t *testing.T) {
	tests := []struct {
		name        string
		err         error
		expected    string
		expectedErr bool
		kind        Kind
	}{
		{
			name:        "basic error",
			err:         New(ResourceNotFound, "key not found"),
			expected:    "key not found",
			expectedErr: true,
			kind:        ResourceNotFound,
		},
		{
			name:        "error concatenation",
			err:         Wrap(ResourceNotFound, "key not found", errors.New("database error")),
			expected:    "key not found: database error",
			expectedErr: true,
			kind:        ResourceNotFound,
		},
		{
			name:        "error concatenation with nil previous error",
			err:         Wrap(ResourceNotFound, "key not found", nil),
			expected:    "key not found",
			expectedErr: true,
			kind:        ResourceNotFound,
		},
		{
			name:        "IsKind with custom error",
			err:         New(ResourceNotFound, "key not found"),
			expected:    "",
			expectedErr: false,
			kind:        BadRequest,
		},
		{
			name:        "IsKind with standard error",
			err:         errors.New("some error"),
			expected:    "",
			expectedErr: false,
			kind:        BadRequest,
		},
		{
			name:        "IsKind with nil error",
			err:         nil,
			expected:    "",
			expectedErr: false,
			kind:        ResourceNotFound,
		},
		{
			name:        "Wrap custom error",
			err:         Wrap(BadRequest, "bad request", errors.New("invalid input")),
			expected:    "bad request: invalid input",
			expectedErr: true,
			kind:        BadRequest,
		},
		{
			name:        "Wrap standard error",
			err:         Wrap(ResourceNotFound, "key not found", errors.New("not found")),
			expected:    "key not found: not found",
			expectedErr: true,
			kind:        ResourceNotFound,
		},
		{
			name:        "Wrap nil error",
			err:         Wrap(Internal, "internal error", nil),
			expected:    "internal error",
			expectedErr: true,
			kind:        Internal,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.expectedErr && tt.err.Error() != tt.expected {
				t.Errorf("Error message mismatch. Expected %s but got %s", tt.expected, tt.err.Error())
			}

			if KindIs(tt.err, tt.kind) != tt.expectedErr {
				t.Errorf("KindIs result mismatch. Expected %v but got %v", tt.expectedErr, !tt.expectedErr)
			}
		})
	}
}
