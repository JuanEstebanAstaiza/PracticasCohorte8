package PruebaTecnicaTres

//ejercicio 2

/*type TaskStatus string

const (
	Pending    TaskStatus = "pendiente"
	InProgress TaskStatus = "en progreso"
	Completed  TaskStatus = "completada"
)

type Task struct {
	Name        string
	Description string
	Assignee    string
	Status      TaskStatus
}

type TaskManager struct {
	tasks []Task
}

func (tm *TaskManager) AddTask(name, description string) {
	task := Task{

		Name:        name,
		Description: description,
		Status:      Pending,
	}
	tm.tasks = append(tm.tasks, task)
}

func (tm *TaskManager) AssignTask(name, assignee string) error {
	for i, task := range tm.tasks {
		if task.Name == name {
			tm.tasks[i].Assignee = assignee
			return nil
		}
	}
	return errors.New("tarea no encontrada")
}

func (tm *TaskManager) UpdateTaskStatus(name string, status TaskStatus) error {
	for i, task := range tm.tasks {
		if task.Name == name {
			tm.tasks[i].Status = status
			return nil
		}
	}
	return errors.New("Tarea no encontrada")
}

func (tm *TaskManager) PendingTasks() []Task {
	var pendingTasks []Task
	for _, task := range tm.tasks {
		if task.Status == Pending {
			pendingTasks = append(pendingTasks, task)
		}
	}
	return pendingTasks
}

func main() {
	tm := &TaskManager{}

	tm.AddTask("Tarea 1", "Descripción de la tarea 1")
	tm.AddTask("Tarea 2", "Descripcion de la tarea 2")

	err := tm.AssignTask("Tarea 1", "Juan")
	if err != nil {
		fmt.Println("Error actualizando estado de tarea: ", err)
	}

	pendingTasks := tm.PendingTasks()
	fmt.Printf("Tareas pendientes: ")
	for _, task := range pendingTasks {
		fmt.Printf("Nombre: %s, Descripcion: %s, Responsable: %s, Estado: %s\n", task.Name, task.Description, task.Assignee, task.Status)
	}
}*/

//Ejercicio 3

/*type BookStatus string

const (
	Available BookStatus = "disponible"
	Borrowed  BookStatus = "prestado"
)

type Book struct {
	Title  string
	Author string
	Genre  string
	Status BookStatus
}

type Library struct {
	books []Book
}

func (l *Library) FindBooks(query string) []Book {
	var foundBooks []Book
	for _, book := range l.books {
		if book.Title == query || book.Author == query {
			foundBooks = append(foundBooks, book)
		}
	}
	return foundBooks
}

func (l *Library) AddBook(title, author, genre string) {
	book := Book{
		Title:  title,
		Author: author,
		Genre:  genre,
		Status: Available,
	}
	l.books = append(l.books, book)
}

func (l *Library) UpdateBookStatus(title string, status BookStatus) error {
	for i, book := range l.books {
		if book.Title == title {
			l.books[i].Status = status
			return nil
		}
	}
	return errors.New("Libro no encontrado")
}

func (l *Library) RemoveBook(title string) error {
	for i, book := range l.books {
		if book.Title == title {
			l.books = append(l.books[:i], l.books[i+1:]...)
			return nil
		}
	}
	return errors.New("libro no encontrado")
}

func main() {
	library := &Library{}

	library.AddBook("El Quijote", "Miguel de Cervantes", "Novela")
	library.AddBook("Cien años de soledad", "Gabriel García Márquez", "Novela")

	err := library.UpdateBookStatus("El Quijote", Borrowed)
	if err != nil {
		fmt.Println("Error actualizando el estado del libro: ", err)
	}

	foundBooks := library.FindBooks("Gabriel García Márquez")
	fmt.Println("Libros encontrados: ")
	for _, book := range foundBooks {
		fmt.Printf("Título: %s, Autor: %s, Género: %s, Estado: %s\n", book.Title, book.Author, book.Genre, book.Status)
	}

	err = library.RemoveBook("Cien años de soledad")
	if err != nil {
		fmt.Println("Error eliminando el libro: ", err)
	}

	fmt.Println("Inventario completo: ")
	for _, book := range library.books {
		fmt.Printf("Título: %s, Autor: %s, Género: %s, Estado: %s\n", book.Title, book.Author, book.Genre, book.Status)
	}

}*/
