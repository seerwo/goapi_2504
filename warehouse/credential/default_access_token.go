package credential

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/seerwo/goapi_2504/cache"
	"github.com/seerwo/goapi_2504/util"
	"sync"
	"time"
)

const (
	// accessTokenURL 获取access_token的接口
	accessTokenURL = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s"
	// stableAccessTokenURL 获取稳定版access_token的接口
	stableAccessTokenURL = "https://api.weixin.qq.com/cgi-bin/stable_token"
	// workAccessTokenURL 企业微信获取access_token的接口
	workAccessTokenURL = "https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=%s&corpsecret=%s"
	// CacheKeyOfficialAccountPrefix 微信公众号cache key前缀
	CacheKeyOfficialAccountPrefix = "gowechat_officialaccount_"
	// CacheKeyMiniProgramPrefix 小程序cache key前缀
	CacheKeyMiniProgramPrefix = "gowechat_miniprogram_"
	// CacheKeyWorkPrefix 企业微信cache key前缀
	CacheKeyWorkPrefix = "gowechat_work_"
)

// ResAccessToken struct
type ResAccessToken struct {
	util.CommonError

	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

// WorkAccessToken 企业微信AccessToken 获取
type WorkAccessToken struct {
	CorpID          string
	CorpSecret      string
	AgentID         string // 可选，用于区分不同应用
	cacheKeyPrefix  string
	cache           cache.Cache
	accessTokenLock *sync.Mutex
}

// NewWorkAccessToken new WorkAccessToken (保持向后兼容)
func NewWorkAccessToken(corpID, corpSecret, agentID, cacheKeyPrefix string, cache cache.Cache) AccessTokenContextHandle {
	// 调用新方法，保持兼容性
	return NewWorkAccessTokenWithAgentID(corpID, corpSecret, agentID, cacheKeyPrefix, cache)
}

// NewWorkAccessTokenWithAgentID new WorkAccessToken with agentID
func NewWorkAccessTokenWithAgentID(corpID, corpSecret, agentID, cacheKeyPrefix string, cache cache.Cache) AccessTokenContextHandle {
	if cache == nil {
		panic("cache is needed")
	}
	return &WorkAccessToken{
		CorpID:          corpID,
		CorpSecret:      corpSecret,
		AgentID:         agentID,
		cache:           cache,
		cacheKeyPrefix:  cacheKeyPrefix,
		accessTokenLock: new(sync.Mutex),
	}
}

// GetAccessToken 企业微信获取access_token,先从cache中获取，没有则从服务端获取
func (ak *WorkAccessToken) GetAccessToken() (accessToken string, err error) {
	return ak.GetAccessTokenContext(context.Background())
}

// GetAccessTokenContext 企业微信获取access_token,先从cache中获取，没有则从服务端获取
func (ak *WorkAccessToken) GetAccessTokenContext(ctx context.Context) (accessToken string, err error) {
	// 加上lock，是为了防止在并发获取token时，cache刚好失效，导致从微信服务器上获取到不同token
	ak.accessTokenLock.Lock()
	defer ak.accessTokenLock.Unlock()

	// 构建缓存key
	var accessTokenCacheKey string

	if ak.AgentID != "" {
		// 如果设置了AgentID，使用新的key格式
		accessTokenCacheKey = fmt.Sprintf("%s_access_token_%s_%s", ak.cacheKeyPrefix, ak.CorpID, ak.AgentID)
	} else {
		// 兼容历史版本的key格式
		accessTokenCacheKey = fmt.Sprintf("%s_access_token_%s", ak.cacheKeyPrefix, ak.CorpID)
	}

	val := ak.cache.Get(accessTokenCacheKey)
	if val != nil {
		accessToken = val.(string)
		return
	}

	// cache失效，从微信服务器获取
	var resAccessToken ResAccessToken
	resAccessToken, err = GetTokenFromServerContext(ctx, fmt.Sprintf(workAccessTokenURL, ak.CorpID, ak.CorpSecret))
	if err != nil {
		return
	}

	expires := resAccessToken.ExpiresIn - 1500
	err = ak.cache.Set(accessTokenCacheKey, resAccessToken.AccessToken, time.Duration(expires)*time.Second)
	if err != nil {
		return
	}

	accessToken = resAccessToken.AccessToken
	return
}

// GetTokenFromServerContext 强制从微信服务器获取token
func GetTokenFromServerContext(ctx context.Context, url string) (resAccessToken ResAccessToken, err error) {
	var body []byte
	body, err = util.HTTPGetContext(ctx, url)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &resAccessToken)
	if err != nil {
		return
	}
	if resAccessToken.ErrCode != 0 {
		err = fmt.Errorf("get access_token error : errcode=%v , errormsg=%v", resAccessToken.ErrCode, resAccessToken.ErrMsg)
		return
	}
	return
}