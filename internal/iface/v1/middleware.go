package v1

import (
	"context"
	"github.com/cold-runner/skylark/internal/domain/entity"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cold-runner/skylark/internal/infrastructure/config"
	"github.com/cold-runner/skylark/internal/infrastructure/errCode"
	"github.com/hertz-contrib/jwt"
)

func (c *userInterface) Jwt() *jwt.HertzJWTMiddleware {
	jwtConfig := config.GetConfig().JwtConfig()
	jwtMiddleware, err := jwt.New(&jwt.HertzJWTMiddleware{
		Realm:            jwtConfig.Realm,
		SigningAlgorithm: jwtConfig.SigningAlgorithm,
		Key:              []byte(jwtConfig.Key),
		KeyFunc:          nil,
		Timeout:          jwtConfig.Timeout * time.Hour,
		MaxRefresh:       jwtConfig.MaxRefresh * time.Hour,
		Authenticator: func(ctx context.Context, requestContext *app.RequestContext) (interface{}, error) {
			hlog.Debug("登陆方法被调用")
			loginUser := &entity.LoginUser{}

			if err := requestContext.BindAndValidate(loginUser); err != nil {
				return nil, err
			}

			// 服务层处理登陆逻辑
			if loggedUserInfo, err := c.application.Authenticator(ctx, requestContext, loginUser); err != nil {
				return nil, err
			} else {
				return loggedUserInfo, nil
			}
		},
		Authorizator: func(data interface{}, ctx context.Context, c *app.RequestContext) bool {
			hlog.Debug("授权方法被调用")
			// TODO 授权
			return true
		},
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*entity.LoggedInUser); ok {
				return jwt.MapClaims{
					jwt.IdentityKey: v,
				}
			}
			return jwt.MapClaims{}
		},
		Unauthorized: func(ctx context.Context, c *app.RequestContext, _ int, _ string) {
			errCode.ResponseWithError(c, errCode.ErrPermissionDenied, nil)
		},
		LoginResponse: func(ctx context.Context, c *app.RequestContext, _ int, message string, _ time.Time) {
			errCode.ResponseOk(c, utils.H{
				"token": message,
			})
		},
		LogoutResponse:  nil,
		RefreshResponse: nil,
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			claims := jwt.ExtractClaims(ctx, c)
			return claims["identity"]
		},
		IdentityKey:                 jwt.IdentityKey,
		TokenLookup:                 "",
		TokenHeadName:               jwtConfig.TokenHeadName,
		WithoutDefaultTokenHeadName: false,
		TimeFunc:                    nil,
		HTTPStatusMessageFunc:       nil,
		PrivKeyFile:                 "",
		PrivKeyBytes:                nil,
		PubKeyFile:                  "",
		PrivateKeyPassphrase:        "",
		PubKeyBytes:                 nil,
		SendCookie:                  false,
		CookieMaxAge:                0,
		SecureCookie:                false,
		CookieHTTPOnly:              false,
		CookieDomain:                "",
		SendAuthorization:           false,
		DisabledAbort:               false,
		CookieName:                  "",
		CookieSameSite:              0,
		ParseOptions:                nil,
	})
	if err != nil {
		panic(err)
	}
	return jwtMiddleware
}
