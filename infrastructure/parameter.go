package infrastructure

import (
	"math/rand"
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
)

const (
	HTTP_SWAGGER      = "HTTP_SWAGGER"
	HOST_PORT         = "HOST_PORT"
	DBNAME            = "DATABASE_NAME"
	DB_HOST_PORT      = "DATABASE_HOST_PORT"
	POST_COLLECTION   = "POST_COLLECTION"
	BANNER_COLLECTION = "BANNER_COLLECTION"
	POPUP_COLLECTION  = "POPUP_COLLECTION"
)

var (
	// develop environment
	DevelopEnv       string
	HostPort         string
	HttpSwagger      string
	DatabaseName     string
	DBHostPort       string
	PostCollection   string
	BannerCollection string
	PopupCollection  string
)

func getStringEnvParameter(envParam string, defaultValue string) string {
	var val string
	if value, ok := os.LookupEnv(envParam); ok {
		val = value
	} else {
		val = defaultValue
	}
	InfoLog.Printf("%s : %s ", envParam, val)
	return val

}

func randomNATSClientID() string {
	return "test" + strconv.Itoa(int(rand.Uint64()))
}

var randD = randomNATSClientID()

func loadEnvParameters() {
	HostPort = getStringEnvParameter(HOST_PORT, ":3000")
	HttpSwagger = getStringEnvParameter(HTTP_SWAGGER, "http://localhost:3000/api/v1/content/swagger/doc.json")
	DatabaseName = getStringEnvParameter(DBNAME, "dev-sanvt-content")
	DBHostPort = getStringEnvParameter(DB_HOST_PORT, "")
	PostCollection = getStringEnvParameter(POST_COLLECTION, "posts")
	BannerCollection = getStringEnvParameter(BANNER_COLLECTION, "banners")
	PopupCollection = getStringEnvParameter(POPUP_COLLECTION, "popups")

}
