/*
 *     Copyright 2020 The Dragonfly Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package middlewares

import (
	"context"
	"net/http"
	"time"

	"d7y.io/dragonfly/v2/manager/model"
	"d7y.io/dragonfly/v2/manager/service"
	"d7y.io/dragonfly/v2/manager/types"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func Jwt(service service.REST) (*jwt.GinJWTMiddleware, error) {
	identityKey := "id"

	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "Dragonfly",
		Key:         []byte("Secret Key"),
		Timeout:     2 * 24 * time.Hour,
		MaxRefresh:  2 * 24 * time.Hour,
		IdentityKey: identityKey,

		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)

			id, ok := claims[identityKey]
			if !ok {
				c.JSON(http.StatusUnauthorized, gin.H{
					"message": "Unavailable token: require user id",
				})
				c.Abort()
				return nil
			}

			c.Set("id", id)
			return id
		},

		Authenticator: func(c *gin.Context) (interface{}, error) {
			// Oauth2 signin
			if rawUser, ok := c.Get("user"); ok {
				user, ok := rawUser.(*model.User)
				if !ok {
					return "", jwt.ErrFailedAuthentication
				}
				return user, nil
			}

			// Normal signin
			var json types.SignInRequest
			if err := c.ShouldBindJSON(&json); err != nil {
				return "", jwt.ErrMissingLoginValues
			}

			user, err := service.SignIn(context.TODO(), json)
			if err != nil {
				return "", jwt.ErrFailedAuthentication
			}

			return user, nil
		},

		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if user, ok := data.(*model.User); ok {
				return jwt.MapClaims{
					identityKey: user.ID,
				}
			}

			return jwt.MapClaims{}
		},

		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"message": http.StatusText(code),
			})
		},

		LoginResponse: func(c *gin.Context, code int, token string, expire time.Time) {
			// Oauth2 signin
			if _, ok := c.Get("user"); ok {
				c.Redirect(http.StatusFound, "/")
				return
			}

			// Normal signin
			c.JSON(code, gin.H{
				"token":  token,
				"expire": expire.Format(time.RFC3339),
			})
		},

		LogoutResponse: func(c *gin.Context, code int) {
			c.Status(code)
		},

		RefreshResponse: func(c *gin.Context, code int, token string, expire time.Time) {
			c.JSON(code, gin.H{
				"token":  token,
				"expire": expire.Format(time.RFC3339),
			})
		},

		TokenLookup:    "cookie: jwt, header: Authorization, query: token",
		TokenHeadName:  "Bearer",
		TimeFunc:       time.Now,
		SendCookie:     true,
		CookieHTTPOnly: false,
	})

	if err != nil {
		return nil, err
	}

	return authMiddleware, nil
}
