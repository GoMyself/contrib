package helper

import (
	"encoding/csv"
	"errors"
	"os"
)

func WriteCsv(data [][]string) error {

	//创建文件
	f, err := os.Create("test.csv")
	if err != nil {
		return errors.New("created csv file failed")
	}
	defer f.Close()

	// 写入UTF-8 BOM
	_, err = f.WriteString("\xEF\xBB\xBF")
	if err != nil {
		return err
	}

	//创建一个新的写入文件流
	w := csv.NewWriter(f)
	//写入数据
	err = w.WriteAll(data)
	if err != nil {
		return err
	}

	w.Flush()

	return nil
}

func ReadCsv(fileName string) ([][]string, error) {

	var data [][]string

	// 读取文件
	fs, err := os.Open(fileName)
	if err != nil {
		return data, err
	}
	defer fs.Close()

	r := csv.NewReader(fs)
	content, err := r.ReadAll()
	if err != nil {
		return data, err
	}

	for _, v := range content {
		data = append(data, v)
	}

	return data, nil
}
