package httperr

import "net/http"

const (
	// 📝 요청 관련 에러 (Request Errors)
	// 요청 파싱 실패 - 400 BadRequest
	ErrorRequestParsingFailed = "ERROR_REQUEST_PARSING_FAILED"
	// 유효하지 않은 Content-Type - 415 Unsupported Media Type
	ErrorRequestInvalidContentType = "ERROR_REQUEST_INVALID_CONTENT_TYPE"

	// 🔐 인증 및 권한 에러 (Authentication and Authorization Errors)
	// 인증 실패 - 401 Unauthorized
	ErrorAuthFailed = "ERROR_AUTH_FAILED"
	// 토큰 만료 - 401 Unauthorized
	ErrorAuthTokenExpired = "ERROR_AUTH_TOKEN_EXPIRED"
	// 권한 없음 - 403 Forbidden
	ErrorAuthPermissionDenied = "ERROR_AUTH_PERMISSION_DENIED"

	// ✅ 유효성 검증 에러 (Validation Errors)
	// 유효성 검증 실패 - 400 BadRequest
	ErrorValidationFailed = "ERROR_VALIDATION_FAILED"
	// 필수 필드 누락 - 400 BadRequest
	ErrorValidationRequiredFieldMissing = "ERROR_VALIDATION_REQUIRED_FIELD_MISSING"
	// 형식 오류 - 400 BadRequest
	ErrorValidationInvalidFormat = "ERROR_VALIDATION_INVALID_FORMAT"

	// 🗃️ 데이터베이스 에러 (Database Errors)
	// 데이터베이스 오류 - 500 Internal Server Error
	ErrorDatabaseFailed = "ERROR_DATABASE_FAILED"
	// 중복 데이터 - 400 BadRequest
	ErrorDatabaseDuplicateEntry = "ERROR_DATABASE_DUPLICATE_ENTRY"
	// 데이터 없음 - 404 NotFound
	ErrorDatabaseRecordNotFound = "ERROR_DATABASE_RECORD_NOT_FOUND"

	// ⚙️ 서버 에러 (Server Errors)
	// 서버 오류 - 500 Internal Server Error
	ErrorServerFailed = "ERROR_SERVER_FAILED"
	// 서버 타임아웃 - 504 Gateway Timeout
	ErrorServerTimeout = "ERROR_SERVER_TIMEOUT"
	// 서버 과부하 - 503 Service Unavailable
	ErrorServerOverload = "ERROR_SERVER_OVERLOAD"

	// 🌐 외부 서비스 에러 (External Service Errors)
	// 외부 서비스 오류 - 502 Bad Gateway
	ErrorExternalServiceFailed = "ERROR_EXTERNAL_SERVICE_FAILED"
	// 외부 서비스 타임아웃 - 504 Gateway Timeout
	ErrorExternalServiceTimeout = "ERROR_EXTERNAL_SERVICE_TIMEOUT"
	// 외부 서비스 사용 불가 - 503 Service Unavailable
	ErrorExternalServiceUnavailable = "ERROR_EXTERNAL_SERVICE_UNAVAILABLE"
)

func statusOf(identifier string) int {
	switch identifier {
	// 📝 요청 관련 에러 (Request Errors)
	case ErrorRequestParsingFailed:
		return http.StatusBadRequest // 400
	case ErrorRequestInvalidContentType:
		return http.StatusUnsupportedMediaType // 415

	// 🔐 인증 및 권한 에러 (Authentication and Authorization Errors)
	case ErrorAuthFailed, ErrorAuthTokenExpired:
		return http.StatusUnauthorized // 401
	case ErrorAuthPermissionDenied:
		return http.StatusForbidden // 403

	// ✅ 유효성 검증 에러 (Validation Errors)
	case ErrorValidationFailed, ErrorValidationRequiredFieldMissing, ErrorValidationInvalidFormat:
		return http.StatusBadRequest // 400

	// 🗃️ 데이터베이스 에러 (Database Errors)
	case ErrorDatabaseFailed:
		return http.StatusInternalServerError // 500
	case ErrorDatabaseDuplicateEntry:
		return http.StatusBadRequest // 400
	case ErrorDatabaseRecordNotFound:
		return http.StatusNotFound // 404

	// ⚙️ 서버 에러 (Server Errors)
	case ErrorServerFailed:
		return http.StatusInternalServerError // 500
	case ErrorServerTimeout:
		return http.StatusGatewayTimeout // 504
	case ErrorServerOverload:
		return http.StatusServiceUnavailable // 503

	// 🌐 외부 서비스 에러 (External Service Errors)
	case ErrorExternalServiceFailed:
		return http.StatusBadGateway // 502
	case ErrorExternalServiceTimeout:
		return http.StatusGatewayTimeout // 504
	case ErrorExternalServiceUnavailable:
		return http.StatusServiceUnavailable // 503

	// 기본값: 알 수 없는 에러 식별자
	default:
		return http.StatusInternalServerError // 500
	}
}
