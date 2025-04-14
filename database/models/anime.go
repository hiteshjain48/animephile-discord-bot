package models

import(
	"github.com/google/uuid"
)

type Anime struct {
	ID          uuid.UUID       
	Title       string 
}