package main

import (
	"runtime"
	"path/filepath"
	"os"
)

type Env struct{
	moduleName string
	rootDir string
	assetsRootDir string
}

func NewEnv() *Env {
	rootDir, ok := getRootDir()
	if !ok {
		return nil
	}

	assetsRootDir, ok := getSubDir(rootDir, "assets")
	if !ok {
		return nil
	}

	env := &Env{}
	env.rootDir = rootDir
	env.assetsRootDir = assetsRootDir
	return env
}

func getRootDir() (string, bool) {
	//curFilePath is: /path/to/go-wxgigo/env.go
	_, curFilePath, _, ok := runtime.Caller(0)
	if !ok {
		return "", false
	}

	return filepath.Dir(curFilePath), true
}

func getSubDir(rootDir string, dirname string) (string, bool) {
	subdir := filepath.Join(rootDir, dirname)
	stat, err := os.Stat(subdir)
	if err != nil {
		return "", false
	}
	if !stat.IsDir() {
		return "", false
	}

	return subdir, true
}


func (self *Env) getAssetPath(assetName string) string {
	extension := filepath.Ext(assetName)
	switch extension {
	case ".html",".htm":
		return filepath.Join(self.assetsRootDir, self.moduleName, "html", assetName)
	case ".js":
		return filepath.Join(self.assetsRootDir, self.moduleName, "js", assetName)
	case ".css":
		return filepath.Join(self.assetsRootDir, self.moduleName, "css", assetName)
	case ".yaml",".yml":
		return filepath.Join(self.assetsRootDir, self.moduleName, "yaml", assetName)
	default:
		return ""
	}
}

