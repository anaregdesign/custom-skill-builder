package main

import (
	"github.com/anaregdesign/custom-skill-builder/model"
	"github.com/anaregdesign/custom-skill-builder/service"
	"github.com/anaregdesign/custom-skill-builder/skill"
	"strings"
)

func main() {
	lowerSkill := skill.NewSkill(func(d *model.Data) (*model.Data, error) {
		if d.Input == "" {
			return nil, model.ErrInputNotFound
		}
		return &model.Data{Output: strings.ToLower(d.Input)}, nil
	})

	book := skill.NewBook()
	book.Register("lower", lowerSkill.Flatten())

	svc := service.NewCustomSkillService(book)
	svc.Run()
}
