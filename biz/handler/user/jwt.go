package user

import (
	"context"
	"github.com/cold-runner/skylark/biz/model/user"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cold-runner/skylark/biz/config"
	"github.com/cold-runner/skylark/biz/infrastructure/errCode"
	"github.com/hertz-contrib/jwt"
)

var JwtMiddleware *jwt.HertzJWTMiddleware

func InitAuthenticatorAndAuthorization(cfg config.Conf) {
	jwtConfig := cfg.GetJwt()
	jwtMw, err := jwt.New(&jwt.HertzJWTMiddleware{
		Realm:            jwtConfig.Realm,
		SigningAlgorithm: jwtConfig.SigningAlgorithm,
		Key:              []byte(jwtConfig.Key),
		Timeout:          jwtConfig.Timeout * time.Hour,
		MaxRefresh:       jwtConfig.MaxRefresh * time.Hour,
		IdentityKey:      jwt.IdentityKey,
		TokenHeadName:    jwtConfig.TokenHeadName,
		// 登陆
		Authenticator: func(c context.Context, ctx *app.RequestContext) (interface{}, error) {
			return ctx.Value("uuid").(string), nil
		},
		// 授权
		Authorizator: func(data interface{}, ctx context.Context, c *app.RequestContext) bool {
			// TODO 授权

			return true
		},
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(string); ok {
				return jwt.MapClaims{
					jwt.IdentityKey: v,
				}
			}
			return jwt.MapClaims{}
		},
		Unauthorized: func(c context.Context, ctx *app.RequestContext, _ int, _ string) {
			errCode.ResponsePermissionDenied(ctx)
		},
		LoginResponse: func(c context.Context, ctx *app.RequestContext, _ int, token string, _ time.Time) {
			ctx.JSON(consts.StatusOK, user.LoginResp{
				Code: errCode.SuccessStatus.Code,
				Msg:  errCode.SuccessStatus.Msg,
				Data: &user.LoginResp_Data{Token: token},
			})
		},
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			claims := jwt.ExtractClaims(ctx, c)
			return claims["identity"]
		},
	})
	if err != nil {
		panic(err)
	}
	JwtMiddleware = jwtMw
}
