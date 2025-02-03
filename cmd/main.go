package main

import (
	"context"
	"fmt"
	noteApplication "github.com/pauloRohling/locknote/internal/application/note"
	"github.com/pauloRohling/locknote/internal/application/token"
	userApplication "github.com/pauloRohling/locknote/internal/application/user"
	noteDomain "github.com/pauloRohling/locknote/internal/domain/note"
	userDomain "github.com/pauloRohling/locknote/internal/domain/user"
	"github.com/pauloRohling/locknote/internal/environment"
	notePersistence "github.com/pauloRohling/locknote/internal/persistence/note"
	"github.com/pauloRohling/locknote/internal/persistence/postgres"
	userPersistence "github.com/pauloRohling/locknote/internal/persistence/user"
	"github.com/pauloRohling/locknote/internal/presentation/rest"
	notePresentation "github.com/pauloRohling/locknote/internal/presentation/rest/note"
	userPresentation "github.com/pauloRohling/locknote/internal/presentation/rest/user"
	"log/slog"
)

func main() {
	env := environment.Env()

	dbPoolBuilder := postgres.NewPoolBuilder(env.GetDatabaseAddress(), env.GetDatabaseUrl())
	dbPool := dbPoolBuilder.Build(context.Background())
	defer dbPool.Close()

	slog.Info("Database connection established")

	tokenVerifier, err := token.NewPasetoVerifier(
		env.Security.Paseto.PublicKey,
		env.Security.Auth.Issuer,
	)

	if err != nil {
		panic(fmt.Errorf("unable to create token verifier: %w", err))
	}

	tokenIssuer, err := token.NewPasetoIssuer(
		env.Security.Paseto.SecretKey,
		env.Security.Paseto.PublicKey,
		env.Security.Auth.Issuer,
		env.Security.Auth.Expiration,
	)

	if err != nil {
		panic(fmt.Errorf("unable to create token issuer: %w", err))
	}

	userFactory := userDomain.NewFactory()
	noteFactory := noteDomain.NewFactory()

	userMapper := userPersistence.NewMapper(userFactory)
	noteMapper := notePersistence.NewMapper(noteFactory)

	userRepository := userPersistence.NewRepository(dbPool, userMapper)
	noteRepository := notePersistence.NewRepository(dbPool, noteMapper)

	createUserUseCase := userApplication.NewCreateUserUseCase(userApplication.CreateUserParams{
		UserFactory:    userFactory,
		UserRepository: userRepository,
	})
	createNoteUseCase := noteApplication.NewCreateNoteUseCase(noteApplication.CreateNoteParams{
		NoteFactory:    noteFactory,
		NoteRepository: noteRepository,
	})
	loginUseCase := userApplication.NewLoginUseCase(userApplication.LoginUsecaseParams{
		TokenIssuer:    tokenIssuer,
		UserRepository: userRepository,
	})

	userService := userApplication.NewService(userApplication.FacadeServiceParams{
		CreateUseCase: createUserUseCase,
		LoginUseCase:  loginUseCase,
	})
	noteService := noteApplication.NewService(noteApplication.FacadeServiceParams{
		CreateNoteUseCase: createNoteUseCase,
	})

	userRestController := userPresentation.NewRestController(userService)
	noteRestController := notePresentation.NewRestController(noteService, tokenVerifier)

	server := rest.NewWebServer(env.Server.Port)
	server.Register(userRestController)
	server.Register(noteRestController)

	if err := server.Start(); err != nil {
		panic(fmt.Errorf("unable to start web server: %w", err))
	}
}
