package offchain

import (
	pb "relay/scripts"
	"sync"
)

type IOPerver struct {
	pb.UnimplementedIOPServer
	IOPInfo []*pb.IOPInfo // read-only after initialized

	mu sync.Mutex // protects routeNotes
}
