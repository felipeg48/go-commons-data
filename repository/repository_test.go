package repository

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"testing"
	"time"
)


type User struct {
	Id		uuid.UUID	`json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Name 	string		`json:"name"`
	Email 	string		`json:"email"`

	CreatedAt             time.Time           `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt             time.Time           `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	//DeletedAt             gorm.DeletedAt      `json:"deleted_at"`
}

var (
	ID uuid.UUID
	REPO *CrudRepository
	USER1 *User = &User{
		Name:  "Felipe",
		Email: "felipeg@email.com",
	}

	USER2 *User = &User{
		Name:  "Emilio",
		Email: "emilio@email.com",
	}

	USER3 *User = &User{
		Name:  "Laura",
		Email: "laura@email.com",
	}

	USER4 *User = &User{
		Name:  "Misael",
		Email: "misael@email.com",
	}
)

func TestNewCrudRepository(t *testing.T) {

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:              time.Second,   // Slow SQL threshold
			LogLevel:                   logger.Info, // Log level
			IgnoreRecordNotFoundError: true,           // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,          // Disable color
		},
	)

	dsn := os.ExpandEnv("host=${DBHOST} user=${DBUSER} password=${DBPASSWORD} dbname=${DBNAME} port=${DBPORT} sslmode=disable")
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	db.Migrator().DropTable(User{})
	db.AutoMigrate(User{})

	REPO = NewCrudRepository(db, User{}, uuid.UUID{})

	assert.NotNil(t, db)
	assert.NotNil(t, REPO)
}

func TestCrudRepository_Save(t *testing.T) {
	result, error := REPO.Save(USER1)

	assert.NoError(t, error)
	assert.NotNil(t,result)

	if user, ok := result.(*User); ok {
		assert.NotNil(t, user.Id)
		ID = user.Id
	}

	result, error = REPO.Save(USER2)

	assert.NoError(t, error)
	assert.NotNil(t,result)

	if user, ok := result.(*User); ok {
		assert.NotNil(t, user.Id)
	}else{
		t.Fail()
	}
}

func TestCrudRepository_FindById(t *testing.T) {
	result,error := REPO.FindById(ID)

	assert.NoError(t, error)
	assert.NotNil(t,result)

	if user, ok := result.(*User); ok {
		assert.NotNil(t, user.Id)
		assert.Equal(t, USER1.Id, user.Id)
	}else{
		t.Fail()
	}
}

func TestCrudRepository_FindAll(t *testing.T) {
	result,error := REPO.FindAll()

	assert.NoError(t, error)
	assert.NotNil(t,result)

	log.Printf(">>> %v",result)
}

func TestCrudRepository_DeleteById(t *testing.T) {
	error := REPO.DeleteById(USER1.Id)
	assert.NoError(t, error)
}

func TestCrudRepository_Status(t *testing.T) {
	error := REPO.Status()
	assert.NoError(t, error)
}



