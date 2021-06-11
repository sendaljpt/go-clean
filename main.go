package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/spf13/viper"

	memberHttpDelivery "github.com/sendaljpt/subscription-service/src/member/delivery/http"
	memberHttpDeliveryMiddleware "github.com/sendaljpt/subscription-service/src/member/delivery/http/middleware"
	memberRepository "github.com/sendaljpt/subscription-service/src/member/repository/mysql"
	memberUsecase "github.com/sendaljpt/subscription-service/src/member/usecase"
)

// load config
func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Subscription service debug mode on!")
	}
}

func main() {
	// database connection
	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPswd := viper.GetString(`database.pswd`)
	dbName := viper.GetString(`database.name`)

	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPswd, dbHost, dbPort, dbName)
	dbConn, err := sql.Open("mysql", connection)

	if err != nil {
		log.Fatal(err)
	}

	err = dbConn.Ping()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	e := echo.New()

	memberMiddleware := memberHttpDeliveryMiddleware.InitMiddleware()
	e.Use(memberMiddleware.CORS)
	memberRepo := memberRepository.NewMysqlMemberRepository(dbConn)

	// timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	mbrUsecase := memberUsecase.NewMemberUsecase(memberRepo)
	memberHttpDelivery.NewMemberHandler(e, mbrUsecase)

	log.Fatal(e.Start(viper.GetString("server.port")))

}
