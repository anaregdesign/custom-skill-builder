package main

import (
	"github.com/anaregdesign/custom-skill-builder/service"
	"github.com/anaregdesign/custom-skill-builder/skill"
	"strings"
)

func main() {
	lowerSkill := skill.NewSkillNoErr("content", "lowerContent", strings.ToLower)
	upperSkill := skill.NewSkillNoErr("content", "upperContent", strings.ToUpper)
	splitSkill := skill.NewSkillNoErr("content", "splitContent", func(s string) []string {
		return strings.Split(s, " ")
	})
	wordCountSkill := skill.NewSkillNoErr("content", "wordCount", func(s string) map[string]int {
		result := make(map[string]int)
		for _, word := range strings.Split(s, " ") {
			result[word]++
		}
		return result
	})

	book := skill.NewBook()
	book.Register("lower", lowerSkill.Flatten())
	book.Register("upper", upperSkill.Flatten())
	book.Register("split", splitSkill.Flatten())
	book.Register("wordcount", wordCountSkill.Flatten())

	svc := service.NewCustomSkillService(book)
	svc.Run()
}
