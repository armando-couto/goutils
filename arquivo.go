package goutils

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func ListarArquivosDaPasta(path string) []os.FileInfo {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		CreateFileDay(fmt.Sprint("a pasta '", path, "' não existe!"))
	}
	return files
}

func LeituraDosArquivos(path string, fileName string) string {
	file, err := os.Open(fmt.Sprint(path, fileName))
	if err != nil {
		log.Fatal(err)
		CreateFileDay(fmt.Sprint("O arquivo '", fileName, "' não está mais na pasta!"))
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)
	return string(b)
}

func RetornaAsLinhas(text string) []string {
	var linhas []string
	scanner := bufio.NewScanner(strings.NewReader(text))
	for scanner.Scan() {
		linhas = append(linhas, scanner.Text())
	}
	return linhas
}

func CriaDiretorio(directory string) bool {
	err := os.MkdirAll(directory, os.ModePerm)
	if isError(err) {
		return false
	}
	return true
}

func CriaArquivoNodiretorio(nomeDoArquivo string, conteudo string) error {
	return ioutil.WriteFile(nomeDoArquivo, []byte(conteudo), 0777)
}

func ExcluiArquivo(diretorio string) error {
	return os.Remove(diretorio)
}

func isError(err error) bool {
	if err != nil {
		CreateFileDay(err.Error())
	}
	return err != nil
}

func Backup(path, name, texto string) {
	path_novo := fmt.Sprint(path, "BACKUP/")
	CriaDiretorio(path_novo)

	path_nome_arquivo_novo := fmt.Sprint(path, "BACKUP/", name)
	CriaArquivoNodiretorio(path_nome_arquivo_novo, texto)

	path_nome_arquivo_antigo := fmt.Sprint(path, name)
	ExcluiArquivo(path_nome_arquivo_antigo)
	CreateFileDay(fmt.Sprint("Concluída a Leitura do Arquivo: ", name))
}
