import {config} from '../config';

/**
 * @param {string} url Path to the page. When overriding, the value must be prefixed with a slash "/".
 */
export function pageview(url: string) {
  if (config.gaTracingId) {
    window.gtag('config', config.gaTracingId, {
      page_path: url,
    });
  }
}
