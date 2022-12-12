package feedback

type FeedbackCore struct {
	ID         uint
	Rating     string
	Feedback   string
	UserId     uint
	UserName   string
	HomestayId uint
}

type ServiceInterface interface {
	CreateFeedback(input FeedbackCore) (err error)
}

type RepositoryInterface interface {
	CreateFeedback(input FeedbackCore) (err error)
}
