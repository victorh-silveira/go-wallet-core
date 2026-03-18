package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/victor-silveira/go-wallet-core/src/internal/domain/entity"
	"github.com/victor-silveira/go-wallet-core/src/internal/infrastructure/repository/postgres"
	"github.com/victor-silveira/go-wallet-core/src/internal/interface/http/handler"
	"github.com/victor-silveira/go-wallet-core/src/internal/usecase/user"
	"github.com/victor-silveira/go-wallet-core/src/internal/usecase/wallet"
)

const AppVersion = "1.0.0"

func main() {
	fmt.Printf("Starting Go Wallet Core Service v%s...\n", AppVersion)

	userRepo := postgres.NewUserRepository()
	walletRepo := postgres.NewWalletRepository()

	createUserUseCase := usecase.NewCreateUserUseCase(userRepo)
	processTrxUseCase := wallet.NewProcessTransactionUseCase(walletRepo, walletRepo)

	userHandler := handler.NewUserHandler(createUserUseCase)
	walletHandler := handler.NewWalletHandler(processTrxUseCase)
	healthHandler := handler.NewHealthHandler(AppVersion)

	ctx := context.Background()
	initialAcc, _ := entity.NewAccount("ACC-001", "USER-001")
	initialAcc.UpdateBalance(500.0)
	_ = walletRepo.SaveAccount(ctx, initialAcc)

	http.HandleFunc("/users", userHandler.CreateUser)
	http.HandleFunc("/wallet/transaction", walletHandler.Transaction)
	http.HandleFunc("/health", healthHandler.HealthCheck)

	http.HandleFunc("/swagger.yaml", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "api/swagger.yaml")
	})

	http.HandleFunc("/docs", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "api/index.html")
	})

	fmt.Println("Server listening on :8080")
	fmt.Println("[TESTE] Conta default carregada: ACC-001 | Saldo: R$ 500,00")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
