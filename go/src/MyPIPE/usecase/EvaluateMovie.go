package usecase

import (
	"MyPIPE/domain/model"
	"MyPIPE/domain/repository"
	"errors"
)

type EvaluateMovie struct{
	UserRepository repository.UserRepository
	MovieRepository repository.MovieRepository
}

func NewEvaluateUsecase(u repository.UserRepository,m repository.MovieRepository)*EvaluateMovie{
	return &EvaluateMovie{
		UserRepository: u,
		MovieRepository: m,
	}
}

func (e EvaluateMovie)EvaluateMovie(evaluateMovieDTO EvaluateMovieDTO)error{
	userID,userIDErr := model.NewUserID(evaluateMovieDTO.UserID)
	if userIDErr != nil{
		return userIDErr
	}
	user,userErr := e.UserRepository.FindById(userID)
	if userErr != nil{
		return userErr
	}
	movieID,movieIDErr := model.NewMovieID(evaluateMovieDTO.MovieID)
	if movieIDErr != nil{
		return movieIDErr
	}
	movie,movieErr := e.MovieRepository.FindById(movieID)
	if movieErr != nil{
		return movieErr
	}

	if movie == nil{
		return errors.New("No Such Movie.")
	}

	evaluater,evaluaterErr :=model.NewEvaluate(evaluateMovieDTO.Evaluation)
	if evaluaterErr != nil{
		return evaluaterErr
	}

	evaluateErr := user.Evaluate(evaluater,movieID)
	if evaluateErr != nil{
		return evaluateErr
	}

	updateUserErr := e.UserRepository.UpdateUser(user)
	if updateUserErr != nil{
		return updateUserErr
	}

	return nil
}

type EvaluateMovieDTO struct{
	UserID uint64 `json:"user_id"`
	MovieID uint64 `json:"movie_id"`
	Evaluation string `json:"evaluation"`
}