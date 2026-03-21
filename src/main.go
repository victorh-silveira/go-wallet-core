package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/victor-silveira/go-wallet-core/src/application/user"
	"github.com/victor-silveira/go-wallet-core/src/application/wallet"
	"github.com/victor-silveira/go-wallet-core/src/domain/entity"
	"github.com/victor-silveira/go-wallet-core/src/infrastructure/repository/memory"
	"github.com/victor-silveira/go-wallet-core/src/interfaces/http/handler"
)

const AppVersion = "1.0.0"

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	slog.SetDefault(logger)

	slog.Info("iniciando servico", "version", AppVersion)

	userRepo := memory.NewUserRepository()
	walletRepo := memory.NewWalletRepository()

	createUserUseCase := user.NewCreateUserUseCase(userRepo)
	processTrxUseCase := wallet.NewProcessTransactionUseCase(walletRepo, walletRepo)

	userHandler := handler.NewUserHandler(createUserUseCase)
	walletHandler := handler.NewWalletHandler(processTrxUseCase)
	healthHandler := handler.NewHealthHandler(AppVersion)

	if os.Getenv("SEED_DEFAULT_ACCOUNT") != "false" {
		ctx := context.Background()
		initialAcc, err := entity.NewAccount("ACC-001", "USER-001")
		if err != nil {
			slog.Error("seed conta", "error", err)
			os.Exit(1)
		}
		const seedBRLCentavos = int64(50_000)
		if err := initialAcc.UpdateBalance(seedBRLCentavos); err != nil {
			slog.Error("saldo seed", "error", err)
			os.Exit(1)
		}
		if err := walletRepo.SaveAccount(ctx, initialAcc); err != nil {
			slog.Error("persistir seed", "error", err)
			os.Exit(1)
		}
		slog.Info("conta seed carregada", "account_id", "ACC-001", "balance_centavos", seedBRLCentavos)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", healthHandler.HealthCheck)
	mux.HandleFunc("POST /users", userHandler.CreateUser)
	mux.HandleFunc("POST /wallet/transaction", walletHandler.Transaction)
	mux.HandleFunc("GET /swagger.yaml", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "api/swagger.yaml")
	})
	mux.HandleFunc("GET /docs", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "api/index.html")
	})

	server := &http.Server{
		Addr:              ":8080",
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       15 * time.Second,
	}

	go func() {
		slog.Info("servidor escutando", "addr", server.Addr)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error("servidor encerrou com erro", "error", err)
			os.Exit(1)
		}
	}()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	<-ctx.Done()
	slog.Info("sinal de encerramento recebido")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(shutdownCtx); err != nil {
		slog.Error("shutdown", "error", err)
		os.Exit(1)
	}
	slog.Info("servidor encerrado com sucesso")
}
