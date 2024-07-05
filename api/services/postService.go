package services

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"html"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"

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

func (s *PostService) PostCreate(c *gin.Context) (int, models.Message) {
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
	post.UserId = GetUserId(c)
	post.Nickname = GetNickname(c)
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

		for index, match := range matches {
			fmt.Println("index ", index)
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
			URL := "http://localhost:7777/static/" + hashString + "." + imageType
			// URL := "/home/bridge71/myTry/contents/" + hashString + "." + imageType
			// URL := "https://store.ymgal.games/topic/content/48/48ac6b8e80c7453cb0d7c0905a85d878.jpg"
			err = os.WriteFile(filePath, imgBytes, 0666)
			if err != nil {
				fmt.Println(err)
				fmt.Println("write to file error")
				return err
			}

			newTag := fmt.Sprintf(`<img width="720" src="%s">`, URL)
			fmt.Println("newTag", newTag)
			postContent.Content = re.ReplaceAllStringFunc(postContent.Content, func(s string) string {
				if s == match[0] {
					return newTag
				}
				return s
			})
			// postContent.Content = re.ReplaceAllString(postContent.Content, newTag)
		}
		// postContent.Content = strings.ReplaceAll(html.EscapeString(postContent.Content), "\n", "<br>")
		err := s.PostRepository.PostCreate(c, post)
		if err != nil {
			fmt.Println("CreateInfo error")
			return err
		}

		postContent.PostId = post.PostId
		err = s.PostRepository.ContentCreate(c, postContent)
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

func (s *PostService) CommentCreate(c *gin.Context) (int, models.Message) {
	comment := &models.Comment{}
	err := c.ShouldBindJSON(comment)
	if err != nil {
		return http.StatusInternalServerError, models.Message{
			RetMessage: "error bind of post",
		}
	}

	comment.UserId = GetUserId(c)
	comment.Nickname = GetNickname(c)
	comment.Content = strings.ReplaceAll(html.EscapeString(comment.Content), "\n", "<br>")
	err = configs.DB.Transaction(func(tx *gorm.DB) error {
		err := s.PostRepository.CommentCreate(c, comment)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return http.StatusInternalServerError, models.Message{
			RetMessage: "something unusual happened when insert comment into database",
		}
	}
	return http.StatusOK, models.Message{
		RetMessage: "comment successfully",
	}
}

func (s *PostService) PostFetch(c *gin.Context) (int, models.Message) {
	var posts []models.Post
	s.PostRepository.PostFetch(c, &posts)
	return http.StatusOK, models.Message{
		RetMessage: "get post successfully",
		Post:       posts,
	}
}

func (s *PostService) CommentReadPostId(c *gin.Context) (int, models.Message) {
	post := &models.Post{}
	err := c.ShouldBindJSON(post)
	if err != nil {
		return http.StatusForbidden, models.Message{RetMessage: "Bind error"}
	}
	var comments []models.Comment
	s.PostRepository.CommentReadPostId(c, &comments, post.PostId)
	return http.StatusOK, models.Message{
		RetMessage: "get post successfully",
		Comment:    comments,
	}
}

func (s *PostService) CommentReadUserId(c *gin.Context) (int, models.Message) {
	user := &models.User{}
	err := c.ShouldBindJSON(user)
	if err != nil {
		return http.StatusForbidden, models.Message{RetMessage: "Bind error"}
	}
	var comments []models.Comment
	s.PostRepository.CommentReadUserId(c, &comments, user.UserId)
	return http.StatusOK, models.Message{
		RetMessage: "get post successfully",
		Comment:    comments,
	}
}

func (s *PostService) ContentReadPostId(c *gin.Context) (int, models.Message) {
	post := &models.Post{}
	err := c.ShouldBindJSON(post)
	if err != nil {
		return http.StatusForbidden, models.Message{RetMessage: "Bind error"}
	}
	postContent := &models.PostContent{}
	s.PostRepository.ContentReadPostId(c, postContent, post.PostId)
	return http.StatusOK, models.Message{
		RetMessage:  "get post successfully",
		PostContent: *postContent,
	}
}

func (s *PostService) PostReadTitle(c *gin.Context) (int, models.Message) {
	post := &models.Post{}
	err := c.ShouldBindJSON(post)
	if err != nil {
		fmt.Println(err)
		return http.StatusForbidden, models.Message{RetMessage: "Bind error"}
	}
	var posts []models.Post
	s.PostRepository.PostReadTitle(c, &posts, post.Title)
	return http.StatusOK, models.Message{
		RetMessage: "get post successfully",
		Post:       posts,
	}
}

func (s *PostService) PostReadNickname(c *gin.Context) (int, models.Message) {
	post := &models.Post{}
	err := c.ShouldBindJSON(post)
	if err != nil {
		fmt.Println(err)
		return http.StatusForbidden, models.Message{RetMessage: "Bind error"}
	}
	var posts []models.Post
	if post.Nickname == "" {
		post.Nickname = GetNickname(c)
	}
	s.PostRepository.PostReadNickname(c, &posts, post.Nickname)
	return http.StatusOK, models.Message{
		RetMessage: "get post successfully",
		Post:       posts,
	}
}

func (s *PostService) LikesChange(c *gin.Context) (int, models.Message) {
	post := &models.Post{}
	err := c.ShouldBindJSON(post)
	if err != nil {
		fmt.Println(err)
		return http.StatusForbidden, models.Message{RetMessage: "post bind error"}
	}
	userId := GetUserId(c)
	likes := &models.Likes{
		UserId: userId,
		PostId: post.PostId,
	}
	fmt.Println("likes count", post.Likes)
	been := false
	err = configs.DB.Transaction(func(tx *gorm.DB) error {
		likesGet := &models.Likes{}
		s.PostRepository.LikeQuery(c, likesGet, userId, post.PostId)
		var err error
		if likesGet.PostId != 0 {
			err = s.PostRepository.LikeDel(c, likes)
			if err != nil {
				return err
			}
			err = s.PostRepository.LikeUpdateDown(c, post)
			been = true
		} else {
			err = s.PostRepository.LikeCreate(c, likes)
			if err != nil {
				return err
			}
			err = s.PostRepository.LikeUpdateUp(c, post)
		}
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return http.StatusInternalServerError, models.Message{
			RetMessage: "something unusual happened when change likes",
		}
	}

	var message string
	if been {
		message = "已取消"
	} else {
		message = "添加成功"
	}
	return http.StatusOK, models.Message{
		RetMessage: message,
	}
}

func (s *PostService) LikesReadUserId(c *gin.Context) (int, models.Message) {
	userId := GetUserId(c)
	var likes []models.Likes
	s.PostRepository.LikeReadUserId(c, &likes, userId)
	if len(likes) == 0 {
		return http.StatusOK, models.Message{
			RetMessage: "ok",
		}
	}
	postId := make([]uint, len(likes))
	for index, like := range likes {
		postId[index] = like.PostId
	}
	var post []models.Post

	s.PostRepository.PostReadId(c, &post, postId)

	return http.StatusOK, models.Message{
		RetMessage: "ok",
		Post:       post,
	}
}

func (s *PostService) CommentMarkCreate(c *gin.Context) (int, models.Message) {
	post := &models.Post{}
	err := c.ShouldBindJSON(post)
	if err != nil {
		fmt.Println(err)
		return http.StatusForbidden, models.Message{RetMessage: "post bind error"}
	}
	userId := GetUserId(c)
	commentMark := &models.CommentMark{
		UserId: userId,
		PostId: post.PostId,
	}
	s.PostRepository.CommentMarkCreate(c, commentMark)

	return http.StatusOK, models.Message{
		RetMessage: "comments",
	}
}

func (s *PostService) CommentMarkReadUserId(c *gin.Context) (int, models.Message) {
	userId := GetUserId(c)
	var commentMark []models.CommentMark
	s.PostRepository.CommentMarkReadUserId(c, &commentMark, userId)

	if len(commentMark) == 0 {
		return http.StatusOK, models.Message{
			RetMessage: "ok",
		}
	}
	postId := make([]uint, len(commentMark))
	for index, comment := range commentMark {
		postId[index] = comment.PostId
		fmt.Println(comment.PostId)
	}
	var post []models.Post

	s.PostRepository.PostReadId(c, &post, postId)

	return http.StatusOK, models.Message{
		RetMessage: "ok",
		Post:       post,
	}
}
