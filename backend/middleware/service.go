package middleware

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"net/http"
	"time"

	"github.com/goawwer/devclash/internal/dto"
	"github.com/goawwer/devclash/pkg/logger"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var (
	ErrInvalidToken       = errors.New("invalid token")
	ErrExpiredToken       = errors.New("expired token")
	ErrReusedRefreshToken = errors.New("refresh token reused")
	signingMethod         = jwt.SigningMethodHS256
)

var auth *Config

func InitAuthConfig(key string, r databaseIntegration) {
	auth = &Config{key, r}
}

func GenerateTokenPair(ctx context.Context, accountID uuid.UUID, role string) (*TokenPair, error) {
	now := time.Now()

	accessExp := now.Add(time.Minute * 5)
	refreshExp := now.Add(time.Hour * 24 * 1)

	accessClaims := CustomClaims{
		AccountID: accountID,
		Role:      role,
		TokenType: "access",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(accessExp),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
		},
	}

	accessToken := jwt.NewWithClaims(signingMethod, accessClaims)
	accessSigned, err := accessToken.SignedString([]byte(auth.Secret))
	if err != nil {
		logger.Error("failed to sign access token", "error", err)
		return nil, err
	}

	refreshID := uuid.New()
	refreshClaims := CustomClaims{
		AccountID: accountID,
		Role:      role,
		TokenType: "refresh",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(refreshExp),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			ID:        refreshID.String(),
		},
	}

	refreshToken := jwt.NewWithClaims(signingMethod, refreshClaims)
	refreshSigned, err := refreshToken.SignedString([]byte(auth.Secret))
	if err != nil {
		logger.Error("failed to sign refresh token", "error", err)
		return nil, err
	}

	hash := hashRefreshToken([]byte(auth.Secret), []byte(refreshSigned))
	tokenHash := base64.StdEncoding.EncodeToString(hash)

	err = auth.Repository.StoreRefreshToken(ctx, dto.RefreshTokenRecord{
		ID:        refreshID,
		AccountID: accountID,
		TokenHash: tokenHash,
		ExpiresAt: refreshExp,
	})
	if err != nil {
		logger.Error("failed to store refresh token: ", err)
		return nil, err
	}

	return &TokenPair{
		AccessToken:     accessSigned,
		RefreshToken:    refreshSigned,
		AccessTokenExp:  accessExp,
		RefreshTokenExp: refreshExp,
	}, nil
}

func ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(t *jwt.Token) (any, error) {
		return []byte(auth.Secret), nil
	})

	if err != nil || !token.Valid {
		return nil, ErrInvalidToken
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok {
		return nil, ErrInvalidToken
	}

	if claims.ExpiresAt != nil && claims.ExpiresAt.Time.Before(time.Now()) {
		return nil, ErrExpiredToken
	}

	return claims, nil
}

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("access")
		if err != nil {
			logger.Error("middleware wrapper cannot take token: ", err)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		claims, err := ParseToken(cookie.Value)
		if err != nil {
			if errors.Is(err, ErrExpiredToken) {
				logger.Info("access token expired:", claims.AccountID)
			} else {
				logger.Error("invalid access token:", err)
			}
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), ClaimsKey, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func SetAuthCookie(w http.ResponseWriter, setup CustomCookie) {
	http.SetCookie(w, &http.Cookie{
		Name:     setup.Name,
		Value:    setup.Value,
		HttpOnly: true,
		Path:     "/",
		Expires:  time.Now().Add(setup.Exp),
	})
}

func RefreshToken(r *http.Request) (*TokenPair, error) {
	cookie, err := r.Cookie("refresh")
	if err != nil {
		logger.Error("no refresh token: ", err)
		return nil, err
	}

	claims, err := ParseToken(cookie.Value)
	if err != nil || claims.TokenType != "refresh" {
		logger.Error("invalid refresh token: ", err)
		return nil, err
	}

	refreshID := claims.ID

	hash := hashRefreshToken([]byte(auth.Secret), []byte(cookie.Value))
	tokenHash := base64.StdEncoding.EncodeToString(hash)

	ok, _, err := auth.Repository.ConsumeRefreshToken(
		r.Context(),
		uuid.MustParse(refreshID),
		claims.AccountID,
		tokenHash,
	)
	if err != nil {
		logger.Error("failed to consume refresh token: ", err)
		return nil, err
	}
	if !ok {
		logger.Error("refresh token reused: ", err)
		return nil, ErrReusedRefreshToken
	}

	pair, err := GenerateTokenPair(r.Context(), claims.AccountID, claims.Role)
	if err != nil {
		logger.Error("failed to generate token pair: ", err)
		return nil, err
	}

	return pair, nil
}

func hashRefreshToken(secret, token []byte) []byte {
	mac := hmac.New(sha256.New, secret)
	mac.Write(token)
	return mac.Sum(nil)
}
