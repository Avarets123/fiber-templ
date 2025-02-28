// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.833
package layout

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func Menu() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = MenuStyle().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<div class=\"menu\"><div class=\"menu__logo\"><img src=\"/public/icons/bag.png\" alt=\"logo\" class=\"menu__logo-img\"><p class=\"menu__logo-text\">Моя работа</p></div><div class=\"menu__right\"><a class=\"menu__right-sign_in\" href=\"/\">Войти</a> <a class=\"menu__right-sign_in menu__right-sign_up\" href=\"/\">Зарегистрироваться</a></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

func MenuStyle() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var2 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var2 == nil {
			templ_7745c5c3_Var2 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "<style>\n    .menu {\n        display: flex;\n        justify-content: space-between;\n    }\n    .menu__right {\n        display: flex;\n        gap: 20px;\n        align-items: center;\n    }\n    .menu__right-link {\n        text-decoration: none;\n        color: var(--color-white);\n        font-size: 16px;\n    }\n    .menu__right-link:hover {\n        color: #9f9f9f;\n        cursor: move;\n    }\n    .menu__logo {\n        display: flex;\n        align-items: center;\n        gap: 10px;\n    }\n    .menu__logo-text {\n        color: var(--color-white);\n        font-size: 20px;\n        font-weight: 600;\n    }\n    .menu__logo-img {\n        width: 28px;\n        height: 28px;\n    }\n    .menu__right-sign_in {\n        font-size: 16px;\n        font-weight: 400;\n        color: var(--color-white);\n        text-decoration: none;\n\n    }\n    .menu__right-sign_up {\n        width: 208px;\n        height: 40px;\n        border-radius: 8px;\n        padding: 14px 20px;\n        background-color: var(--color-primary);\n        font-weight: 600;\n    }\n    .menu__right-sign_up:hover {\n        background-color: #19514a;\n    }\n\n    </style>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
