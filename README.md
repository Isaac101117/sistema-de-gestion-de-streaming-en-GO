# 🎬 GoStream - Sistema de Gestión de Streaming (Versión Final)

![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)
![Database](https://img.shields.io/badge/Database-SQLite-003B57?style=flat&logo=sqlite)
![Frontend](https://img.shields.io/badge/Frontend-HTML%20%2B%20Bootstrap-563D7C?style=flat&logo=bootstrap)
![Status](https://img.shields.io/badge/Estado-Finalizado-success)

> **Proyecto Académico:** Programación Orientada a Objetos (UIDE)  
> **Estudiante:** Isaac Hernández  
> **Etapa:** Entrega Final (Implementación Web y SQL)

---

## 📋 Descripción del Proyecto

**GoStream** es una aplicación web transaccional diseñada para administrar un servicio de streaming.
A diferencia de las versiones anteriores basadas en consola, esta versión final implementa una arquitectura **Cliente-Servidor** con persistencia de datos real.

El sistema permite gestionar (Crear, Leer, Actualizar) catálogos de películas, series y usuarios a través de una interfaz gráfica moderna, almacenando toda la información en una base de datos **SQLite**.

---

## 🚀 Características Técnicas

Este proyecto integra los conocimientos adquiridos durante todo el curso:

* **Persistencia de Datos (SQL):** Uso de `SQLite` para almacenar datos permanentemente. El archivo `gostream.db` se genera automáticamente.
* **Servidor Web en Go:** Uso del paquete `net/http` para manejar peticiones y respuestas.
* **Driver Pure Go:** Implementación de `modernc.org/sqlite` para asegurar compatibilidad total en Windows/Mac/Linux sin necesidad de compiladores C (CGO).
* **Frontend Dinámico:** Plantillas `html/template` integradas con **Bootstrap 5** para una UI responsiva.
* **Gestión de Errores:** Manejo robusto de conexiones a base de datos y errores HTTP.

---

## 🛠️ Tecnologías Utilizadas

| Componente | Tecnología | Descripción |
| :--- | :--- | :--- |
| **Backend** | Go (Golang) | Lógica del servidor y manejo de rutas. |
| **Base de Datos** | SQLite | Motor SQL ligero y serverless. |
| **Driver** | modernc.org/sqlite | Conector SQL que no requiere GCC. |
| **Frontend** | HTML5 / CSS3 | Estructura y diseño (Bootstrap CDN). |

---

## ⚙️ Instalación y Ejecución

Sigue estos pasos para correr el proyecto en tu máquina local:

### 1. Clonar el repositorio
```bash
git clone [https://github.com/Isaac101117/sistema-de-gestion-de-streaming-en-GO.git](https://github.com/Isaac101117/sistema-de-gestion-de-streaming-en-GO.git)
cd sistema-de-gestion-de-streaming-en-GO
2. Instalar dependencias
Es importante descargar el driver de SQLite:

Bash

go mod tidy
3. Ejecutar el servidor
Bash

go run main.go
4. Abrir en el navegador
Una vez veas el mensaje 🚀 Servidor corriendo..., abre tu navegador favorito y ve a:

http://localhost:8080

📂 Estructura del Proyecto
main.go: (Backend) Contiene toda la lógica del servidor, conexión a BD, creación de tablas y controladores (Handlers).

index.html: (Frontend) Interfaz gráfica con formularios y tablas dinámicas.

go.mod / go.sum: Gestión de dependencias del proyecto.

gostream.db: Archivo de base de datos (se crea automáticamente al ejecutar el programa).

✅ Funcionalidades Implementadas
1. Catálogo de Películas
Registro de Título, Género, Duración y Sinopsis.

Visualización en tabla dinámica.

2. Catálogo de Series
Registro diferenciado con número de Temporadas y Capítulos.

3. Gestión de Usuarios
Alta de nuevos usuarios.

Actualización de datos: Modificación de correo electrónico mediante búsqueda de username (UPDATE SQL).

👤 Autor
Desarrollado por Isaac para la asignatura de POO.
