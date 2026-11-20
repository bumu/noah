package dbmod

import (
	"testing"

	"github.com/jackc/pgx/v5"
	"gorm.io/driver/postgres"
)

func TestDatabaseDialector(t *testing.T) {
	for _, scheme := range []string{"postgres", "postgresql", "pgx", "pq", "mysql"} {
		t.Run(scheme, func(t *testing.T) {
			dialect, err := databaseDialector(scheme + "://test:example@localhost/testdb?sslmode=disable")
			if err != nil {
				t.Fatal(err)
			}
			want := "postgres"
			if scheme == "mysql" {
				want = "mysql"
			}
			if dialect.Name() != want {
				t.Fatalf("dialect = %q, want %q", dialect.Name(), want)
			}
			if p, ok := dialect.(*postgres.Dialector); ok {
				cfg, err := pgx.ParseConfig(p.DSN)
				if err != nil {
					t.Fatal(err)
				}
				if cfg.Database != "testdb" || cfg.User != "test" {
					t.Fatalf("unexpected parsed connection settings")
				}
			}
		})
	}
}

func TestDatabaseDialectorRejectsUnsupportedDatabase(t *testing.T) {
	if _, err := databaseDialector("sqlite:///tmp/test.db"); err == nil {
		t.Fatal("expected unsupported database error")
	}
}
