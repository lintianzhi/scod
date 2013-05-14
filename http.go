package main

import (
    "net/http"
    "regexp"
    "io/ioutil"
    "text/template"
    "fmt"
)

var template_dir = "templates/"
var infile_dir  = "highlight_py/infiles/"
var outfile_dir = "highlight_py/outfiles/"
var fileValidator = regexp.MustCompile("^[0-9a-zA-Z]+$")

var templates = template.Must(template.ParseFiles(template_dir+"edit.html",template_dir+"view.html"))

type Page struct {
    Title string
    Body  []byte
}

func loadPage(title string) (*Page, error) {
    body, err := ioutil.ReadFile(outfile_dir+title)
    if err != nil {
        return nil, err
    }
    return &Page{title,body}, nil
}

const lenPath0 = len("/sharecode/")
func readCodeHandler(w http.ResponseWriter,r *http.Request) {
    fn := r.URL.Path[lenPath0:]
    if len(fn)==0 {
        // for edit
    }
    if !fileValidator.MatchString(fn) {
        http.NotFound(w, r)
        return
    }
    page, err := loadPage(fn)
    if err != nil {
        fmt.Println(err)
        http.NotFound(w,r)
        return
    }
    renderTemplate(w, "view", page)
}

func editCodeHandler(w http.ResponseWriter, r *http.Request) {
    renderTemplate(w, "edit", nil)
}

func saveCodeHandler(w http.ResponseWriter, r *http.Request) {
    body := r.FormValue("code")
    code_tp := r.FormValue("code_type")
    fmt.Println(code_tp)
    title, err := process_code(body, code_tp)  // process_code should return title
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    http.Redirect(w, r, "/sharecode/"+title, http.StatusFound)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
    http.Redirect(w, r, "/sharecode", http.StatusFound)
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
    err := templates.ExecuteTemplate(w, tmpl+".html", p)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}
