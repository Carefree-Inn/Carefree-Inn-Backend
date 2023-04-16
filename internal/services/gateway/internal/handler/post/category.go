package post

import (
	"context"
	"gateway/internal"
	"gateway/internal/handler/user"
	"gateway/pkg"
	"gateway/pkg/errno"
	"gateway/pkg/log"
	"github.com/gin-gonic/gin"
	pbPost "github.com/jackj-ohn1/package/proto/post"
	pbUser "github.com/jackj-ohn1/package/proto/user"
	"github.com/pkg/errors"
	"strconv"
)

//  GetCategory getCategory
//	@Summary		获取分区信息 api
//	@Tags			post
//	@Description	获取分区信息
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}	internal.Response
//	@Router			/post/category/all [get]
func (p *postHandler) GetCategory(c *gin.Context) {
	ctx := context.WithValue(c.Request.Context(), "X-Request-Id", pkg.GetUUid(c))
	
	resp, err := p.PostService.GetCategory(ctx, &pbPost.Request{})
	if err != nil {
		internal.ServerError(c, errno.GetCategoryInfoError.Error())
		log.Panic(log.WithField("X-Request-Id", c.MustGet("uuid")),
			errors.WithStack(err), errno.GetCategoryInfoError.Error())
		return
	}
	
	internal.Success(c, resp.Categories)
}

//  GetPostCategory getPostCategory
//	@Summary		获取分区帖子信息 api
//	@Tags			post
//	@Description	获取分区帖子信息
//	@Accept			json
//	@Produce		json
//  @Param          page            query       int false  "页码"
//  @Param          limit           query       int false  "条数"
//	@Success		200		{object}	internal.Response
//	@Router			/post/category [get]
func (p *postHandler) GetPostOfCategory(c *gin.Context) {
	ctx := context.WithValue(c.Request.Context(), "X-Request-Id", pkg.GetUUid(c))
	
	var category categoryInfo
	if err := c.ShouldBindJSON(&category); err != nil {
		internal.Error(c, errno.JsonDataError)
		log.Warn(
			log.WithField("X-Request-Id", c.MustGet("uuid")),
			errors.WithStack(err), errno.JsonDataError.Error(),
		)
		return
	}
	
	page, errPage := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, errLimit := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if errPage != nil || errLimit != nil {
		page = 1
		limit = 10
	}
	
	resp, err := p.PostService.GetPostOfCategory(ctx, &pbPost.CategoryRequest{
		Category: &pbPost.CategoryInfo{
			CategoryId: category.CategoryId,
		},
		Limit: uint32(limit),
		Page:  uint32(page),
	})
	if err != nil {
		internal.ServerError(c, errno.GetCategoryCategoryPostError.Error())
		log.Panic(log.WithField("X-Request-Id", c.MustGet("uuid")),
			errors.WithStack(err), errno.GetCategoryCategoryPostError.Error())
		return
	}
	
	var accounts = make([]string, 0, len(resp.Posts))
	for _, v := range resp.Posts {
		accounts = append(accounts, v.Account)
	}
	
	data, err := p.GetUserInfoWithAny(ctx, accounts, resp.Posts, p.assemble)
	if err != nil {
		internal.ServerError(c, errno.GetCategoryCategoryPostError.Error())
		log.Panic(log.WithField("X-Request-Id", c.MustGet("uuid")),
			errors.WithStack(err), errno.GetCategoryCategoryPostError.Error())
		return
	}
	
	internal.Success(c, data.([]*PostInfo))
}

func (p *postHandler) assemble(ref map[string]*user.UserInfo, unknown any) (any, error) {
	posts, ok := unknown.([]*pbPost.PostInfo)
	if !ok {
		return nil, nil
	}
	
	var data = make([]*PostInfo, 0, len(posts))
	for _, v := range posts {
		tag := make([]*tagInfo, 0, len(v.Tag))
		for _, val := range v.Tag {
			tag = append(tag, &tagInfo{
				TagId: val.TagId,
				Title: v.Title,
			})
		}
		
		data = append(data, &PostInfo{
			PostId:  v.PostId,
			Title:   v.Title,
			Content: v.Content,
			Category: categoryInfo{
				CategoryId: v.Category.CategoryId,
				Title:      v.Category.Title,
			},
			Tags:       tag,
			UserInfo:   ref[v.Account],
			CreateTime: v.CreateTime,
			Star:       v.Star,
		})
	}
	
	return data, nil
}

func (p *postHandler) GetUserInfoWithAny(ctx context.Context, accounts []string, data any, assemble func(map[string]*user.UserInfo, any) (any, error)) (any, error) {
	resp, err := p.UserService.GetBatchUserProfile(ctx, &pbUser.BatchUserProfileRequest{
		Accounts: accounts,
	})
	if err != nil {
		return nil, err
	}
	
	var ref = make(map[string]*user.UserInfo, len(resp.Data))
	for _, v := range resp.Data {
		ref[v.Account] = &user.UserInfo{
			Account:  v.Account,
			Nickname: v.Nickname,
			Avatar:   v.Avatar,
			Sex:      int8(v.Sex),
		}
	}
	
	return assemble(ref, data)
	
}
