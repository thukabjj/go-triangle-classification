package entity

type NotAuthorizedError struct{}

func (m *NotAuthorizedError) Error() string {
	return "User not authorized!"
}
