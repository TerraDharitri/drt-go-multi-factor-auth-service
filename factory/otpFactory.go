package factory

import (
	"crypto"

	"github.com/TerraDharitri/drt-go-multi-factor-auth-service/config"
	"github.com/TerraDharitri/drt-go-multi-factor-auth-service/handlers"
	"github.com/TerraDharitri/drt-go-multi-factor-auth-service/handlers/secureOtp"
	"github.com/TerraDharitri/drt-go-multi-factor-auth-service/handlers/twofactor"
	"github.com/TerraDharitri/drt-go-multi-factor-auth-service/handlers/twofactor/sec51"
	"github.com/TerraDharitri/drt-go-multi-factor-auth-service/redis"
)

const hashType = crypto.SHA1

// CreateOTPHandler will create a new otp handler instance
func CreateOTPHandler(configs *config.Configs) (handlers.TOTPHandler, error) {
	otpProvider := sec51.NewSec51Wrapper(configs.GeneralConfig.TwoFactor.Digits, configs.GeneralConfig.TwoFactor.Issuer)
	return twofactor.NewTwoFactorHandler(otpProvider, hashType)
}

// CreateSecureOTPHandler will create a new otp handler instance
func CreateSecureOTPHandler(configs *config.Configs) (handlers.SecureOtpHandler, error) {
	rateLimiter, err := redis.CreateRedisRateLimiter(configs.ExternalConfig.Redis, configs.GeneralConfig.TwoFactor)
	if err != nil {
		return nil, err
	}

	secureOtpArgs := secureOtp.ArgsSecureOtpHandler{
		RateLimiter: rateLimiter,
	}
	return secureOtp.NewSecureOtpHandler(secureOtpArgs)
}
