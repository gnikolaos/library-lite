// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.778
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func LoginForm() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"auth h-full flex flex-col gap-10 lg:flex-row-reverse mt-8 lg:mt-0 justify-start lg:justify-center items-center\"><div hx-get=\"/\" hx-target=\"body\" hx-push-url=\"true\" class=\"w-64 h-fit cursor-pointer\"><img src=\"/img/logo/lili-high-resolution-logo-transparent.webp\" alt=\"logo\" class=\"h-full w-full\"></div><div class=\"flex-shrink-0 w-full max-w-md px-1 lg:px-0\" hx-ext=\"response-targets\"><form class=\"px-4 py-8 lg:px-8 lg:py-14 m-0 rounded-2xl shadow-xl bg-cream\" hx-post=\"/auth/login\" hx-target-error=\"#errorMessage\" hx-swap=\"innerHTML\" hx-indicator=\"#spinner\" hx-on:htmx:after-request=\"this.reset()\"><h1 class=\"text-5xl font-bold pb-6 text-center\">Log in now!</h1><div class=\"relative flex items-center py-2\"><label class=\"sr-only\" for=\"email\"><span class=\"label-text\">Email</span></label> <input id=\"email\" type=\"email\" name=\"email\" placeholder=\"Enter email\" class=\"appearance-none rounded-sm border-none w-full text-sm px-4 py-3 text-black leading-tight focus:outline-none focus:ring-accent focus:ring-2\" autocomplete=\"username\" required></div><div class=\"relative flex items-center py-2\"><label class=\"sr-only\" for=\"password\"><span class=\"label-text\">Password</span></label> <input id=\"password\" type=\"password\" name=\"password\" placeholder=\"Enter password\" class=\"appearance-none rounded-sm border-none w-full text-sm px-4 py-3 text-black leading-tight focus:outline-none focus:ring-accent focus:ring-2\" autocomplete=\"current-password\" required></div><div class=\"py-2 text-sm\"><label class=\"flex gap-2 items-center\" for=\"rememberMe\"><span>Remember me</span> <input id=\"rememberMe\" type=\"checkbox\" name=\"rememberMe\" class=\"w-4 rounded-sm focus:ring-accent focus:ring-4\"></label></div><div class=\"mt-6 flex flex-col gap-2\"><i id=\"errorMessage\" class=\"text-red-600 h-6 max-h-6 mb-1 text-sm\"></i> <button class=\"px-8 py-2.5 w-fit bg-accent text-sm text-white rounded-md hover:bg-accent focus:outline-white\" type=\"submit\" value=\"submit\">Submit <span id=\"spinner\" class=\"loading loading-spinner loading-sm htmx-indicator my-indicator\"></span></button></div></form></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func SignupForm() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"auth h-full flex flex-col gap-10 lg:flex-row-reverse mt-8 lg:mt-0 justify-start lg:justify-center items-center\"><div hx-get=\"/\" hx-target=\"body\" hx-push-url=\"true\" class=\"w-64 h-fit cursor-pointer\"><img src=\"/img/logo/lili-high-resolution-logo-transparent.webp\" alt=\"logo\" class=\"h-full w-full\"></div><div class=\"flex-shrink-0 w-full max-w-md px-1 lg:px-0\" hx-ext=\"response-targets\"><form class=\"px-4 py-8 lg:px-8 lg:py-14 m-0 rounded-2xl shadow-xl bg-cream\" hx-post=\"/auth/signup/\" hx-target-error=\"#errorMessage\" hx-swap=\"innerHTML\" hx-indicator=\"#spinner\" hx-on:htmx:after-request=\"this.reset()\"><h1 class=\"text-5xl font-bold pb-6 text-center\">Sing up now!</h1><div class=\"relative flex items-center py-2\"><label class=\"sr-only\" for=\"email\"><span class=\"label-text\">Email</span></label> <input id=\"email\" type=\"email\" name=\"email\" placeholder=\"Enter email\" class=\"appearance-none rounded-sm border-none w-full text-sm px-4 py-3 text-black leading-tight focus:outline-none focus:ring-accent focus:ring-2\" autocomplete=\"email\" required></div><div class=\"relative flex items-center py-2\"><label class=\"sr-only\" for=\"password\"><span class=\"label-text\">Password</span></label> <input id=\"password\" type=\"password\" name=\"password\" placeholder=\"Enter password\" class=\"appearance-none rounded-sm border-none w-full text-sm px-4 py-3 text-black leading-tight focus:outline-none focus:ring-accent focus:ring-2\" autocomplete=\"new-password\" required></div><div class=\"relative flex items-center py-2\"><label class=\"sr-only\" for=\"confirm-password\"><span class=\"label-text\">Password</span></label> <input id=\"confirm-password\" type=\"confirm-password\" name=\"confirm-password\" placeholder=\"Confirm password\" class=\"appearance-none rounded-sm border-none w-full text-sm px-4 py-3 text-black leading-tight focus:outline-none focus:ring-accent focus:ring-2\" required></div><div class=\"py-2 text-sm\"><label class=\"flex gap-2 items-center\" for=\"rememberMe\"><span>Remember me</span> <input id=\"rememberMe\" type=\"checkbox\" name=\"rememberMe\" class=\"w-4 rounded-sm focus:ring-accent focus:ring-4\"></label></div><div class=\"mt-6\"><i id=\"errorMessage\" class=\"text-red-600 h-6 max-h-6 mb-1\"></i> <button class=\"px-8 py-2.5 bg-accent text-sm text-white rounded-md hover:bg-accent/75 focus:outline-white\" type=\"submit\" value=\"submit\">Submit <span id=\"spinner\" class=\"loading loading-spinner loading-sm htmx-indicator my-indicator\"></span></button></div></form></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
