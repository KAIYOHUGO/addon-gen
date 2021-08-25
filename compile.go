package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/evanw/esbuild/pkg/api"
)

func compile(cfg *Config, path string) error {
	// check
	p := filepath.Join(path, cfg.Addonpy.EntryPoint)
	{
		if info, err := os.Stat(p); err != nil || info.IsDir() {
			return nil
		}
	}
	// run command
	{
		if o, err := exec.Command("transcrypt", "-p", ".none", "-n", "-od", jscache, p).Output(); err != nil {
			fmt.Println(string(o))
			return err
		}
		r := api.Build(api.BuildOptions{
			Bundle:            true,
			Write:             true,
			MinifyWhitespace:  true,
			MinifyIdentifiers: true,
			MinifySyntax:      true,
			TreeShaking:       api.TreeShakingIgnoreAnnotations,
			EntryPoints:       []string{filepath.Join(path, jscache, strings.TrimSuffix(cfg.Addonpy.EntryPoint, ".py")+".js")},
			Outdir:            filepath.Join(cfg.Addonpy.OutputPath, path),
		})
		if len(r.Errors) > 0 {
			err := errors.New(r.Errors[0].Text)
			for i := 1; i < len(r.Errors); i++ {
				err = fmt.Errorf("%w =>build error: %s", err, r.Errors[i].Text)
			}
			return err
		}
	}
	return nil
}
