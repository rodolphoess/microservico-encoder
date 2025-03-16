package services

import (
	"domain/application/repositories"
	"domain/domain"
	"domain/framework/database"
	"github.com/joho/godotenv"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"log"
	"testing"
	"time"
)

func init() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func prepare() (*domain.Video, repositories.VideoRepository) {
	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "test.mp4"
	video.CreatedAt = time.Now()

	repo := repositories.VideoRepositoryDB{Db: db}

	return video, repo
}

func TestVideoServiceDownload(t *testing.T) {
	video, repo := prepare()

	videoService := NewVideoService()
	videoService.Video = video
	videoService.VideoRepository = repo

	err := videoService.Download("codeeducationrodolphotest")
	require.Nil(t, err)

	err = videoService.Fragment()
	require.Nil(t, err)
}
