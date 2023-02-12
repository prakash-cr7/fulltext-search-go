package model

import "gorm.io/gorm"

type User1 struct {
	gorm.Model
	Name string `gorm:"index:,class:FULLTEXT,option:WITH PARSER ngram"`
  }