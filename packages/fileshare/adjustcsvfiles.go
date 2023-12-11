package fileshare

import (
	"encoding/csv"
	"os"
	"sort"
)

func WriteCSVFile(path string, expenseMap map[string]float64) (err error) {
	//Create a sorted slice of the headers
	headers := []string{}
		for key := range expenseMap{
			headers = append(headers, key)
		}
		sort.Strings(headers)
	
	
	file, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil{
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	fileInfo, err := os.Stat(path)
	if err != nil{
		return err
	}

	if fileInfo.Size() == 0 {
		writer.Write(headers)
	}

	something := []string{"hello", "how", "are", "you", "today?"}
	writer.Write(something)
	return nil
}

