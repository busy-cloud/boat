package web

import (
	"embed"
	"errors"
	"io/fs"
	"net/http"
	"path"
	"strings"
)

type item struct {
	fs    http.FileSystem
	path  string
	base  string
	index string
}

var items []*item

// Static 静态目录
// path url路径。
// base 基本路径（主要用于zip文件）。
// index 默认首页，index.html 代表SPA应用。
func Static(fs http.FileSystem, path, base, index string) {
	items = append(items, &item{fs: fs, path: path, base: base, index: index})
}

func StaticFS(fs fs.FS, path, base, index string) {
	items = append(items, &item{fs: http.FS(fs), path: path, base: base, index: index})
}

func StaticDir(dir string, path, base, index string) {
	items = append(items, &item{fs: http.Dir(dir), path: path, base: base, index: index})
}

func StaticZip(zip string, path, base, index string) {
	items = append(items, &item{fs: &zipFS{filename: zip}, path: path, base: base, index: index})
}

func StaticEmbedFS(fs embed.FS, path, base, index string) {
	items = append(items, &item{fs: http.FS(fs), path: path, base: base, index: index})
}

func OpenStaticFile(name string) (file http.File, err error) {
	//低效
	for _, f := range items {
		//fn := path.Join(fbase, name)
		// && !strings.HasPrefix(name, "/$")
		if f.path == "" || f.path != "" && strings.HasPrefix(name, f.path) {
			//去除前缀
			fn := path.Join(f.base, strings.TrimPrefix(name, f.path))

			//查找文件
			file, err = f.fs.Open(fn)
			if file != nil {
				fi, _ := file.Stat()
				if !fi.IsDir() {
					return
				}
			}

			//尝试默认页
			if f.index != "" {
				file, err = f.fs.Open(path.Join(f.base, f.index))
				if file != nil {
					fi, _ := file.Stat()
					if !fi.IsDir() {
						return
					}
				}
			}

			return nil, errors.New("not found")
		}
	}
	return nil, errors.New("not found")
}
