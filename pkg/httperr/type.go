package httperr

import "net/http"

const (
	// 📝 요청 관련 에러 (Request Errs)
	// 요청 파싱 실패 - 400 BadRequest
	RequestParsingFailed = "ERROR_REQUEST_PARSING_FAILED"
	// 유효하지 않은 Content-Type - 415 Unsupported Media Type
	RequestInvalidContentType = "ERROR_REQUEST_INVALID_CONTENT_TYPE"

	// 🔐 인증 및 권한 에러 (Authentication and Authorization Errs)
	// 인증 실패 - 401 Unauthorized
	AuthFailed = "ERROR_AUTH_FAILED"
	// 권한 없음 - 403 Forbidden
	AuthPermissionDenied = "ERROR_AUTH_PERMISSION_DENIED"

	// ✅ 유효성 검증 에러 (Validation Errs)
	// 유효성 검증 실패 - 400 BadRequest
	ValidationFailed = "ERROR_VALIDATION_FAILED"
	// 필수 필드 누락 - 400 BadRequest
	ValidationRequiredFieldMissing = "ERROR_VALIDATION_REQUIRED_FIELD_MISSING"
	// 형식 오류 - 400 BadRequest
	ValidationInvalidFormat = "ERROR_VALIDATION_INVALID_FORMAT"

	// 토큰 에러 (Token Errs)

	// 🗃️ 데이터베이스 에러 (Database Errs)
	// 데이터베이스 오류 - 500 Internal Server Err
	DatabaseFailed = "ERROR_DATABASE_FAILED"
	// 중복 데이터 - 400 BadRequest
	DatabaseDuplicateEntry = "ERROR_DATABASE_DUPLICATE_ENTRY"
	// 데이터 없음 - 404 NotFound
	DatabaseRecordNotFound = "ERROR_DATABASE_RECORD_NOT_FOUND"

	// ⚙️ 서버 에러 (Server Errs)
	// 서버 오류 - 500 Internal Server Err
	ServerInternalError = "ERROR_SERVER_INTERNAL_ERROR"
	// 서버 타임아웃 - 504 Gateway Timeout
	ServerTimeout = "ERROR_SERVER_TIMEOUT"
	// 서버 과부하 - 503 Service Unavailable
	ServerOverload = "ERROR_SERVER_OVERLOAD"

	// 🌐 외부 서비스 에러 (External Service Errs)
	// 외부 서비스 오류 - 502 Bad Gateway
	ExternalServiceFailed = "ERROR_EXTERNAL_SERVICE_FAILED"
	// 외부 서비스 타임아웃 - 504 Gateway Timeout
	ExternalServiceTimeout = "ERROR_EXTERNAL_SERVICE_TIMEOUT"
	// 외부 서비스 사용 불가 - 503 Service Unavailable
	ExternalServiceUnavailable = "ERROR_EXTERNAL_SERVICE_UNAVAILABLE"
)

func statusOf(identifier string) int {
	switch identifier {
	// 📝 요청 관련 에러 (Request Errs)
	case RequestParsingFailed:
		return http.StatusBadRequest // 400
	case RequestInvalidContentType:
		return http.StatusUnsupportedMediaType // 415

	// 🔐 인증 및 권한 에러 (Authentication and Authorization Errs)
	case AuthFailed:
		return http.StatusUnauthorized // 401
	case AuthPermissionDenied:
		return http.StatusForbidden // 403

	// ✅ 유효성 검증 에러 (Validation Errs)
	case ValidationFailed, ValidationRequiredFieldMissing, ValidationInvalidFormat:
		return http.StatusBadRequest // 400

	// 토큰 에러 (Token Errs)

	// 🗃️ 데이터베이스 에러 (Database Errs)
	case DatabaseFailed:
		return http.StatusInternalServerError // 500
	case DatabaseDuplicateEntry:
		return http.StatusBadRequest // 400
	case DatabaseRecordNotFound:
		return http.StatusNotFound // 404

	// ⚙️ 서버 에러 (Server Errs)
	case ServerInternalError:
		return http.StatusInternalServerError // 500
	case ServerTimeout:
		return http.StatusGatewayTimeout // 504
	case ServerOverload:
		return http.StatusServiceUnavailable // 503

	// 🌐 외부 서비스 에러 (External Service Errs)
	case ExternalServiceFailed:
		return http.StatusBadGateway // 502
	case ExternalServiceTimeout:
		return http.StatusGatewayTimeout // 504
	case ExternalServiceUnavailable:
		return http.StatusServiceUnavailable // 503

	// 기본값: 알 수 없는 에러 식별자
	default:
		return http.StatusInternalServerError // 500
	}
}
