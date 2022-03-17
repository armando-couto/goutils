package goutils

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

/*
	ListFolderFiles o antigo nome era: ListarArquivosDaPasta
*/
func ListFolderFiles(path string) []os.FileInfo {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		CreateFileDayError(FormatMessage(MessageError{File: "ListFolderFiles(path)", Error: fmt.Sprint("a pasta '", path, "' não existe!")}))
	}
	return files
}

/*
	ReadingFiles o antigo nome era: LeituraDosArquivos
*/
func ReadingFiles(path string, fileName string) string {
	file, err := os.Open(fmt.Sprint(path, fileName))
	if err != nil {
		log.Fatal(err)
		CreateFileDayError(FormatMessage(MessageError{File: "ReadingFiles(path, fileName)", Error: fmt.Sprint("O arquivo '", fileName, "' não está mais na pasta!")}))
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)
	return string(b)
}

/*
	RemoveFilesOfPath o antigo nome era: RemoveFile
*/
func RemoveFilesOfPath(path string) {
	if os.Remove(path) != nil {
		CreateFileDayError(FormatMessage(MessageError{Error: "Problema ao remover do diretório local"}))
	} else {
		CreateFileDayError(FormatMessage(MessageError{Error: "Removido do local"}))
	}
}

/*
	ReturnsTheRows o antigo nome era: RetornaAsLinhas
*/
func ReturnsTheRows(text string) []string {
	var linhas []string
	scanner := bufio.NewScanner(strings.NewReader(text))
	for scanner.Scan() {
		linhas = append(linhas, scanner.Text())
	}
	return linhas
}

/*
	CreateDirectory o antigo nome era: CriaDiretorio
*/
func CreateDirectory(directory string) bool {
	err := os.MkdirAll(directory, os.ModePerm)
	if IsError(err) {
		return false
	}
	return true
}

/*
	CreateArchiveNodirectory o antigo nome era: CreateArchiveNodirectory
*/
func CreateArchiveNodirectory(nomeDoArquivo string, conteudo string) error {
	return ioutil.WriteFile(nomeDoArquivo, []byte(conteudo), 0777)
}

/*
	IsError o antigo nome era: IsError
*/
func IsError(err error) bool {
	if err != nil {
		CreateFileDayError(FormatMessage(MessageError{Error: err.Error()}))
	}
	return err != nil
}

/*
	Backup
*/
func Backup(path, name, texto string) {
	path_novo := fmt.Sprint(path, "BACKUP/")
	CreateDirectory(path_novo)

	path_nome_arquivo_novo := fmt.Sprint(path, "BACKUP/", name)
	CreateArchiveNodirectory(path_nome_arquivo_novo, texto)

	path_nome_arquivo_antigo := fmt.Sprint(path, name)
	os.Remove(path_nome_arquivo_antigo)
	CreateFileDayInfo(fmt.Sprint("Concluída a Leitura do Arquivo: ", name))
}
