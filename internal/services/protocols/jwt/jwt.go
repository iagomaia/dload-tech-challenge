package jwtprotocols

type IJwt interface {
	Generate(userId string, claims map[string]any) (string, error)
	Verify(token string) (map[string]any, error)
}
