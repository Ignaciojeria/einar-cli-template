package container

import (
	"archetype/app/shared/archetype/slog"
	"errors"
	"os"

	"github.com/google/uuid"
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
	ConfigurationContainer   = make(map[string]DependencyContainer)
	InstallationsContainer   = make(map[string]DependencyContainer)
	InboundAdapterContainer  = make(map[string]DependencyContainer)
	OutboundAdapterContainer = make(map[string]DependencyContainer)
	HTTPServerContainer      DependencyContainer
)

func Inject(dependency LoadDependency, props InjectionProps, container map[string]DependencyContainer) error {
	if props.DependencyID == "" {
		err := errors.New("container injector error on InjectionProps. DependencyID can't be empty")
		slog.Logger().Error(err.Error())
		return err
	}

	if _, exists := container[props.DependencyID]; exists {
		err := errors.New("container injector error. Next dependency already exists: " + props.DependencyID)
		slog.Logger().Error(err.Error())
		return err
	}

	container[props.DependencyID] = DependencyContainer{LoadDependency: dependency, InjectionProps: props, isPresent: true}

	return nil
}

func InjectInboundAdapter(dependency LoadDependency, props ...InjectionProps) error {
	return Inject(dependency, InjectionProps{uuid.NewString()}, InboundAdapterContainer)
}

func InjectOutboundAdapter(dependency LoadDependency, props ...InjectionProps) error {
	return Inject(dependency, InjectionProps{uuid.NewString()}, OutboundAdapterContainer)
}

func InjectInstallation(dependency LoadDependency, props ...InjectionProps) error {
	return Inject(dependency, InjectionProps{uuid.NewString()}, InstallationsContainer)
}

func InjectConfiguration(dependency LoadDependency, props ...InjectionProps) error {
	return Inject(dependency, InjectionProps{uuid.NewString()}, InstallationsContainer)
}

func InjectHTTPServer(dependency LoadDependency, props ...InjectionProps) error {
	HTTPServerContainer = DependencyContainer{LoadDependency: dependency, InjectionProps: InjectionProps{uuid.NewString()}, isPresent: true}
	return nil
}

type IExit func() error

var Exit IExit = func() error {
	// Implement any cleanup tasks here
	os.Exit(0)
	return nil
}
