package datatransfers

type TodoInsert struct {
	Activity string `json:"activity" binding:"required"`
	Priority string `json:"priority" binding:"required"`
}

type TodoUpdate struct {
	Activity string `json:"activity,omitempty" binding:"-"`
	Priority string `json:"priority,omitempty" binding:"-"`
	IsDone   bool   `json:"is_done,omitempty" binding:"-"`
}
