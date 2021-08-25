package main

const (
	jscache = "__jscache__"
)

type Config struct {
	Addonpy *Addonpy `json:"addon_py" default:"{}"`
}
type Addonpy struct {
	OutputPath string   `json:"outputpath" default:"../dest"`
	OmitList   []string `json:"omit_file_extension" default:"[\".py\",\".pyc\",\"__pycache__\",\"__jscache__\"]"`
	EntryPoint string   `json:"entrypoint" default:"main.py"`
}

func has(array []string, key ...string) bool {
	for _, a := range array {
		for _, k := range key {
			if a == k {
				return true
			}
		}
	}
	return false
}
