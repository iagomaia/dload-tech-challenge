package hashprotocols

type IHash interface {
	MakeHash(input string) (string, error)
	CompareHash(source, target string) error
}
