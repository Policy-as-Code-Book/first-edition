package utils

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"strconv"
	"strings"
	"time"
)

func RemoveFile(path string) bool {
	e := os.Remove(path)
	if e != nil {
		fmt.Errorf("could not remove %s: %w", path, e)
		return false
	}

	return true
}

func WriteFile(path string, bytes []byte) error {
	var i int
	if _, err := os.Stat(path); err == nil {
		f, e := os.Open(path)
		defer f.Close()
		if e == nil {
			i, e = f.Write(bytes)
			if e != nil {
				return e
			}

			fmt.Println(fmt.Sprintf("wrote %d bytes to %s", i, path))
		}
	} else if errors.Is(err, os.ErrNotExist) {
		f, e := os.Create(path)
		defer f.Close()
		if e == nil {
			i, e = f.Write(bytes)
		}
		if e != nil {
			return e
		}
		fmt.Sprintf("wrote %d bytes to %s", i, path)
	} else {
		return err
	}

	return nil
}

func ReadFile(path string) ([]byte, error) {
	var bytes []byte
	var err error
	var fi fs.FileInfo

	fi, err = os.Stat(path)
	if fi != nil && err == nil {
		fmt.Println(fmt.Sprintf("File info: %+v", fi))
	}

	if err != nil {
		fmt.Println(fmt.Sprintf("File info issues: %+v", err))
	}

	// Read file
	bytes, err = os.ReadFile(path)
	if err != nil {
		return bytes, err
	}

	return bytes, nil
}

func FileExists(fileName string) bool {
	_, err := os.Stat(fileName)

	// check if error is "file not exists"
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func FileSearch(filePath string, search string) ([]int, error) {
	var lines = make([]int, 0)

	f, err := os.Open(filePath)
	if err != nil {
		return lines, err
	}
	defer f.Close()

	// Splits on newlines by default.
	scanner := bufio.NewScanner(f)

	line := 1
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), search) {
			lines = append(lines, line)
		}
		line++
	}

	return lines, nil
}

func TsString() string {
	now := time.Now() // current local time
	sec := now.Unix()

	return strconv.FormatInt(sec, 10)
}

func SplitParse(name string, splitter string, pos int) string {
	if strings.Contains(name, splitter) {
		return strings.Split(name, splitter)[pos]
	}

	return name
}

func MD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func SearchArray(a []string, s string) bool {
	for _, v := range a {
		if s == v {
			return true
		}
	}

	return false
}
