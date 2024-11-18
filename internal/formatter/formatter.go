package formatter

type Formatter interface {
	Format(any) (string, error)
}
