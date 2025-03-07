package web

import (
	"embed"
	"io/fs"
	"net/http"
	"os"
	"path"
)

type storeItem struct {
	fs   http.FileSystem
	base string
}

type Store struct {
	items []*storeItem
}

func (s *Store) Web(fs http.FileSystem, base string) {
	s.items = append(s.items, &storeItem{fs: fs, base: base})
}

func (s *Store) FS(fs fs.FS, base string) {
	s.items = append(s.items, &storeItem{fs: http.FS(fs), base: base})
}

func (s *Store) Dir(dir string, base string) {
	s.items = append(s.items, &storeItem{fs: http.Dir(dir), base: base})
}

func (s *Store) Zip(zip string, base string) {
	s.items = append(s.items, &storeItem{fs: &ZipFS{Filename: zip}, base: base})
}

func (s *Store) EmbedFS(fs embed.FS, base string) {
	s.items = append(s.items, &storeItem{fs: http.FS(fs), base: base})
}

func (s *Store) Open(name string) (file http.File, err error) {
	//低效
	for _, f := range s.items {
		fn := path.Join(f.base, name)

		//查找文件
		file, err = f.fs.Open(fn)
		if file != nil {
			fi, _ := file.Stat()
			if !fi.IsDir() {
				return
			}
		}
		return nil, os.ErrNotExist
	}
	return nil, os.ErrNotExist
}
