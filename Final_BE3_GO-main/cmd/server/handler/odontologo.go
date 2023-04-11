package handler

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/luisaribotto/Final_BE3_GO/internal/domain"
	"github.com/luisaribotto/Final_BE3_GO/internal/odontologo"
	"github.com/luisaribotto/Final_BE3_GO/pkg/web"
)

type odontologoHandler struct {
	s odontologo.Service
}

func NewOdontologoHandler(s odontologo.Service) *odontologoHandler {
	return &odontologoHandler{
		s: s,
	}
}

func (h *odontologoHandler) CreateOdontologo() gin.HandlerFunc {
	return func(c *gin.Context) {
		var odontologo domain.Odontologo

		err := c.ShouldBindJSON(&odontologo)
		if err != nil {
			web.Failure(c, 400, errors.New("JSON DE ODONTOLOGO INVALIDO"))
			return
		}
		p, err := h.s.Create(odontologo)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 201, p, "ODONTOLOGO CREADO CORRECTAMENTE")
	}
}

func (h *odontologoHandler) GetOdontologoByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("idOdontologo")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("ID INVALIDO"))
			return
		}
		odontologo, err := h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("ODONTOLOGO NO ENCONTRADO"))
			return
		}
		web.Success(c, 200, odontologo, "ODONTOLOGO OBTENIDO POR ID")
	}
}

func validateEmptys(odontologo *domain.Odontologo) (bool, error) {
	switch {
	case odontologo.NombreOdontologo == "" || odontologo.ApellidoOdontologo == "" || odontologo.MatriculaOdontologo == "":
		return false, errors.New("LOS CAMPOS NO PUEDEN SER VACIOS")
	}
	return true, nil
}

func (h *odontologoHandler) UpdateOdontologo() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("idOdontologo")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("ID INVALIDO"))
			return
		}
		_, err = h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("ID DE ODONTOLOGO INEXISTENTE"))
			return
		}
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		var odontologo domain.Odontologo
		err = c.ShouldBindJSON(&odontologo)
		if err != nil {
			web.Failure(c, 400, errors.New("JSON INVALIDO"))
			return
		}
		p, err := h.s.Update(id, odontologo)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, p, "ODONTOLOGO ACTUALIZADO CON EXITO")
	}
}

func (h *odontologoHandler) UpdateOdontologoForField() gin.HandlerFunc {
	type Request struct {
		NombreOdontologo    string `json:"nombreOdontologo,omitempty"`
		ApellidoOdontologo  string `json:"apellidoOdontologo,omitempty"`
		MatriculaOdontologo string `json:"matriculaOdontologo,omitempty"`
	}
	return func(c *gin.Context) {
		var r Request
		idParam := c.Param("idOdontologo")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("ID INVALIDO"))
			return
		}
		_, err = h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("ID DE ODONTOLOGO INEXISTENTE"))
			return
		}
		if err := c.ShouldBindJSON(&r); err != nil {
			web.Failure(c, 400, errors.New("JSON INVALIDO"))
			return
		}
		update := domain.Odontologo{
			NombreOdontologo:    r.NombreOdontologo,
			ApellidoOdontologo:  r.ApellidoOdontologo,
			MatriculaOdontologo: r.MatriculaOdontologo,
		}
		p, err := h.s.Update(id, update)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, p, "ODONTOLOGO ACTUALIZADO CON EXITO")
	}
}

func (h *odontologoHandler) DeleteOdontologo() gin.HandlerFunc {

	return func(c *gin.Context) {

		idParam := c.Param("idOdontologo")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("ID INVALIDO"))
			return
		}
		_, err = h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("ID DE ODONTOLOGO INEXISTENTE"))
			return
		}
		err = h.s.Delete(id)
		if err != nil {
			web.Failure(c, 404, err)
			return
		}
		web.Success(c, 200, nil, "ODONTOLOGO ELIMINADO CORRECTAMENTE")
	}
}
