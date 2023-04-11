package paciente

import (
	"errors"

	"github.com/luisaribotto/Final_BE3_GO/internal/domain"
	"github.com/luisaribotto/Final_BE3_GO/pkg/store"
)

type Repository interface {
	GetPacienteByID(id int) (domain.Paciente, error)
	CreatePaciente(p domain.Paciente) (domain.Paciente, error)
	UpdatePaciente(id int, p domain.Paciente) (domain.Paciente, error)
	DeletePaciente(id int) error
}

type repository struct {
	storage store.StoreInterface
}

func NewRepository(storage store.StoreInterface) Repository {
	return &repository{storage}
}

func (r *repository) CreatePaciente(p domain.Paciente) (domain.Paciente, error) {

	err := r.storage.CreatePaciente(p)
	if err != nil {
		return domain.Paciente{}, errors.New("ERROR CREANDO PACIENTE")
	}
	return p, nil
}

func (r *repository) GetPacienteByID(id int) (domain.Paciente, error) {
	paciente, err := r.storage.ReadPaciente(id)
	if err != nil {
		return domain.Paciente{}, errors.New("PACIENTE INEXISTENTE")
	}
	return paciente, nil
}

func (r *repository) UpdatePaciente(id int, p domain.Paciente) (domain.Paciente, error) {
	err := r.storage.UpdatePaciente(p)
	if err != nil {
		return domain.Paciente{}, errors.New("ERROR ACTUALIZANDO AL PACIENTE")
	}
	return p, nil
}

func (r *repository) DeletePaciente(id int) error {
	err := r.storage.DeletePaciente(id)
	if err != nil {
		return err
	}
	return nil
}
