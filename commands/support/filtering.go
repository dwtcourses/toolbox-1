package support

func GetFilter(args []string) string {
	filter := ""
	if len(args) > 0 {
		filter = args[0]
	}
	return filter
}