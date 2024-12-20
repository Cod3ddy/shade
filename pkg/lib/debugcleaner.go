package lib

import (
	"bufio"
	"os"
	"regexp"
)

func CleanFile(path string) error {

	f, err := os.Open(path)
	if err != nil{
		return err
	}

	defer f.Close()

	var lines []string
	scnr := bufio.NewScanner(f)

	reg := regexp.MustCompile(`fmt\.Printf\(`)

	for scnr.Scan(){
		line := scnr.Text()
		if !reg.MatchString(line){
			lines = append(lines, line)
		}
	}

	if err := scnr.Err(); err != nil{
		return err
	}

	f, err = os.Create(path)
	if err != nil{
		return err 
	}

	defer f.Close()

	writer := bufio.NewWriter(f)

	for _, line := range lines{
		_, err := writer.WriteString(line + "\n")
		if err != nil{
			return err
		}
	}

	writer.Flush()

	return nil
}
