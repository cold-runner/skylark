package social

import (
	"github.com/pkg/errors"
	"golang.org/x/oauth2"
	"net/http"
)

type UserInfo struct {
	Id          string
	Username    string
	DisplayName string
	UnionId     string
	Email       string
	Phone       string
	CountryCode string
	AvatarUrl   string
}

type ProviderInfo struct {
	Type         string
	SubType      string
	ClientId     string
	ClientSecret string
	AppId        string
	HostUrl      string
	RedirectUrl  string

	TokenURL    string
	AuthURL     string
	UserInfoURL string
	UserMapping map[string]string
}

type IdProvider interface {
	SetHttpClient(client *http.Client)
	GetToken(code string) (*oauth2.Token, error)
	GetUserInfo(token *oauth2.Token) (*UserInfo, error)
}

func GetIdProvider(idpInfo *ProviderInfo, redirectUrl string) (IdProvider, error) {
	switch idpInfo.Type {
	case "QQ":
		return NewQqIdProvider(idpInfo.ClientId, idpInfo.ClientSecret, redirectUrl), nil
	default:
		return nil, errors.New("unsupported social provider")
	}
}
