package SS2023CI_02_MS_Go

import "os"

func main() {
	a := App{}
	a.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"))

	//export APP_DB_USERNAME=postgres
	//export APP_DB_PASSWORD=<whatever password you use>
	//export APP_DB_NAME=postgres

	a.Run(":8010")
}
