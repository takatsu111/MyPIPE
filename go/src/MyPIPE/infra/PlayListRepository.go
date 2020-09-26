package infra

import (
	"MyPIPE/domain/model"
	"github.com/jinzhu/gorm"
	"time"
)

type PlayListPersistence struct{
	databaseAccessor *gorm.DB
}

func NewPlayListPersistence()*PlayListPersistence{
	return &PlayListPersistence{
		databaseAccessor: ConnectGorm(),
	}
}

func (p PlayListPersistence) FindByID(playListID model.PlayListID) (*model.PlayList, error) {
	var playLists model.PlayList
	result := p.databaseAccessor.Where("id = ?",playListID).Take(&playLists)
	if result.Error != nil{
		return nil,result.Error
	}
	result = p.databaseAccessor.Table("play_list_items").Where("play_list_id = ?",playListID).Pluck("movie_id",&playLists.PlayListItems)
	if result.Error != nil{
		return nil,result.Error
	}
	return &playLists,nil
}

func (p PlayListPersistence) FindByName(playListName model.PlayListName) ([]model.PlayList, error) {
	panic("implement me")
}

func (p PlayListPersistence) FindByUserID(playListUserID model.UserID) ([]model.PlayList, error) {
	panic("implement me")
}

func (p PlayListPersistence) FindByUserIDAndName(playListUserID model.UserID, playListName model.PlayListName) ([]model.PlayList, error) {
	var playLists []model.PlayList
	resultFindPlayList := p.databaseAccessor.Where("user_id = ? and name = ?",playListUserID,playListName).Find(&playLists)
	if resultFindPlayList.Error != nil{
		return nil,resultFindPlayList.Error
	}
	return playLists,nil
}

func (p *PlayListPersistence) Save(playList *model.PlayList) error {
	if playList.ID == 0{
		createResult := p.databaseAccessor.Create(&playList)
		if createResult.Error != nil{
			return createResult.Error
		}
		return nil
	}

	var playListItem playListItem

	transactionErr := p.databaseAccessor.Transaction(func(tx *gorm.DB) error {

		deleteResult :=tx.Exec("Delete From play_list_items Where play_list_id = ?",playList.ID)
		if deleteResult.Error != nil{
			return deleteResult.Error
		}

		for _,movieId := range playList.PlayListItems{
			playListItem.PlayListID = playList.ID
			playListItem.MovieID = movieId
			insertResult := tx.Create(&playListItem)
			if insertResult.Error != nil{
				return insertResult.Error
			}
		}

		saveResult := tx.Save(playList)
		if saveResult.Error != nil{
			return saveResult.Error
		}
		return nil
	})

	if transactionErr != nil{
		return transactionErr
	}


	return nil
}

type playListItem struct{
	PlayListID model.PlayListID
	MovieID model.MovieID
	CreatedAt     time.Time
	UpdatedAt     time.Time
}