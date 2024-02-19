package main

import "fmt"

func main() {

	usuario := UserT{"Sebastian", "Torres", 30, "st@gmail.com", "12"}
	Users = append(Users, usuario);

	usuario.cambiarEdad(22)
	fmt.Printf("\nLa nueva edad es : %d\n",usuario.edad);

	usuario.cambiarNombre("Isildur")
	fmt.Printf("\nEl nuevo nombre es : %s\n",usuario.nombre);

	usuario.cambiarCorreo("str9393@gmail.com")
	fmt.Printf("\nEl nuevo correo es : %s\n",usuario.correo);

	usuario.cambiarContrasenia("555")
	fmt.Printf("\nLa nueva contraseña es : %s\n",usuario.contraseña);


	fmt.Printf("\nMostrar todos los usuarios\n")

	u := UserT{}
    u.GetAll()
	
}

type User interface {
	cambiarNombre() string
	cambiarEdad() string
	cambiarCorreo() string
	cambiarContrasenia() string
}

type UserT struct {
	nombre     string
	apellido   string
	edad       int
	correo     string
	contraseña string
}

var Users []UserT

func (u *UserT) cambiarNombre(nombre string) string {
	u.nombre = nombre
	return u.nombre
}

func (u *UserT) cambiarEdad(edad int) int {
	u.edad = edad
	return u.edad

}

func (u *UserT) cambiarCorreo(correo string) string {
			u.correo = correo
			return u.correo
}

func (u *UserT) cambiarContrasenia(pass string) string {
			u.contraseña = pass
			return u.contraseña
}

func (u *UserT) GetAll() {
	for _, user := range Users {
		fmt.Printf("\nNombre %s, Apellido: %s,Correo: %s, Constrasenia: %s\n",user.nombre, user.apellido,user.correo,user.contraseña);
	}
	
}
