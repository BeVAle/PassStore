package command

var Command interface {
	Run(arguments []string)
}
