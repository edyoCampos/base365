import { useState } from "react";
import { useTranslation } from "react-i18next";
import { ChevronDown, ChevronRight, Shield } from "lucide-react";
import { requiredScopes } from "./channel-schemas";

interface ChannelScopesInfoProps {
  channelType: string;
}

export function ChannelScopesInfo({ channelType }: ChannelScopesInfoProps) {
  const { t } = useTranslation("channels");
  const scopes = requiredScopes[channelType];
  const [expanded, setExpanded] = useState(false);

  if (!scopes || scopes.length === 0) return null;

  return (
    <div className="rounded-md border border-warning/50 bg-badge-warning text-sm">
      <button
        type="button"
        onClick={() => setExpanded(!expanded)}
        className="flex items-center gap-2 w-full px-3 py-2 text-left text-accent-foreground dark:text-primary hover:bg-accent rounded-md transition-colors"
      >
        <Shield className="h-4 w-4 shrink-0" />
        <span className="flex-1 font-medium">{t("scopes.title")}</span>
        {expanded
          ? <ChevronDown className="h-4 w-4 shrink-0" />
          : <ChevronRight className="h-4 w-4 shrink-0" />}
      </button>
      {expanded && (
        <div className="px-3 pb-3 space-y-2">
          <p className="text-xs text-accent-foreground dark:text-primary">
            {t("scopes.description")}
          </p>
          <div className="space-y-0.5">
            {scopes.map((s) => (
              <div key={s.scope} className="flex items-baseline gap-2 text-xs font-mono">
                <code className="text-accent-foreground dark:text-primary">{s.scope}</code>
                {s.note && (
                  <span className="text-primary dark:text-primary font-sans text-xs-plus">
                    — {s.note}
                  </span>
                )}
              </div>
            ))}
          </div>
          <p className="text-xs text-primary pt-1">
            {t("scopes.publishReminder")}
          </p>
        </div>
      )}
    </div>
  );
}
