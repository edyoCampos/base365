//go:build !sqlite && !sqliteonly

package cmd

import (
	"database/sql"

	"github.com/edyoCampos/base365/internal/store"
	"github.com/edyoCampos/base365/internal/store/pg"
)

func newVaultGraphStore(db *sql.DB) store.VaultGraphStore { return pg.NewPGVaultGraphStore(db) }
func newKGGraphStore(db *sql.DB) store.KGGraphStore       { return pg.NewPGKGGraphStore(db) }
