package models

type Department struct {
	ID        int    `json:"dep_id"`
	NameOfDep string `json:"name_of_dep"`
	HeadID    int    `json:"dep_head"`
	MemberIDs []int  `json:"member_ids"`
}
