package Config

import (
	//"encoding/json"
	"fmt"
	//"io/ioutil"
	//"log"
	"strings"
	//"reflect"
	"bufio"
	"os"
	//"github.com/spf13/viper"
)

type Config struct {
	Key1 string
	Key2 string
	Key3 Key
}

type Key struct {
	key4 int
	key5 string
}

func SearchKey(filename string, key string) string {
	readFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	readFile.Close()

	for _, line := range fileLines {
		res := strings.Split(line, ":")
		value := res[0]
		if len(value) > 1 {
			storedkey := value[1 : len(value)-1]
			if storedkey == key {
				return line
			}
		}

	}
	return ""
}

// func main() {
// 	line := SearchKey("app.env", "Key1")
// 	fmt.Println(line)
// }
