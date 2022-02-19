package models

import "time"

// 有効期限切れかどうかを調べます。
// 期限切れの場合、trueが返る
func CheckExpired(entry *Period) bool {
	now := time.Now()
	var periodTime time.Time

	// 有効期限の日時
	if entry.PeriodMinute != 0 {
		periodTime = entry.CreateDate.Add(time.Duration(entry.PeriodMinute) * time.Minute)
	} else if entry.PeriodHour != 0 {
		periodTime = entry.CreateDate.Add(time.Duration(entry.PeriodHour) * time.Hour)
	} else if entry.PeriodDay != 0 {
		periodTime = entry.CreateDate.Add(time.Duration(entry.PeriodDay*24) * time.Hour)
	} else {
		// hour, minuteどちらも定義されていない場合、createDateで比較する = かならずtrueが返る
		periodTime = entry.CreateDate
	}

	// 有効期限より前に今の時間がある場合Falseを返す
	return now.After(periodTime)
}
