package user

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/json"
	"time"

	"github.com/cold-runner/skylark/biz/infrastructure/cache"
	"github.com/cold-runner/skylark/biz/infrastructure/errCode"
	"github.com/cold-runner/skylark/biz/infrastructure/store"
	"github.com/cold-runner/skylark/biz/infrastructure/store/orm_gen"
	"github.com/cold-runner/skylark/biz/model/user"
	"github.com/cold-runner/skylark/biz/util"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/errors"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/google/uuid"
	stdErr "github.com/pkg/errors"
	"gorm.io/gorm"
)

type LoginUser struct {
	Uuid uuid.UUID `json:"uuid"`
}

func (l *LoginUser) PasswordLogin(c context.Context, ctx *app.RequestContext, req *user.PasswordLoginReq) *errors.Error {
	storeIns := store.GetIns()
	lark, err := storeIns.GetLark(c, storeIns.LarkByStuNum(req.StuNum))

	switch {
	case err == nil:
		if util.VerifyPassword(req.Password, lark.Password) != nil {
			return errCode.WrapBizErr(ctx, stdErr.New("password is incorrect!"), errCode.ErrPasswordIncorrect)
		}
		ctx.Set("uuid", lark.ID.String())
		return nil
	case stdErr.Is(err, gorm.ErrRecordNotFound):
		errMsg := "user not exist!" + "recv stuNum: " + req.GetStuNum()
		return errCode.WrapBizErr(ctx, stdErr.New(errMsg), errCode.ErrUserNotFound)
	default:
		return errCode.WrapBizErr(ctx, err, errCode.ErrUnknown)
	}
}

func (l *LoginUser) PhoneLogin(c context.Context, ctx *app.RequestContext, req *user.PhoneLoginReq) *errors.Error {
	storeIns := store.GetIns()
	cacheIns := cache.GetIns()
	lark, err := storeIns.GetLark(c, storeIns.LarkByPhone(req.Phone))

	switch {
	case stdErr.Is(err, gorm.ErrRecordNotFound):
		errMsg := "user not exist!" + "recv phone: " + req.Phone
		return errCode.WrapBizErr(ctx, stdErr.New(errMsg), errCode.ErrUserNotFound)
	case err != nil:
		return errCode.WrapBizErr(ctx, err, errCode.ErrUnknown)
	}

	storedSmsCode, err := cacheIns.Get(c, req.Phone)
	if err != nil {
		return errCode.WrapBizErr(ctx, err, errCode.ErrUnknown)
	}
	sCode := storedSmsCode.(string)
	if sCode != req.SmsCode {
		errMsg := "invalid smsCode! recv: " + req.SmsCode + " stored: " + sCode
		return errCode.WrapBizErr(ctx, stdErr.New(errMsg), errCode.ErrSmsCode)
	}
	err = cacheIns.Del(c, req.Phone)
	if err != nil {
		hlog.Warnf("delete smsCode failed! err: %v", err)
	}
	ctx.Set("uuid", lark.ID.String())
	return nil
}

type Lark struct {
	orm_gen.Lark
	FollowerCount uint64 `json:"follower_count"`
	FolloweeCount uint64 `json:"followee_count"`
	PostCount     uint64 `json:"post_count"`
	EssayCount    uint64 `json:"essay_count"`
}

func (l *Lark) GetById(c context.Context, ctx *app.RequestContext, uuid string) *errors.Error {
	storeIns := store.GetIns()
	lark, err := storeIns.GetLarkAllDetail("id", uuid)
	switch {
	case stdErr.Is(err, gorm.ErrRecordNotFound):
		errMsg := "user not exist!" + "recv uuid: " + uuid
		return errCode.WrapBizErr(ctx, stdErr.New(errMsg), errCode.ErrUserNotFound)
	case err != nil:
		return errCode.WrapBizErr(ctx, err, errCode.ErrUnknown)
	}
	data, _ := json.Marshal(lark)
	_ = json.Unmarshal(data, l)
	return nil
}

func (l *Lark) GetByStuNum(c context.Context, ctx *app.RequestContext, stuNum string) *errors.Error {
	storeIns := store.GetIns()
	lark, err := storeIns.GetLarkAllDetail("stu_num", stuNum)
	switch {
	case stdErr.Is(err, gorm.ErrRecordNotFound):
		errMsg := "user not exist!" + "recv stu_num: " + stuNum
		return errCode.WrapBizErr(ctx, stdErr.New(errMsg), errCode.ErrUserNotFound)
	case err != nil:
		return errCode.WrapBizErr(ctx, err, errCode.ErrUnknown)
	}
	data, _ := json.Marshal(lark)
	_ = json.Unmarshal(data, l)
	return nil
}

func (l *Lark) Format() *user.Basic {
	return &user.Basic{
		UserId:            l.ID.String(),
		StuNum:            l.StuNum,
		AvatarUrl:         l.Avatar,
		StuName:           l.Name,
		BriefIntroduction: l.Introduce,
		College:           l.College,
		Major:             l.Major,
		Grade:             l.Grade,
		FolloweeCount:     l.FolloweeCount,
		FollowerCount:     l.FollowerCount,
		PostArticleCount:  l.PostCount,
		EssayArticleCount: l.EssayCount,
		Power:             1,
	}

}

type LarkWithAllDetail struct {
	Id            string    `gorm:"column:id" json:"id"`
	CreatedAt     time.Time `gorm:"column:created_at" json:"created_at"`
	StuNum        string    `gorm:"column:stu_num" json:"stu_num"`
	Name          string    `gorm:"column:name" json:"name"`
	Gender        string    `gorm:"column:gender" json:"gender"`
	College       string    `gorm:"column:college" json:"college"`
	Major         string    `gorm:"column:major" json:"major"`
	Grade         string    `gorm:"column:grade" json:"grade"`
	Province      string    `gorm:"column:province" json:"province"`
	Email         string    `gorm:"column:email" json:"email"`
	Introduce     string    `gorm:"column:introduce" json:"introduce"`
	Avatar        string    `gorm:"column:avatar" json:"avatar"`
	State         int       `gorm:"column:state" json:"state"`
	FolloweeCount uint64    `gorm:"column:followee_count" json:"followee_count"`
	FollowerCount uint64    `gorm:"column:follower_count" json:"follower_count"`
	PostCount     uint64    `gorm:"column:post_count" json:"post_count"`
	EssayCount    uint64    `gorm:"column:essay_count" json:"essay_count"`
}
