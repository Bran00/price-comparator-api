package request

type RequestProductParams struct {
  Name  string `query:"name" validate:"required"`
}

