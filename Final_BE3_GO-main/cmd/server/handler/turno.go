package handler

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/luisaribotto/Final_BE3_GO/internal/domain"
	"github.com/luisaribotto/Final_BE3_GO/internal/turno"
	"github.com/luisaribotto/Final_BE3_GO/pkg/web"
)

type turnoHandler struct {
	s turno.Service
}

func NewTurnoHandler(s turno.Service) *turnoHandler {
	return &turnoHandler{
		s: s,
	}
}

func (h *turnoHandler) CreateTurno() gin.HandlerFunc {
	return func(c *gin.Context) {
		var turno domain.Turno

		err := c.ShouldBindJSON(&turno)
		if err != nil {
			web.Failure(c, 400, errors.New("JSON DE TURNO INVALIDO"))
			return
		}
		p, err := h.s.CreateTurno(turno)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 201, p, "TURNO CREADO CORRECTAMENTE")
	}
}

func (h *turnoHandler) GetTurnoByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("idTurno")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("ID INVALIDO"))
			return
		}
		turno, err := h.s.GetTurnoByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("TURNO NO ENCONTRADO"))
			return
		}
		web.Success(c, 200, turno, "TURNO OBTENIDO POR ID")
	}
}

func validateFieldsTurno(turno *domain.Turno) (bool, error) {
	switch {
	case turno.DescripcionTurno == "" || turno.FechaTurno == "" || turno.IdOdontologo == "" || turno.IdPaciente == "":
		return false, errors.New("LOS CAMPOS NO PUEDEN SER VACIOS")
	}
	return true, nil
}

func (h *turnoHandler) UpdateTurno() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("idTurno")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("ID INVALIDO"))
			return
		}
		_, err = h.s.GetTurnoByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("ID DE TURNO INEXISTENTE"))
			return
		}
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		var turno domain.Turno
		err = c.ShouldBindJSON(&turno)
		if err != nil {
			web.Failure(c, 400, errors.New("JSON INVALIDO"))
			return
		}
		valid, err := validateFieldsTurno(&turno)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		p, err := h.s.UpdateTurno(id, turno)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, p, "TURNO ACTUALIZADO CON EXITO")
	}
}

func (h *turnoHandler) UpdateTurnoForField() gin.HandlerFunc {
	type Request struct {
		DescripcionTurno string `json:"descripcionTurno,omitempty"`
		FechaTurno       string `json:"fechaTurno,omitempty"`
		IdOdontologo     string `json:"idOdontologo,omitempty"`
		IdPaciente       string `json:"idPaciente,omitempty"`
	}
	return func(c *gin.Context) {
		var r Request
		idParam := c.Param("idTurno")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("ID INVALIDO"))
			return
		}
		_, err = h.s.GetTurnoByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("ID DE TURNO INEXISTENTE"))
			return
		}
		if err := c.ShouldBindJSON(&r); err != nil {
			web.Failure(c, 400, errors.New("JSON INVALIDO"))
			return
		}
		update := domain.Turno{
			DescripcionTurno: r.DescripcionTurno,
			FechaTurno:       r.FechaTurno,
			IdOdontologo:     r.IdOdontologo,
			IdPaciente:       r.IdPaciente,
		}
		p, err := h.s.UpdateTurno(id, update)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, p, "TURNO ACTUALIZADO CON EXITO")
	}
}

func (h *turnoHandler) DeleteTurno() gin.HandlerFunc {

	return func(c *gin.Context) {
		idParam := c.Param("idTurno")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("ID INVALIDO"))
			return
		}
		_, err = h.s.GetTurnoByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("ID DE TURNO INEXISTENTE"))
			return
		}
		err = h.s.DeleteTurno(id)
		if err != nil {
			web.Failure(c, 404, err)
			return
		}
		web.Success(c, 200, nil, "TURNO ELIMINADO CORRECTAMENTE")
	}
}