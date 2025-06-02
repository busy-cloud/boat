package store

import (
	"embed"
	"github.com/busy-cloud/boat/lib"
	"io/fs"
	"net/http"
	"os"
	"path"
	"strings"
)

type Item struct {
	fs   FS
	base string
}

func (s *Item) Open(name string) (fs.File, error) {
	return s.fs.Open(path.Join(s.base, name))
}

func (s *Item) ReadDir(name string) ([]fs.DirEntry, error) {
	return s.fs.ReadDir(path.Join(s.base, name))
}

func (s *Item) ReadFile(name string) ([]byte, error) {
	return s.fs.ReadFile(path.Join(s.base, name))
}

type Store struct {
	items lib.Map[Item]
}

func (s *Store) Exists(key string) bool {
	return s.items.Load(key) != nil
}

func (s *Store) Dir(key string, root string, base string) {
	s.items.Store(key, &Item{fs: Dir(root), base: base})
}

func (s *Store) Zip(key string, zip string, base string) {
	s.items.Store(key, &Item{fs: &ZipFS{Filename: zip}, base: base})
}

func (s *Store) EmbedFS(key string, fs embed.FS, base string) {
	s.items.Store(key, &Item{fs: fs, base: base})
}

func (s *Store) Open(name string) (http.File, error) {
	if name[0:1] == "/" {
		name = name[1:]
	}

	index := strings.Index(name, "/")
	if index == -1 {
		return nil, os.ErrNotExist
	}

	//分割，取key
	key := name[:index]
	name = name[index+1:]
	file, err := s.OpenFile(key, name)
	if err != nil {
		return nil, err
	}
	return HttpFile(file), err
}

func (s *Store) OpenFile(key, name string) (file fs.File, err error) {
	item := s.items.Load(key)
	if item == nil {
		return nil, os.ErrNotExist
	}

	fn := path.Join(item.base, name)

	//查找文件
	file, err = item.fs.Open(fn)
	if err == nil {
		fi, e := file.Stat()
		if e != nil {
			return nil, e
		}
		if fi != nil && !fi.IsDir() {
			return
		}
	}

	return nil, os.ErrNotExist
}
