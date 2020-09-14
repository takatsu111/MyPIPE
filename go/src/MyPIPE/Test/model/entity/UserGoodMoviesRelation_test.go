package test

import (
	"MyPIPE/domain/model"
	"MyPIPE/infra"
	"testing"
)

func TestRelationBetweenUserAndGoodMovies(t *testing.T){
	q := &model.User{ID:model.UserID(1012)}
	infra.ConnectGorm().Preload("GoodMovies").Find(q).QueryExpr()
}
