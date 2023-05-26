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

	book := skill.NewBook()
	book.Register("lower", lowerSkill.Flatten())
	book.Register("upper", upperSkill.Flatten())

	svc := service.NewCustomSkillService(book)
	svc.Run()
}
