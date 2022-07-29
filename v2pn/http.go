package v2pn

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// https://tutorialedge.net/golang/creating-simple-web-server-with-golang/

func Http_server_start(dataDir string, externalFilesDir string) {
	f, err := os.OpenFile(externalFilesDir+"/httpLog.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	multiWriter := io.MultiWriter(os.Stdout, f)
	log.SetOutput(multiWriter)

	log.Println("Starting http server")

	// http.HandleFunc("/", http_handler)
	// http.Handle("/", http.FileServer(http.Dir(".")))
	http.Handle("/dataDir/", http.StripPrefix("/dataDir/", http.FileServer(http.Dir(dataDir))))
	http.Handle("/externalFilesDir/", http.StripPrefix("/externalFilesDir/", http.FileServer(http.Dir(externalFilesDir))))

	http.HandleFunc("/v2pn/uploadConfig", func(writer http.ResponseWriter, request *http.Request) {
		log.Println("request.URL.Path: ", request.URL.Path)
		switch request.Method {
		case "GET":
			log.Println("No GET Method")
			// http.ServeFile(writer, request, "D:\\Documents\\notes\\vscode\\go\\form.html")
		case "POST":
			// Parse our multipart form, 10 << 20 specifies a maximum
			// upload of 10 MB files.
			request.ParseMultipartForm(10 << 20)
			// FormFile returns the first file for the given key `myFile`
			// it also returns the FileHeader so we can get the Filename,
			// the Header and the size of the file
			file, handler, err := request.FormFile("filename")
			if err != nil {
				log.Println("Error Retrieving the File")
				log.Println(err)
				return
			}
			defer file.Close()
			log.Println("Uploaded File: ", handler.Filename)
			log.Println("File Size: ", handler.Size)
			log.Println("MIME Header: ", handler.Header)

			tempFile, err := os.OpenFile(dataDir+"/"+handler.Filename, os.O_RDWR|os.O_CREATE, 0666)
			if err != nil {
				log.Fatalf("error opening file: %v", err)
			}
			defer tempFile.Close()

			// read all of the contents of our uploaded file into a
			// byte array
			fileBytes, err := io.ReadAll(file)
			if err != nil {
				log.Println(err)
			}
			// write this byte array to our temporary file
			tempFile.Write(fileBytes)
			// return that we have successfully uploaded our file!
			fmt.Fprintf(writer, "Successfully Uploaded File\n")
		default:
			fmt.Fprintf(writer, "Sorry, only GET and POST methods are supported.")
		}
	})

	http.HandleFunc("/v2pn/allowedList", func(writer http.ResponseWriter, request *http.Request) {
		log.Println("request.URL.Path: ", request.URL.Path)
		switch request.Method {
		case "GET":
			allowedList, err := os.OpenFile(dataDir+"/allowedList.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
			if err != nil {
				log.Fatalf("error opening file: %v", err)
			}
			defer allowedList.Close()
			data, _ := io.ReadAll(allowedList)
			fmt.Fprintf(writer, string(data))
		case "POST":
			allowedList, err := os.OpenFile(dataDir+"/allowedList.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
			if err != nil {
				log.Fatalf("error opening file: %v", err)
			}
			defer allowedList.Close()
			body, _ := io.ReadAll(request.Body)
			allowedList.WriteString(string(body) + "\n")
			fmt.Fprintf(writer, "Successfully POST %s\n", string(body))
		default:
			fmt.Fprintf(writer, "Sorry, only GET and POST methods are supported.")
		}
	})

	http.HandleFunc("/v2pn/globalvpn", func(writer http.ResponseWriter, request *http.Request) {
		log.Println("request.URL.Path: ", request.URL.Path)
		switch request.Method {
		case "GET":
			globalvpn, err := os.OpenFile(dataDir+"/globalvpn.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
			if err != nil {
				log.Fatalf("error opening file: %v", err)
			}
			defer globalvpn.Close()
			data, _ := io.ReadAll(globalvpn)
			fmt.Fprintf(writer, string(data))
		case "POST":
			globalvpn, err := os.OpenFile(dataDir+"/globalvpn.txt", os.O_RDWR|os.O_CREATE, 0666)
			if err != nil {
				log.Fatalf("error opening file: %v", err)
			}
			defer globalvpn.Close()
			body, _ := io.ReadAll(request.Body)
			globalvpn.WriteString(string(body) + "\n")
			fmt.Fprintf(writer, "Successfully POST %s\n", string(body))
		default:
			fmt.Fprintf(writer, "Sorry, only GET and POST methods are supported.")
		}
	})

	http.HandleFunc("/v2pn/running", func(writer http.ResponseWriter, request *http.Request) {
		log.Println("request.URL.Path: ", request.URL.Path)
		switch request.Method {
		case "GET":
			running, err := os.OpenFile(dataDir+"/running.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
			if err != nil {
				log.Fatalf("error opening file: %v", err)
			}
			defer running.Close()
			data, _ := io.ReadAll(running)
			fmt.Fprintf(writer, string(data))
		case "POST":
			log.Println("No POST Method")
		default:
			fmt.Fprintf(writer, "Sorry, only GET and POST methods are supported.")
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
