package ports

import "context"

type Out func(ctx context.Context, REPLACE_BY_YOUR_DOMAIN map[string]string) error
