package pagination

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPagination(t *testing.T) {
	testCases := map[string]struct {
		page           int32
		size           int32
		expected       Pagination
		expectedLimit  int32
		expectedOffset int32
	}{
		"should create a pagination with the provided page and size": {
			page:           1,
			size:           10,
			expected:       Pagination{Page: 1, Size: 10},
			expectedLimit:  10,
			expectedOffset: 0,
		},
		"should create a pagination with the default page and size if the provided page is less than 1": {
			page:           0,
			size:           10,
			expected:       Pagination{Page: 1, Size: 10},
			expectedLimit:  10,
			expectedOffset: 0,
		},
		"should create a pagination with the default page and size if the provided size is less than 1": {
			page:           1,
			size:           0,
			expected:       Pagination{Page: 1, Size: 10},
			expectedLimit:  10,
			expectedOffset: 0,
		},
		"should create a pagination with the default page and size if the provided size is greater than 100": {
			page:           1,
			size:           101,
			expected:       Pagination{Page: 1, Size: 10},
			expectedLimit:  10,
			expectedOffset: 0,
		},
		"should increase the offset if the provided page is greater than 1": {
			page:           2,
			size:           10,
			expected:       Pagination{Page: 2, Size: 10},
			expectedLimit:  10,
			expectedOffset: 10,
		},
	}

	for testName, testCase := range testCases {
		t.Run(testName, func(t *testing.T) {
			result := NewPagination(testCase.page, testCase.size)

			assert.Equal(t, testCase.expected.Page, result.Page)
			assert.Equal(t, testCase.expected.Size, result.Size)
		})
	}
}
