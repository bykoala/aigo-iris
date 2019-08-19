package repositories

import "datamodels"

type MovieRepository interface {
	GetMovieName() string
}

type MovieManager struct{
}

func NewMovieManager() MovieRepository{
	return &MovieManager{}
}

func (m *MovieManager) GetMovieName() string{
	//模拟赋值给模型，一般是从数据库获取数据
	movie := &datamodels.Movie{Name:"风语咒"}
	return movie.Name
}