package main

import "generators/internal/openrtb"

func main() {
	openrtb.WriteBidRequests(1_000_000, "bid_requests.json.gz")
}
