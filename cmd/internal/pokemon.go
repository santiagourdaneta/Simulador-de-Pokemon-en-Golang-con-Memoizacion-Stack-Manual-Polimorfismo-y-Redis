package pokemon

import "fmt"

// Interfaz para el polimorfismo
type Pokemon interface {
	GetNombre() string
	Atacar() int
	GetID() string // Necesario para la Key de Redis
}

type Pikachu struct { Nombre string }
func (p Pikachu) GetNombre() string { return p.Nombre }
func (p Pikachu) GetID() string     { return "pika_" + p.Nombre }
func (p Pikachu) Atacar() int       { return 40 } // Daño fijo 

type Charmander struct { Nombre string }
func (c Charmander) GetNombre() string { return c.Nombre }
func (c Charmander) GetID() string     { return "char_" + c.Nombre }
func (c Charmander) Atacar() int       { return 50 } 