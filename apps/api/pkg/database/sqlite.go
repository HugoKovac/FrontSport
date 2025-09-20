package database

import (
	"GoNext/base/ent"
	"GoNext/base/pkg/config"
	"context"
	"fmt"
	"log"

	"ariga.io/atlas/sql/migrate"
	atlas "ariga.io/atlas/sql/schema"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
	_ "github.com/mattn/go-sqlite3"
)

// NewEntClient creates a new Ent client connected to PostgreSQL
func NewSQLiteEntClient(config *config.Config) *ent.Client {

	dsn := fmt.Sprintf("file:%s?cache=shared&_fk=1", config.Db.Sqlite.Path)

	client, err := ent.Open(dialect.Postgres, dsn)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	// Run the migrations
	ctx := context.Background()
	// Run migration.
	err = client.Schema.Create(
		ctx,
		// Hook into Atlas Diff process.
		schema.WithDiffHook(func(next schema.Differ) schema.Differ {
			return schema.DiffFunc(func(current, desired *atlas.Schema) ([]atlas.Change, error) {
				// Before calculating changes.
				changes, err := next.Diff(current, desired)
				if err != nil {
					return nil, err
				}
				// After diff, you can filter
				// changes or return new ones.
				return changes, nil
			})
		}),
		// Hook into Atlas Apply process.
		schema.WithApplyHook(func(next schema.Applier) schema.Applier {
			return schema.ApplyFunc(func(ctx context.Context, conn dialect.ExecQuerier, plan *migrate.Plan) error {
				// Example to hook into the apply process, or implement
				// a custom applier. For example, write to a file.
				//
				//  for _, c := range plan.Changes {
				//      fmt.Printf("%s: %s", c.Comment, c.Cmd)
				//      if err := conn.Exec(ctx, c.Cmd, c.Args, nil); err != nil {
				//          return err
				//      }
				//  }
				//
				return next.Apply(ctx, conn, plan)
			})
		}),
	)

	return client
}
