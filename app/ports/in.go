package ports

import "context"

type In func(ctx context.Context, REPLACE_BY_YOUR_DOMAIN map[string]string) error
