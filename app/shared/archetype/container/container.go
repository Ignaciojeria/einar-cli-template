package container

import (
	"errors"
	"os"
	"sync"

	"github.com/rs/zerolog/log"
)

type DependencyContainer struct {
	InjectionProps InjectionProps
	LoadDependency LoadDependency
	isPresent      bool
}

type InjectionProps struct {
	DependencyID string // name of injected dependency should be unique and required
}

type LoadDependency func() error

var (
	InstallationsContainer   = make(map[string]DependencyContainer)
	InboundAdapterContainer  = make(map[string]DependencyContainer)
	OutboundAdapterContainer = make(map[string]DependencyContainer)
	HTTPServerContainer      DependencyContainer
)

// Mutexes for each container
var (
	containerMutex = &sync.Mutex{}
)

func Inject(dependency LoadDependency, props InjectionProps, container map[string]DependencyContainer, mutex *sync.Mutex) error {
	if props.DependencyID == "" {
		err := errors.New("container injector error on InjectionProps. DependencyID can't be empty")
		log.Error().Err(err).Send()
		return err
	}

	mutex.Lock()
	defer mutex.Unlock()

	if _, exists := container[props.DependencyID]; exists {
		err := errors.New("container injector error. Next dependency already exists: " + props.DependencyID)
		log.Error().Err(err).Send()
		return err
	}

	container[props.DependencyID] = DependencyContainer{LoadDependency: dependency, InjectionProps: props, isPresent: true}

	return nil
}

func InjectInBoundAdapter(dependency LoadDependency, props InjectionProps) error {
	return Inject(dependency, props, InboundAdapterContainer, containerMutex)
}

func InjectOutBoundAdapter(dependency LoadDependency, props InjectionProps) error {
	return Inject(dependency, props, OutboundAdapterContainer, containerMutex)
}

func InjectInstallation(dependency LoadDependency, props InjectionProps) error {
	return Inject(dependency, props, InstallationsContainer, containerMutex)
}

func InjectHTTPServer(dependency LoadDependency, props InjectionProps) error {
	HTTPServerContainer = DependencyContainer{LoadDependency: dependency, InjectionProps: props, isPresent: true}
	return nil
}

type IExit func() error

var Exit IExit = func() error {
	// Implement any cleanup tasks here
	os.Exit(0)
	return nil
}
