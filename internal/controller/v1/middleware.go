package v1

import (
	"context"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cold-runner/skylark/internal/model/user"
	"github.com/cold-runner/skylark/internal/pkg/code"
	"github.com/cold-runner/skylark/internal/pkg/config"
	"github.com/hertz-contrib/jwt"
	bizErr "github.com/marmotedu/errors"
)

func (c *controllerV1) Jwt() *jwt.HertzJWTMiddleware {
	jwtConfig := config.GetConfig().JwtConfig()
	jwtMiddleware, err := jwt.New(&jwt.HertzJWTMiddleware{
		Realm:            jwtConfig.Realm,
		SigningAlgorithm: jwtConfig.SigningAlgorithm,
		Key:              []byte(jwtConfig.Key),
		KeyFunc:          nil,
		Timeout:          jwtConfig.Timeout * time.Hour,
		MaxRefresh:       jwtConfig.MaxRefresh * time.Hour,
		Authenticator: func(ctx context.Context, requestContext *app.RequestContext) (interface{}, error) {
			loginUser := &user.LoginUser{}
			// 参数校验
			if err := requestContext.BindAndValidate(loginUser); err != nil {
				return nil, err
			}

			// 服务层处理登陆逻辑
			if loggedUserInfo, err := c.serviceIns.Authenticator(ctx, loginUser); err != nil {
				return nil, err
			} else {
				return loggedUserInfo, nil
			}
		},
		Authorizator: func(data interface{}, ctx context.Context, c *app.RequestContext) bool {
			// TODO 授权
			return true
		},
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*user.LoggedUser); ok {
				return jwt.MapClaims{
					jwt.IdentityKey: v,
				}
			}
			return jwt.MapClaims{}
		},
		Unauthorized: func(ctx context.Context, c *app.RequestContext, _ int, _ string) {
			code.WriteResponse(c, bizErr.WithCode(code.ErrPermissionDenied, "", nil), nil)
		},
		LoginResponse: func(ctx context.Context, c *app.RequestContext, _ int, message string, _ time.Time) {
			code.WriteResponse(c, bizErr.WithCode(code.ErrSuccess, "", nil), utils.H{
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
