// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.663
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func PostNew() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"container flex flex-col items-center\"><form action=\"/posts/new\" method=\"POST\" class=\"shadow-md rounded gap-4\"><div class=\"flex-col mb-4\"><label for=\"title\" class=\"text-gray-700 text-sm font-bold my-2\">Title</label> <input required type=\"text\" class=\"shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline\" name=\"title\"></div><div class=\"flex-col mb-4\"><label for=\"desc\" class=\"text-gray-700 text-sm font-bold my-2\">Description</label> <input required type=\"text\" class=\"shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline\" name=\"desc\"></div><div class=\"flex-col mb-4\"><label for=\"content\" class=\"text-gray-700 text-sm font-bold my-2\">Content</label> <textarea required class=\"shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline\" name=\"content\"></textarea></div><div class=\"flex-col mb-4\"><label for=\"imgurl\" class=\"text-gray-700 text-sm font-bold my-2\">Image Url</label> <input required type=\"text\" class=\"shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline\" name=\"imgurl\"></div><div class=\"flex-col mb-4\"><label class=\"text-gray-700 text-sm font-bold my-2\">Tags</label><div name=\"tag-list\"></div><div class=\"flex flex-col\"><input class=\"shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline\" name=\"tag\"> <button hx-post=\"/tag-add\" hx-target=\"[name=&#39;tag-list&#39;]\" hx-swap=\"beforeend\">Add</button></div></div><div class=\"flex items-center justify-center\"><button type=\"submit\" class=\"bg-transparent hover:bg-blue-500 text-blue-700 font-semibold hover:text-white py-2 px-4 border border-blue-500 hover:border-transparent rounded\">Submit</button></div></form></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
