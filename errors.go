package go_jwt

type ParseError struct{}

func (e *ParseError) Error() string {
	return "failed to parse jwt"
}

type InvalidError struct{}

func (e *InvalidError) Error() string {
	return "invalid jwt"
}

type ExpError struct{}

func (e *ExpError) Error() string {
	return "expired jwt"
}

type NbfError struct{}

func (e *NbfError) Error() string {
	return "not activated jwt"
}
