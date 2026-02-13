package entity

// Question ساخت استراکت و مدل سازی بخش سوالات
type Question struct {
	ID              uint
	QuestionText    string            // متن سوال
	PossibleAnswers []PossibleAnswers // پاسخ ثبت شده توسط پلیر
	CorrectAnswerID uint              // شناسه پاسخ صحیح
	DifficultyLevel DifficultyLevel   // دریجه سختی سوالات
	CategoryID      uint              // ای دی دسته بندی
}

// PossibleAnswers ساخت استراکت و مدل سازی بخش پاسخ انتخابی کاربر
type PossibleAnswers struct {
	ID                  uint
	PossibleAnswersText string                //متن پاسخ
	Choice              PossibleAnswersChoice // انتخاب پلیر
}

// PossibleAnswersChoice متغیر یا تایپ انتخاب پاسخ احتمالی کاربر
type PossibleAnswersChoice uint8

// IsValid تعریف یک فانکشن برای اعتبار سنجی اینکه کاربر پاسخ ها را به درستی انتخاب کند
func (p PossibleAnswersChoice) IsValid() bool {
	//پاسخ انتخابی کاربر باید بزرگ تر مساوی A و کوچک تر مساوی D باشد یعنی فقط یکی از 4 گزینه A تا D را انتخاب کند تا پاسخ به درستی ثبت شود
	if p >= PossibleAnswerA && p <= PossibleAnswerD {
		//اگر شرط درست بود ان را ولید قرار بده
		return true
	}
	//در غیر این صورت ولید قرار نده
	return false
}

// چند ثابت تعریف می کنیم
const (
	PossibleAnswerA PossibleAnswersChoice = iota + 1 // مقدار iota یا ایوتا مقدار پیش فرض 0 است و ما چون نمیخواهیم پیشفرض 0 باشد ان را + 1 می کنیم تا از 1 شروع شود
	PossibleAnswerB                                  // 2
	PossibleAnswerC                                  //3
	PossibleAnswerD                                  //3
)

// DifficultyLevel متغیر یا تایپ درجه سختی سوالات
type DifficultyLevel uint8

// چند ثابت تعریف می کنیم
const (
	DifficultyLevelEasy   DifficultyLevel = iota + 1 // مقدار iota یا ایوتا مقدار پیش فرض 0 است و ما چون نمیخواهیم پیشفرض 0 باشد ان را + 1 می کنیم تا از 1 شروع شود 1 آسان
	DifficultyLevelMedium                            //2 متوسظ
	DifficultyLevelHard                              //3 سخت
)

// IsValid ساخت یک فانکشن برای اعبار سنجی سختی سوالات
func (q DifficultyLevel) IsValid() bool {
	//اگر پاسخ بزرگ تر مساوی اسان بود و کوچک تر مساوی سخت بود
	if q >= DifficultyLevelEasy && q >= DifficultyLevelHard {
		// درجه سختی ولید است
		return true
	}
	// در غیر اینصورت ولید نمی شود
	return false
}
