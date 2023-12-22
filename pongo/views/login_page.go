package views

import (
	"example.com/pongodemo/models"
	"example.com/pongodemo/views/widget"
)

func NewLoginForm(value models.Login) widget.Form {
	return widget.Form{
		Disabled: false,
		Method:   widget.MethodPost,
		Action:   "",
		Fields: []widget.FormControl{
			{
				Label:       "邮箱",
				ID:          "email",
				Type:        widget.ControlTypeEmail,
				Name:        "email",
				Value:       value.Email,
				Placeholder: "admin@example.org",
				Required:    true,
			},
			{
				Label:    "密码",
				ID:       "password",
				Type:     widget.ControlTypePassword,
				Name:     "password",
				Value:    "",
				Required: true,
			},
		},
		SubmitBtn: widget.PrimaryBlockBtn.
			SetName("登录").
			SetDisabledText("正在登录..."),
		CancelBtn: widget.Link{},
		DeleteBtn: widget.Link{},
	}
}
