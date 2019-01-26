package dream

type flags struct {
	ConfigPath string `json:"c"`
}

func ParseFlags(args []string) *flags {

	flags := args[1:]
	flagMap := make(map[string]string)
	for i, flag := range flags {
		if isFlag(flag) {
			if isFlag(flag) {
				flagMap[flag] = ""
			} else {
				flagMap[flag] = args[i+1]
			}
		}
	}
		
}

func isFlag(s string) bool {
	prefix := s[0:1]
	if prefix == "-" || prefix == "--" {
		return true
	}
	return false
}
