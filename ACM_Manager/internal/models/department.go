package models

type Department struct {
	NameOfDep string `db:"name_of_dep" json:"name_of_dep"`
	HeadID    int    `db:"head_id" json:"dep_head"`
}
