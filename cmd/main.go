package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	peep "peep/pkg"
	"peep/pkg/beers"
	"peep/pkg/log/logrus"
	"peep/pkg/server"
	mysql "peep/pkg/storage/mysql"
	"strconv"
)

func main() {
	var (
		hostName, _     = os.Hostname()
		defaultServerID = fmt.Sprintf("%s-%s", os.Getenv("API_NAME"), hostName)
		defaultHost     = os.Getenv("API_SERVER_HOST")
		defaultPort, _  = strconv.Atoi(os.Getenv("API_SERVER_PORT"))
	)

	fmt.Println("hostname is:", hostName)
	fmt.Println("defaultServerID:", defaultServerID)
	fmt.Println("defaultHost:", defaultHost)
	fmt.Println("defaultPort:", defaultPort)

	host := flag.String("host", defaultHost, "define host of the server")
	port := flag.Int("port", defaultPort, "define port of the server")
	serverID := flag.String("server-id", defaultServerID, "define server identifier")
	flag.Parse()

	logger := logrus.NewLogger()

	repo := newMySQLRepository()
	beerServices := beers.NewService(repo, logger)

	httpAddr := fmt.Sprintf("%s:%d", *host, *port)

	s := server.New(
		*serverID,
		beerServices,
	)

	fmt.Println("The server is on tap now:", httpAddr)
	log.Fatal(http.ListenAndServe(httpAddr, s.Router()))
}

func newMySQLRepository() peep.Repository {
	mysqlConn, err := mysql.NewConn()
	if err != nil {
		log.Fatal(err)
	}
	return mysql.NewRepository("gophers", mysqlConn)
}
