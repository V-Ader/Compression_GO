package file

import "os"

func Save(filePath string, data string) error {
	file, err := os.Create(filePath)
	check(err)
	defer file.Close()

	_, err = file.Write([]byte(data))
	check(err)

	return nil
}

func Load(filename string, binary bool) []byte {
	dat, err := os.ReadFile(filename)
	check(err)
	return dat
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
