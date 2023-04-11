package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"
	"github.com/luisaribotto/Final_BE3_GO/cmd/server/handler"
	"github.com/luisaribotto/Final_BE3_GO/internal/odontologo"
	"github.com/luisaribotto/Final_BE3_GO/internal/paciente"
	"github.com/luisaribotto/Final_BE3_GO/internal/turno"
	"github.com/luisaribotto/Final_BE3_GO/pkg/store"
)

func main() {

	db, err := sql.Open("mysql", "root:root@/my_db")
	if err != nil {
		panic(err.Error())
	}
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	storage := store.NewSqlStore(db)

	repo := odontologo.NewRepository(storage)
	service := odontologo.NewService(repo)
	odontologoHandler := handler.NewOdontologoHandler(service)

	repoPaciente := paciente.NewRepository(storage)
	servicePaciente := paciente.NewService(repoPaciente)
	pacienteHandler := handler.NewPacienteHandler(servicePaciente)

	repoTurno := turno.NewRepository(storage)
	serviceTurno := turno.NewService(repoTurno)
	turnoHandler := handler.NewTurnoHandler(serviceTurno)

	engine := gin.Default()

	engine.SetTrustedProxies([]string{"127.0.0.1"})

	engine.GET("/api/v1/ping", func(c *gin.Context) { c.String(200, "pong") })

	odontologos := engine.Group("/api/v1/odontologos")
	{
		odontologos.POST("", odontologoHandler.CreateOdontologo())
		odontologos.GET(":idOdontologo", odontologoHandler.GetOdontologoByID())
		odontologos.PUT(":idOdontologo", odontologoHandler.UpdateOdontologo())
		odontologos.PATCH(":idOdontologo", odontologoHandler.UpdateOdontologoForField())
		odontologos.DELETE(":idOdontologo", odontologoHandler.DeleteOdontologo())

	}

	// determino las uris de paciente
	pacientes := engine.Group("/api/v1/pacientes")
	{
		pacientes.POST("", pacienteHandler.CreatePaciente())
		pacientes.GET(":idPaciente", pacienteHandler.GetPacienteByID())
		pacientes.PUT(":idPaciente", pacienteHandler.UpdatePaciente())
		pacientes.PATCH(":idPaciente", pacienteHandler.UpdatePacienteForField())
		pacientes.DELETE(":idPaciente", pacienteHandler.DeletePaciente())

	}

	turnos := engine.Group("/api/v1/turnos")
	{
		turnos.POST("", turnoHandler.CreateTurno())
		turnos.GET(":idTurno", turnoHandler.GetTurnoByID())
		turnos.PUT(":idTurno", turnoHandler.UpdateTurno())
		turnos.PATCH(":idTurno", turnoHandler.UpdateTurnoForField())
		turnos.DELETE(":idTurno", turnoHandler.DeleteTurno())

	}

	engine.Run(":8080")
}
