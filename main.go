package main

import (
	"fmt"

	"./multimedia"
)

func main() {
	cW := multimedia.ContenidoWeb{Multimedias: []multimedia.Multimedia{}}
	op := 1
	var auxTitulo string
	var auxFormato string
	var auxEntero int64
	for op != 0 {
		fmt.Println("1) Capturar Imagen")
		fmt.Println("2) Capturar Audio")
		fmt.Println("3) Capturar Video")
		fmt.Println("4) Mostrar Contenido Web")
		fmt.Println("0) Salir")
		fmt.Println("Opción: ")
		fmt.Scan(&op)
		switch op {
		case 1:
			fmt.Println("Capture titulo: ")
			fmt.Scan(&auxTitulo)
			fmt.Println("Capture formato: ")
			fmt.Scan(&auxFormato)
			fmt.Println("Capture canales: ")
			fmt.Scan(&auxEntero)
			cW.Multimedias = append(cW.Multimedias, &multimedia.Imagen{
				Titulo:  auxTitulo,
				Formato: auxFormato,
				Canales: auxEntero,
			})

		case 2:
			fmt.Println("Capture titulo: ")
			fmt.Scan(&auxTitulo)
			fmt.Println("Capture formato: ")
			fmt.Scan(&auxFormato)
			fmt.Println("Capture duracion(segundos): ")
			fmt.Scan(&auxEntero)
			cW.Multimedias = append(cW.Multimedias, &multimedia.Audio{
				Titulo:   auxTitulo,
				Formato:  auxFormato,
				Duracion: auxEntero,
			})
		case 3:
			fmt.Println("Capture titulo: ")
			fmt.Scan(&auxTitulo)
			fmt.Println("Capture formato: ")
			fmt.Scan(&auxFormato)
			fmt.Println("Capture frames: ")
			fmt.Scan(&auxEntero)
			cW.Multimedias = append(cW.Multimedias, &multimedia.Video{
				Titulo:  auxTitulo,
				Formato: auxFormato,
				Frames:  auxEntero,
			})
		case 4:
			fmt.Println("Opción 3")
		case 0:
			fmt.Println("Adios")

		default:
			fmt.Println("Opción no existe")
		}
	}
}
