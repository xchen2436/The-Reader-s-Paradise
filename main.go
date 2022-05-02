package main

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
	//"io/ioutil"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
	http.HandleFunc("/", Index)
	http.HandleFunc("/Login.html", Login)
	http.HandleFunc("/Register.html", Register)
	http.HandleFunc("/Success.html", Success)
	http.HandleFunc("/Store.html", Store)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	log.Println("Starting the web server at port 8080")
	http.ListenAndServe(":8080", nil)

}
func Index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "Index.html", nil)
}
func Login(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "Login.html", nil)
}
func Register(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "Register.html", nil)
}
func Store(w http.ResponseWriter, r *http.Request) {
	lusername := r.FormValue("loginUsername")
	readUsers, err := os.ReadFile("Users.txt")
	if err != nil {
		panic(err)
	}
	s := string(readUsers)
	CreateValidate, err := os.Create("UsersValidate.txt")
	if err != nil {
		panic(err)
	}
	if strings.Contains(s, lusername) == true {
		OpenUsers, err := os.Open("Users.txt")
		if err != nil {
			panic(err)
		}
		Userscanner := bufio.NewScanner(OpenUsers)
		for Userscanner.Scan() {
			line := Userscanner.Text()
			items := strings.Split(line, ",")
			loginusername := items[0]
			loginpassword := items[1]
			lusername := r.FormValue("loginUsername")
			if lusername == loginusername {
				CreateValidate.WriteString(loginusername + "," + loginpassword + "\n")
				if err != nil {
					fmt.Println(err)
				} else {
					ReadValidate, err := os.Open("UsersValidate.txt")
					if err != nil {
						panic(err)
					}
					Validatescanner := bufio.NewScanner(ReadValidate)
					for Validatescanner.Scan() {
						line := Validatescanner.Text()
						items := strings.Split(line, ",")
						validateusername := items[0]
						validatepassword := items[1]
						lusername := r.FormValue("loginUsername")
						lpassword := r.FormValue("loginPassword")
						if lusername == validateusername && lpassword == validatepassword {
							tpl.ExecuteTemplate(w, "Store.html", nil)
						} else {
							loginfail := r.FormValue("loginfail") == "error"
							err := tpl.ExecuteTemplate(w, "Login.html", loginfail)
							if err != nil {
								http.Error(w, "Something wrong with your username or password, please try again!", 500)
							}
						}
					}
				}
			} else {
			}

		}
	} else {
		tpl.ExecuteTemplate(w, "Login.html", nil)
	}
}
func Cart(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "Cart.html", nil)
}
func Success(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	uname := r.FormValue("registerusername")
	pword := r.FormValue("registerpassword")

	d := struct {
		Username string
		Password string
	}{
		Username: uname,
		Password: pword,
	}
	tpl.ExecuteTemplate(w, "Success.html", d)
	file, err := os.OpenFile("Users.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	} else {
		file.WriteString(uname + "," + pword + "\n")
		fmt.Println("User Added")
	}
	file.Close()
}
