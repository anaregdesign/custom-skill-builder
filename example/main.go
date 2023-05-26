package main

import (
	"github.com/anaregdesign/custom-skill-builder/model"
	"github.com/anaregdesign/custom-skill-builder/service"
	"github.com/anaregdesign/custom-skill-builder/skill"
	"strings"
)

func main() {
	lowerSkill := skill.NewSkill(func(d *model.StringData) (*model.StringData, error) {
		if d.Input == "" {
			return nil, model.ErrInputNotFound
		}
		return &model.StringData{Output: strings.ToLower(d.Input)}, nil
	})

	upperSkill := skill.NewSkill(func(d *model.StringData) (*model.StringData, error) {
		if d.Input == "" {
			return nil, model.ErrInputNotFound
		}
		return &model.StringData{Output: strings.ToUpper(d.Input)}, nil
	})

	splitSkill := skill.NewSkill(func(d *model.StringData) (*model.CollectionStringData, error) {
		if d.Input == "" {
			return nil, model.ErrInputNotFound
		}
		return &model.CollectionStringData{Output: strings.Split(d.Input, " ")}, nil
	})

	countWordsSkill := skill.NewSkill(func(d *model.StringData) (*model.IntData, error) {
		if d.Input == "" {
			return nil, model.ErrInputNotFound
		}
		return &model.IntData{Output: len(strings.Split(d.Input, " "))}, nil
	})

	book := skill.NewBook()
	book.Register("lower", lowerSkill.Flatten())
	book.Register("upper", upperSkill.Flatten())
	book.Register("split", splitSkill.Flatten())
	book.Register("countWords", countWordsSkill.Flatten())

	svc := service.NewCustomSkillService(book)
	svc.Run()
}
