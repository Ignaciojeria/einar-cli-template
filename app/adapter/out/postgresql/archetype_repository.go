package postgresql

import (
	einar "archetype/app/shared/archetype/postgres"
	"context"

	"gorm.io/gorm"
)

var ArchetypeRepository = func(ctx context.Context, REPLACE_BY_YOUR_DOMAIN map[string]string) error {
	_, span := tracer.Start(ctx, "ArchetypeRepository")
	defer span.End()

	var _ *gorm.DB = einar.DB
	//PUT YOUR POSTGRESL OPERATION USING EINAR HERE :
	//....einar.DB....
	return nil
}
