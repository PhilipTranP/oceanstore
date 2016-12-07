package tapestry

/*
	This is a utility class tacked on to the tapestry DOLR.
*/
type BlobStore struct {
	blobs map[string]Blob
}

type Blob struct {
	bytes []byte
	done  chan bool
}

type BlobStoreRPC struct {
	store *BlobStore
}

/*
	Create a new blobstore
*/
func NewBlobStore() *BlobStore {
	bs := new(BlobStore)
	bs.blobs = make(map[string]Blob)
	return bs
}

/*
	For RPC server registration
*/
func NewBlobStoreRPC(store *BlobStore) *BlobStoreRPC {
	rpc := new(BlobStoreRPC)
	rpc.store = store
	return rpc
}