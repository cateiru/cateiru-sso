/**
 * @param {string} key - cookie key
 * @returns {string} - cookie value
 */
export default function cookieValue(key: string): string {
  return ((document.cookie + ';').match(key + '=([^¥S;]*)') || [])[1];
}
