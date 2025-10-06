package api

import (
	"bws_microservice_url/main/src/core/url_manager/register"
	"bws_microservice_url/main/src/core/url_manager/update"
	"bws_microservice_url/main/src/core/urls"
	"bws_microservice_url/main/src/dto"
	"net/http"

	bw_helper "github.com/bindways/bw_microservice_share2/bw_gin"
	"github.com/bindways/bw_microservice_share2/bw_oauth"
	"github.com/bindways/bws_microservice_share/enum"
	"github.com/bindways/bws_microservice_share/router"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BwsUrlController struct {
	urlRegisterService *register.BwsUrlRegisterService
	urlUpdateService   *update.BwsUrlUpdate
	urlTrackService    *urls.BwsUrlTrackService
}

func (t *BwsUrlController) Constructor1() *BwsUrlController {
	t.urlUpdateService = new(update.BwsUrlUpdate).Constructor()
	t.urlRegisterService = new(register.BwsUrlRegisterService).Constructor()
	t.urlTrackService = new(urls.BwsUrlTrackService).Constructor1()
	return t
}

func (t *BwsUrlController) Controller(engine *gin.Engine) {
	engine.GET(router.BwsMicroserviceUrl.GetPublicPath("/track/all"),
		bw_oauth.CheckHasRoles(enum.BwsRoleCustomer),
		func(context *gin.Context) {
			idCustomer, err := primitive.ObjectIDFromHex(context.GetString("idToken"))
			if err != nil {
				bw_helper.NewErrorResponse400(context, err)
				return
			}
			urlTrack, err := t.urlTrackService.GetUrlTrackByCustomer(idCustomer)
			if err != nil {
				bw_helper.NewErrorResponse400(context, err)
				return
			}
			context.JSON(http.StatusOK, urlTrack)
		},
	)

	engine.POST(router.BwsMicroserviceUrl.GetPublicPath("/creation"),
		bw_oauth.CheckHasRoles(enum.BwsRoleCustomer),
		func(context *gin.Context) {
			urlTrackCreator := dto.BwsUrlTrackCreator{}
			if err := context.ShouldBindJSON(&urlTrackCreator); err != nil {
				return
			}
			idCustomer, err := primitive.ObjectIDFromHex(context.GetString("idToken"))
			if err != nil {
				bw_helper.NewErrorResponse400(context, err)
				return
			}
			urlTrack, err := t.urlRegisterService.RegisterUrl(idCustomer, urlTrackCreator)
			if err != nil {
				bw_helper.NewErrorResponse400(context, err)
				return
			}
			context.JSON(http.StatusOK, urlTrack)
		},
	)

	engine.PUT(router.BwsMicroserviceUrl.GetPublicPath("/track/:idUrl"),
		bw_oauth.CheckHasRoles(enum.BwsRoleCustomer),
		func(context *gin.Context) {
			idCustomer, err := primitive.ObjectIDFromHex(context.GetString("idToken"))
			if err != nil {
				bw_helper.NewErrorResponse400(context, err)
				return
			}
			idUrl, err := primitive.ObjectIDFromHex(context.Param("idUrl"))
			if err != nil {
				bw_helper.NewErrorResponse400(context, err)
				return
			}
			urlTrack, err := t.urlUpdateService.UpdateUrl(idUrl, idCustomer)
			if err != nil {
				bw_helper.NewErrorResponse400(context, err)
				return
			}
			context.JSON(http.StatusOK, urlTrack)
		},
	)
}
