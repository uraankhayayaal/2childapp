package seed

import (
	"log"
	"time"

	"github.com/uraankhayayaal/2childapp/models"
	"gorm.io/gorm"
)

var users = []models.User{
	models.User{
		Name:             "Ayaal",
		Email:            "uraankhayayaal@yandex.ru",
		Password:         "$2a$10$qND/uzaiU2SXPoEjQtdOAObGjz2AaX99Zpgxm5foLOCLi809bJUkm", // 00000000
		Role:             "user",
		Photo:            "https://i.pravatar.cc/300?img=30",
		VerificationCode: "",
		Verified:         true,
		Provider:         "local",
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	},
	models.User{
		Name:             "NotVerifiedUser",
		Email:            "not_verified_user@example.ru",
		Password:         "$2a$10$qND/uzaiU2SXPoEjQtdOAObGjz2AaX99Zpgxm5foLOCLi809bJUkm", // 00000000
		Role:             "user",
		Photo:            "https://i.pravatar.cc/300?img=31",
		VerificationCode: "MUEyMWVUTXJkM0lzMjhWa0ZCQlM=", // 1A21eTMrd3Is28VkFBBS
		Verified:         false,
		Provider:         "local",
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	},
}

var posts = []models.Post{
	models.Post{
		Title:     "Title 1",
		Content:   "Hello world 1",
		Image:     "https://i.pravatar.cc/300?img=61",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	models.Post{
		Title:     "Title 2",
		Content:   "Hello world 2",
		Image:     "https://i.pravatar.cc/300?img=62",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
}

func Load(db *gorm.DB) {
	db.Debug().Migrator().DropTable(&models.Post{}, &models.User{})
	db.Debug().AutoMigrate(&models.User{}, &models.Post{})
	db.Debug().Migrator().CreateConstraint(&models.Post{}, "User")
	db.Debug().Migrator().CreateConstraint(&models.Post{}, "fk_posts_id")
	for i := range users {
		userId := db.Debug().Model(&models.User{}).Create(&users[i])
		if userId == nil {
			log.Fatalf("cannot seed users table: %v", userId)
		}
		posts[i].User = users[i].ID

		postId := db.Debug().Model(&models.Post{}).Create(&posts[i])
		if postId == nil {
			log.Fatalf("cannot seed posts table: %v", postId)
		}
	}
}
