package cmd

import (
	"os"
	"time"

	"github.com/edyoCampos/base365/internal/config"
	"github.com/edyoCampos/base365/internal/edition"
	"github.com/edyoCampos/base365/internal/hooks"
	"github.com/edyoCampos/base365/internal/hooks/budget"
	hookhandlers "github.com/edyoCampos/base365/internal/hooks/handlers"
	"github.com/edyoCampos/base365/internal/providers"
	"github.com/edyoCampos/base365/internal/security"
	"github.com/edyoCampos/base365/internal/store"
	"github.com/edyoCampos/base365/internal/store/pg"
	usagecaps "github.com/edyoCampos/base365/internal/usage/caps"
)

// sharedHookHandlers is populated by wireExtras so the gateway.go router
// wiring can reuse the same handler instances for the `hooks.test` runner.
// nil when hook store is absent (hooks disabled).
var sharedHookHandlers map[hooks.HandlerType]hooks.Handler

// buildHookHandlers constructs the production handler map used by both the
// dispatcher (sync + async chain) and the `hooks.test` test runner. Keeping
// this factory single-source ensures test-panel behavior mirrors production.
//
// Budget wiring (C1 fix): the PromptHandler receives a budget.Store bound
// to pg.NewPGHookBudget so token spend is atomically deducted per tenant.
// When the DB handle is unavailable, budget falls back to nil (Lite desktop).
func buildHookHandlers(stores *store.Stores, providerReg *providers.Registry, hooksCfg config.HooksConfig, usageCapSvc *usagecaps.Service) map[hooks.HandlerType]hooks.Handler {
	encryptKey := os.Getenv("BASE365_ENCRYPTION_KEY")

	var budgetStore *budget.Store
	if stores != nil && stores.DB != nil {
		budgetStore = budget.New(pg.NewPGHookBudget(stores.DB), nil)
	}

	promptHandler := &hookhandlers.PromptHandler{
		Resolver:     hookhandlers.NewRegistryResolver(providerReg, stores.SystemConfigs),
		Budget:       budgetStore,
		UsageCaps:    usageCapSvc,
		DefaultModel: "haiku",
	}

	// ScriptHandler: bounded by cfg.Hooks caps; zero values fall back to
	// handler defaults (10 / 3 / 500). Safe for concurrent reuse across
	// dispatcher + hooks.test runner (each Execute allocates its own runtime).
	scriptHandler := hookhandlers.NewScriptHandler(
		hooksCfg.ScriptConcurrency,
		hooksCfg.ScriptPerTenantConcurrency,
		hooksCfg.ScriptCacheSize,
	)

	return map[hooks.HandlerType]hooks.Handler{
		hooks.HandlerCommand: &hookhandlers.CommandHandler{Edition: edition.Current()},
		hooks.HandlerHTTP: &hookhandlers.HTTPHandler{
			EncryptKey: encryptKey,
			Client:     security.NewSafeClient(10 * time.Second),
		},
		hooks.HandlerPrompt: promptHandler,
		hooks.HandlerScript: scriptHandler,
	}
}
