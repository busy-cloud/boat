package store

import (
	"embed"
	"io/fs"
	"net/http"
	"os"
	"path"
	"sync"
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
	Items []*Item
	lock  sync.RWMutex
}

func (s *Store) Dir(dir string, base string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.Items = append(s.Items, &Item{fs: Dir(dir), base: base})
}

func (s *Store) Zip(zip string, base string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.Items = append(s.Items, &Item{fs: &ZipFS{Filename: zip}, base: base})
}

func (s *Store) EmbedFS(fs embed.FS, base string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.Items = append(s.Items, &Item{fs: fs, base: base})
}

func (s *Store) Open(name string) (http.File, error) {
	file, err := s.OpenFile(name)
	if err != nil {
		return nil, err
	}
	return HttpFile(file), err
}

func (s *Store) OpenFile(name string) (file fs.File, err error) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	//低效
	//for _, item := range s.Items {
	//逆序，优先用后来者
	for i := len(s.Items) - 1; i >= 0; i-- {
		item := s.Items[i]

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
		//return nil, os.ErrNotExist
	}
	return nil, os.ErrNotExist
}
