# Module goutils
Functions used in everyday life that help a lot.

## Installation
```shell script
$ go get github.com/armando-couto/goutils
```

## Methods
### arquivo.go
```go
ListFolderFiles(path string) Return []os.FileInfo
ReadingFiles(path string, fileName string) Return string
RemoveFilesOfPath(path string) Return void
ReturnsTheRows(text string) Return []string
CreateDirectory(directory string) Return bool
CreateArchiveNodirectory(nomeDoArquivo string, conteudo string) Return error
IsError(err error) Return bool
Backup(path, name, texto string) Return void
```

### array.go
````go
ContainsToStringInArray(a []string, x string) Return int
````

### aws_sqs.go
```go
ConectionSQS() Return *sqs.SQS
TokenGeneratorMessageId() Return string
```

### base64.go
```go
EncodeStringToBase64(value string) Return string
DecodeBase64ToString(value string) Return string
```

### bool.go
```go
ContainsToStringInArrayReturnBool(a []string, x string) Return bool
```

### cpf_cnpj.go
```go
ValidationCPF(cpf string) Return error
ValidationCNPJ(cnpj string) Return error
```

### csv.go
```go
CSVGZExport(StructCsv [][]string, nomeArquivo string) Return error
CSVRead(pathFile string) Return [][]string
GeneratorFilesGen(payload []interface{}, f func([][]string) [][]string, nomeArq string) Return error
```

### date.go
```go
ConvertStringToTimeLayoutDDMMYYYY(value string) Return time.Time
ConvertTimeToStringLayoutDD_MM_YYYY(date time.Time) Return string
ConvertStringToTimeLayoutDD_MM_YYYY(value string) Return time.Time
ConvertStringToTimeLayoutDD_MM_YYYY_HH_MM_SS(value string) Return time.Time
ConvertStringDD_MM_YYYY(date string) Return time.Time
ConverTimeToStrinLayoutYYYY_MM_DD(data time.Time) Return string
ConvertStringToTimeLayoutYYYYMMDD(value string) Return time.Time
ConvertStringToTimeLayoutYYYY_MM_DD(value string) Return time.Time
ConverTimeToStrinLayoutYYYYMMDD(data time.Time) Return string
ConvertStringToTimeLayoutYYMMDDHHMMSS(value string) Return time.Time
ConvertStringToTimeLayoutYYYYMMDDHHMMSS(value string) Return time.Time
ConvertStringToTimeLayout_YYYY_MM_DD_HH_MM_SS(date time.Time) Return string
ConvertStringToTimeLayout_YYYY_MM_DD(date time.Time) Return string
ConvertStringToTimeLayoutYYYY_MM_DDTHH_MM_SS_000Z(value string) Return time.Time
ConvertStringToTimeLayoutYYYY_MM_DDTHH_MM_SS_000(value string) Return time.Time
ConvertStringToTimeLayoutHHMMSS(value string) Return time.Time
ConvertStringToTimeLayoutHH_MM_SS(value string) Return time.Time
ConvertStringToTimeLayoutDDMMYYYYHHMMSS(d time.Time, h time.Time) Return time.Time
RangeDate(end, start time.Time) func() Return time.Time
DatePlusTime(date, timeOfDay time.Time) Returns (time.Time, error)
```

### environment.go
```go
Godotenv(key string) Return string
```

### float.go
```go
Subtract(valo1, valor2 float64) Return float64
ConvertFloatToFloatScale2(valor float64) Return float64
ConvertFloat64ToString(value float64) Return string
ConvertStringToFloat64(value string) Return float64
ConvertStringToFloatScale2Comma(value string) Return float64
ConvertStringToFloatScale2(value string) Return float64
```

### ftp.go
```go
ListOfAvailableFilesNoFTP(remote string, port string, user string, pass string) Returns (*ftp.ServerConn, error, []string)
GetHostKey(host string, isProduction bool) Return ssh.PublicKey
```

### integer.go
```go
ContainsInt(a []int, x int) Return bool
ConvertStringToInt5Digits(value string) Return int
ConvertStringToInt(value string) Return int
ConvertIntToString(value int) Return string
```

### logger_file.go
```go
CreateFileDay(message string) Return void
```

### mysql_connection.go
```go
ConnectionBDMySQL() Return *sql.DB
```

### psql_connection.go
```go
ConnectionBDPostgreSQL() Return *sql.DB
ConnectionBDPostgreSQLORM() Return (DB *gorm.DB)
```

### string.go
```go
RemoveZerosInLeft(value string) Return string
RemoveSpaceString(value string) Return string
RemoveCharacters(value string) Return string
StandardizesMasksByCard(numeroCartao string) Return string
StringTrim(value string) Return string
RemoveSpecialCharacters(value string) Return string
RandSeq(n int) Return string
TokenGeneratorNLength(length int) Return string
TokenGeneratorOrderReferenceId() Return string
RemoveHeadHyphen(s string) Return string
ValidateIfNotEmptyNumber(valor string) Return string
ValidateIfNotEmptyDate(data string) Return string
RemoveCNPJMask(cnpj string) Return string
ParseBinToHex(s string) Return string
```

### struct.go
```go
CheckIfIdIfNotZero(objetoId int) Return interface{}
```

### try_catch.go
```go
Throw(up Exception) Return void
(tcf Block) Do() Return void
```

### validacoes.go
```go
MaskCard6Initials(cardString string) Return string
MaskLastDigits(card string) Return string
MaskCard(cardString string) Return string
ExpiryDate(valor string) Returns (string, string)
ValidateTelephone(telefone string) Return string
```

