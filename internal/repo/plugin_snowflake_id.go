package repo

import (
	"reflect"

	"exchange-backend/internal/pkg/snowflake"
	"gorm.io/gorm"
)

func NewSnowflakeIDPlugin(snowflakeNode *snowflake.Node) *SnowflakeIDPlugin {
	return &SnowflakeIDPlugin{
		snowflakeNode: snowflakeNode,
	}
}

type SnowflakeIDPlugin struct {
	snowflakeNode *snowflake.Node
}

func (p *SnowflakeIDPlugin) Name() string {
	return "snowflake_id"
}

func (p *SnowflakeIDPlugin) Initialize(db *gorm.DB) error {
	if err := db.Callback().Create().Before("gorm:create").Register("snowflake_id:inject", p.beforeCreate); err != nil {
		return err
	}

	return nil
}

func (p *SnowflakeIDPlugin) beforeCreate(db *gorm.DB) {
	if db.Error != nil {
		return
	}

	ctx := db.Statement.Context

	idField := db.Statement.Schema.LookUpField("id")
	if idField == nil {
		return
	}

	switch db.Statement.ReflectValue.Kind() {
	case reflect.Slice, reflect.Array:
		{
			ids, err := p.snowflakeNode.BulkGenerate(ctx, db.Statement.ReflectValue.Len())
			if err != nil {
				db.Error = err
				return
			}

			for i := 0; i < db.Statement.ReflectValue.Len(); i++ {
				id := ids[i]
				value := db.Statement.ReflectValue.Index(i)

				if _, isZero := idField.ValueOf(ctx, value); isZero {
					db.Error = idField.Set(ctx, value, id.Int64())
					if db.Error != nil {
						return
					}
				}
			}
		}
	case reflect.Struct:
		{
			id, err := p.snowflakeNode.Generate(ctx)
			if err != nil {
				db.Error = err
				return
			}

			value := db.Statement.ReflectValue

			if _, isZero := idField.ValueOf(ctx, value); isZero {
				db.Error = idField.Set(ctx, value, id.Int64())
				if db.Error != nil {
					return
				}
			}
		}
	}
}
