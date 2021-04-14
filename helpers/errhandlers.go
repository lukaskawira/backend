package helpers

import (
	"errors"
)

//http code
const BadRequest = 400
const Unauthorized = 401
const PageNotFound = 404
const InternalServerError = 500
const NotImplemented = 501

//command string
const Post = "post"
const Put = "put"
const Get = "get"
const Delete = "delete"

const PostMessage = "Error inserting data to database"
const PutMessage = "Error updating data to database"
const GetMessage = "Error getting data from database"
const DeleteMessage = "Error deleting data from database"

func ErrorMessage(command string) error {
	switch command {
	case Post:
		return errors.New(PostMessage)
	case Put:
		return errors.New(PutMessage)
	case Get:
		return errors.New(GetMessage)
	case Delete:
		return errors.New(DeleteMessage)
	default:
		return errors.New("Bad Request")
	}
}

func ErrorCode(err string) int {
	switch err {
	case PostMessage:
		return 500
	case PutMessage:
		return 500
	case GetMessage:
		return 500
	case DeleteMessage:
		return 500
	default:
		return 400
	}
}
