package methods

import (
	"context"
	"slices"

	"github.com/edyoCampos/base365/internal/config"
	"github.com/edyoCampos/base365/internal/gateway"
	"github.com/edyoCampos/base365/internal/i18n"
	"github.com/edyoCampos/base365/internal/permissions"
	"github.com/edyoCampos/base365/internal/store"
	"github.com/edyoCampos/base365/pkg/protocol"
)

// canSeeAll checks if user has full data visibility (admin role OR owner user).
func canSeeAll(role permissions.Role, ownerIDs []string, userID string) bool {
	if permissions.HasMinRole(role, permissions.RoleAdmin) {
		return true
	}
	if userID != "" && slices.Contains(ownerIDs, userID) {
		return true
	}
	return false
}

// requireSessionOwner verifies the caller owns the session identified by key.
// Returns true if the caller is authorized (owner, admin, or matching user).
// On failure, sends an error response and returns false.
func requireSessionOwner(ctx context.Context, sessions store.SessionStore, cfg *config.Config, client *gateway.Client, reqID string, key string) bool {
	if canSeeAll(client.Role(), cfg.Gateway.OwnerIDs, client.UserID()) {
		return true
	}
	locale := store.LocaleFromContext(ctx)
	sess := sessions.Get(ctx, key)
	if sess == nil {
		client.SendResponse(protocol.NewErrorResponse(reqID, protocol.ErrNotFound, i18n.T(locale, i18n.MsgNotFound, "session", key)))
		return false
	}
	if sess.UserID != client.UserID() {
		client.SendResponse(protocol.NewErrorResponse(reqID, protocol.ErrUnauthorized, i18n.T(locale, i18n.MsgPermissionDenied, "session")))
		return false
	}
	return true
}
