package main

import (
	"context"
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
	restError "github.com/pauloRohling/locknote/internal/presentation/rest/error"
	notePresentation "github.com/pauloRohling/locknote/internal/presentation/rest/note"
	tokenPresentation "github.com/pauloRohling/locknote/internal/presentation/rest/token"
	userPresentation "github.com/pauloRohling/locknote/internal/presentation/rest/user"
	"go.uber.org/zap"
)

func main() {
	env := environment.Env()

	applicationLogger := environment.GetApplicationLogger()
	persistenceLogger := environment.GetPersistenceLogger()
	presentationLogger := environment.GetPresentationLogger()
	securityLogger := environment.GetSecurityLogger()
	defer func(applicationLogger *zap.Logger) { _ = applicationLogger.Sync() }(applicationLogger)
	defer func(persistenceLogger *zap.Logger) { _ = persistenceLogger.Sync() }(persistenceLogger)
	defer func(presentationLogger *zap.Logger) { _ = presentationLogger.Sync() }(presentationLogger)
	defer func(securityLogger *zap.Logger) { _ = securityLogger.Sync() }(securityLogger)

	dbPoolBuilder := postgres.NewPoolBuilder(env.GetDatabaseAddress(), env.GetDatabaseUrl(), persistenceLogger)
	dbPool := dbPoolBuilder.Build(context.Background())
	defer dbPool.Close()

	persistenceLogger.Info(
		"Database connection established",
		zap.String("address", env.GetDatabaseAddress()),
	)

	tokenVerifier, err := token.NewPasetoVerifier(
		env.Security.Paseto.PublicKey,
		env.Security.Auth.Issuer,
	)

	if err != nil {
		securityLogger.Fatal("unable to create token verifier", zap.Error(err))
	}

	tokenIssuer, err := token.NewPasetoIssuer(
		env.Security.Paseto.SecretKey,
		env.Security.Paseto.PublicKey,
		env.Security.Auth.Issuer,
		env.Security.Auth.Expiration,
	)

	if err != nil {
		securityLogger.Fatal("unable to create token issuer", zap.Error(err))
	}

	tokenVerifierMiddleware := tokenPresentation.VerifierMiddleware(tokenVerifier)

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
	getNoteUseCase := noteApplication.NewGetNoteUseCase(noteApplication.GetNoteParams{
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
		GetNoteUseCase:    getNoteUseCase,
	})

	userRestController := userPresentation.NewRestController(userService)
	noteRestController := notePresentation.NewRestController(noteService, tokenVerifierMiddleware)

	server := rest.NewWebServer(env.Server.Port)
	server.SetErrorHandler(restError.NewErrorHandler(presentationLogger))
	server.Register(userRestController)
	server.Register(noteRestController)

	if err := server.Start(); err != nil {
		presentationLogger.Fatal("unable to start web server", zap.Error(err))
	}
}
