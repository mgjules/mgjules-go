package static

type static struct{}

func New() *static {
	return &static{}
}

func (*static) Cleanup() error {
	return nil
}
