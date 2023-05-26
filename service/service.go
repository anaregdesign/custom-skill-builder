package service

import (
	"github.com/anaregdesign/custom-skill-builder/skill"
	"github.com/gin-gonic/gin"
	"io"
)

type CustomSkillService struct {
	book *skill.Book
}

func NewCustomSkillService(book *skill.Book) *CustomSkillService {
	return &CustomSkillService{book: book}
}

func (s *CustomSkillService) Run() error {
	r := gin.Default()

	r.POST("/skills/:name", func(c *gin.Context) {
		name := c.Param("name")
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		result, err := s.book.Apply(name, body)
		if err == skill.ErrNotFound {
			c.JSON(404, gin.H{"error": err.Error()})
			return
		} else if err == skill.ErrParse {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		} else if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.Data(200, "application/json", result)
	})
	return r.Run()
}
