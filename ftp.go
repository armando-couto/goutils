package goutils

import (
	"bufio"
	"github.com/jlaffaye/ftp"
	"golang.org/x/crypto/ssh"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func ListaDeArquivosDisponiveisNoFTP(remote string, port string, user string, pass string) (*ftp.ServerConn, error, []string) {
	c, err := ftp.Dial(remote+":"+port, ftp.DialWithTimeout(60*time.Second))
	if err != nil {
		CreateFileDay(err.Error())
	}
	err = c.Login(user, pass)
	if err != nil {
		CreateFileDay(err.Error())
	}

	//var nome_arquivo string
	list, _ := c.NameList("/")
	if err := c.Quit(); err != nil {
		CreateFileDay(err.Error())
	}
	return c, err, list
}

// parse OpenSSH known_hosts file
// ssh or use ssh-keyscan to get initial key
func GetHostKey(host string, isProduction bool) ssh.PublicKey {
	var file *os.File
	var err error

	if isProduction {
		file, err = os.Open(filepath.Join("/", "home", "ubuntu", ".ssh", "known_hosts")) // Busca na pasta do usuario ubuntu
	} else {
		file, err = os.Open(filepath.Join(os.Getenv("HOME"), ".ssh", "known_hosts")) // Não busca no root, quando o crontab executa busca no root
	}

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var hostKey ssh.PublicKey
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), " ")
		if len(fields) != 3 {
			continue
		}
		if strings.Contains(fields[0], host) {
			var err error
			hostKey, _, _, _, err = ssh.ParseAuthorizedKey(scanner.Bytes())
			if err != nil {
				log.Fatalf("error parsing %q: %v", fields[2], err)
			}
			break
		}
	}

	if hostKey == nil {
		log.Fatalf("no hostkey found for %s", host)
	}

	return hostKey
}
