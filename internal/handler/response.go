package handler

type ErrorResponse struct {
	Message string `example:"Something went wrong try again later!" json:"message"`
} // @name Error

type StatusResponse struct {
	Status string `example:"Everything is OK" json:"status"`
} // @name Status

type SignInRequest struct {
	Email    string `example:"john.doe@example.com" json:"email"`
	Password string `example:"qwerty" json:"password"`
} // @name SignInRequest

type SignUpRequest struct {
	Id                   uint   `example:"14" json:"id"`
	Name                 string `example:"John Doe" json:"name"`
	Email                string `example:"john.doe@example.com" json:"email"`
	Password             string `example:"qwerty" json:"password"`
	PasswordConfirmation string `example:"qwerty" json:"password_confirmation"`
} // @name SignUpRequest

type AuthResponse struct {
	Type         string `example:"Bearer" json:"type"`
	AccessToken  string `example:"eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJPbmxpbmUgSldUIEJ1aW.xkZXIiLCJpYXQiOjE3MjE4N" json:"access_token"`
	RefreshToken string `example:"eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJPbmxpbmUgSldUIEJ1aW.xkZXIiLCJpYXQiOjE3MjE4N" json:"refresh_token"`
} // @name AuthResponse

type RoleResponseRequest struct {
	DisplayTitle string   `example:"UserTitle" json:"display_title"`
	Permissions  []string `example:"prefix_create,prefix_read" json:"permissions"`
} // @name Role

type RoleResponse struct {
	Id           uint                  `example:"12" json:"id"`
	Title        string                `example:"RoleTitle" json:"title"`
	DisplayTitle string                `example:"UserTitle" json:"display_title"`
	Permissions  []PermissionsResponse `json:"permissions"`
} // @name RoleResponse

type RolesResponse struct {
	Data []RoleResponse `json:"data"`
} // @name RolesResponse

type AssignRoleToUserRequest struct {
	RoleId uint `example:"14" json:"role_id"`
} // @name AssignRoleToUserRequest

type PermissionsResponse struct {
	Data []string `example:"prefix_create,prefix_read,..." json:"data"`
} // @name PermissionsResponse

type CategoryRequest struct {
	Name string `example:"Transport" json:"name"`
} // @name CategoryRequest

type CategoryResponse struct {
	Name string `example:"Transport" json:"name"`
	Type string `example:"income|expense" json:"type"`
} // @name CategoryResponse

type CategoriesResponse struct {
	Data []CategoryResponse `json:"data"`
} // @name CategoriesResponse

type BudgetCreateRequest struct {
	Title        string  `example:"New car" json:"title"`
	Type         string  `example:"income" json:"type"`
	Amount       float64 `example:"10.25" json:"amount"`
	UserId       uint    `example:"12" json:"user_id"`
	CategorySlug string  `example:"test-category" json:"category_slug"`
} // @name BudgetCreateRequest

type BudgetUpdateRequest struct {
	Title  string  `example:"New car" json:"title"`
	Type   string  `example:"income" json:"type"`
	Amount float64 `example:"10.25" json:"amount"`
	UserId uint    `example:"12" json:"user_id"`
} // @name BudgetUpdateRequest

type BudgetResponse struct {
	Id     uint    `example:"10" json:"id"`
	Title  string  `example:"New budget" json:"title"`
	Type   string  `example:"income" json:"type"`
	Amount float64 `example:"1000" json:"amount"`
	UserID uint    `example:"10" json:"user_id"`
} //@name BudgetResponse

type BudgetsResponse struct {
	Data []BudgetResponse `json:"data"`
} //@name BudgetsResponse

type UserUpdateRequest struct {
	TgId     int64  `example:"131231231231323" json:"tg_id"`
	Name     string `example:"Jon Doe" json:"name"`
	Email    string `example:"jon_doe@gmail.com" json:"email"`
	Password string `example:"qwerty" json:"password"`
} // @name UserUpdateRequest

type UserResponse struct {
	Id    uint         `example:"14" json:"id"`
	TgId  int64        `example:"131231231231323" json:"tg_id"`
	Name  string       `example:"Jon Doe" json:"name"`
	Email string       `example:"jon_doe@gmail.com" json:"email"`
	Role  RoleResponse `json:"role"`
} // @name UserResponse
