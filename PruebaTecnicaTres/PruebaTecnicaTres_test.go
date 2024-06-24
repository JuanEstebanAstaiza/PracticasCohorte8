package PruebaTecnicaTres

import "testing"

//------------------------------------------------------------------------------------

/*func TestAddTask(t *testing.T) {
	tm := &TaskManager{}
	tm.AddTask("Tarea 1", "Descripcion de la tarea 1")

	if len(tm.tasks) != 1 {
		t.Errorf("Se esperaba 1 tarea, se obtuvo %d", len(tm.tasks))
	}

	if tm.tasks[0].Name != "Tarea 1" {
		t.Errorf("Nombre de la tarea incorrecto, se esperaba 'Tarea 1', se obutuvo %s", tm.tasks[0].Name)
	}
}

func TestAssignTask(t *testing.T) {
	tm := &TaskManager{}
	tm.AddTask("Tarea 1", "Descripción de la tarea 1")
	err := tm.AssignTask("Tarea 1", "Juan")

	if err != nil {
		t.Errorf("Error asignando tarea: %v", err)
	}

	if tm.tasks[0].Assignee != "Juan" {
		t.Errorf("Responsable incorrecto, se esperaba 'Juan', se obtuvo '%s' ", tm.tasks[0].Assignee)
	}
}

func TestUpdateTaskStatus(t *testing.T) {
	tm := &TaskManager{}
	tm.AddTask("Tarea 1", "Descripción de la tarea 1")
	err := tm.UpdateTaskStatus("Tarea 1", InProgress)

	if err != nil {
		t.Errorf("Error actualizando estado de tarea: %v", err)
	}

	if tm.tasks[0].Status != InProgress {
		t.Errorf("Estado incorrecto, se esperaba '%s', se obtuvo '%s'", InProgress, tm.tasks[0].Status)
	}
}

func TestPendingTasks(t *testing.T) {
	tm := &TaskManager{}
	tm.AddTask("Tarea 1", "Descripción de la tarea 1")
	tm.AddTask("Tarea 2", "Descripción de la tarea 2")
	tm.UpdateTaskStatus("Tarea 1", InProgress)

	pendingTasks := tm.PendingTasks()
	if len(pendingTasks) != 1 {
		t.Errorf("Se esperaba 1 tarea pendiente, se obtuvo %d", len(pendingTasks))
	}

	if pendingTasks[0].Name != "Tarea 2" {
		t.Errorf("Nombre de la tarea pendiente incorrecto, se esperaba 'Tarea 2', se obtuvo '%s'", pendingTasks[0].Name)
	}
}

func TestTaskManagerIntegration(t *testing.T) {
	tm := &TaskManager{}

	tm.AddTask("Tarea 1", "Descripción de la tarea 1")
	tm.AddTask("Tarea 2", "Descripción de la tarea 2")

	err := tm.AssignTask("Tarea 1", "Juan")
	if err != nil {
		t.Errorf("Error asignando tarea: %v", err)
	}

	err = tm.UpdateTaskStatus("Tarea 2", InProgress)
	if err != nil {
		t.Errorf("Error actualizando estado de tarea: %v", err)
	}

	pendingTasks := tm.PendingTasks()
	if len(pendingTasks) != 1 {
		t.Errorf("Se esperaba 1 tarea pendiente, se obtuvo %d", len(pendingTasks))
	}

	if pendingTasks[0].Name != "Tarea 1" {
		t.Errorf("Nombre de la tarea pendiente incorrecto, se esperaba 'Tarea 1', se obtuvo '%s'", pendingTasks[0].Name)
	}

	if pendingTasks[0].Assignee != "Juan" {
		t.Errorf("Responsable incorrecto, se esperaba 'Juan', se obtuvo '%s'", pendingTasks[0].Assignee)
	}
}*/

// Tests ejercicio 3

func TestAddBook(t *testing.T) {

	library := &Library{}
	library.AddBook("El Quijote", "Miguel de Cervantes", "Novela")
	library.AddBook("Cien años de soledad", "Gabriel García Márquez", "Novela")

	foundBooks := library.FindBooks("Miguel de Cervantes")
	if len(foundBooks) != 1 {
		t.Errorf("Se esperaba 1 libro, se obtuvo %d", len(foundBooks))
	}

	if foundBooks[0].Title != "El Quijote" {
		t.Errorf("Título de libro encontrado incorrecto, se esperaba 'El Quijote', se obtuvo '%s'", foundBooks[0].Title)
	}
}

func TestUpdateBooksStatus(t *testing.T) {
	library := &Library{}
	library.AddBook("El Quijote", "Miguel de Cervantes", "Novela")
	err := library.UpdateBookStatus("El Quijote", Borrowed)

	if err != nil {
		t.Errorf("Error actualizando estado del libro: %v", err)
	}

	if library.books[0].Status != Borrowed {
		t.Errorf("Estado del libro incorrecto, se esperaba '%s', se obtuvo '%s'", Borrowed, library.books[0].Status)
	}
}

func TestRemoveBook(t *testing.T) {
	library := &Library{}

	library.AddBook("El Quijote", "Miguel de Cervantes", "Novela")
	err := library.RemoveBook("El Quijote")

	if err != nil {
		t.Errorf("Error eliminando libro: %v", err)
	}

	if len(library.books) != 0 {
		t.Errorf("Se esperaba 0 libros, se obtuvo %d", len(library.books))
	}

}

func TestLibraryIntegration(t *testing.T) {
	library := &Library{}

	library.AddBook("El Quijote", "Miguel de Cervantes", "Novela")
	library.AddBook("Cien años de soledad", "Gabriel García Márquez", "Novela")

	err := library.UpdateBookStatus("El Quijote", Borrowed)
	if err != nil {
		t.Errorf("Error actualizando estado del libro: %v", err)
	}

	foundBooks := library.FindBooks("Gabriel García Márquez")
	if len(foundBooks) != 1 {
		t.Errorf("Se esperaba 1 libro encontrado, se obtuvo %d", len(foundBooks))
	}

	if foundBooks[0].Title != "Cien años de soledad" {
		t.Errorf("Título del libro encontrado incorrecto, se esperaba 'Cien años de soledad', se obtuvo '%s'", foundBooks[0].Title)
	}

	err = library.RemoveBook("Cien años de soledad")
	if err != nil {
		t.Errorf("Error eliminando libro: %v", err)
	}

	if len(library.books) != 1 {
		t.Errorf("Se esperaba 1 libro en la colección, se obtuvo %d", len(library.books))
	}

	if library.books[0].Title != "El Quijote" {
		t.Errorf("Título del libro restante incorrecto, se esperaba 'El Quijote', se obtuvo '%s'", library.books[0].Title)
	}
}
