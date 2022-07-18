package jwt

import (
	"strings"
	"time"

	"github.com/ajikamaludin/go-grpc_basic/pkg/v1/utils/errors"

	"github.com/ajikamaludin/go-grpc_basic/pkg/v1/config"
	"github.com/ajikamaludin/go-grpc_basic/pkg/v1/utils/constants"
	"github.com/dgrijalva/jwt-go"
)

type CustomClaims struct {
	User      string `json:"user,omitempty"`
	IsRefresh bool   `json:"isRefresh,omitempty"`
	jwt.StandardClaims
}

type Token struct {
	Type           string
	Access         string
	ExpiredPeriode int64
	Refresh        string
}

func GenerateToken(config *config.Config, user string) (*Token, error) {
	atoken, err := set(user, false, constants.Jwt_Token_Expired_Periode, config)
	if err != nil {
		return nil, err
	}

	arefresh, err := set(user, true, constants.Jwt_Refresh_Expired_Periode, config)
	if err != nil {
		return nil, err
	}

	return &Token{
		Type:           config.Jwt.Type,
		Access:         atoken,
		ExpiredPeriode: constants.Jwt_Token_Expired_Periode,
		Refresh:        arefresh,
	}, nil
}

func set(user string, isrefresh bool, exp time.Duration, config *config.Config) (string, error) {
	// create refresh token
	claims := CustomClaims{
		User:      user,
		IsRefresh: isrefresh,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(exp * time.Second).Unix(),
			Issuer:    config.Jwt.Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	atoken, err := token.SignedString([]byte(config.Jwt.Key))
	if err != nil {
		return "", err
	}

	return atoken, nil
}

func ClaimToken(config *config.Config, auth string, isrefresh bool) (jwt.MapClaims, error) {
	var token *jwt.Token
	var err error

	if !isrefresh {
		// Bearer token as  RFC 6750 standard
		if strings.Split(auth, " ")[0] != config.Jwt.Type {
			return nil, errors.New("Invalid token")
		}

		token, err = claim(auth, config.Jwt.Key, false)
		if err != nil {
			return nil, err
		}
	} else {
		token, err = claim(auth, config.Jwt.Key, true)
		if err != nil {
			return nil, err
		}
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("failed to claim token")
	}

	// validate issuer
	if claims["iss"] != config.Jwt.Issuer {
		return nil, errors.New("Invalid token")
	}

	// validate refresh token
	if isrefresh {
		if claims["IsRefresh"] == false {
			return nil, errors.New("Invalid token")
		}
	} else {
		if claims["IsRefresh"] == true {
			return nil, errors.New("Invalid token")
		}
	}

	return claims, nil
}

func claim(auth, key string, isrefresh bool) (*jwt.Token, error) {
	if !isrefresh {
		auth = strings.Split(auth, " ")[1]
	}

	token, err := jwt.Parse(auth, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			errors.New("Signing method invalid")
		} else if method != jwt.SigningMethodHS256 {
			errors.New("Signing method invalid")
		}

		return []byte(key), nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}
