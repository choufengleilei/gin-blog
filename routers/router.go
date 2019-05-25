package routers

import (
	"gin/gin-blog/docs"
	"gin/gin-blog/middleware/jwt"
	"gin/gin-blog/pkg/export"
	"gin/gin-blog/pkg/qrcode"
	"gin/gin-blog/pkg/setting"
	"gin/gin-blog/pkg/upload"
	"gin/gin-blog/routers/api"
	"gin/gin-blog/routers/api/v1"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
)

func InitRouter() *gin.Engine {

	// programatically set swagger info
	docs.SwaggerInfo.Title = "gin-blog Swagger API"
	docs.SwaggerInfo.Description = "This is a sample server Petstore server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "127.0.0.1:8000"
	docs.SwaggerInfo.BasePath = "/"

	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.ServerSetting.RunMode)

	r.GET("/auth", api.GetAuth)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.POST("/upload", api.UploadImage)

	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))

	r.StaticFS("/export", http.Dir(export.GetExcelFullPath()))

	r.StaticFS("/qrcode", http.Dir(qrcode.GetQrCodeFullPath()))

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		//获取标签列表
		apiv1.GET("/tag/list", v1.GetTags)
		//新建标签
		apiv1.POST("/tag/add", v1.AddTag)
		//更新指定标签
		apiv1.POST("/tag/edit/:id", v1.EditTag)
		//删除指定标签
		apiv1.POST("/tag/delete/:id", v1.DeleteTag)
		//导出标签
		r.POST("/tag/export", v1.ExportTag)
		//导入标签
		r.POST("/tag/import", v1.ImportTag)

		//获取文章列表
		apiv1.GET("/article/list", v1.GetArticles)
		//获取指定文章
		apiv1.GET("/article/getOne/:id", v1.GetArticle)
		//新建文章
		apiv1.POST("/article/add", v1.AddArticle)
		//更新指定文章
		apiv1.POST("/article/edit/:id", v1.EditArticle)
		//删除指定文章
		apiv1.POST("/article/delete/:id", v1.DeleteArticle)
		//生成二维码
		apiv1.POST("/article/poster/generate", v1.GenerateArticlePoster)

	}

	return r
}
