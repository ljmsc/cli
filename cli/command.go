package cli

type Command interface {
	Run(param map[string]string) int
}
