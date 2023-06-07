package app

// EmptyBaseAppOptions is a stub implementing AppOptions
type EmptyBaseAppOptions struct{}

// Get implements AppOptions
func (ao EmptyBaseAppOptions) Get(o string) interface{} {
	return nil
}
