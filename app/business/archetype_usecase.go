package business

import (
	"context"
)

var ArchetypeUseCase = func(ctx context.Context, REPLACE_BY_YOUR_DOMAIN map[string]string) error {
	_, span := tracer.Start(ctx, "ArchetypeUseCase")
	defer span.End()
	//IMPLEMENT YOUR BUSINESS USECASE HERE
	return nil
}
