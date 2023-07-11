package crossingLevels

import (
	"fmt"
	"go.uber.org/dig"
)

type DatabaseConnection struct {
	ConnectionString string
}

type UserRepository struct {
	Database *DatabaseConnection
}

func NewDatabaseConnection() *DatabaseConnection {
	return &DatabaseConnection{
		ConnectionString: "localhost:5432",
	}
}

func NewUserRepository(database *DatabaseConnection) *UserRepository {
	return &UserRepository{
		Database: database,
	}
}

func GetUser(database *DatabaseConnection, userRepository *UserRepository) {
	fmt.Println("User:", userRepository)
	fmt.Println("Database:", database)
}

func T2() {
	container := dig.New()

	err := container.Provide(NewDatabaseConnection)
	if err != nil {
		fmt.Println("Error providing DatabaseConnection:", err)
		return
	}

	err = container.Provide(NewUserRepository)
	if err != nil {
		fmt.Println("Error providing UserRepository:", err)
		return
	}

	err = container.Invoke(GetUser)
	if err != nil {
		fmt.Println("Error invoking GetUser:", err)
		return
	}
}
