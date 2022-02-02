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
