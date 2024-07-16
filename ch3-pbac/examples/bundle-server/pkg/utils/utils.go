package utils

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/rs/zerolog/log"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"time"
)

// Namespacing functions with structs, overcoming the lack of overloading

type strng struct{}
type intgr struct{}

func (z strng) ArrayContains(a []string, s string) bool {
	for _, x := range a {
		if x == s {
			return true
		}
	}
	return false
}

var Strng strng

func (z intgr) ArrayContains(a []int, i int) bool {
	for _, x := range a {
		if x == i {
			return true
		}
	}
	return false
}

var Intgr intgr

func test() {
	var sa []string
	sa = append(sa, "Hello")
	sa = append(sa, "World")

	r := Strng.ArrayContains(sa, "Hello")
	fmt.Println(r)

	var ia []int
	ia = append(ia, 0)
	ia = append(ia, 1)
	r = Intgr.ArrayContains(ia, 0)
	fmt.Println(r)
}

func ArrayContains(a []string, s string) bool {
	for _, x := range a {
		if x == s {
			return true
		}
	}
	return false
}

func MD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func FileInfos(path string) (map[string]os.FileInfo, error) {
	m := make(map[string]os.FileInfo)
	err := filepath.Walk(path,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			m[path] = info
			return nil
		})
	if err != nil {
		return m, err
	}

	return m, nil
}

func TsString() string {
	now := time.Now() // current local time
	sec := now.Unix()

	return strconv.FormatInt(sec, 10)
}

func RemoveFiles(dir string) error {
	f, err := os.ReadDir(dir)
	if err != nil {
		return err
	}
	for _, d := range f {
		err = os.RemoveAll(path.Join([]string{dir, d.Name()}...))
		Logger.Error().Err(err).Msgf("could not delete %s: ", d.Name())
	}

	return nil
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

			Logger.Debug().Msgf("wrote %d bytes to %s", i, path)
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

		Logger.Debug().Msgf("wrote %d bytes to %s", i, path)
	} else {
		return err
	}

	return nil
}

func ReadFile(path string) ([]byte, error) {
	var bytes []byte
	var err error
	var fi fs.FileInfo

	fi, err = os.Stat(path) //};  errors.Is(err, os.ErrNotExist) {
	if fi != nil && err == nil {
		log.Debug().Msgf("File info: %+v", fi)
	} else {
		log.Error().Err(err).Msg("File info issues: %+v")
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
	} else {
		return true
	}
}
