package cmd

import (
	"github.com/edyoCampos/base365/internal/agent"
	"github.com/edyoCampos/base365/internal/audio"
	"github.com/edyoCampos/base365/internal/bus"
	"github.com/edyoCampos/base365/internal/cache"
	"github.com/edyoCampos/base365/internal/channelmemory"
	"github.com/edyoCampos/base365/internal/channels"
	"github.com/edyoCampos/base365/internal/config"
	"github.com/edyoCampos/base365/internal/eventbus"
	"github.com/edyoCampos/base365/internal/gateway"
	httpapi "github.com/edyoCampos/base365/internal/http"
	"github.com/edyoCampos/base365/internal/memory"
	"github.com/edyoCampos/base365/internal/providers"
	"github.com/edyoCampos/base365/internal/skills"
	"github.com/edyoCampos/base365/internal/store"
	"github.com/edyoCampos/base365/internal/tools"
	usagecaps "github.com/edyoCampos/base365/internal/usage/caps"
	"github.com/edyoCampos/base365/internal/vault"
)

// gatewayDeps holds shared dependencies used across the extracted gateway setup functions.
// It is populated in runGateway() and passed to helper methods to avoid long parameter lists.
type gatewayDeps struct {
	cfg              *config.Config
	server           *gateway.Server
	msgBus           *bus.MessageBus
	pgStores         *store.Stores
	providerRegistry *providers.Registry
	channelMgr       *channels.Manager
	channelMemorySvc *channelmemory.Service
	agentRouter      *agent.Router
	toolsReg         *tools.Registry
	skillsLoader     *skills.Loader         // optional: enables skill creation in evolution approval
	permCache        *cache.PermissionCache // nil if no tenant store; closed on shutdown to stop sweep goroutines
	enrichProgress   *vault.EnrichProgress  // nil if enrichment worker not registered
	enrichWorker     *vault.EnrichWorker    // nil if enrichment worker not registered; for stop/enqueue
	workspace        string
	dataDir          string
	domainBus        eventbus.DomainEventBus
	usageCapSvc      *usagecaps.Service
	audioMgr         *audio.Manager      // nil if TTS not configured; used by TTSHandler
	ttsHandler       *httpapi.TTSHandler // nil if TTS not configured; for hot-reload
	teamWorkEmbedder memory.EmbeddingProvider
}
