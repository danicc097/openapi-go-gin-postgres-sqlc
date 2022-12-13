package main

import (
	"flag"
	"fmt"
	"log"
	"path"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/envvar"
	"github.com/go-jet/jet/v2/generator/metadata"
	"github.com/go-jet/jet/v2/generator/postgres"
	"github.com/go-jet/jet/v2/generator/template"
	postgres2 "github.com/go-jet/jet/v2/postgres"
	_ "github.com/lib/pq"
)

// clear && go run cmd/cli/main.go -env .env.dev
func main() {
	var out, env, schema, dbname string

	flag.StringVar(&dbname, "dbname", "public", "Database name to generate from")
	flag.StringVar(&schema, "schema", "public", "Database schema to generate from")
	flag.StringVar(&env, "env", ".env", "Environment Variables filename")
	flag.StringVar(&out, "out", "", "Out dir for generated files")
	flag.Parse()

	if out == "" {
		log.Fatal("--out flag is required")
	}

	if err := envvar.Load(env); err != nil {
		log.Fatalf("envvar.Load: %s\n", err)
	}

	appEnv := envvar.GetEnv("APP_ENV", "dev")
	if err := envvar.Load(path.Join(".env." + appEnv)); err != nil {
		log.Fatalf("envvar.Load: %s\n", err)
	}

	cfg := internal.Config()
	dbConnection := postgres.DBConnection{
		Host:       cfg.Postgres.Server,
		Port:       cfg.Postgres.Port,
		User:       cfg.Postgres.User,
		Password:   cfg.Postgres.Password,
		DBName:     dbname,
		SchemaName: schema,
		SslMode:    "disable",
	}

	err := postgres.Generate(
		out,
		dbConnection,
		template.Default(postgres2.Dialect).
			UseSchema(func(schemaMetaData metadata.Schema) template.Schema {
				return template.DefaultSchema(schemaMetaData).
					UsePath("/" + schema).
					// UseSQLBuilder(template.DefaultSQLBuilder().UsePath("/" + schema)).
					UseModel(template.DefaultModel().
						UseTable(func(table metadata.Table) template.TableModel {
							return template.DefaultTableModel(table).
								UseField(func(columnMetaData metadata.Column) template.TableModelField {
									defaultTableModelField := template.DefaultTableModelField(columnMetaData)

									return defaultTableModelField.UseTags(
										// fmt.Sprintf(`json:"%s"`, snaker.ForceLowerCamelIdentifier(snaker.SnakeToCamel(columnMetaData.Name))),
										fmt.Sprintf(`db:"%s"`, columnMetaData.Name),
									)
								})
						}).
						UseView(func(table metadata.Table) template.ViewModel {
							return template.DefaultViewModel(table).
								UseField(func(columnMetaData metadata.Column) template.TableModelField {
									defaultTableModelField := template.DefaultTableModelField(columnMetaData)
									if table.Name == "actor_info" && columnMetaData.Name == "actor_id" {
										return defaultTableModelField.UseTags(`sql:"primary_key"`)
									}
									return defaultTableModelField
								})
						}),
					)
			}),
	)
	if err != nil {
		log.Fatalf("jet generation failed: %v", err)
	}
}
