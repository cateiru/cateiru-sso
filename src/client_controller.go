package src

import (
	"context"
	"database/sql"
	"errors"
	"net/http"

	"github.com/cateiru/cateiru-sso/src/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// そのユーザーはクライアントにアクセスできる権限を持っているか見る
func checkCanAccessToClient(ctx context.Context, db boil.ContextExecutor, client *models.Client, u *models.User) error {
	if client.OrgID.Valid {
		// orgが設定されている場合
		orgUser, err := models.OrganizationUsers(
			models.OrganizationUserWhere.OrganizationID.EQ(client.OrgID.String),
			models.OrganizationUserWhere.UserID.EQ(u.ID),
		).One(ctx, db)
		if errors.Is(err, sql.ErrNoRows) {
			// orgのアクセス権限がない場合
			return NewHTTPError(http.StatusForbidden, "you are not member of this org")
		}
		if err != nil {
			return err
		}

		if orgUser.Role == "guest" {
			return NewHTTPUniqueError(http.StatusForbidden, ErrNoAuthority, "you are not authority to access this organization")
		}
	} else {
		// orgが設定されていない場合は、clientの作成者のみがアクセス可能
		if client.OwnerUserID != u.ID {
			return NewHTTPError(http.StatusForbidden, "you are not owner of this client")
		}
	}
	return nil
}

// クライアントの詳細を取得する
func getClientDetails(ctx context.Context, db *sql.DB, clientId string, u *models.User) (*ClientDetailResponse, error) {
	client, err := models.Clients(
		models.ClientWhere.ClientID.EQ(clientId),
	).One(ctx, db)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, NewHTTPError(http.StatusNotFound, "client not found")
	}
	if err != nil {
		return nil, err
	}

	if err := checkCanAccessToClient(ctx, db, client, u); err != nil {
		return nil, err
	}

	redirectUrlRecords, err := models.ClientRedirects(
		models.ClientRedirectWhere.ClientID.EQ(client.ClientID),
	).All(ctx, db)
	if err != nil {
		return nil, err
	}
	redirectUrls := make([]string, len(redirectUrlRecords))
	for i, redirect := range redirectUrlRecords {
		redirectUrls[i] = redirect.URL
	}

	referrerUrlRecords, err := models.ClientReferrers(
		models.ClientReferrerWhere.ClientID.EQ(client.ClientID),
	).All(ctx, db)
	if err != nil {
		return nil, err
	}
	referrerUrls := make([]string, len(referrerUrlRecords))
	for i, referrer := range referrerUrlRecords {
		// referrerはホストのみを見るので
		referrerUrls[i] = referrer.Host
	}

	scopesRecords, err := models.ClientScopes(
		models.ClientScopeWhere.ClientID.EQ(client.ClientID),
	).All(ctx, db)
	if err != nil {
		return nil, err
	}
	scopes := make([]string, len(scopesRecords))
	for i, scope := range scopesRecords {
		scopes[i] = scope.Scope
	}

	return &ClientDetailResponse{
		ClientSecret: client.ClientSecret,

		RedirectUrls: redirectUrls,
		ReferrerUrls: referrerUrls,
		Scopes:       scopes,

		OrgId: client.OrgID,

		ClientResponse: ClientResponse{
			ClientID: client.ClientID,

			Name:        client.Name,
			Description: client.Description,
			Image:       client.Image,

			IsAllow: client.IsAllow,
			Prompt:  client.Prompt,

			OrgMemberOnly: client.OrgMemberOnly,

			CreatedAt: client.CreatedAt,
			UpdatedAt: client.UpdatedAt,
		},
	}, nil
}

func getClientAllowRules(ctx context.Context, db *sql.DB, clientId string) ([]ClientAllowUserRuleResponse, error) {
	rules, err := models.ClientAllowRules(
		models.ClientAllowRuleWhere.ClientID.EQ(clientId),
		qm.Limit(100),
	).All(ctx, db)
	if err != nil {
		return nil, err
	}

	// ユーザーIDのリストを作る
	userIds := []string{}
	for _, rule := range rules {
		if rule.UserID.Valid {
			userIds = append(userIds, rule.UserID.String)
		}
	}

	// WHERE IN で一気にユーザー引いてくる
	// n+1 対策
	users, err := models.Users(
		models.UserWhere.ID.IN(userIds),
	).All(ctx, db)
	if err != nil {
		return nil, err
	}

	roleResponse := make([]ClientAllowUserRuleResponse, len(rules))
	for i, rule := range rules {
		var user *PublicUserResponse = nil
		if rule.UserID.Valid {
			// ユーザーを探す
			for _, u := range users {
				if u.ID == rule.UserID.String {
					user = &PublicUserResponse{
						ID:       u.ID,
						UserName: u.UserName,
						Avatar:   u.Avatar,
					}
					break
				}
			}
		}

		roleResponse[i] = ClientAllowUserRuleResponse{
			Id:          rule.ID,
			User:        user,
			EmailDomain: rule.EmailDomain,
		}
	}

	return roleResponse, nil
}
