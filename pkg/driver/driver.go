package driver

import (
	"exam_service/pkg/domain/models"
	"github.com/go-redis/redis"
	"github.com/nitishm/go-rejson"
	"github.com/spf13/viper"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

func GetDbConnection() (*redis.Client, *rejson.Handler, *gorm.DB) {
	redisDb, redisJsonDb := getRedisDbConnection()
	return redisDb, redisJsonDb, getPGDbConnetion()
}

func getRedisDbConnection() (*redis.Client, *rejson.Handler) {
	redisDb := redis.NewClient(&redis.Options{
		Addr:     viper.GetString("REDIS_DB_HOST") + viper.GetString("REDIS_DB_PORT"),
		Password: viper.GetString("REDIS_DB_PASSWORD"),
		DB:       viper.GetInt("REDIS_DB_NAME"),
	})
	_, err := redisDb.Ping().Result()
	if err != nil {
		log.Println("Can't connect with Redis database", "Errors:", err)
	} else {
		log.Println("Redis Db is connected", "Errors:", err)
	}
	redisJsonDb := rejson.NewReJSONHandler()
	redisJsonDb.SetGoRedisClient(redisDb)

	return redisDb, redisJsonDb
}

func getPGDbConnetion() *gorm.DB {
	gormDb, err := gorm.Open(sqlite.Open("studentGrade.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	err = gormDb.AutoMigrate(&models.StudentGrade{}, &models.Report{})
	if err != nil {
		panic(err)
	}

	//dataSourceName := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s",
	//	viper.Get("PG_DB_HOST"),
	//	viper.Get("PG_DB_PORT"),
	//	viper.Get("PG_DB_NAME"),
	//	viper.Get("PG_DB_USER"),
	//	viper.Get("PG_DB_PASSWORD"))
	//pgDb, err := sql.Open("pgx", dataSourceName)
	//if err != nil {
	//	log.Fatal(fmt.Sprintf("unable to conect to db"))
	//	panic(err)
	//}
	//log.Println("pinged pg db")
	//
	//gormDB, err := gorm.Open("postgres", pgDb)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("pg db connected successfully")
	////gormDB.DropTableIfExists(&models.User{}, &models.Type{}) // For development purposes
	//err = gormDB.AutoMigrate(&models.StudentGrade{}, &models.Report{})
	//if err != nil {
	//	log.Println(err)
	//}
	//
	return gormDb
}
