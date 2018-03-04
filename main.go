package main

import "os"
import "fmt"
import "io/ioutil"
import "log"
import "path"

func processDir(root string, dir string, sizes map[string]int64) {
	files, err := ioutil.ReadDir(dir)

	if err != nil {
		log.Println(err)
		return
	}

	for _, f := range files {
		p := path.Join(dir, f.Name())
		if f.IsDir() {
			processDir(root, p, sizes)
		} else {
			subpathsAddSize(root, sizes, p, f.Size())
		}
	}

}

func subpathsAddSize(dir string, sizes map[string]int64, p string, size int64) {
	for ; p != dir ; p = path.Join(p, "..") {
		sizes[p] += size
	}
}

// http://yourbasic.org/golang/formatting-byte-size-to-human-readable-format/
func ByteCountBinary(b int64) string {
    const unit = 1000
    if b < unit {
        return fmt.Sprintf("%dB", b)
    }
    div, exp := int64(unit), 0
    for n := b / unit; n >= unit; n /= unit {
        div *= unit
        exp++
    }
    return fmt.Sprintf("%.1f%cB",
        float64(b)/float64(div), "KMGTPE"[exp])
}


func displayMap(sizes map[string]int64) {
	for k, v := range sizes { 
		fmt.Printf("%s: %s\n", k, ByteCountBinary(v))
	}
}

func main() {
	if (len(os.Args) != 2) {
		return
	}

	dir := os.Args[1]
	dir = path.Join(dir, ".")

	var sizes map[string]int64 = make(map[string]int64)
	processDir(dir, dir, sizes)
	displayMap(sizes)
}
