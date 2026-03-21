package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/victor-silveira/go-wallet-core/src/application/user"
	"github.com/victor-silveira/go-wallet-core/src/application/wallet"
	"github.com/victor-silveira/go-wallet-core/src/domain/entity"
	"github.com/victor-silveira/go-wallet-core/src/infrastructure/repository/postgres"
	"github.com/victor-silveira/go-wallet-core/src/interfaces/http/handler"
)

const AppVersion = "1.0.0"

func main() {
	fmt.Printf("Starting Go Wallet Core Service v%s...\n", AppVersion)

	userRepo := postgres.NewUserRepository()
	walletRepo := postgres.NewWalletRepository()

	createUserUseCase := user.NewCreateUserUseCase(userRepo)
	processTrxUseCase := wallet.NewProcessTransactionUseCase(walletRepo, walletRepo)

	userHandler := handler.NewUserHandler(createUserUseCase)
	walletHandler := handler.NewWalletHandler(processTrxUseCase)
	healthHandler := handler.NewHealthHandler(AppVersion)

	if os.Getenv("SEED_DEFAULT_ACCOUNT") != "false" {
		ctx := context.Background()
		initialAcc, err := entity.NewAccount("ACC-001", "USER-001")
		if err != nil {
			log.Fatalf("failed to seed default account: %v", err)
		}
		if err := initialAcc.UpdateBalance(500.0); err != nil {
			log.Fatalf("failed to set seed account balance: %v", err)
		}
		if err := walletRepo.SaveAccount(ctx, initialAcc); err != nil {
			log.Fatalf("failed to persist seed account: %v", err)
		}
		fmt.Println("[TESTE] Conta default carregada: ACC-001 | Saldo: R$ 500,00")
	}

	http.HandleFunc("/users", userHandler.CreateUser)
	http.HandleFunc("/wallet/transaction", walletHandler.Transaction)
	http.HandleFunc("/health", healthHandler.HealthCheck)

	http.HandleFunc("/swagger.yaml", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "api/swagger.yaml")
	})

	http.HandleFunc("/docs", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "api/index.html")
	})

	server := &http.Server{
		Addr:              ":8080",
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       15 * time.Second,
	}

	fmt.Println("Server listening on :8080")
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server failed to start: %v", err)
	}
}
