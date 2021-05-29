package routes

import (
	"Shield/api"
	"Shield/middleWare"
	"Shield/model"
	_"Shield/tool"
	"encoding/gob"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleWare.Cors())

	gob.Register(model.User{})
	store := cookie.NewStore([]byte("secret"))

	//路由上加入session中间件
	r.Use(sessions.Sessions("session", store))

	//文章
	articleRoute := r.Group("api/article")
	{
		articleRoute.GET("", api.Article)
		articleRoute.POST("", api.NewArticle)
		articleRoute.PUT("", api.UpdateArticle)
		articleRoute.DELETE("", api.DeleteArticle)
		articleRoute.GET("gift", api.GiftArticle)
		articleRoute.GET("search", api.FindArticle)
		articleRoute.GET("like", api.LikeArticle)
		articleRoute.GET("read", api.ReadArticle)
	}

	//文章列表
	articlesRoute := r.Group("api/articles")
	{
		articlesRoute.GET("", api.Articles)
	}

	//用户
	userRoute := r.Group("api/user")
	{
		userRoute.POST("login", api.VerifyUser)
		userRoute.GET("", api.UserById)
		userRoute.POST("", api.NewUser)
		userRoute.PUT("", api.UpdateUser)
		userRoute.DELETE("", api.DeleteUser)
		userRoute.GET("auth", api.Auth)
	}

	//频道
	channelRoute := r.Group("api/channel")
	{
		channelRoute.POST("", api.NewChannel)
		channelRoute.PUT("", api.UpdateChannel)
		channelRoute.DELETE("", api.DeleteChannel)
	}

	//频道列表
	channelsRoute := r.Group("api/channels")
	{
		channelsRoute.GET("", api.Channels)
	}

	//音乐API
	musicRoute := r.Group("api/music")
	{
		musicRoute.GET("", api.Music)
		musicRoute.POST("", api.NewMusic)
		musicRoute.DELETE("", api.DeleteMusic)
		musicRoute.GET("gift", api.GiftMusic)
		//musicRoute.GET("search", api.FindMusic)
	}

	//音乐列表API
	musicsRoute := r.Group("api/musics")
	{
		musicsRoute.GET("", api.Musics)
	}

	//杂项API
	sundryRoute := r.Group("api/sundry")
	{
		sundryRoute.GET("", api.Sundry)
		sundryRoute.POST("", api.NewSundry)
		sundryRoute.DELETE("", api.DeleteSundry)
		sundryRoute.PUT("", api.UpdateSundry)
		sundryRoute.GET("gift", api.GiftSundry)
		sundryRoute.GET("search", api.FindSundry)
		sundryRoute.GET("like", api.LikeSundry)
		sundryRoute.GET("read", api.ReadSundry)
	}

	//杂项列表API
	sundriesRoute := r.Group("api/sundries")
	{
		sundriesRoute.GET("", api.Sundries)
	}

	//影相API
	photoRoute := r.Group("api/photo")
	{
		photoRoute.GET("", api.Photo)
		photoRoute.POST("", api.NewPhoto)
		photoRoute.DELETE("", api.DeletePhoto)
		photoRoute.GET("gift", api.GiftPhoto)
	}

	//杂项列表API
	photosRoute := r.Group("api/photos")
	{
		photosRoute.GET("", api.Photos)
	}


	_ = r.Run("0.0.0.0:8888")

}
