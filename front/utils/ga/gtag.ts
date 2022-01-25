export const GA_TRACKING_ID = process.env.GOOGLE_ANALYTICS_ID || '';

/**
 * @param {string} url Path to the page. When overriding, the value must be prefixed with a slash "/".
 */
export function pageview(url: string) {
  if (GA_TRACKING_ID) {
    window.gtag('config', GA_TRACKING_ID, {
      page_path: url,
    });
  }
}
