package response

import (
	"altaStore/api/paginator"
	"altaStore/business/user"
)

type getAllUserResponse struct {
	Meta  paginator.Meta    `json:"meta"`
	Users []GetUserResponse `json:"users"`
}

//NewGetAllUserResponse construct GetAllUserResponse
func NewGetAllUserResponse(users []user.User, page int, rowPerPage int) getAllUserResponse {

	var (
		lenUsers = len(users)
	)

	getAllUserResponse := getAllUserResponse{}
	getAllUserResponse.Meta.BuildMeta(lenUsers, page, rowPerPage)

	for index, value := range users {
		if index == getAllUserResponse.Meta.RowPerPage {
			continue
		}

		var getUserResponse GetUserResponse

		getUserResponse.ID = value.ID
		getUserResponse.Name = value.Name
		getUserResponse.Email = value.Email
		getUserResponse.PhoneNumber = value.PhoneNumber
		getUserResponse.CreatedAt = value.CreatedAt
		getUserResponse.UpdatedAt = value.UpdatedAt
		getUserResponse.DeletedAt = value.DeletedAt

		getAllUserResponse.Users = append(getAllUserResponse.Users, getUserResponse)
	}

	if getAllUserResponse.Users == nil {
		getAllUserResponse.Users = []GetUserResponse{}
	}

	return getAllUserResponse
}
