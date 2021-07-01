package proxy

type Subject interface {
	Do() string
}

type RealSubject struct {
	a, b string
}

func (r RealSubject) Do() string {
	return r.a + r.b
}

type Proxy struct {
	real RealSubject
}

func NewProxy(a, b string) *Proxy {
	p := &Proxy{}
	real := RealSubject{
		a: a,
		b: b,
	}
	p.real = real
	return p
}

func (p Proxy) Do() string {
	var res string

	// preprocessing, such as cache checking, instance initialization, authentication...
	res += "pre:"

	res += p.real.Do()

	// process the result, such as writing to cache...
	res += ":after"

	return res
}
