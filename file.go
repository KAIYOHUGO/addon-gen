package main

import (
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func copyfolder(cfg *Config) error {
	return filepath.Walk("./", func(path string, info fs.FileInfo, err error) error {
		// check
		{
			s := make([]string, 0)
			s = append(s, filepath.Ext(info.Name()))
			s = append(s, strings.Split(path, "/")...)
			if has(cfg.Addonpy.OmitList, s...) {
				return nil
			}
		}
		// create file path
		pc := filepath.Join(cfg.Addonpy.OutputPath, path)
		{
			ic, err := os.Stat(pc)
			if err == nil && ic.IsDir() == info.IsDir() && ic.Name() == info.Name() && ic.ModTime() == info.ModTime() {
				return nil
			}
		}

		if info.IsDir() {
			os.Mkdir(filepath.Join(cfg.Addonpy.OutputPath, path), 0777)
			return nil
		}

		// copy file
		{
			fo, err := os.Open(path)
			if err != nil {
				return err
			}
			defer fo.Close()
			fc, err := os.Create(pc)
			if err != nil {
				return err
			}
			defer fc.Close()
			if _, err := io.Copy(fc, fo); err != nil {
				return err
			}
		}
		return nil
	})
}
