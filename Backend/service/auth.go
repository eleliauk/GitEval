package service

import (
	"context"
	llmv1 "github.com/GitEval/GitEval-Backend/client/gen"
	"github.com/GitEval/GitEval-Backend/model"
	"github.com/google/go-github/v50/github"
)

type GitHubAPIProxy interface {
	GetLoginUrl() string
	SetClient(userID int64, client *github.Client)
	GetClientFromMap(userID int64) (*github.Client, bool)
	GetClientByCode(code string) (*github.Client, error)
	CalculateScore(ctx context.Context, id int64, name string) float64
	GetUserInfo(ctx context.Context, client *github.Client, username string) (*github.User, error)
}

type UserServiceProxy interface {
	InitUser(ctx context.Context, u model.User) (err error)
	GetUserById(ctx context.Context, id int64) (model.User, error)
	CreateUser(ctx context.Context, u model.User) error
}

type AuthService struct {
	githubAPI GitHubAPIProxy
	u         UserServiceProxy
	l         llmv1.LLMServiceClient
}

func NewAuthService(u UserServiceProxy, api GitHubAPIProxy, l llmv1.LLMServiceClient) *AuthService {
	return &AuthService{
		u: u,
		//因为让其成为中枢，必然要依赖注入到这个authService
		githubAPI: api,
		l:         l,
	}
}

func (s *AuthService) Login(ctx context.Context) (url string, err error) {
	url = s.githubAPI.GetLoginUrl()
	return url, nil
}

func (s *AuthService) CallBack(ctx context.Context, code string) (userId int64, err error) {
	client, err := s.githubAPI.GetClientByCode(code)
	if err != nil {
		return 0, err
	}

	userInfo, err := s.githubAPI.GetUserInfo(ctx, client, "")
	if err != nil {
		return 0, err
	}

	// 根据用户 ID 查找用户
	user, err := s.u.GetUserById(ctx, userInfo.GetID())
	// 如果用户不存在，创建新用户,如果存在
	if (user == model.User{}) {
		user = model.User{
			LoginName:         userInfo.GetLogin(),
			ID:                userInfo.GetID(),
			AvatarURL:         userInfo.GetAvatarURL(),
			Name:              userInfo.GetName(),
			Company:           userInfo.GetName(),
			Blog:              userInfo.GetBlog(),
			Location:          userInfo.GetLocation(),
			Email:             userInfo.GetEmail(),
			Bio:               userInfo.GetBio(),
			PublicRepos:       userInfo.GetPublicRepos(),
			Followers:         userInfo.GetFollowers(),
			Following:         userInfo.GetFollowing(),
			TotalPrivateRepos: userInfo.GetTotalPrivateRepos(),
			Collaborators:     userInfo.GetCollaborators(),
			Score:             s.githubAPI.CalculateScore(ctx, userInfo.GetID(), userInfo.GetLogin()), //获取用户的分数
		}

		//首次创建用户
		err = s.u.CreateUser(ctx, user)
		if err != nil {
			return 0, err
		}

	}

	//这里做异步主要是为了保证用户体验,否则等待时间过长了
	go func() {
		// 每次都尝试初始化用户关系网
		err = s.u.InitUser(context.Background(), user)
		if err != nil {
			return
		}
	}()

	//存储用户到内存中去
	s.githubAPI.SetClient(user.ID, client)

	return user.ID, nil
}
