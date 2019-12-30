package compass

type Server interface {
	Run(args ...interface{}) error
}
