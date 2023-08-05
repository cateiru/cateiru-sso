export function formatRedirectUrl(url: string): string {
  const base = window.location.href;
  const u = new URL(url, base);
  return u.pathname;
}
