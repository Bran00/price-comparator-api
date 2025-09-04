package request

type RequestProductParams struct {
  Name  string `form:"name" validate:"required"`
}

