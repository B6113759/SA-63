package main

import (
	"context"
	"fmt"
	"log"

	"github.com/B6113759/app/controllers"
	_ "github.com/B6113759/app/docs"
	"github.com/B6113759/app/ent"
	"github.com/B6113759/app/ent/coolroomtype"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Users struct {
	User []User
}

type User struct {
	USERNAME  string
	USEREMAIL string
}

type Patients struct {
	Patient []Patient
}

type Patient struct {
	PATIENTNAME string
	PATIENTAGE  int
}

type Coolrooms struct {
	Coolroom []Coolroom
}

type Coolroom struct {
	COOLROOMNAME   string
	COOLROOMTYPEID int
}

type CoolroomTypes struct {
	CoolroomType []CoolroomType
}

type CoolroomType struct {
	COOLROOMTYPE_NAME string
}

type Relatives struct {
	Relative []Relative
}

type Relative struct {
	RELATIVENAME string
}

// @title SUT SA Example API
// @version 1.0
// @description This is a sample server for SUT SE 2563
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationUrl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationUrl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information
func main() {
	router := gin.Default()
	router.Use(cors.Default())

	client, err := ent.Open("sqlite3", "file:ent.db?cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("fail to open sqlite3: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	v1 := router.Group("/api/v1")
	controllers.NewUserController(v1, client)
	controllers.NewRelativeController(v1, client)
	controllers.NewCoolroomController(v1, client)
	controllers.NewCoolroomTypeController(v1, client)
	controllers.NewDeceasedReceiveController(v1, client)
	controllers.NewPatientController(v1, client)

	// Set Users Data
	users := Users{
		User: []User{
			User{"นพ.กำลัง นอนพอดี", "kumlang@gmail.com"},
			User{"นพ.รักษา อาการ", "raksaa@example.com"},
		},
	}

	for _, u := range users.User {
		client.User.
			Create().
			SetUseremail(u.USEREMAIL).
			SetUserName(u.USERNAME).
			Save(context.Background())
	}

	// Set Patients Data
	patients := Patients{
		Patient: []Patient{
			Patient{"นาย โชคช่วย ไม่ช่วย", 42},
			Patient{"นาย แกง หม้อใหญ่", 28},
			Patient{"นาย รักษา เป็นโรค", 35},
		},
	}

	for _, p := range patients.Patient {
		client.Patient.
			Create().
			SetPatientName(p.PATIENTNAME).
			SetPatientAge(p.PATIENTAGE).
			Save(context.Background())
	}

	// Set CoolroomTypes Data
	coolroomtypes := []string{"ศพรอร่างจากญาติ", "อาจารย์ใหญ่"}
	for _, ct := range coolroomtypes {
		client.CoolroomType.
			Create().
			SetCoolroomtypeName(ct).
			Save(context.Background())
	}

	// Set Coolrooms Data
	coolrooms := Coolrooms{
		Coolroom: []Coolroom{
			Coolroom{"001", 1},
			Coolroom{"002", 1},
			Coolroom{"003", 2},
		},
	}

	for _, cr := range coolrooms.Coolroom {

		ct, err := client.CoolroomType.
			Query().
			Where(coolroomtype.IDEQ(int(cr.COOLROOMTYPEID))).
			Only(context.Background())

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		client.Coolroom.
			Create().
			SetCoolroomName(cr.COOLROOMNAME).
			SetCoolroomtype(ct).
			Save(context.Background())
	}

	// Set Relatives Data
	relatives := []string{"ขจร กำลังยืน", "ณยศ กำลังตื่น"}
	for _, r := range relatives {
		client.Relative.
			Create().
			SetRelativeName(r).
			Save(context.Background())
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()
}
