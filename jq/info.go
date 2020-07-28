package jq

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"regexp"
)

type JQInfo struct {
	Path    string
	Version string
}

func GetInfo() (*JQInfo, error) {
	path, err := exec.LookPath("jq")
	if err != nil {
		return nil, err
	}
	_, err = os.Stat(path)
	if err != nil {
		return nil, err
	}

	version, err := getVersion(path)
	if err != nil {
		return nil, err
	}
	return &JQInfo{Path: path, Version: version}, nil
}

func getVersion(path string) (string, error) {
	// get version from `jq --help`
	// since `jq --version` diffs between versions
	// e.g., 1.3 & 1.4
	var b bytes.Buffer
	cmd := exec.Command(path, "--help")
	cmd.Stdout = &b
	cmd.Stderr = &b
	cmd.Run()

	out := bytes.TrimSpace(b.Bytes())
	r := regexp.MustCompile(`\[version (.+)\]`)
	if r.Match(out) {
		m := r.FindSubmatch(out)[1]
		return string(m), nil
	}

	return "", fmt.Errorf("can't find jq version: %s", out)
}
