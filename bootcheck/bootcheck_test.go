package bootcheck

import "testing"

func TestEnvironmentOK(t *testing.T) {
	tests := []struct {
		name    string
		version string
		want    bool
	}{
		{name: "valid", version: "go1.22.0", want: true},
		{name: "empty", version: "", want: false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := EnvironmentOK(tc.version)

			if got != tc.want {
				t.Fatalf("EnvironmentOK(%q) = %v, want %v", tc.version, got, tc.want)
			}
		})
	}
}
