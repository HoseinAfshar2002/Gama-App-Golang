package entity

// User ساخت یک استراکت یا همان مدل سازی  کاربر در گیم اپلیکیشن
type User struct {
	ID          uint   // ای دی یوزر
	PhoneNumber string //شماره تلفن یوزر
	Name        string // نام یوزر
	Password    string
}
