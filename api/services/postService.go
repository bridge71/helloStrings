package services

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"

	"github.com/bridge71/helloStrings/api/configs"
	"github.com/bridge71/helloStrings/api/models"
	"github.com/bridge71/helloStrings/api/repositories"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PostService struct {
	PostRepository *repositories.PostRepository
}

func NewPostService(postRepository *repositories.PostRepository) *PostService {
	return &PostService{PostRepository: postRepository}
}

func (s *PostService) CreatePost(c *gin.Context) (int, models.Message) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println("error json:", err)
		return http.StatusForbidden, models.Message{RetMessage: "error json"}
	}
	post := &models.Post{}
	err = json.Unmarshal(body, post)
	if err != nil {
		fmt.Println("error json:", err)
		return http.StatusForbidden, models.Message{RetMessage: "error json"}
	}
	postContent := &models.PostContent{}
	err = json.Unmarshal(body, postContent)
	if err != nil {
		fmt.Println("error json:", err)
		return http.StatusForbidden, models.Message{RetMessage: "error json"}
	}

	fmt.Println(post)
	// fmt.Println(postContent)

	err = configs.DB.Transaction(func(tx *gorm.DB) error {
		re := regexp.MustCompile(`<img src="data:image/([^;]+);base64,([^"]+)"[^>]*>`)
		matches := re.FindAllStringSubmatch(postContent.Content, -1)

		for _, match := range matches {
			if len(match) < 3 {
				fmt.Println("match error")
				continue
			}

			imageType := match[1]
			base64Data := match[2]
			imgBytes, err := base64.StdEncoding.DecodeString(base64Data)
			if err != nil {
				fmt.Println("base decode error")
				return err
			}

			hash := md5.New()
			hash.Write(imgBytes)
			hashBytes := hash.Sum(nil)
			hashString := hex.EncodeToString(hashBytes)

			fmt.Println(hashString)
			filePath := "/home/bridge71/myTry/contents/" + hashString + "." + imageType
			URL := "http://localhost:7777/contents/" + hashString + "." + imageType
			err = os.WriteFile(filePath, imgBytes, 0666)
			if err != nil {
				fmt.Println(err)
				fmt.Println("write to file error")
				return err
			}

			newTag := fmt.Sprintf(`<n-img src="%s"></n-img>`, URL)
			postContent.Content = re.ReplaceAllString(postContent.Content, newTag)
		}
		err := s.PostRepository.CreateInfo(c, post)
		if err != nil {
			fmt.Println("CreateInfo error")
			return err
		}

		postContent.PostId = post.PostId
		err = s.PostRepository.InsertContent(c, postContent)
		if err != nil {
			fmt.Println("InsertContent error")
			return err
		}

		return nil
	})
	if err != nil {
		return http.StatusInternalServerError, models.Message{
			RetMessage: "something unusual happened when insert post into database",
		}
	}
	return http.StatusOK, models.Message{
		RetMessage: "post successfully",
	}
}
