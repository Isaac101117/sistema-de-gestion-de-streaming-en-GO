// Autor: Isaac Hernandez
// Fecha: 15 de diciembre de 2025
// Descripción: Sistema de Gestion de streaming.
package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "modernc.org/sqlite" // <--- CAMBIO IMPORTANTE: Driver Pure Go
)

// --- ESTRUCTURAS ---
type Pelicula struct {
	ID       int
	Nombre   string
	Genero   string
	Duracion int
	Sinopsis string
}

type Serie struct {
	ID         int
	Nombre     string
	Genero     string
	Temporadas int
	Capitulos  int
	Sinopsis   string
}

type Usuario struct {
	ID       int
	Username string
	Correo   string
}

type PageData struct {
	Peliculas []Pelicula
	Series    []Serie
	Usuarios  []Usuario
	Error     string // Para mostrar errores en pantalla si ocurren
}

var db *sql.DB

func initDB() {
	var err error
	// CAMBIO IMPORTANTE: Usamos "sqlite"
	db, err = sql.Open("sqlite", "./gostream.db")
	if err != nil {
		log.Fatal(err)
	}

	sqlStmt := `
	CREATE TABLE IF NOT EXISTS peliculas (id INTEGER PRIMARY KEY AUTOINCREMENT, nombre TEXT, genero TEXT, duracion INTEGER, sinopsis TEXT);
	CREATE TABLE IF NOT EXISTS series (id INTEGER PRIMARY KEY AUTOINCREMENT, nombre TEXT, genero TEXT, temporadas INTEGER, capitulos INTEGER, sinopsis TEXT);
	CREATE TABLE IF NOT EXISTS usuarios (id INTEGER PRIMARY KEY AUTOINCREMENT, username TEXT, correo TEXT);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("Error creando tablas: %q\n", err)
		return
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	data := PageData{}

	// 1. Consultar Películas (Con manejo de error)
	rows, err := db.Query("SELECT id, nombre, genero, duracion, sinopsis FROM peliculas")
	if err != nil {
		log.Println("Error leyendo peliculas:", err)
	} else {
		defer rows.Close()
		for rows.Next() {
			var p Pelicula
			rows.Scan(&p.ID, &p.Nombre, &p.Genero, &p.Duracion, &p.Sinopsis)
			data.Peliculas = append(data.Peliculas, p)
		}
	}

	// 2. Consultar Series
	rowsSeries, err := db.Query("SELECT id, nombre, genero, temporadas, capitulos, sinopsis FROM series")
	if err != nil {
		log.Println("Error leyendo series:", err)
	} else {
		defer rowsSeries.Close()
		for rowsSeries.Next() {
			var s Serie
			rowsSeries.Scan(&s.ID, &s.Nombre, &s.Genero, &s.Temporadas, &s.Capitulos, &s.Sinopsis)
			data.Series = append(data.Series, s)
		}
	}

	// 3. Consultar Usuarios
	rowsUsers, err := db.Query("SELECT id, username, correo FROM usuarios")
	if err != nil {
		log.Println("Error leyendo usuarios:", err)
	} else {
		defer rowsUsers.Close()
		for rowsUsers.Next() {
			var u Usuario
			rowsUsers.Scan(&u.ID, &u.Username, &u.Correo)
			data.Usuarios = append(data.Usuarios, u)
		}
	}

	// Renderizar
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, "Error cargando index.html: "+err.Error(), 500)
		return
	}
	tmpl.Execute(w, data)
}

func addMovieHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		stmt, _ := db.Prepare("INSERT INTO peliculas(nombre, genero, duracion, sinopsis) VALUES(?, ?, ?, ?)")
		stmt.Exec(r.FormValue("nombre"), r.FormValue("genero"), r.FormValue("duracion"), r.FormValue("sinopsis"))
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func addSerieHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		stmt, _ := db.Prepare("INSERT INTO series(nombre, genero, temporadas, capitulos, sinopsis) VALUES(?, ?, ?, ?, ?)")
		stmt.Exec(r.FormValue("nombre"), r.FormValue("genero"), r.FormValue("temporadas"), r.FormValue("capitulos"), r.FormValue("sinopsis"))
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func addUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		stmt, _ := db.Prepare("INSERT INTO usuarios(username, correo) VALUES(?, ?)")
		stmt.Exec(r.FormValue("username"), r.FormValue("correo"))
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func updateUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		stmt, _ := db.Prepare("UPDATE usuarios SET correo = ? WHERE username = ?")
		stmt.Exec(r.FormValue("new_email"), r.FormValue("username_target"))
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func main() {
	initDB()
	// No cerramos db en main para que siga viva mientras corre el servidor

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/add-movie", addMovieHandler)
	http.HandleFunc("/add-serie", addSerieHandler)
	http.HandleFunc("/add-user", addUserHandler)
	http.HandleFunc("/update-user", updateUserHandler)

	fmt.Println("🚀 Servidor corriendo en: http://localhost:8080")
	fmt.Println("   (Presiona Ctrl+C para detener)")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
