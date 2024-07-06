export function hideEmail(email: string) {
  const matched = email.match(
    /^(?<prefix>[A-Z0-9._%+-])(?<host>[A-Z0-9._%+-]+)@(?<tdomain>[A-Z0-9.-]+)\.(?<domain>[A-Z]{2,})/i
  );
  if (!matched) return 'メールアドレス';

  const {prefix, host, domain} = matched.groups!;

  return `${prefix}${Array(host.length).fill('*').join('')}@****.${domain}`;
}
