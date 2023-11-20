export const hawManyDaysAgo = (date: Date): string => {
  const now = new Date();
  const diffSec = Math.floor((now.getTime() - date.getTime()) / 1000);

  const rtf = new Intl.RelativeTimeFormat('ja', {
    numeric: 'always',
  });

  if (diffSec < 0) {
    return '0秒前';
  } else if (diffSec < 60) {
    return rtf.format(-diffSec, 'second');
  } else if (diffSec < 3600) {
    return rtf.format(-Math.floor(diffSec / 60), 'minute');
  } else if (diffSec < 86400) {
    return rtf.format(-Math.floor(diffSec / 3600), 'hour');
  } else if (diffSec < 86400 * 7) {
    return rtf.format(-Math.floor(diffSec / 86400), 'day');
  } else if (diffSec < 86400 * 30) {
    return rtf.format(-Math.floor(diffSec / (86400 * 7)), 'week');
  } else {
    return rtf.format(-Math.floor(diffSec / (86400 * 30)), 'month');
  }
};
