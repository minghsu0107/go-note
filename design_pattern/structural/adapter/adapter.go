package adapter

// Target is the adapter interface
type Target interface {
	Request() string
}

// Adaptee is the adapted interface
type Adaptee interface {
	SpecificRequest() string
}

// NewCustomAdaptee is the Adaptee factory
func NewCustomAdaptee() Adaptee {
	return &customAdapteeImpl{}
}

// adapteeImpl is the adapted concrete implementation
type customAdapteeImpl struct{}

// SpecificRequest is the method of adapteeImpl
func (*customAdapteeImpl) SpecificRequest() string {
	return "adaptee method"
}

// NewAdapter is the adapter factory
func NewAdapter(adaptee Adaptee) Target {
	return &adapter{
		Adaptee: adaptee,
	}
}

// Adapter convert Adaptee to Target interface
// use annoymous structure field
type adapter struct {
	Adaptee
}

// implement Target interface
func (a *adapter) Request() string {
	return a.SpecificRequest()
}
