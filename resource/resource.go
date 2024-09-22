package resource

import (
	"context"
	"errors"
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/meidomx/mxartifact/config"
	"github.com/meidomx/mxartifact/resource/httplistener"
)

type ResourceManager struct {
	httpResources map[string]*HttpResource
}

func NewResourceManager(cfg *config.Config) *ResourceManager {
	r := &ResourceManager{
		httpResources: make(map[string]*HttpResource),
	}
	for _, elem := range cfg.Shared.Listeners {
		var hr HttpResource
		{
			hr.r = chi.NewRouter()
			hr.r.Use(middleware.Recoverer)
			hr.r.Use(middleware.Logger)
		}
		hr.pathMapping = map[string]*struct {
			hostnameMapping map[string]*struct {
				handler httplistener.HttpHandler
			}
		}{}
		hr.addresses = elem.Addresses
		r.addHttpResource(elem.Name, &hr)
	}

	return r
}

func (r *ResourceManager) addHttpResource(name string, hr *HttpResource) {
	if _, ok := r.httpResources[name]; ok {
		panic("duplicated http resource: " + name)
	}
	r.httpResources[name] = hr
}

func (r *ResourceManager) Startup() (err error) {
	defer func() {
		if recovered := recover(); recovered != nil || err != nil {
			for _, v := range r.httpResources {
				if err := v.Shutdown(); err != nil {
					//FIXME log shutdown error
				}
			}
		}
	}()
	for _, v := range r.httpResources {
		if err = v.Startup(); err != nil {
			return
		} else {
			log.Println("starting http services on:", v.addresses)
		}
	}
	return
}

func (r *ResourceManager) Shutdown() error {
	for _, v := range r.httpResources {
		if err := v.Shutdown(); err != nil {
			//FIXME log shutdown error
		}
	}
	return nil
}

type HttpResource struct {
	r         chi.Router
	addresses []string

	pathMapping map[string]*struct {
		hostnameMapping map[string]*struct {
			handler httplistener.HttpHandler
		}
	}

	listeners []struct {
		net.Listener
		*http.Server
	}
}

func (r *HttpResource) Startup() error {
	for _, addr := range r.addresses {
		l, err := net.Listen("tcp", addr)
		if err != nil {
			return err
		}
		srv := &http.Server{Handler: r.r}
		r.listeners = append(r.listeners, struct {
			net.Listener
			*http.Server
		}{l, srv})
		go func() {
			if err := srv.Serve(l); err != nil { // expected not to have listener error since the listener is already created
				log.Println("http resource serve error:", err)
			}
		}()
	}
	return nil
}

func (r *HttpResource) Shutdown() error {
	var errs []error
	for _, l := range r.listeners {
		if err := l.Listener.Close(); err != nil {
			errs = append(errs, err)
		}
		if err := l.Server.Shutdown(context.Background()); err != nil {
			errs = append(errs, err)
		}
	}
	return errors.Join(errs...)
}

func (r *HttpResource) addHttpHandler(hostname string, baseUri string, handler httplistener.HttpHandler) {
	if strings.HasSuffix(baseUri, "/") {
		baseUri += "*"
	} else {
		baseUri += "/*"
	}

	mapping, ok := r.pathMapping[baseUri]
	if !ok {
		mapping = &struct {
			hostnameMapping map[string]*struct {
				handler httplistener.HttpHandler
			}
		}{
			hostnameMapping: map[string]*struct {
				handler httplistener.HttpHandler
			}{},
		}
		r.pathMapping[baseUri] = mapping
		r.r.HandleFunc(baseUri, func(w http.ResponseWriter, r *http.Request) {
			host := r.Host
			if strings.Contains(host, ":") {
				host = strings.TrimSpace(strings.Split(host, ":")[0])
			}
			h, ok := mapping.hostnameMapping[host]
			if !ok {
				http.NotFound(w, r)
				return
			}
			h.handler(w, r)
		})
	}
	h, ok := mapping.hostnameMapping[hostname]
	if !ok {
		h = &struct {
			handler httplistener.HttpHandler
		}{
			handler: handler,
		}
		mapping.hostnameMapping[hostname] = h
	} else {
		panic(errors.New("duplicated hostname with same baseurl in http resource: " + hostname))
	}
}
