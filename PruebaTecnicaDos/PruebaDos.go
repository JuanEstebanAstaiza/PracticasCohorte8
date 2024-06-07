package main

import "fmt"

//ejercicio 2

/*type Tarea struct {
	Nombre      string
	Descripcion string
	Responsable string
	Estado      string // "pendiente" , "en progreso" o "completada"
}

func crearTarea(nombre, descripcion, responsable string) *Tarea {
	nuevaTarea := &Tarea{
		Nombre:      nombre,
		Descripcion: descripcion,
		Responsable: responsable,
		Estado:      "pendiente", //Por defecto

	}
	return nuevaTarea
}

func actualizarEstado(tarea *Tarea, nuevoEstado string) {
	tarea.Estado = nuevoEstado
}

func mostrarTareasPendientes(tareas []*Tarea) {
	fmt.Println("Tareas pendientes: ")
	for _, tarea := range tareas {
		if tarea.Estado == "pendiente" {
			fmt.Printf("- Nombre: %s, Descripción: %s, Responsable: %s\n", tarea.Nombre, tarea.Descripcion, tarea.Responsable)
		}
	}
}

func main() {
	// Crear algunas tareas

	tarea1 := crearTarea("Implementar inicio de sesion", "Implementar la funcionalidad de inicio de sesion a mi aplicacion", "Juan")
	tarea2 := crearTarea("Crear un banco en Java", "Diseñar e implementar una aplicacion bancaria en Java 8", "André")
	tarea3 := crearTarea("Crear sitio web para un restaurante", "Crear un sitio web usando React para un restaurante", "Ignacio")

	//Mostrar tareas
	tareas := []*Tarea{tarea1, tarea2, tarea3}
	mostrarTareasPendientes(tareas)

	//actualizar los estados
	actualizarEstado(tarea1, "Completada")
	fmt.Println("\nDespués de actualizar el estado de una tarea: ")
	mostrarTareasPendientes(tareas)
}*/

//ejercicio 4

//Estructura para un contacto

/*type Contacto struct {
	Nombre    string
	Telefono  string
	Correo    string
	Direccion string
}

// Funcion de agregacion

func agregarContacto(agenda map[string]Contacto, nombre, telefono, correo, direccion string) {
	nuevoContacto := Contacto{

		Nombre:    nombre,
		Telefono:  telefono,
		Correo:    correo,
		Direccion: direccion,
	}
	agenda[nombre] = nuevoContacto
	fmt.Printf("Contacto '%s' agregado a la agenda. \n", nombre)
}

func buscarPorNombre(agenda map[string]Contacto, nombre string) (Contacto, bool) {
	fmt.Println("Buscar usuario por nombre: ")
	contacto, encontrado := agenda[nombre]
	if !encontrado {
		fmt.Printf("Contacto '%s' no encontrado.\n", nombre)
	}
	return contacto, encontrado
}

func buscarPorTelefono(agenda map[string]Contacto, telefono string) []Contacto {

	var contactosPorTelefono []Contacto
	for _, contacto := range agenda {
		if contacto.Telefono == telefono {
			contactosPorTelefono = append(contactosPorTelefono, contacto)
		}
	}
	return contactosPorTelefono
}

func actualizarContacto(agenda map[string]Contacto, nombre, telefono, correo, direccion string) {

	if _, encontrado := agenda[nombre]; encontrado {
		contactoActualizado := Contacto{
			Nombre:    nombre,
			Telefono:  telefono,
			Correo:    correo,
			Direccion: direccion,
		}
		agenda[nombre] = contactoActualizado
		fmt.Printf("Contacto '%s' actualizado en la agenda.\n", nombre)

	} else {
		fmt.Printf("Contacto '%s' no encontrado.\n", nombre)
	}
}

func eliminarContacto(agenda map[string]Contacto, nombre string) {
	if _, encontrado := agenda[nombre]; encontrado {
		delete(agenda, nombre)
		fmt.Printf("Contacto '%s' eliminado de la agenda.\n", nombre)
	} else {
		fmt.Printf("Contacto '%s' no encontrado.\n", nombre)
	}
}

func mostrarAgenda(agenda map[string]Contacto) {
	fmt.Println("Agenda de contactos: ")
	for _, contacto := range agenda {
		fmt.Printf("- Nombre: %s, Telefono %s, Correo %s, Dirección: %s", contacto.Nombre, contacto.Telefono, contacto.Correo, contacto.Direccion)
	}
}

func main() {
	agenda := make(map[string]Contacto)

	// Agregar algunos contactos a la agenda
	agregarContacto(agenda, "Juan Pérez", "123456789", "juan@example.com", "Calle 123")
	agregarContacto(agenda, "María López", "987654321", "maria@example.com", "Avenida 456")

	// Mostrar la agenda completa
	mostrarAgenda(agenda)

	// Actualizar la información de un contacto
	actualizarContacto(agenda, "Juan Pérez", "123456789", "juan@gmail.com", "Calle 1234")

	// Mostrar la agenda después de actualizar un contacto
	mostrarAgenda(agenda)

	// Eliminar un contacto de la agenda
	eliminarContacto(agenda, "María López")

	// Mostrar la agenda después de eliminar un contacto
	mostrarAgenda(agenda)
}*/

//ejercicio 3

type Libro struct {
	Titulo     string
	Autor      string
	Genero     string
	Disponible bool
}

func agregarLibro(coleccion map[string]Libro, titulo, autor, genero string) {
	nuevoLibro := Libro{
		Titulo:     titulo,
		Autor:      autor,
		Genero:     genero,
		Disponible: true,
	}
	coleccion[titulo] = nuevoLibro
	fmt.Printf("Libro '%s' agregado a la coleccion \n", titulo)

}

func buscarPorTitulo(coleccion map[string]Libro, titulo string) (Libro, bool) {
	libro, encontrado := coleccion[titulo]
	if !encontrado {
		fmt.Printf("Libro '%s' no encontrado. \n", titulo)
	}
	return libro, encontrado
}

func buscarPorAutor(coleccion map[string]Libro, autor string) []Libro {
	var librosPorAutor []Libro
	for _, libro := range coleccion {
		if libro.Autor == autor {
			librosPorAutor = append(librosPorAutor, libro)
		}
	}
	return librosPorAutor
}

func actualizarEstado(coleccion map[string]Libro, titulo string, disponible bool) {
	if libro, encontrado := coleccion[titulo]; encontrado {
		libro.Disponible = disponible
		coleccion[titulo] = libro
		estado := "disponible"
		if !disponible {
			estado = "prestado"
		}
		fmt.Printf("Estado del libro '%s' actualizado a %s. \n", titulo, estado)
	} else {
		fmt.Printf("Libro '%s' no encontrado. \n", titulo)
	}
}

func eliminarLibro(coleccion map[string]Libro, titulo string) {
	if _, encontrado := coleccion[titulo]; encontrado {
		delete(coleccion, titulo)
		fmt.Printf("Libro '%s' eliminado de la coleccion.\n ", titulo)
	} else {
		fmt.Printf("Libro '%s' no encontrado. \n", titulo)
	}
}

func mostrarColeccion(coleccion map[string]Libro) {
	fmt.Println("Coleccion de libros: ")
	for _, libro := range coleccion {
		estado := "disponible"
		if !libro.Disponible {
			estado = "prestado"
		}
		fmt.Printf("- Titulo: %s, Autor: %s, Genero: %s, Estado: %s\n", libro.Titulo, libro.Autor, libro.Genero, estado)
	}
}

func main() {
	coleccion := make(map[string]Libro)

	// Agregar algunos libros a la colección
	agregarLibro(coleccion, "El principito", "Antoine de Saint-Exupéry", "Fábula")
	agregarLibro(coleccion, "Cien años de soledad", "Gabriel García Márquez", "Realismo mágico")
	agregarLibro(coleccion, "1984", "George Orwell", "Distopía")

	// Mostrar la colección completa
	mostrarColeccion(coleccion)

	// Actualizar el estado de un libro
	actualizarEstado(coleccion, "El principito", false)

	// Mostrar la colección después de actualizar el estado de un libro
	mostrarColeccion(coleccion)

	// Eliminar un libro de la colección
	eliminarLibro(coleccion, "1984")

	// Mostrar la colección después de eliminar un libro
	mostrarColeccion(coleccion)
}
