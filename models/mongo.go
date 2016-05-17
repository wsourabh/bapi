package models
import(
"gopkg.in/mgo.v2"
	"fmt"
)

var (
	// Reference to the singleton.
	singleton mongoManager
)

type (
	// mongoManager manages a map of session.
	mongoManager struct {
		sessions map[string]*mgo.Session
	}
)

func CreateSession( host string, databaseName string) error {

	session, err := mgo.Dial(host)
	// Establish the master session.
	if err != nil {
		return err
	}
	// Add the database to the map.
	singleton = mongoManager{
		sessions: make(map[string]*mgo.Session),
	}
	singleton.sessions["dev"] = session

	return nil
}

func init(){
	if err := CreateSession("127.0.0.1", "oz_development"); err != nil {
		fmt.Println("error while connecting")
	}
}