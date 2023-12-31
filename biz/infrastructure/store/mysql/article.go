package mysql

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/json"
	"github.com/cold-runner/skylark/biz/infrastructure/store/orm_gen"
	"github.com/cold-runner/skylark/biz/infrastructure/store/other"
	"gorm.io/gen"
	"gorm.io/gen/field"
)

func (m *mysqlIns) GetPost(c context.Context, conds ...gen.Condition) (*orm_gen.Post, error) {
	return m.Query.Post.WithContext(c).Where(conds...).First()
}

func (m *mysqlIns) GetCategoryList(c context.Context, conds ...gen.Condition) ([]*orm_gen.Categorie, error) {
	return m.Query.Categorie.WithContext(c).Where(conds...).Find()
}

func (m *mysqlIns) CreateDraft(c context.Context, draft *orm_gen.Draft) error {
	return m.Query.Draft.WithContext(c).Create(draft)
}

func (m *mysqlIns) GetDraft(c context.Context, conds ...gen.Condition) (*orm_gen.Draft, error) {
	return m.Query.Draft.WithContext(c).Where(conds...).First()
}

func (m *mysqlIns) GetDraftList(c context.Context, conds ...gen.Condition) ([]*orm_gen.Draft, error) {
	return m.Query.Draft.WithContext(c).Where(conds...).Find()
}

func (m *mysqlIns) UpdateDraftSelect(c context.Context, selectScopes []field.Expr, whereScopes []gen.Condition, draft *orm_gen.Draft) error {
	_, err := m.Query.Draft.WithContext(c).Select(selectScopes...).Where(whereScopes...).Updates(draft)
	return err
}

func (m *mysqlIns) UpdateDraftOmit(c context.Context, omitScopes []field.Expr, whereScopes []gen.Condition, draft *orm_gen.Draft) error {
	_, err := m.Query.Draft.WithContext(c).Omit(omitScopes...).Where(whereScopes...).Updates(draft)
	return err
}

func (m *mysqlIns) GetPostList(c context.Context, conds ...gen.Condition) ([]*orm_gen.Post, error) {
	return m.Query.Post.WithContext(c).Where(conds...).Find()
}

func (m *mysqlIns) GetTagList(c context.Context, conds ...gen.Condition) ([]*orm_gen.Tag, error) {
	return m.Query.Tag.WithContext(c).Where(conds...).Find()
}

func (m *mysqlIns) GetPostTagList(c context.Context, conds ...gen.Condition) ([]*orm_gen.PostTag, error) {
	return m.Query.PostTag.WithContext(c).Where(conds...).Find()
}

func (m *mysqlIns) GetPostListAllDetail(c context.Context) ([]byte, error) {
	lark, pt, post, tag := m.Query.Lark, m.Query.PostTag, m.Query.Post, m.Query.Tag
	var res []other.PostWithAllDetail
	// TODO 多表
	err := post.WithContext(c).
		Select(post.ALL, post.UserID, lark.Name).
		LeftJoin(lark, lark.ID.EqCol(post.UserID)).
		Scan(&res)
	if err != nil {
		return nil, err
	}

	var pwt []other.PostWithTag
	err = pt.WithContext(c).
		Select(pt.PostID, tag.Name).
		LeftJoin(tag, tag.ID.EqCol(pt.TagID)).
		Order(pt.PostID).
		Scan(&pwt)
	if err != nil {
		return nil, err
	}
	pwtMap := map[string][]string{}
	// 双指针
	if len(pwt) != 0 {
		left, right := 0, 0
		for right < len(pwt) {
			if pwt[left].PostId != pwt[right].PostId {
				for left < right {
					pwtMap[pwt[left].PostId] = append(pwtMap[pwt[left].PostId], pwt[left].TagName)
					left++
				}
			}
			right++
		}
		for left < right {
			pwtMap[pwt[left].PostId] = append(pwtMap[pwt[left].PostId], pwt[left].TagName)
			left++
		}
	}
	for i := range res {
		res[i].Tags = pwtMap[res[i].ID]
	}

	data, _ := json.Marshal(res)
	return data, err
}

// ExistDraft 查询是否存在目标草稿
func (m *mysqlIns) ExistDraft(c context.Context, conds ...gen.Condition) (bool, error) {
	count, err := m.Query.Draft.WithContext(c).Where(conds...).Count()
	if err != nil {
		return false, err
	}
	if count != 0 {
		return true, nil
	}
	return false, nil
}
