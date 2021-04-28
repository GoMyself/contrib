package helper

import (
	"bytes"
	"encoding/csv"
	"os"
)

func WriteCsv(data [][]string) (*bytes.Buffer, error) {

	b := new(bytes.Buffer)

	// 写入UTF-8 BOM
	_, err := b.WriteString("\xEF\xBB\xBF")
	if err != nil {
		return b, err
	}

	//创建一个新的写入文件流
	w := csv.NewWriter(b)
	//写入数据
	err = w.WriteAll(data)
	if err != nil {
		return b, err
	}

	w.Flush()

	return b, nil
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
