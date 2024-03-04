export function badgeColor(role: string) {
  if (role === 'owner') {
    return 'red';
  }

  if (role === 'member') {
    return 'blue';
  }

  return 'gray';
}
