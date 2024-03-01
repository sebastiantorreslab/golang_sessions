package main

import "fmt"

type Bootcamp struct {
	Organizacion string
	Lenguaje     string
	Alumnos      []Alumno
}

type Alumno struct {
	Nombre       string
	FechaIngreso string
}

func createBootcamp(org, lng string) Bootcamp {
	boot := Bootcamp{
		Organizacion: org,
		Lenguaje:     lng,
	}

	return boot
}

func (b *Bootcamp) saveAlumni(alumno Alumno) {
	b.Alumnos = append(b.Alumnos, alumno)
}

func main() {

	bootcampGO := createBootcamp("MELI", "GO")

	alm1 := Alumno {
		Nombre:"Agustin",
		FechaIngreso: "22/02/22",
	}
	alm2 := Alumno {
		Nombre:"Micca",
		FechaIngreso: "22/02/22",
	}
	alm3 := Alumno {
		Nombre:"Sebastian",
		FechaIngreso: "22/02/22",
	}

	bootcampGO.saveAlumni(alm1)
	bootcampGO.saveAlumni(alm2)
	bootcampGO.saveAlumni(alm3)


	fmt.Print(bootcampGO);

}