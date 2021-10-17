package guard

import "testing"

func TestCheck(t *testing.T) {
	type testcase struct {
		name     string
		allow    roleName
		roles    []roleName
		expected bool
	}

	testcases := [...]testcase{
		{
			name:     "invalid",
			allow:    HeadAdmin,
			roles:    []roleName{Market},
			expected: false,
		},
		{
			name:  "valid",
			allow: HeadAdmin,
			roles: []roleName{
				HeadSupportManager,
			},
			expected: false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			res := check(tc.allow, tc.roles...)
			if res != tc.expected {
				t.Errorf("expected: %v, got: %v", tc.expected, res)
			}
		})
	}
}

func BenchmarkCheck(b *testing.B) {
	b.ResetTimer()
	for i := b.N; i > 0; i-- {
		check(HeadAdmin, []roleName{
			HeadSupportManager,
		}...)
	}
}
