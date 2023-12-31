// Code generated by hertz generator. DO NOT EDIT.

package comment

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	comment "github.com/cold-runner/skylark/biz/handler/comment"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	{
		_comment := root.Group("/comment", _commentMw()...)
		_comment.POST("/create", append(_createcommentMw(), comment.CreateComment)...)
		_comment.DELETE("/delete", append(_deletecommentMw(), comment.DeleteComment)...)
		_comment.GET("/get", append(_getcommentsMw(), comment.GetComments)...)
		_comment.GET("/getReply", append(_getreplycommentsMw(), comment.GetReplyComments)...)
		_comment.PUT("/update", append(_updatecommentMw(), comment.UpdateComment)...)
	}
}
