package core

type Service interface{}

// ServiceRegistry holds the services in a map.
var services = make(map[string]Service)

// Bind adds a service to the registry with the given name.
func ServiceRegistryBind(name string, service Service) {
	services[name] = service
}

// Unbind removes the service from the registry by name.
func ServiceRegistryUnbind(name string) {
	delete(services, name)
}

// Lookup retrieves a service from the registry by name.
func ServiceRegistryLookup(name string) Service {
	return services[name]
}

// List returns a slice of ServiceTuple containing all services in the registry.
func ServiceRegistryList() []string {
	serviceList := []string{}
	for k := range services {
		serviceList = append(serviceList, k)
	}
	return serviceList
}
