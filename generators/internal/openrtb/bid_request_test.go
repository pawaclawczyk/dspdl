package openrtb

import (
	"os"
	"testing"
)

func BenchmarkRandomBidRequest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		randomBidRequest()
	}
}

func tmpPath() string {
	f, err := os.CreateTemp("", "")
	if err != nil {
		panic(err)
	}
	err = f.Close()
	if err != nil {
		panic(err)
	}
	return f.Name()
}

func BenchmarkWriteBidRequests(b *testing.B) {
	tests := []struct {
		name string
		n    int
		path string
	}{
		{"1", 1, tmpPath()},
		{"1,000", 1_000, tmpPath()},
		{"1,000,000", 1_000_000, tmpPath()},
	}

	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				WriteBidRequests(tt.n, tt.path)
			}
		})
	}
}
