package helper

import (
	"bytes"
	"encoding/csv"
	"io"
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

func ReadCsv(fs io.Reader) ([][]string, error) {

	var data [][]string

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
