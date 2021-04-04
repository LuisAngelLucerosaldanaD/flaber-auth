package logger

import (
	"flaber-auth/internal/env"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

var (
	Trace             *log.Logger
	Info              *log.Logger
	Warning           *log.Logger
	Error             *log.Logger
	logReviewInterval int
)

// Configura el logger para escritura en disco
func init() {
	c := env.NewConfiguration()

	AppPathLog := c.App.PathLog
	AppServiceName := c.App.ServiceName
	logReviewInterval = c.App.LogReviewInterval

	// crea la ruta donde se almacenarán los log
	path := fmt.Sprintf("%s/%s", AppPathLog, AppServiceName)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// Se crea folder para almacenar el log
		if err = os.MkdirAll(path, 0777); err != nil {
			log.Fatalf("no se puede crear la ruta de los log: %v", err)
		}
	}

	fileTrace, err := os.OpenFile(path+"/trace.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening fileTrace file: %v", err)
	}

	fileInfo, err := os.OpenFile(path+"/info.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening fileInfo file: %v", err)
	}

	fileWarning, err := os.OpenFile(path+"/warning.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening fileWarning file: %v", err)
	}

	fileError, err := os.OpenFile(path+"/error.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening fileError file: %v", err)
	}

	logFiles := map[string]*os.File{
		"trace":   fileTrace,
		"info":    fileInfo,
		"warning": fileWarning,
		"error":   fileError,
	}

	Trace = log.New(fileTrace,
		"TRACE: ",
		log.Ldate|log.Ltime|log.Llongfile)

	Info = log.New(fileInfo,
		"INFO: ",
		log.Ldate|log.Ltime|log.Llongfile)

	Warning = log.New(fileWarning,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Llongfile)

	Error = log.New(fileError,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Llongfile)

	watchFiles(logFiles)
}

// watchFiles observa los archivos log
// los trunca de ser necesario y genera un backup
func watchFiles(files map[string]*os.File) {
	go func() {
		for {
			for _, f := range files {
				fi, err := f.Stat()
				if err != nil {
					log.Printf("revisando las estadísticas del archivo: %s", f.Name())
					Error.Printf("revisando las estadísticas del archivo: %s", f.Name())
				}

				// Si el tamaño es de una mega
				if fi.Size() >= 1024*1024 {
					archiveFile(f)
				}
			}
			time.Sleep(time.Second * time.Duration(logReviewInterval))
		}
	}()
}

// archiveFile Genera un backup del archivo
func archiveFile(file *os.File) {
	var err error
	fPath := fmt.Sprintf(
		"%s-%s.log",
		file.Name()[:len(file.Name())-4],
		time.Now().Format("20060102150405"),
	)

	af, err := os.OpenFile(fPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Printf("error creating %s file: %v\n", fPath, err)
	}
	defer af.Close()

	fContent, err := ioutil.ReadFile(file.Name())
	if err != nil {
		log.Printf("error al leer el contenido actual del archivo %s: %v\n", file.Name(), err)
	}
	af.Write(fContent)

	err = ioutil.WriteFile(file.Name(), []byte("Se creó el archivo "+fPath+"\n"), 0666)
	if err != nil {
		log.Println("no se pudo limpiar el archivo de log actual")
	}
}
