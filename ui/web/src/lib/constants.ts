// Barrel re-exports for backward compatibility.
// Import directly from sub-modules for new code.
export { ROUTES } from "./routes";
export {
  TIMEZONE_OPTIONS,
  getAllIanaTimezones,
  isValidIanaTimezone,
} from "./timezone-utils";

export const LOCAL_STORAGE_KEYS = {
  TOKEN: "base365:token",
  USER_ID: "base365:userId",
  SENDER_ID: "base365:senderID",
  TENANT_ID: "base365:tenant_id",
  TENANT_HINT: "base365:tenant_hint",
  SETUP_SKIPPED: "base365:setup_skipped",
  THEME: "base365:theme",
  SIDEBAR_COLLAPSED: "base365:sidebarCollapsed",
  LANGUAGE: "base365:language",
  TIMEZONE: "base365:timezone",
} as const;

export const SUPPORTED_LANGUAGES = ["en", "vi", "zh", "ko", "ru"] as const;
export type Language = (typeof SUPPORTED_LANGUAGES)[number];

export const LANGUAGE_LABELS: Record<Language, string> = {
  en: "English",
  vi: "Tiếng Việt",
  zh: "中文",
  ko: "한국어",
  ru: "Русский",
};
