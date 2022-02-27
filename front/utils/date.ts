export const formatDate = (date: Date): string => {
  const week = ['日', '月', '火', '水', '木', '金', '土'];

  const year = date.getFullYear();
  const month = date.getMonth() + 1;
  const weekDay = week[date.getDay()];
  const day = date.getDate();
  const hour = date.getHours();
  const minutes = date.getMinutes();

  return `${year}年${month}月${day}日${weekDay}曜日 ${hour}:${(
    '00' + minutes
  ).slice(-2)}`;
};

export const hawManyDaysAgo = (date: Date): string => {
  const now = new Date();
  const diffSec = Math.floor((now.getTime() - date.getTime()) / 1000);

  if (diffSec < 3600) {
    return `${Math.floor(diffSec / 60)}分前`;
  } else if (diffSec < 86400) {
    return `${Math.floor(diffSec / 3600)}時間前`;
  } else if (diffSec < 86400 * 7) {
    return `${Math.floor(diffSec / 86400)}日前`;
  } else if (diffSec < 86400 * 30) {
    return `${Math.floor(diffSec / (86400 * 7))}週間前`;
  } else if (diffSec < 86400 * 90) {
    // 3ヶ月前（90日）まで表示する
    return `${Math.floor(diffSec / (86400 * 30))}ヶ月前`;
  } else {
    return '3ヶ月以上前';
  }
};
