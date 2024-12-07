package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/bentenison/microservice/foundation/logger"
	"github.com/gin-gonic/gin"
)

func formatCode(lang, code string) (string, error) {
	var cmd *exec.Cmd
	var out, errb bytes.Buffer

	switch lang {
	case "go":
		cmd = exec.Command("gofmt")
		cmd.Stdin = strings.NewReader(code)
	case "java":
		cmd = exec.Command("google-java-format", "--replace", "-")
		cmd.Stdin = strings.NewReader(code)
	case "c", "cpp":
		cmd = exec.Command("clang-format")
		cmd.Stdin = strings.NewReader(code)
	case "php":
		cmd = exec.Command("php-cs-fixer", "fix", "--format=none", "--dry-run", "--diff")
		cmd.Stdin = strings.NewReader(code)
	case "python":
		cmd = exec.Command("black", "--quiet", "-")
		cmd.Stdin = strings.NewReader(code)
	default:
		return "", fmt.Errorf("unsupported language: %s", lang)
	}

	// Run the command and capture its output
	cmd.Stdout = &out
	cmd.Stderr = &errb

	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("error formatting code: %v, %s", err, errb.String())
	}

	return out.String(), nil
}

type FormatPayload struct {
	Language string `json:"language,omitempty" db:"language"`
	Code     string `json:"code,omitempty" db:"code"`
}

func formatHandler(c *gin.Context) {

	var payload FormatPayload
	if err := c.Bind(&payload); err != nil {
		log.Println("error binding data", err.Error())
		c.JSON(http.StatusExpectationFailed, gin.H{"error": err.Error()})
		return
	}
	// 	lang := "python"
	// 	code := `import time

	// start_time = time.time()

	// # printing all even numbers till 20
	// for i in range(20):
	//   if i % 2 == 0:
	//     print(i, end = " ")

	// end_time = time.time()
	// time_taken = end_time - start_time
	// print("\nTime: ", time_taken)`
	// Format the code using the formatCode function
	formattedCode, err := formatCode(payload.Language, payload.Code)
	if err != nil {
		log.Println("error binding data", err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"formatted_code": formattedCode})
}

func main() {
	log := logger.NewCustomLogger(map[string]interface{}{
		"service": "formatter-tool",
		"env":     "production",
		"build":   "1.0.0",
	})
	// Initialize the Gin router
	r := gin.Default()

	// Define the route for code formatting
	r.POST("/format", formatHandler)

	if err := run(log, r); err != nil {
		log.Errorc(context.TODO(), "error while running server", map[string]interface{}{
			"error": err.Error(),
		})
	}
}
func run(log *logger.CustomLogger, r *gin.Engine) error {
	api := http.Server{
		Addr:    ":8010",
		Handler: r,
		// ReadTimeout:  cfg.Web.ReadTimeout,
		// WriteTimeout: cfg.,
		// IdleTimeout:  cfg.Web.IdleTimeout,
		// ErrorLog: lo,
	}

	// Start the server
	serverErrors := make(chan error, 1)
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)
	ctx := context.Background()
	go func() {
		log.Infoc(context.TODO(), "formatter-api router started", map[string]interface{}{
			"port": ":8010",
		})
		serverErrors <- api.ListenAndServe()
	}()

	// // -------------------------------------------------------------------------
	// // Shutdown

	select {
	case err := <-serverErrors:
		return fmt.Errorf("server error: %w", err)

	case sig := <-shutdown:
		log.Infoc(context.TODO(), "shutdown started", map[string]interface{}{
			"signal": sig,
		})
		defer log.Infoc(context.TODO(), "shutdown completed", map[string]interface{}{})

		ctx, cancel := context.WithTimeout(ctx, time.Duration(5*time.Second))
		defer cancel()

		if err := api.Shutdown(ctx); err != nil {
			api.Close()
			return fmt.Errorf("could not stop server gracefully: %w", err)
		}
	}
	return nil
}
