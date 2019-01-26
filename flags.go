package dream

import "encoding/json"

type flags struct {
	ConfigPath string `json:"c"`
}

func ParseFlags(args []string) *flags {

	f := args[1:]
	flagMap := make(map[string]string)
	for i, flag := range f {
		if isFlag(flag) {
			if isFlag(f[i+1]) {
				flagMap[flag] = ""
			} else {
				flagMap[flag] = f[i+1]
			}
		}
	}

	tmp, err := json.Marshal(flagMap)
	if err != nil {
		panic(err)
	}
	result := new(flags)
	err = json.Unmarshal(tmp, result)
	if err != nil {
		panic(err)
	}

	return result
}

func isFlag(s string) bool {
	prefix := s[0:1]
	if prefix == "-" || prefix == "--" {
		return true
	}
	return false
}
