package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/princetheprogrammerbtw/flowforge/internal/config"
	"github.com/princetheprogrammerbtw/flowforge/internal/database"
	"github.com/princetheprogrammerbtw/flowforge/internal/logger"
	"github.com/princetheprogrammerbtw/flowforge/internal/redis"
	"go.uber.org/zap"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	logger.InitLogger(cfg.LogLevel)
	defer logger.Log.Sync()

	pool, err := database.InitDB(cfg.DBURL)
	if err != nil {
		logger.Log.Fatal("Could not connect to database", zap.Error(err))
	}
	defer pool.Close()

	rdb, err := redis.InitRedis(cfg.RedisAddr)
	if err != nil {
		logger.Log.Fatal("Could not connect to redis", zap.Error(err))
	}
	defer rdb.Close()

	r := chi.NewRouter()

	userRepo := repository.NewUserRepository(db.New(pool))
	authHandler := api.NewAuthHandler(userRepo, cfg)

	workflowRepo := repository.NewWorkflowRepository(db.New(pool), database.NewStore(pool))
	workflowHandler := api.NewWorkflowHandler(workflowRepo)

	asynqClient := worker.NewAsynqClient(cfg.RedisAddr)
	distributor := worker.NewTaskDistributor(asynqClient)
	webhookHandler := api.NewWebhookHandler(distributor)

	// Start Worker Server in goroutine
	srvWorker := worker.StartWorkerServer(cfg.RedisAddr)
	engine := worker.NewEngine(db.New(pool), distributor)
	muxWorker := asynq.NewServeMux()
	muxWorker.HandleFunc(worker.TypeExecuteWorkflow, engine.HandleWorkflowExecution)
	muxWorker.HandleFunc(worker.TypeExecuteNode, engine.HandleNodeExecution)
	
	go func() {
		if err := srvWorker.Run(muxWorker); err != nil {
			logger.Log.Fatal("worker server", zap.Error(err))
		}
	}()

	r.Use(middleware.RequestID)
...
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	r.Post("/webhooks/{workflow_id}", webhookHandler.HandleWebhook)

	r.Route("/auth", func(r chi.Router) {
...
		r.Use(middleware.Throttle(10))
		r.Post("/register", authHandler.Register)
		r.Post("/login", authHandler.Login)
		r.Group(func(r chi.Router) {
			r.Use(api.AuthMiddleware(cfg))
			r.Get("/me", authHandler.GetMe)
		})
	})

	r.Route("/workflows", func(r chi.Router) {
		r.Use(api.AuthMiddleware(cfg))
		r.Get("/", workflowHandler.List)
		r.Post("/", workflowHandler.Create)
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", workflowHandler.GetByID)
			r.Delete("/", workflowHandler.Delete)
			r.Get("/canvas", workflowHandler.GetCanvas)
			r.Post("/canvas", workflowHandler.SaveCanvas)
			r.Patch("/toggle", workflowHandler.Toggle)
			r.Get("/history", workflowHandler.GetHistory)
		})
	})

	srv := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: r,
	}

	go func() {
		logger.Log.Info("Server starting", zap.String("port", cfg.Port))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Log.Fatal("listen", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Log.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Log.Fatal("Server forced to shutdown", zap.Error(err))
	}

	logger.Log.Info("Server exited gracefully")
}
