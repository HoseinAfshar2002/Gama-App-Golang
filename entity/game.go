package entity

import "time"

/*
تفاوت entity پلیر و یوزر این است
که پلیر همان یوزر است که در یک گیم خاص بازی می کند
،مثلا یوزر ای دی شماره 8 در گیم شماره 47 بازی می کند
با یک پلیر ای دی مشخص مثلا 210
*/

// Game ساخت استراکت و مدل سازی بخش بازی
type Game struct {
	ID          uint      // ای دی گیم
	CategoryID  uint      // ای دی دسته بندی ان گیم
	QuestionIDs []uint    // یک لیست از ای دی سوالات
	PlayersIDs  []uint    // یک لیست از ای ذی پلیر های بازی
	StartTime   time.Time //تایم شروع بازی
}

// Player ساخت یک استراکت و مدل سازی بازیکنان
type Player struct {
	ID      uint
	UserID  uint           // ای دی بازیکن
	GamaId  uint           // ای دی گیم
	Score   uint           // امتیاز کاربر در گیم
	Answers []PlayerAnswer // یک لیست پاسخ های کاربر که در بازی داده است
}

// PlayerAnswer ساخت استراکت پاسخ های کاربر
type PlayerAnswer struct {
	Id         uint
	PlayerId   uint                  // ای دی پلیر در ان گیم که شرکت کرده
	QuestionID uint                  // ای دی سوالی که پلیر به ان پاسخ می دهد
	Choice     PossibleAnswersChoice // پاسخ انتخابی که پلیر ان را انتخاب می کند
}
