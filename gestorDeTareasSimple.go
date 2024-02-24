package main

import (
	"fmt"
	)

type Tarea struct {
	nombre string
	descripcion string
	estado string
}

func ActualizarEstado(tarea *Tarea){
	valido := false
	for !valido {
		fmt.Println("Seleccione el estado de la tarea: ")
		fmt.Println("1. En proceso")
		fmt.Println("2. Terminada")
		fmt.Println("3. Abandonada")
		var opcion int
		fmt.Scanln(&opcion)
		switch opcion {
		case 1:
			tarea.estado = "en proceso"
			valido = true
		case 2:
			tarea.estado = "terminada"
			valido = true
		case 3:
			tarea.estado = "abanadonada"
			valido = true
		default:
			fmt.Println("Opcion no valida")
		}
	}
}


func MostrarTareas(tareas []Tarea){
	for i := 0; i < len(tareas); i++ {
		fmt.Println("Tarea ", i+1)
		fmt.Println("Nombre: ", tareas[i].nombre)
		fmt.Println("Descripcion: ", tareas[i].descripcion)
		fmt.Println("Estado: ", tareas[i].estado)
		fmt.Println("-------------------------------")
	}
}

func main(){
	opcion := 0
	var tareas []Tarea;

	for opcion != 5 {
		fmt.Println("1. Agregar tarea")
		fmt.Println("2. Mostrar tareas")
		fmt.Println("3. Eliminar tarea")
		fmt.Println("4. Actualizar estado de tarea")
		fmt.Println("5. Salir")
		fmt.Scanln(&opcion)
		switch opcion {
		case 1:
			var nombre string
			var descripcion string
			fmt.Println("Ingrese el nombre de la tarea")
			fmt.Scanln(&nombre)
			fmt.Println("Ingrese la descripcion de la tarea")
			fmt.Scanln(&descripcion)
			tarea := Tarea{nombre, descripcion, "pendiente"}
			tareas = append(tareas, tarea)
		case 2:
			MostrarTareas(tareas)
		case 3:
			fmt.Println("Seleccione la tarea a eliminar:")
			MostrarTareas(tareas)
			fmt.Print("Opcion: ")
			var opcion int
			fmt.Scanln(&opcion)
			tareas = append(tareas[:opcion-1], tareas[opcion:]...)
			fmt.Println("Tarea eliminada")
		case 4:
			fmt.Println("Seleccione la tarea a actualizar:")
			MostrarTareas(tareas)
			fmt.Print("Opcion: ")
			var opcion int
			fmt.Scanln(&opcion)
			ActualizarEstado(&tareas[opcion-1])
		default:
			fmt.Println("Opcion no valida")
		}
	}
}
