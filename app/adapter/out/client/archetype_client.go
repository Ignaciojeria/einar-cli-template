package client

import (
	"archetype/app/domain/ports/out"
	einar "archetype/app/shared/archetype/resty"
	"context"

	"github.com/go-resty/resty/v2"
)

var ArchetypeRestyClient out.ArchetypeOutBoundPort = func(ctx context.Context, REPLACE_BY_YOUR_DOMAIN map[string]string) (err error) {
	var _ *resty.Client = einar.Client
	//PUT YOUR HTTP OPERATION USING EINAR HERE :
	//....einar.Client
	return nil
}
