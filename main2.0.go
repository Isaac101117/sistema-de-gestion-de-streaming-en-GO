package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// --- 1. ESTRUCTURAS (TEMAS UNIDAD 1 Y 3) ---

type Pelicula struct {
	Nombre   string
	Genero   string
	Duracion int
	Sinopsis string
}

type Serie struct {
	Nombre     string
	Genero     string
	Temporadas int
	Capitulos  int
	Sinopsis   string
}

type Usuario struct {
	Username string
	Correo   string
}

// --- 2. "BASE DE DATOS" EN MEMORIA (SLICES) ---

var peliculas = []Pelicula{
	{"Inception", "Sci-Fi", 148, "Un ladrón que roba secretos a través de los sueños."},
}

var series = []Serie{
	{"Breaking Bad", "Drama", 5, 62, "Un profesor de química se vuelve productor de metanfetamina."},
}

var usuarios = []Usuario{
	{"isaac10", "isaac@gmail.com"},
}

// --- 3. FUNCIONES DE LECTURA ---

func leerTexto(reader *bufio.Reader) string {
	texto, _ := reader.ReadString('\n')
	return strings.TrimSpace(texto)
}

// --- 4. SISTEMA PRINCIPAL ---

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n================================")
		fmt.Println("   SISTEMA GOSTREAM - GESTIÓN")
		fmt.Println("================================")
		fmt.Println("1. Ver Catálogo de Películas")
		fmt.Println("2. Ver Catálogo de Series")
		fmt.Println("3. Ver Lista de Usuarios")
		fmt.Println("4. Registrar Nueva Película")
		fmt.Println("5. Registrar Nueva Serie")
		fmt.Println("6. Registrar Nuevo Usuario")
		fmt.Println("7. Actualizar Datos de Usuario (Punteros)")
		fmt.Println("8. Salir")
		fmt.Print("Elija una opción: ")

		opcionStr := leerTexto(reader)
		opcion, _ := strconv.Atoi(opcionStr)

		switch opcion {
		case 1:
			fmt.Println("\n--- PELÍCULAS ---")
			for _, p := range peliculas {
				fmt.Printf("🎬 %s | Género: %s | %d min\n   Sinopsis: %s\n", p.Nombre, p.Genero, p.Duracion, p.Sinopsis)
			}

		case 2:
			fmt.Println("\n--- SERIES ---")
			for _, s := range series {
				fmt.Printf("📺 %s | Género: %s | %d Temp / %d Caps\n   Sinopsis: %s\n", s.Nombre, s.Genero, s.Temporadas, s.Capitulos, s.Sinopsis)
			}

		case 3:
			fmt.Println("\n--- USUARIOS REGISTRADOS ---")
			for i, u := range usuarios {
				fmt.Printf("%d. Usuario: %s | Email: %s\n", i+1, u.Username, u.Correo)
			}

		case 4:
			fmt.Println("\n--- REGISTRAR PELÍCULA ---")
			fmt.Print("Nombre: ")
			n := leerTexto(reader)
			fmt.Print("Género: ")
			g := leerTexto(reader)
			fmt.Print("Duración (min): ")
			dStr := leerTexto(reader)
			d, _ := strconv.Atoi(dStr)
			fmt.Print("Sinopsis: ")
			s := leerTexto(reader)
			peliculas = append(peliculas, Pelicula{n, g, d, s})
			fmt.Println("✅ Película guardada.")

		case 5:
			fmt.Println("\n--- REGISTRAR SERIE ---")
			fmt.Print("Nombre: ")
			n := leerTexto(reader)
			fmt.Print("Género: ")
			g := leerTexto(reader)
			fmt.Print("Temporadas: ")
			tStr := leerTexto(reader)
			t, _ := strconv.Atoi(tStr)
			fmt.Print("Capítulos: ")
			cStr := leerTexto(reader)
			c, _ := strconv.Atoi(cStr)
			fmt.Print("Sinopsis: ")
			s := leerTexto(reader)
			series = append(series, Serie{n, g, t, c, s})
			fmt.Println("✅ Serie guardada.")

		case 6:
			fmt.Println("\n--- REGISTRAR USUARIO ---")
			fmt.Print("Nombre de Usuario: ")
			u := leerTexto(reader)
			fmt.Print("Gmail/Email: ")
			e := leerTexto(reader)
			usuarios = append(usuarios, Usuario{u, e})
			fmt.Println("✅ Usuario registrado con éxito.")

		case 7:
			fmt.Println("\n--- ACTUALIZAR USUARIO ---")
			fmt.Print("Número de usuario a editar: ")
			idxStr := leerTexto(reader)
			idx, _ := strconv.Atoi(idxStr)

			if idx > 0 && idx <= len(usuarios) {
				// USO DE PUNTEROS (UNIDAD 3)
				// Obtenemos la dirección de memoria del usuario seleccionado
				usuarioRef := &usuarios[idx-1]

				fmt.Printf("Editando a: %s\n", usuarioRef.Username)
				fmt.Print("Nuevo Nombre de Usuario: ")
				usuarioRef.Username = leerTexto(reader)
				fmt.Print("Nuevo Gmail: ")
				usuarioRef.Correo = leerTexto(reader)

				fmt.Println("✅ Datos actualizados mediante referencia de memoria.")
			} else {
				fmt.Println("❌ El usuario no existe.")
			}

		case 8:
			fmt.Println("Saliendo del sistema...")
			return

		default:
			fmt.Println("Opción inválida.")
		}
	}
}
