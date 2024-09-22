package resource

import (
	"errors"

	"github.com/meidomx/mxartifact/resource/httplistener"
)

func (r *ResourceManager) AddHttpResourceConsumer(resourceName, hostname, baseUri string, handler httplistener.HttpHandler) {
	res, ok := r.httpResources[resourceName]
	if !ok {
		panic(errors.New("resource " + resourceName + " not exist"))
	}
	res.addHttpHandler(hostname, baseUri, handler)
}
