package ports

// TemplatesPort is the port interface for template interactions.
type TemplatesPort interface {
	Get(name string, data any) (out []byte, err error)
}
