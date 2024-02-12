package exception

type NotFoundError struct {
	NotFoundDetails string
}

func (err NotFoundError) Error() string {
	return err.NotFoundDetails
}
