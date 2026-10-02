package cmd

import (
	"sync"

	"github.com/edyoCampos/base365/internal/agent"
	"github.com/edyoCampos/base365/internal/bus"
	"github.com/edyoCampos/base365/internal/channels"
	"github.com/edyoCampos/base365/internal/config"
	"github.com/edyoCampos/base365/internal/memory"
	"github.com/edyoCampos/base365/internal/providers"
	"github.com/edyoCampos/base365/internal/scheduler"
	"github.com/edyoCampos/base365/internal/store"
	"github.com/edyoCampos/base365/internal/tools"
	usagecaps "github.com/edyoCampos/base365/internal/usage/caps"
)

// ConsumerDeps bundles shared dependencies for consumer message handlers.
// Replaces 11+ positional params with a single injectable struct.
type ConsumerDeps struct {
	Cfg              *config.Config
	Agents           *agent.Router
	Sched            *scheduler.Scheduler
	ChannelMgr       *channels.Manager
	MsgBus           *bus.MessageBus
	TeamStore        store.TeamStore
	AgentLinkStore   store.AgentLinkStore
	AgentStore       store.AgentStore
	SessStore        store.SessionStore
	PostTurn         tools.PostTurnProcessor
	QuotaChecker     *channels.QuotaChecker
	ContactCollector *store.ContactCollector
	TaskRunSessions  sync.Map
	SubagentMgr      *tools.SubagentManager
	UsageCaps        *usagecaps.Service
	ProviderReg      *providers.Registry
	TeamWorkEmbedder memory.EmbeddingProvider
	BgWg             sync.WaitGroup
	GetAnnounceMu    func(string) *sync.Mutex
}
