package conf

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	MaxIdleConns    int
	MaxOpenConns    int
	ShutdownTimeout int
	ResponseTimeOut int
	RequestTimeOut  int
	BookAPIPort     string
	AdminAPIPort    string
	AuthAPIPort     string
	BrokerAPIPort   string
	ExecutorAPIPort string
	CreatorAPIPort  string
	ExamAPIPort     string
	GRPCPort        string
	AuthGRPCPort    string
	AdminGRPCPort   string
	DebugPort       string
	DBDSN           string
	DBName          string
	User            string
	Password        string
	Host            string
	Environment     string
	MongoHost       string
	MongoPort       string
	MongoUser       string
	MongoPassword   string
	MongoAuth       string
	MongoDbName     string
	Language        string
	JWTKey          string
	TracerHost      string
	RabbitQueues    string
	RabbitURL       string
	TracerProb      float64
	AllowDirect     bool
	AllowGRPC       bool
}

// LoadConfig loads configuration from environment variables and optional .env file
func LoadConfig() (*Config, error) {
	// Load .env file if it exists
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables only")
	}

	config := &Config{
		BookAPIPort:     getEnv("BOOKAPI_PORT", ":8000"),
		AuthAPIPort:     getEnv("AUTHAPI_PORT", ":8001"),
		DebugPort:       getEnv("DEBUG_PORT", ":8002"),
		BrokerAPIPort:   getEnv("BROKER_PORT", ":8003"),
		ExecutorAPIPort: getEnv("EXECUTOR_PORT", ":8004"),
		CreatorAPIPort:  getEnv("CREATOR_PORT", ":8005"),
		AdminAPIPort:    getEnv("ADMIN_PORT", ":8006"),
		ExamAPIPort:     getEnv("EXAM_PORT", ":8008"),
		GRPCPort:        getEnv("GRPC_PORT", ":50001"),
		AuthGRPCPort:    getEnv("AUTH_GRPC_PORT", ":50002"),
		AdminGRPCPort:   getEnv("ADMIN_GRPC_PORT", ":50003"),
		DBDSN:           getEnv("DBDSN", "postgres://epic:password@localhost:5432"),
		User:            getEnv("DBUSER", "epic"),
		Password:        getEnv("DBPASSWORD", "admin#123"),
		Host:            getEnv("HOST", "localhost"),
		DBName:          getEnv("DBNAME", "epic"),
		Environment:     getEnv("ENV", "devlopment"),
		MongoHost:       getEnv("MONGO_HOST", "localhost"),
		MongoPort:       getEnv("MONGO_PORT", "27017"),
		MongoUser:       getEnv("MONGO_USER", "admin"),
		MongoPassword:   getEnv("MONGI_PASS", "admin#123"),
		MongoAuth:       getEnv("MONGO_AUTH", "admin"),
		MongoDbName:     getEnv("MONGO_DBNAME", "EXECUTOR"),
		Language:        getEnv("CONTAINER_LANGUAGE", "python"),
		JWTKey:          getEnv("JWT_KEY", "mysupersecret"),
		TracerHost:      getEnv("TRACER_HOST", "http://localhost:14268/api/traces"),
		RabbitQueues:    getEnv("RABBIT_QUEUES", "programming_questions,challenge_data,code_execution_stats"),
		RabbitURL:       getEnv("RABBIT_URL", "amqp://guest:guest@localhost:5672/"),
		// TracerProb:      getEnv("TRACER_PROB", "mysupersecret"),
		// AllowDirect:   getEnv("ENV", "devlopment"),
	}
	idleConns, _ := strconv.Atoi(getEnv("MAXIDLECONNS", "10"))
	openConns, _ := strconv.Atoi(getEnv("MAXOPENCONNS", "10"))
	shutdownTimeout, _ := strconv.Atoi(getEnv("SHUTDOWN_TIMEOUT", "5"))
	requestTimeout, _ := strconv.Atoi(getEnv("REQUESTTIMEOUT_SEC", "1"))
	responseTimeout, _ := strconv.Atoi(getEnv("RESPONSETIMEOUT_SEC", "1"))
	allowDirect, _ := strconv.ParseBool(getEnv("ALLOW_DIRECT", "true"))
	allowGRPC, _ := strconv.ParseBool(getEnv("ALLOW_GRPC", "true"))
	tracerProb, _ := strconv.ParseFloat(getEnv("TRACER_PROB", "0.5"), 64)
	config.AllowGRPC = allowGRPC
	config.AllowDirect = allowDirect
	config.MaxIdleConns = idleConns
	config.MaxOpenConns = openConns
	config.ResponseTimeOut = responseTimeout
	config.RequestTimeOut = requestTimeout
	config.ShutdownTimeout = shutdownTimeout
	config.TracerProb = tracerProb
	return config, nil
}

// Helper function to get environment variables with a fallback default value
func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
