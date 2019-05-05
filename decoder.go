package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	xj "github.com/basgys/goxml2json"
	y "github.com/ghodss/yaml"
)

type yamlDecoder struct {
}

func (i *yamlDecoder) Decode(data []byte, v interface{}) error {
	bb, err := y.YAMLToJSON(data)
	if err != nil {
		return err
	}
	return json.Unmarshal(bb, &v)
}

type xmlDecoder struct {
}

func (i *xmlDecoder) Decode(data []byte, v interface{}) error {
	buf, err := xj.Convert(bytes.NewReader(data))
	if err != nil {
		return err
	}
	return json.Unmarshal(buf.Bytes(), &v)
}

type csvDecoder struct {
}

func (i *csvDecoder) Decode(data []byte, v interface{}) error {
	reader := csv.NewReader(bytes.NewReader(data))
	rr, err := reader.ReadAll()
	if err != nil {
		return errors.New("gojsonq: " + err.Error())
	}
	if len(rr) < 1 {
		return errors.New("gojsonq: csv data can't be empty! At least contain the header row")
	}
	var arr = make([]map[string]interface{}, 0)
	header := rr[0] // assume the very first row as header
	for i := 1; i <= len(rr)-1; i++ {
		if rr[i] == nil { // if a row is empty, skip it
			continue
		}
		mp := map[string]interface{}{}
		for j := 0; j < len(header); j++ {
			// convert data to different types
			// if header contains field like, ID|NUMBER,Name|String,IsStudent|BOOLEAN
			t := strings.Split(header[j], "|")
			var typ string
			if len(t) > 1 {
				typ = t[1]
			}
			hdr := strings.TrimSpace(t[0])
			switch typ {
			default:
				mp[hdr] = rr[i][j]

			case "STRING":
				mp[hdr] = rr[i][j]

			case "NUMBER":
				if fv, err := strconv.ParseFloat(rr[i][j], 64); err == nil {
					mp[hdr] = fv
				} else {
					mp[hdr] = 0.0
				}

			case "BOOLEAN":
				if strings.ToLower(rr[i][j]) == "true" ||
					rr[i][j] == "1" {
					mp[hdr] = true
				} else {
					mp[hdr] = false
				}

			}
		}
		arr = append(arr, mp)
	}
	bb, err := json.Marshal(arr)
	if err != nil {
		return fmt.Errorf("gojsonq: %v", err)
	}
	return json.Unmarshal(bb, &v)
}
