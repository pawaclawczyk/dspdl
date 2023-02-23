package openrtb

import (
	"compress/gzip"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/prebid/openrtb/v17/openrtb2"
	"os"
)

// randomBidRequest returns absolute minimal bid requests (only required fields) for a banner ad.
func randomBidRequest() *openrtb2.BidRequest {
	// uuid.New panics if error occurs
	// uuid.New() is equivalent for uuid.Must(uuid.NewRandom())
	imp := openrtb2.Imp{
		ID:     uuid.New().String(),
		Banner: &openrtb2.Banner{},
	}
	return &openrtb2.BidRequest{
		ID:  uuid.New().String(),
		Imp: []openrtb2.Imp{imp},
	}
}

// WriteBidRequests creates a file with n random bid requests.
func WriteBidRequests(n int, path string) {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()
	compressor := gzip.NewWriter(f)
	defer func() {
		if err := compressor.Close(); err != nil {
			panic(err)
		}
	}()
	encoder := json.NewEncoder(compressor)
	for i := 0; i < n; i++ {
		if err = encoder.Encode(randomBidRequest()); err != nil {
			panic(err)
		}
	}
}
