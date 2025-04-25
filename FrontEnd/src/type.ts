/**
 * @给发后端发送获取到的code
 */
export interface ResponseLogin {
  data: ResponseCallBack;
  msg?: string;
}
/**
 * response.CallBack
 */
export interface ResponseCallBack {
  token?: string;
}
/**
 * response.Success Rank
 */
export interface ResponseRank {
  data: ResponseRanking;
  msg?: string;
}

/**
 * response.Ranking
 */
export interface ResponseRanking {
  leaderboard: Ranks[];
}

/**
 * model.Leaderboard
 */
export interface Ranks {
  avatar_url?: string;
  score?: number;
  user_id?: number;
  user_name?: string;
}
/**
 * response.Success获取用户信息
 */
export interface ResponseUserInfo {
  data: ResponseUser;
  msg?: string;
}

/**
 * response.User
 */
export interface ResponseUser {
  domain?: string[];
  user: UserInfo;
}
/**
 * model.User
 */
export interface UserInfo {
  /**
   * 用户头像的 URL
   */
  avatar_url: string;
  /**
   * 用户的个人简介
   */
  Bio?: string;
  /**
   * 博客连接
   */
  blog?: string;
  /**
   * 协作者的数量
   */
  collaborators?: number;
  /**
   * 用户所属的公司
   */
  company?: string;
  /**
   * 邮箱
   */
  email?: string;
  /**
   * 评估
   */
  evaluation?: string;
  /**
   * 粉丝数
   */
  followers?: number;
  /**
   * 关注数
   */
  following?: number;
  id?: number;
  /**
   * 地区
   */
  location?: string;
  /**
   * 用户的登录名
   */
  login_name?: string;
  /**
   * 真实姓名
   */
  name?: string;
  /**
   * 国籍
   */
  nationality?: string;
  /**
   * 用户公开的仓库的数量
   */
  public_repos?: number;
  /**
   * 评分
   */
  score?: number;
  /**
   * 用户的私有仓库总数
   */
  total_private_repos?: number;
}
/**
 * response.Success -Eval
 */
export interface ResponseEval {
  data: ResponseEvaluationResp;
  msg?: string;
}

/**
 * response.EvaluationResp
 */
export interface ResponseEvaluationResp {
  evaluation: string;
}
/**
 * response.Success
 */
export interface ResponseSearchs {
  data: ResponseSearchResp;
  msg?: string;
}

/**
 * response.SearchResp
 */
export interface ResponseSearchResp {
  users: ModelUser[];
}

/**
 * model.User
 */
export interface ModelUser {
  /**
   * 用户头像的 URL
   */
  avatar_url?: string;
  /**
   * 用户的个人简介
   */
  Bio?: string;
  /**
   * 博客连接
   */
  blog?: string;
  /**
   * 协作者的数量
   */
  collaborators?: number;
  /**
   * 用户所属的公司
   */
  company?: string;
  /**
   * 邮箱
   */
  email?: string;
  /**
   * 评估
   */
  evaluation?: string;
  /**
   * 粉丝数
   */
  followers?: number;
  /**
   * 关注数
   */
  following?: number;
  id?: number;
  /**
   * 地区
   */
  location?: string;
  /**
   * 用户的登录名
   */
  login_name?: string;
  /**
   * 真实姓名
   */
  name?: string;
  /**
   * 国籍
   */
  nationality?: string;
  /**
   * 用户公开的仓库的数量
   */
  public_repos?: number;
  /**
   * 评分
   */
  score?: number;
  /**
   * 用户的私有仓库总数
   */
  total_private_repos?: number;
}
