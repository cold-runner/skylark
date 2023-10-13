package social

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"

	"golang.org/x/oauth2"
)

type QqIdProvider struct {
	Client *http.Client
	Config *oauth2.Config
}

func NewQqIdProvider(clientId string, clientSecret string, redirectUrl string) *QqIdProvider {
	idp := &QqIdProvider{}

	config := idp.getConfig()
	config.ClientID = clientId
	config.ClientSecret = clientSecret
	config.RedirectURL = redirectUrl
	idp.Config = config

	return idp
}

func (idp *QqIdProvider) SetHttpClient(client *http.Client) {
	idp.Client = client
}

func (idp *QqIdProvider) getConfig() *oauth2.Config {
	endpoint := oauth2.Endpoint{
		TokenURL: "https://graph.qq.com/oauth2.0/token",
	}

	config := &oauth2.Config{
		Scopes:   []string{"get_user_info"},
		Endpoint: endpoint,
	}

	return config
}

func (idp *QqIdProvider) GetToken(code string) (*oauth2.Token, error) {
	params := url.Values{}
	params.Add("grant_type", "authorization_code")
	params.Add("client_id", idp.Config.ClientID)
	params.Add("client_secret", idp.Config.ClientSecret)
	params.Add("code", code)
	params.Add("redirect_uri", idp.Config.RedirectURL)

	accessTokenUrl := fmt.Sprintf("https://graph.qq.com/oauth2.0/token?%s", params.Encode())
	resp, err := idp.Client.Get(accessTokenUrl)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	tokenContent, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	re := regexp.MustCompile("token=(.*?)&")
	matched := re.FindAllStringSubmatch(string(tokenContent), -1)
	accessToken := matched[0][1]
	token := &oauth2.Token{
		AccessToken: accessToken,
		TokenType:   "Bearer",
	}
	return token, nil
}

//{
//	"ret": 0,
//	"msg": "",
//	"is_lost": 0,
//	"nickname": "飞翔的企鹅",
//	"gender": "男",
//	"gender_type": 1,
//	"province": "",
//	"city": "安道尔城",
//	"year": "1968",
//	"constellation": "",
//	"figureurl": "http:\/\/qzapp.qlogo.cn\/qzapp\/101896710\/C0D022F92B604AA4B1CDF92CC79463A4\/30",
//	"figureurl_1": "http:\/\/qzapp.qlogo.cn\/qzapp\/101896710\/C0D022F92B604AA4B1CDF92CC79463A4\/50",
//	"figureurl_2": "http:\/\/qzapp.qlogo.cn\/qzapp\/101896710\/C0D022F92B604AA4B1CDF92CC79463A4\/100",
//	"figureurl_qq_1": "http://thirdqq.qlogo.cn/g?b=oidb&k=QtAu5OiaSfqGD0kfclwvxJA&s=40&t=1557635654",
//	"figureurl_qq_2": "http://thirdqq.qlogo.cn/g?b=oidb&k=QtAu5OiaSfqGD0kfclwvxJA&s=100&t=1557635654",
//	"figureurl_qq": "http://thirdqq.qlogo.cn/g?b=oidb&k=QtAu5OiaSfqGD0kfclwvxJA&s=640&t=1557635654",
//	"figureurl_type": "1",
//	"is_yellow_vip": "0",
//	"vip": "0",
//	"yellow_vip_level": "0",
//	"level": "0",
//	"is_yellow_year_vip": "0"
//}

type QqUserUnionId struct {
	ClientId string `json:"client_id"`
	OpenId   string `json:"open_id"`
	UnionId  string `json:"union_id"`
}

type QqUserInfo struct {
	Ret             int    `json:"ret"`
	Msg             string `json:"msg"`
	IsLost          int    `json:"is_lost"`
	Nickname        string `json:"nickname"`
	Gender          string `json:"gender"`
	GenderType      int    `json:"gender_type"`
	Province        string `json:"province"`
	City            string `json:"city"`
	Year            string `json:"year"`
	Constellation   string `json:"constellation"`
	Figureurl       string `json:"figureurl"`
	Figureurl1      string `json:"figureurl_1"`
	Figureurl2      string `json:"figureurl_2"`
	FigureurlQq1    string `json:"figureurl_qq_1"`
	FigureurlQq2    string `json:"figureurl_qq_2"`
	FigureurlQq     string `json:"figureurl_qq"`
	FigureurlType   string `json:"figureurl_type"`
	IsYellowVip     string `json:"is_yellow_vip"`
	Vip             string `json:"vip"`
	YellowVipLevel  string `json:"yellow_vip_level"`
	Level           string `json:"level"`
	IsYellowYearVip string `json:"is_yellow_year_vip"`
}

func (idp *QqIdProvider) getUserUnionIdStruct(token *oauth2.Token) (unionId *QqUserUnionId, err error) {
	openIdUrl := fmt.Sprintf("https://graph.qq.com/oauth2.0/me?access_token=%s&unionid=1&fmt=json", token.AccessToken)
	resp, err := idp.Client.Get(openIdUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	qqUserUnionId := &QqUserUnionId{}
	err = json.Unmarshal(body, qqUserUnionId)
	if err != nil {
		return nil, err
	}

	return qqUserUnionId, nil
}

func (idp *QqIdProvider) GetUserInfo(token *oauth2.Token) (*UserInfo, error) {
	qqUserUnionId, err := idp.getUserUnionIdStruct(token)
	if err != nil {
		return nil, err
	}

	userInfoUrl := fmt.Sprintf("https://graph.qq.com/user/get_user_info?access_token=%s&oauth_consumer_key=%s&openid=%s", token.AccessToken, idp.Config.ClientID, qqUserUnionId.OpenId)
	resp, err := idp.Client.Get(userInfoUrl)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	userInfoBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var qqUserInfo QqUserInfo
	err = json.Unmarshal(userInfoBody, &qqUserInfo)
	if err != nil {
		return nil, err
	}

	if qqUserInfo.Ret != 0 {
		return nil, fmt.Errorf("ret expected 0, got %d", qqUserInfo.Ret)
	}

	userInfo := UserInfo{
		Id:          qqUserUnionId.UnionId,
		Username:    qqUserInfo.Nickname,
		DisplayName: qqUserInfo.Nickname,
		AvatarUrl:   qqUserInfo.FigureurlQq1,
	}
	return &userInfo, nil
}

func (idp *QqIdProvider) GetUserUnionId(token *oauth2.Token) (string, error) {
	idStruct, err := idp.getUserUnionIdStruct(token)
	if err != nil {
		return "", err
	}
	return idStruct.UnionId, nil
}
