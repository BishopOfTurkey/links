package linkstore

import (
	"encoding/csv"
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
)

type IndexView struct {
	Links []string
}

type AddLinkJSON struct {
	Url  string
	Code string
}

func Server(password string) {
	file, err := os.OpenFile("links.csv", os.O_RDONLY|os.O_CREATE, 0644)

	if err != nil {
		log.Fatalf("Failed to open file: %v\n", err)
	}

	r := csv.NewReader(file)
	linksArr, err := r.ReadAll()
	file.Close()
	links := make([]string, len(linksArr))
	for i, l := range linksArr {
		links[i] = l[0]
	}

	if err != nil {
		log.Fatalf("Failed to parse csv: %v\n", err)
	}

	fmap := template.FuncMap{
		"reverse": reverse,
	}

	t, err := template.New("index.html").Funcs(fmap).ParseFiles("index.html")
	if err != nil {
		log.Fatalf("Failed to parse html template: %v\n", err)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		view := IndexView{
			Links: links,
		}
		t.Execute(w, view)
	})

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	mux.HandleFunc("/add", func(w http.ResponseWriter, req *http.Request) {
		dec := json.NewDecoder(req.Body)
		in := &AddLinkJSON{}
		err := dec.Decode(in)
		if err != nil {
			log.Printf("Failed to decode JSON: %v", err)
			w.WriteHeader(400)
			io.WriteString(w, "Failed to decode JSON.")
			return
		}
		if password != in.Code {
			w.WriteHeader(403)
			io.WriteString(w, "Code wrong.")
			return
		}
		links = append(links, in.Url)
		file, err = os.OpenFile("links.csv", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
		if err != nil {
			log.Printf("Couldn't open csv file %v", err)
		}
		csvW := csv.NewWriter(file)
		err = csvW.Write([]string{in.Url})
		if err != nil {
			log.Printf("Failed to write new record: %v", err)
		}
		w.WriteHeader(201)
		io.WriteString(w, "Ok.")
		csvW.Flush()

		err = file.Sync()
		if err != nil {
			log.Printf("Failed to save csv: %v", err)
		}
	})

	log.Fatalln(http.ListenAndServe(":8002", mux))
}

func reverse(strs []string) []string {
	rev := make([]string, len(strs))
	for i, str := range strs {
		rev[len(strs)-i-1] = str
	}
	return rev
}
