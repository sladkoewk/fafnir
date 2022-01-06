package boltdb

type Bucket string

const (
	Email Bucket = "email"
	Sheet Bucket = "sheet"
)

type AppStorage interface {
	Save(id int64, email string, bucket Bucket) error
	Get(id int64, bucket Bucket) (string, error)
}
