//go:generate embedfiles -pkg=example 1.jpg 2.jpg
package example

func LoadImage(name string) []byte {
	switch name {
	case "1.jpg", "2.jpg":
		return files[name]
	}
	return nil
}
