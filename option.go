package navio

import "github.com/opentracing/opentracing-go"

type options struct {
	name string
	isSecure bool
	certFile string
	keyFile string
	isTracing bool
	tracer opentracing.Tracer
	block bool
}
type ClientOption func(*options)
// WithInsecure returns a ClientOption that configure the client connection to be insecure
func WithInsecure() ClientOption {
	return func(o *options) {
		o.isSecure = false
	}
}

// WithCredentials returns a ClientOption that configure the client connection to be TLS secured with cert and key files locations
func WithCredentials (certFile, keyFile string) ClientOption{
	return func(o *options) {
		o.isSecure =true
		o.certFile=certFile
		o.keyFile=keyFile
	}
}

// WithTracer returns a ClientOption that configure the client connection to be traced with OpenTracing global tracer
func WithTracer (tracer  opentracing.Tracer) ClientOption{
	return func(o *options) {
		o.isTracing = true
		o.tracer = tracer
	}
}

// WithBlock returns a ClientOption that configure the client connection to be blocked until the underline connection is established
func WithBlock () ClientOption{
	return func(o *options) {
		o.block = true
	}
}
// WithName returns a ClientOption that set a name for the client connection
// setting name is important for identification connection to navio server
// if name will not be set , a random name will be created
func WithName (name string) ClientOption{
	return func(o *options) {
		o.name = name
	}
}

//