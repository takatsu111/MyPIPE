package repository

import "MyPIPE/domain/model"

type PlayListRepository interface {
	FindByID(playListID model.PlayListID) (*model.PlayList,error)
	FindByName(playListName model.PlayListName) ([]model.PlayList,error)
	FindByUserID(playListUserID model.UserID) ([]model.PlayList,error)
	FindByUserIDAndName(playListUserID model.UserID,playListName model.PlayListName) ([]model.PlayList,error)
	Save(playList *model.PlayList) error
}
