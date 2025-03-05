package domain

//Enrollent sera una especia de tabla intermedia entre course y user

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Enrollment struct {
	ID        string       `json:"id" gorm:"type:char(36);not null;primaryKey;uniqueIndex"` //Char 36 (pq uuid), no nulo, PK, unico
	UserID    string       `json:"user_id,omitempty" gorm:"type:char(36)"`                  //COMO ES MICROSERVICIO Y BBDD SEPARADA ESTARÁ PERO NO COMO FK
	User      *User        `json:"user,omitempty" gorm:"-"`                                 //OJO AQUI LE PONEMOS gorm - PARA QUE NO SEA CREE COMO FK (PQ SERAN BBDD SEPARADAS)
	CourseID  string       `json:"course_id,omitempty" gorm:"type:char(36)"`                //COMO ES MICROSERVICIO Y BBDD SEPARADA ESTARÁ PERO NO COMO FK
	Course    *Course      `json:"course,omitempty" gorm:"-"`                               //OJO AQUI LE PONEMOS gorm - PARA QUE NO SEA CREE COMO FK (PQ SERAN BBDD SEPARADAS)
	Status    EnrollStatus `json:"status" gorm:"type:char(2)"`                              //Estado del curso sera de tipo enrollstatus y cmo teine valores PREDEFINIDOS DEBERIA solo permitir estos
	CreatedAt *time.Time   `json:"-"`                                                       // `json:"-"` PARA QUE NO se incluya este campo en las respuestas JSON
	UpdatedAt *time.Time   `json:"-"`
}

// Definiremos un tipo de dato que representa un String y que tendra teimpos de datos (como enum?)
type EnrollStatus string

const (
	Pending  EnrollStatus = "P"
	Active   EnrollStatus = "A"
	Studying EnrollStatus = "S"
	Inactive EnrollStatus = "I"
)

func (e *Enrollment) BeforeCreate(tx *gorm.DB) (err error) {
	if e.ID == "" {
		e.ID = uuid.New().String()
	}
	return
}
