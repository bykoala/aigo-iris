package controllers

import (
	"github.com/kataras/iris/mvc"
	"repositories"
	"services"
)

type MovieController struct {

}

func (c *MovieController) Get() mvc.View{
	movieReposiroty := repositories.NewMovieManager()
	movieService := services.NewMovieServiceManger(movieReposiroty)
	MoiveResult := movieService.ShowMovieName()
	return mvc.View{
		Name:"movie/index.html",
		Data:MoiveResult,
	}
}
