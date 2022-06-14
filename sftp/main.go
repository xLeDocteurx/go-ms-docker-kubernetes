package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/url"
	"os"
	"strings"
	"time"
	
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	
	// _ "net/http/pprof" // only for profiling
	// "net/http"
)

func main() {
	// // dev : profiling
	// go func() {
	// 	log.Println(http.ListenAndServe("localhost:6060", nil))
	// }()

	// config
	host := "51.178.193.133"
	port := "22"
	username := "jth-company"
	password := "LKKqbEaQ0gApvmthPuw6n_gG"
	dirName := "test-files"

	//----------------

	// Create a url
	rawurl := fmt.Sprintf("sftp://%v:%v@%v", username, password, host)
	// normally : parse url etc ... not for poc
	// Parse the URL
	parsedUrl, err := url.Parse(rawurl)
	if err != nil {
		log.Fatalf("Failed to parse SFTP To Go URL: %v", err)
	}
	fmt.Println("parsedUrl :", parsedUrl)

	// Get hostkey
	// hostKey := getHostKey(host) // should be implemented in prod

	auths := []ssh.AuthMethod{ssh.Password(password)}

	// Initialize ssh client configuration
	config := ssh.ClientConfig{
		User:            username,
		Auth:            auths,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		// HostKeyCallback: ssh.FixedHostKey(hostKey),
		// Timeout: 30 * time.Second,
	}

	addr := fmt.Sprintf("%v:%v", host, port)

	//----------------

	treatedFilesGlobal := map[string]bool{} // use map as set in go

	for {
		treatedFilesForLoop := map[string]bool{} // use map as set in go

		// Connect to server (ssh client)
		conn, err := ssh.Dial("tcp", addr, &config)
		if err != nil {
			log.Fatalf("Failed to connec to host [%v]: %v", addr, err)
		}
		// normally use : defer.conn.CLose() , but for now close manually (defer only déallocates at end of function)

		// Create new SFTP client
		sc, err := sftp.NewClient(conn)
		if err != nil {
			log.Fatalf("Unable to start SFTP subsystem: %v", err)
		}
		// defer sc.Close()

		// TODO : config db_client + connect to db endpoint

		// list files from dir dirName
		files, err := sc.ReadDir(dirName)
		if err != nil {
			log.Fatalf("Unable to list remote dir: %v", err)
		}

		nbFiles := len(files)

		for _, file := range files {
			fileDescription := fmt.Sprintf("%v -- modified : %v", file.Name(), file.ModTime()) // to store in set

			// check if file already treated
			_, fileDescriptionInSet := treatedFilesGlobal[fileDescription]
			if fileDescriptionInSet {
				continue // file already treated => go directly to next iteration
			}

			openedFile, err := sc.OpenFile(fmt.Sprintf("%v/%v", dirName, file.Name()), (os.O_RDONLY))
			if err != nil {
				log.Fatalf("unable to open remote file: %v", err)
			}

			// log
			fmt.Println()
			fmt.Println("file description :", fileDescription)

			// READ AS CSV
			headerIndexMap := map[string]int{}
			csvReader := csv.NewReader(openedFile)

			// read line by line (first is header) + create values for sql
			valuesStrList := []string{}
			isFirstRow := true
			for {
				record, err := csvReader.Read()
				if err != nil {
					if err == io.EOF {
						break
					} else {
						log.Fatalf("Unable to read line of file as csv : %v", err)
					}
				}

				// handle in function if it is first row (header) or others :
				if isFirstRow {
					// header
					// Add mapping: Column/property name --> record index
					for i, v := range record {
						headerIndexMap[v] = i
					}
					isFirstRow = false

				} else {
					// other rows
					// variables that should be configurable
					defaultValue := "blabla"
					colName1 := "MSG-NOM-ASS"
					colName2 := "MSG-NUM-POL"
					var colVal1 string
					var colVal2 string

					recordIndex, colIsPresent := headerIndexMap[colName1]
					if colIsPresent {
						colVal1 = record[recordIndex]
					} else {
						colVal1 = defaultValue
					}

					recordIndex, colIsPresent = headerIndexMap[colName2]
					if colIsPresent {
						colVal2 = record[recordIndex]
					} else {
						colVal2 = defaultValue
					}

					valuesStr := fmt.Sprintf("('%v', '%v---%v') ", colVal1, colVal2, colVal1)
					valuesStrList = append(valuesStrList, valuesStr)
				}
			}
			openedFile.Close()

			// create sql request
			sqlRequest := fmt.Sprintf("INSERT INTO client_1.files (\"MSG-NOM-ASS\", customhash) VALUES %v;", strings.Join(valuesStrList, ","))

			fmt.Println(sqlRequest)
			// PUT IT IN THE DB !!!
			// TODO

			// update treated files
			treatedFilesGlobal[fileDescription] = true  // global state
			treatedFilesForLoop[fileDescription] = true // for logging for this iteration

		}

		// free connection to sftp
		sc.Close()
		conn.Close()
		// free connection to db
		// TODO

		// log
		if len(treatedFilesForLoop) > 0 {
			fmt.Println()
			fmt.Println("Files treated in this iteration :", len(treatedFilesForLoop))
			fmt.Println("--------------------------------------------------------------")

		} else {
			fmt.Printf("-%v", nbFiles)
			time.Sleep(1 * time.Second)
		}

	}

	println("blabla")
}
