package tools

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

func FilenameWithCurrentTime(filename string) string {
	baseFilename := FilenameWithoutExtension(filename)
	currentTime := time.Now().Format("2006_Jan_02_15_04_05")
	extension := filepath.Ext(filename)
	return fmt.Sprintf("%s-%s%s", baseFilename, currentTime, extension)
}

func FilenameWithoutExtension(filename string) string {
	return strings.TrimSuffix(filename, filepath.Ext(filename))
}

func FilepathFromHome(path string) string {
	return filepath.Join(HomePath(), path)
}

func HomePath() string {
	switch runtime.GOOS {
	case "windows":
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	case "linux":
		home := os.Getenv("XDG_CONFIG_HOME")
		if home != "" {
			return home
		}
	}
	return os.Getenv("HOME")
}

func WriteToBuff(r io.Reader) (*bytes.Buffer, error) {
	buff := bytes.Buffer{}
	if _, err := io.Copy(&buff, r); err != nil {
		return nil, err
	}
	return &buff, nil
}

func SliceToUpper(s []string) []string {
	for i := range s {
		s[i] = strings.ToUpper(s[i])
	}
	return s
}
