package utils

import "net/http"

type AppError struct {
	errMessage string
	statusCode int
}

func (e *AppError) Error() string {
	return e.errMessage
}

func (e *AppError) GetStatusCode() int {
	return e.statusCode
}

func (e *AppError) SetErrMessage(msg string) {
	e.errMessage = msg
}

func (e *AppError) SetStatusCode(code int) {
	e.statusCode = code
}

var (
	EncodingResponseError = &AppError{
		errMessage: "error encoding response",
		statusCode: http.StatusInternalServerError}
	ConnectingToDatabaseError = &AppError{
		errMessage: "error connecting to database",
		statusCode: http.StatusInternalServerError}
	DatabaseQueryError = &AppError{
		errMessage: "database query error",
		statusCode: http.StatusInternalServerError}
	UnknownInternalServerError = &AppError{
		errMessage: "unknown internal server error",
		statusCode: http.StatusInternalServerError}
	InvalidRequestPayloadError = &AppError{
		errMessage: "invalid request payload",
		statusCode: http.StatusBadRequest}
	UnitNotFoundError = &AppError{
		errMessage: "unit not found",
		statusCode: http.StatusNotFound}
	StartingTransactionError = &AppError{
		errMessage: "error starting database transaction",
		statusCode: http.StatusInternalServerError}
	CommitingTransactionError = &AppError{
		errMessage: "error commiting transaction",
		statusCode: http.StatusInternalServerError}
)
