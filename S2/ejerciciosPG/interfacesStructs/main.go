package main

import "fmt"

type User struct {
	nombre   string
	apellido string
	edad     int
	correo   string
	pass     string
}

type Iuser interface {
	cambiarNombre()
	cambiarApellido()
	cambiarEdad()
	cambiarCorreo()
	cambiarPass()
}

func main() {

	user1 := User{
		nombre:   "Opi",
		apellido: "Haimer",
		edad:     55,
		correo:   "manhattan55@gmail.com",
		pass:     "3-2-1",
	}

	user1.getUser()
	fmt.Println("cambiarNombre")
	user1.cambiarNombre("Kylian")

	fmt.Println()
	fmt.Println("cambiarApellido")
	user1.cambiarApellido("Mbappe")

	fmt.Println()
	fmt.Println("cambiarEdad")
	user1.cambiarEdad(23)

	fmt.Println()
	fmt.Println("cambiarCorreo")
	user1.cambiarCorreo("km@gmail.com")

	fmt.Println()
	fmt.Println("cambiarPass")
	user1.cambiarPass("310")

	fmt.Println()
	fmt.Println("user modified")
	user1.getUser()
	fmt.Println()

}

func (u *User) cambiarNombre(nombre string) {
	u.nombre = nombre
}

func (u *User) cambiarApellido(apellido string) {
	u.apellido = apellido
}

func (u *User) cambiarEdad(edad int) {
	u.edad = edad
}

func (u *User) cambiarCorreo(correo string) {
	u.correo = correo
}

func (u *User) cambiarPass(pass string) {
	u.pass = pass
}

func (u *User) getUser() {
	fmt.Println(u.nombre, " ", u.apellido, " ", u.edad, " ", u.correo, " ")
}
