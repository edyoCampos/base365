# Revisão do âmbar e do laranja (RF-20)

Regra (PDF de tokens, seção 06): se remover a cor faz o usuário perder a informação de que algo precisa de atenção, é **aviso** (token `warning`). Senão é **decoração** (violeta: `primary`/`accent`). Âmbar nunca é série de gráfico de status.

A classificação automática olha 4 linhas antes e depois de cada ocorrência (e o nome do arquivo) à procura de termos de atenção (warn, alert, limit, expire, pending, retry, offline, etc.). Revisar a coluna **aviso**; o que sobrar como decoração foi convertido para a paleta.

## Totais

| Superfície | Cor | Aviso | Decoração |
|---|---|---|---|
| web | amber | 258 | 201 |
| web | orange | 45 | 122 |
| desktop | amber | 40 | 27 |
| desktop | orange | 0 | 36 |

## Por arquivo (aviso / decoração)

| Arquivo | Aviso | Decoração |
|---|---|---|
| `ui/desktop/frontend/src/components/agents/AgentCard.tsx` | 2 | 10 |
| `ui/desktop/frontend/src/components/agents/evolution-section-expanded.tsx` | 0 | 4 |
| `ui/desktop/frontend/src/components/agents/evolution-suggestions-list.tsx` | 0 | 5 |
| `ui/desktop/frontend/src/components/agents/sandbox-section.tsx` | 4 | 0 |
| `ui/desktop/frontend/src/components/builtin-tools/stt-provider-form.tsx` | 3 | 0 |
| `ui/desktop/frontend/src/components/channels/ChannelList.tsx` | 4 | 0 |
| `ui/desktop/frontend/src/components/chat/ActivityIndicator.tsx` | 2 | 0 |
| `ui/desktop/frontend/src/components/chat/ThinkingBlock.tsx` | 0 | 1 |
| `ui/desktop/frontend/src/components/chat/ToolCallBlock.tsx` | 0 | 2 |
| `ui/desktop/frontend/src/components/common/Toaster.tsx` | 4 | 0 |
| `ui/desktop/frontend/src/components/cron/cron-list-item.tsx` | 0 | 6 |
| `ui/desktop/frontend/src/components/layout/sidebar/SidebarFooter.tsx` | 1 | 0 |
| `ui/desktop/frontend/src/components/mcp/McpServerList.tsx` | 2 | 0 |
| `ui/desktop/frontend/src/components/onboarding/SummoningModal.tsx` | 0 | 4 |
| `ui/desktop/frontend/src/components/onboarding/summoning-progress-steps.tsx` | 0 | 6 |
| `ui/desktop/frontend/src/components/skills/SkillList.tsx` | 4 | 4 |
| `ui/desktop/frontend/src/components/skills/skill-card.tsx` | 2 | 6 |
| `ui/desktop/frontend/src/components/storage/file-tree-node.tsx` | 0 | 2 |
| `ui/desktop/frontend/src/components/teams/KanbanCard.tsx` | 2 | 1 |
| `ui/desktop/frontend/src/components/teams/TeamCreateDialog.tsx` | 0 | 1 |
| `ui/desktop/frontend/src/components/teams/TeamSettingsModal.tsx` | 1 | 0 |
| `ui/desktop/frontend/src/components/teams/task-detail-meta.tsx` | 4 | 0 |
| `ui/desktop/frontend/src/components/teams/team-member-list.tsx` | 0 | 6 |
| `ui/desktop/frontend/src/components/traces/trace-span-tree.tsx` | 0 | 2 |
| `ui/desktop/frontend/src/types/team.ts` | 4 | 3 |
| `ui/desktop/frontend/src/utils/channel-status.ts` | 1 | 0 |
| `ui/web/src/adapters/vault-graph-adapter.ts` | 0 | 4 |
| `ui/web/src/components/agents/v3-capabilities-modal/pipeline-tab.tsx` | 0 | 3 |
| `ui/web/src/components/agents/v3-info-modal/v3-feature-card.tsx` | 0 | 2 |
| `ui/web/src/components/agents/v3-info-modal/v3-info-tab-memory.tsx` | 0 | 1 |
| `ui/web/src/components/agents/v3-info-modal/v3-info-tab-orchestration.tsx` | 0 | 1 |
| `ui/web/src/components/chat/activity-indicator.tsx` | 2 | 1 |
| `ui/web/src/components/chat/chat-top-bar.tsx` | 0 | 2 |
| `ui/web/src/components/chat/rich-content.tsx` | 6 | 2 |
| `ui/web/src/components/chat/thinking-block.tsx` | 0 | 1 |
| `ui/web/src/components/chat/tool-call-card.tsx` | 0 | 2 |
| `ui/web/src/components/layout/about-dialog.tsx` | 0 | 3 |
| `ui/web/src/components/layout/system-settings-compaction-card.tsx` | 6 | 2 |
| `ui/web/src/components/layout/system-settings-embedding-card.tsx` | 0 | 2 |
| `ui/web/src/components/layout/system-settings-modal.tsx` | 9 | 0 |
| `ui/web/src/components/layout/topbar.tsx` | 0 | 1 |
| `ui/web/src/components/shared/drag-preview.tsx` | 0 | 1 |
| `ui/web/src/components/shared/file-tree-file-icon.tsx` | 0 | 4 |
| `ui/web/src/components/shared/markdown-callout-block.tsx` | 8 | 0 |
| `ui/web/src/components/ui/badge.tsx` | 6 | 0 |
| `ui/web/src/components/ui/toaster.tsx` | 4 | 0 |
| `ui/web/src/pages/agents/agent-card.tsx` | 3 | 7 |
| `ui/web/src/pages/agents/agent-detail/agent-files-tab.tsx` | 4 | 0 |
| `ui/web/src/pages/agents/agent-detail/agent-header.tsx` | 0 | 7 |
| `ui/web/src/pages/agents/agent-detail/agent-overview-tab.tsx` | 2 | 0 |
| `ui/web/src/pages/agents/agent-detail/agent-permissions-tab.tsx` | 6 | 1 |
| `ui/web/src/pages/agents/agent-detail/chatgpt-oauth-quota-strip.tsx` | 1 | 0 |
| `ui/web/src/pages/agents/agent-detail/codex-pool-page-header.tsx` | 4 | 0 |
| `ui/web/src/pages/agents/agent-detail/codex-pool-request-accent.ts` | 0 | 35 |
| `ui/web/src/pages/agents/agent-detail/config-sections/chatgpt-oauth-pool-sections.tsx` | 0 | 11 |
| `ui/web/src/pages/agents/agent-detail/config-sections/workspace-sharing-section.tsx` | 9 | 26 |
| `ui/web/src/pages/agents/agent-detail/evolution-tab/evolution-suggestions-table.tsx` | 0 | 6 |
| `ui/web/src/pages/agents/agent-detail/file-sections/system-prompt-preview.tsx` | 0 | 3 |
| `ui/web/src/pages/agents/agent-detail/heartbeat-config-dialog.tsx` | 0 | 1 |
| `ui/web/src/pages/agents/agent-detail/heartbeat-schedule-section.tsx` | 0 | 1 |
| `ui/web/src/pages/agents/agent-detail/overview-sections/chatgpt-oauth-routing-summary-section.tsx` | 0 | 4 |
| `ui/web/src/pages/agents/agent-detail/overview-sections/evolution-section.tsx` | 12 | 2 |
| `ui/web/src/pages/agents/agent-detail/overview-sections/heartbeat-card.tsx` | 0 | 2 |
| `ui/web/src/pages/agents/agent-detail/overview-sections/hooks-summary-card.tsx` | 0 | 8 |
| `ui/web/src/pages/agents/agent-detail/overview-sections/pinned-skills-section.tsx` | 0 | 1 |
| `ui/web/src/pages/agents/agent-detail/overview-sections/skills-section.tsx` | 0 | 1 |
| `ui/web/src/pages/agents/agent-detail/prompt-mode-badge-utils.ts` | 0 | 4 |
| `ui/web/src/pages/agents/agent-list-row.tsx` | 0 | 10 |
| `ui/web/src/pages/agents/summoning-modal.tsx` | 0 | 10 |
| `ui/web/src/pages/api-keys/api-key-code-dialog.tsx` | 0 | 2 |
| `ui/web/src/pages/backup-restore/backup-preflight-panel.tsx` | 18 | 0 |
| `ui/web/src/pages/backup-restore/system-restore-panel.tsx` | 4 | 0 |
| `ui/web/src/pages/builtin-tools/builtin-tool-settings-dialog.tsx` | 0 | 4 |
| `ui/web/src/pages/builtin-tools/builtin-tools-page.tsx` | 8 | 2 |
| `ui/web/src/pages/builtin-tools/stt-provider-form.tsx` | 6 | 0 |
| `ui/web/src/pages/builtin-tools/web-search-chain-form.tsx` | 0 | 1 |
| `ui/web/src/pages/channels/bitrix24/bitrix-portal-authorize-step.tsx` | 9 | 0 |
| `ui/web/src/pages/channels/bitrix24/bitrix-portal-help-section.tsx` | 1 | 0 |
| `ui/web/src/pages/channels/bitrix24/bitrix-portal-select.tsx` | 1 | 0 |
| `ui/web/src/pages/channels/channel-instance-form-step.tsx` | 0 | 1 |
| `ui/web/src/pages/channels/channel-scopes-info.tsx` | 4 | 12 |
| `ui/web/src/pages/channels/channels-status-utils.ts` | 5 | 0 |
| `ui/web/src/pages/cli-credentials/cli-credential-env-vars-section.tsx` | 2 | 0 |
| `ui/web/src/pages/cli-credentials/cli-credential-grant-env-row.tsx` | 2 | 0 |
| `ui/web/src/pages/cli-credentials/cli-user-credentials-dialog.tsx` | 6 | 0 |
| `ui/web/src/pages/config/config-page.tsx` | 4 | 0 |
| `ui/web/src/pages/config/sections/behavior-pending-compaction-card.tsx` | 8 | 0 |
| `ui/web/src/pages/config/sections/behavior-security-card.tsx` | 6 | 1 |
| `ui/web/src/pages/config/sections/behavior-ux-card.tsx` | 6 | 1 |
| `ui/web/src/pages/config/sections/tools-browser-section.tsx` | 0 | 1 |
| `ui/web/src/pages/contacts/contacts-page.tsx` | 0 | 2 |
| `ui/web/src/pages/cron/cron-detail/cron-overview-tab.tsx` | 8 | 2 |
| `ui/web/src/pages/events/event-sections/agent-event-cards.tsx` | 4 | 4 |
| `ui/web/src/pages/events/event-sections/event-categories.ts` | 0 | 2 |
| `ui/web/src/pages/events/event-sections/event-detail-dialog.tsx` | 0 | 4 |
| `ui/web/src/pages/hooks/components/hook-history-table.tsx` | 2 | 0 |
| `ui/web/src/pages/hooks/components/hook-list-row.tsx` | 0 | 8 |
| `ui/web/src/pages/hooks/components/hook-overview-tab.tsx` | 0 | 8 |
| `ui/web/src/pages/hooks/components/hook-test-panel.tsx` | 4 | 0 |
| `ui/web/src/pages/import-export/agent-import-panel.tsx` | 1 | 0 |
| `ui/web/src/pages/import-export/import-export-page.tsx` | 6 | 0 |
| `ui/web/src/pages/import-export/team-import-panel.tsx` | 1 | 0 |
| `ui/web/src/pages/login/pairing-form.tsx` | 0 | 1 |
| `ui/web/src/pages/login/tenant-selector.tsx` | 4 | 0 |
| `ui/web/src/pages/mcp/mcp-oauth-dialog.tsx` | 3 | 0 |
| `ui/web/src/pages/mcp/mcp-user-credentials-dialog.tsx` | 9 | 0 |
| `ui/web/src/pages/overview/channel-attention-panel.tsx` | 4 | 0 |
| `ui/web/src/pages/overview/quota-usage-card.tsx` | 1 | 0 |
| `ui/web/src/pages/overview/system-health-card.tsx` | 0 | 3 |
| `ui/web/src/pages/packages/components/source-pill.tsx` | 0 | 4 |
| `ui/web/src/pages/packages/components/update-all-modal.tsx` | 2 | 0 |
| `ui/web/src/pages/packages/github-binaries-section.tsx` | 12 | 4 |
| `ui/web/src/pages/pending-messages/message-list-dialog.tsx` | 0 | 2 |
| `ui/web/src/pages/providers/provider-cli-section.tsx` | 8 | 10 |
| `ui/web/src/pages/providers/provider-detail/provider-embedding-section.tsx` | 0 | 2 |
| `ui/web/src/pages/providers/provider-list-row.tsx` | 2 | 0 |
| `ui/web/src/pages/providers/provider-oauth-section.tsx` | 0 | 4 |
| `ui/web/src/pages/providers/provider-utils.tsx` | 2 | 0 |
| `ui/web/src/pages/sessions/run-timeline-panel.tsx` | 0 | 3 |
| `ui/web/src/pages/sessions/sessions-page.tsx` | 0 | 1 |
| `ui/web/src/pages/setup/step-agent.tsx` | 6 | 0 |
| `ui/web/src/pages/skills/missing-deps-panel.tsx` | 22 | 0 |
| `ui/web/src/pages/skills/skill-health-summary.tsx` | 3 | 0 |
| `ui/web/src/pages/skills/skill-upload-entry.tsx` | 2 | 0 |
| `ui/web/src/pages/teams/board/board-header.tsx` | 0 | 4 |
| `ui/web/src/pages/teams/board/board-utils.ts` | 1 | 0 |
| `ui/web/src/pages/teams/board/kanban-card.tsx` | 2 | 1 |
| `ui/web/src/pages/teams/board/team-info-dialog.tsx` | 0 | 4 |
| `ui/web/src/pages/teams/board/team-members-dialog.tsx` | 0 | 8 |
| `ui/web/src/pages/teams/task-sections/task-detail-dialog.tsx` | 1 | 4 |
| `ui/web/src/pages/teams/task-sections/task-list.tsx` | 0 | 4 |
| `ui/web/src/pages/teams/team-audit-logs-modal.tsx` | 8 | 0 |
| `ui/web/src/pages/teams/team-card.tsx` | 0 | 1 |
| `ui/web/src/pages/teams/team-create-dialog.tsx` | 0 | 2 |
| `ui/web/src/pages/teams/team-features-modal.tsx` | 0 | 2 |
| `ui/web/src/pages/teams/team-notifications-section.tsx` | 2 | 1 |
| `ui/web/src/pages/teams/team-orchestration-section.tsx` | 3 | 5 |
| `ui/web/src/pages/tenants-admin/tenant-detail-page.tsx` | 0 | 8 |
| `ui/web/src/pages/traces/trace-span-tree-node.tsx` | 0 | 2 |
| `ui/web/src/pages/vault/components/vault-tree.tsx` | 1 | 4 |
| `ui/web/src/pages/vault/vault-document-sidebar.tsx` | 0 | 8 |
| `ui/web/src/pages/vault/vault-documents-table.tsx` | 0 | 4 |
| `ui/web/src/pages/webhooks/webhook-secret-dialog.tsx` | 3 | 0 |
| `ui/web/src/pages/webhooks/webhook-test-dialog.tsx` | 4 | 0 |
