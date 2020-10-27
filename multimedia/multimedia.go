package multimedia

import (
	"fmt"
)

type ContenidoWeb struct {
	Multimedias []Multimedia
}

func (cW *ContenidoWeb) Mostrar() {
	for _, m := range cW.Multimedias {
		m.Mostrar()
	}
}

type Multimedia interface {
	Mostrar()
}
type Imagen struct {
	Titulo  string
	Formato string
	Canales int64
}
type Audio struct {
	Titulo   string
	Formato  string
	Duracion int64
}
type Video struct {
	Titulo  string
	Formato string
	Frames  int64
}

func (i *Imagen) Mostrar() {
	fmt.Println("--Imagen--")
	fmt.Println("Titulo: ", i.Titulo)
	fmt.Println("Formato: ", i.Formato)
	fmt.Println("Canales: ", i.Canales)
}
func (a *Audio) Mostrar() {
	fmt.Println("--Audio--")
	fmt.Println("Titulo: ", a.Titulo)
	fmt.Println("Formato: ", a.Formato)
	fmt.Println("Duración: ", a.Duracion)
}
func (v *Video) Mostrar() {
	fmt.Println("--Video--")
	fmt.Println("Titulo: ", v.Titulo)
	fmt.Println("Formato: ", v.Formato)
	fmt.Println("Frames: ", v.Frames)
}
