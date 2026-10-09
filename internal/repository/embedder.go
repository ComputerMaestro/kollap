package repository

type Embedder interface {
	Embed(string) ([]float32, error)
}
