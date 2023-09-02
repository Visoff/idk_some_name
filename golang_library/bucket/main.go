package bucket

import "os"

func Write(name string, content []byte) error {
	return os.WriteFile("bucket/static/"+name, content, 0666)
}

func Read(name string) ([]byte, error) {
	return os.ReadFile("bucket/static/" + name)
}
