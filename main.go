package main

import (
	"fmt"
	"strings"
)

// 1. DEFINICIÓN DE ESTRUCTURAS (Sintaxis básica)
// En lugar de clases, Go usa 'structs' para definir modelos de datos.
type Pelicula struct {
	Titulo    string
	Categoria string
	Duracion  int // en minutos
