package goutils

import (
	"bytes"
	"compress/gzip"
	"encoding/csv"
	"errors"
	"fmt"
	"github.com/fatih/structs"
	"os"
)

/*
	CSVGZExport o antigo nome era: CsvGzExport
*/
func CSVGZExport(StructCsv [][]string, nomeArquivo string) error {
	if len(StructCsv) < 1 {
		return errors.New("ERRO: Dados vazios")
	}
	var buf bytes.Buffer

	file, err := os.Create(nomeArquivo + ".csv.gz")
	if err != nil {
		return err
	}
	defer file.Close()
	wgzp := gzip.NewWriter(file)
	wgzp.Write(buf.Bytes())

	for _, v := range StructCsv {
		var a string
		for j, x := range v {
			a += x
			if j < len(v)-1 {
				a += ";"
			}
		}

		wgzp.Write([]byte(string(a + "\n")))
		defer wgzp.Close()
	}

	CreateFileDay(Message{Info: fmt.Sprint("Arquivo ", nomeArquivo, " gerado CSV.GZ")})
	return nil
}

/*
	CSVRead o antigo nome era: CsvRead
*/
func CSVRead(pathFile string) [][]string {
	f, err := os.Open(pathFile)
	if err != nil {
		CreateFileDay(Message{Error: "ERRO: Leitura arquivo CSV:" + err.Error()})
	}
	defer f.Close()

	arqReader := csv.NewReader(f)
	arqReader.LazyQuotes = true
	lines, err := arqReader.ReadAll()
	if err != nil {
		CreateFileDay(Message{Error: "ERRO: Leitura arquivo CSV:" + err.Error()})
	}
	return lines
}

/*
	GeneratorFilesGen o antigo nome era: GeradorArquivosGen
*/
func GeneratorFilesGen(payload []interface{}, f func([][]string) [][]string, nomeArq string) error {
	structToCsv := make([][]string, 0)
	if len(payload) < 1 {
		return errors.New("Arquivo vazio !!! ---->" + nomeArq)
	}

	for _, a := range payload {
		b := make([]string, 0)

		for _, v := range structs.Values(a) {
			if _, ok := v.(string); ok {
				b = append(b, fmt.Sprintf("%v", v))
			}
		}
		structToCsv = append(structToCsv, b)
	}
	if f != nil {
		structToCsv = f(structToCsv)
	}

	err := CSVGZExport(structToCsv, nomeArq)
	return err
}
