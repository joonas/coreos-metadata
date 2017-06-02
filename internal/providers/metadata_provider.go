package providers

type MetadataProvider interface {
	FetchMetadata() (Metadata, error)
}
