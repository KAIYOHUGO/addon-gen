package main

type Config struct {
	Addonpy *Addonpy `json:"addon_py" default:"{}"`
}
type Addonpy struct {
	Debug      bool     `json:"debug" default:"false"`
	OutputPath string   `json:"outputpath" default:"../dest"`
	OmitList   []string `json:"omit_file_extension" default:"[\".py\",\".pyc\",\"__pycache__\",\"__jscache__\"]"`
	EntryPoint string   `json:"entrypoint" default:"main.py"`
}
