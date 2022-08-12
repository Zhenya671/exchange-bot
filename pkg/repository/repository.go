package repository

type Bucket string

const (
	UserData Bucket = "user_data"
)

type DataUsers interface {
	Save(chatID int64, firstName string, bucket Bucket) error
	Get(chatID int64, bucket Bucket) (string, error)
}
