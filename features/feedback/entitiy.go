package feedback

type FeedbackCore struct {
	ID         uint
	Rating     string
	Feedback   string
	UserId     uint
	UserName   string
	HomestayID uint
}

type Homestay struct {
	ID       uint
	Name     string
	Feedback []FeedbackCore
}

type ServiceInterface interface {
	CreateFeedback(input FeedbackCore) (err error)
}

type RepositoryInterface interface {
	CreateFeedback(input FeedbackCore) (err error)
}
