// Code generated by templ@(devel) DO NOT EDIT.

package httpdebug

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func list(uris []string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_1 := templ.GetChildren(ctx)
		if var_1 == nil {
			var_1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		// Element (standard)
		_, err = templBuffer.WriteString("<table>")
		if err != nil {
			return err
		}
		// Element (standard)
		_, err = templBuffer.WriteString("<tr>")
		if err != nil {
			return err
		}
		// Element (standard)
		_, err = templBuffer.WriteString("<th>")
		if err != nil {
			return err
		}
		// Text
		var_2 := `File`
		_, err = templBuffer.WriteString(var_2)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</th>")
		if err != nil {
			return err
		}
		// Element (standard)
		_, err = templBuffer.WriteString("<th>")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</th>")
		if err != nil {
			return err
		}
		// Element (standard)
		_, err = templBuffer.WriteString("<th>")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</th>")
		if err != nil {
			return err
		}
		// Element (standard)
		_, err = templBuffer.WriteString("<th>")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</th>")
		if err != nil {
			return err
		}
		// Element (standard)
		_, err = templBuffer.WriteString("<th>")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</th>")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</tr>")
		if err != nil {
			return err
		}
		// For
		for _, uri := range uris {
			// Element (standard)
			_, err = templBuffer.WriteString("<tr>")
			if err != nil {
				return err
			}
			// Element (standard)
			_, err = templBuffer.WriteString("<td>")
			if err != nil {
				return err
			}
			// StringExpression
			var var_3 string = uri
			_, err = templBuffer.WriteString(templ.EscapeString(var_3))
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</td>")
			if err != nil {
				return err
			}
			// Element (standard)
			_, err = templBuffer.WriteString("<td>")
			if err != nil {
				return err
			}
			// Element (standard)
			_, err = templBuffer.WriteString("<a")
			if err != nil {
				return err
			}
			// Element Attributes
			_, err = templBuffer.WriteString(" href=")
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("\"")
			if err != nil {
				return err
			}
			var var_4 templ.SafeURL = getMapURL(uri)
			_, err = templBuffer.WriteString(templ.EscapeString(string(var_4)))
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("\"")
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString(">")
			if err != nil {
				return err
			}
			// Text
			var_5 := `Mapping`
			_, err = templBuffer.WriteString(var_5)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</a>")
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</td>")
			if err != nil {
				return err
			}
			// Element (standard)
			_, err = templBuffer.WriteString("<td>")
			if err != nil {
				return err
			}
			// Element (standard)
			_, err = templBuffer.WriteString("<a")
			if err != nil {
				return err
			}
			// Element Attributes
			_, err = templBuffer.WriteString(" href=")
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("\"")
			if err != nil {
				return err
			}
			var var_6 templ.SafeURL = getSourceMapURL(uri)
			_, err = templBuffer.WriteString(templ.EscapeString(string(var_6)))
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("\"")
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString(">")
			if err != nil {
				return err
			}
			// Text
			var_7 := `Source Map`
			_, err = templBuffer.WriteString(var_7)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</a>")
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</td>")
			if err != nil {
				return err
			}
			// Element (standard)
			_, err = templBuffer.WriteString("<td>")
			if err != nil {
				return err
			}
			// Element (standard)
			_, err = templBuffer.WriteString("<a")
			if err != nil {
				return err
			}
			// Element Attributes
			_, err = templBuffer.WriteString(" href=")
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("\"")
			if err != nil {
				return err
			}
			var var_8 templ.SafeURL = getTemplURL(uri)
			_, err = templBuffer.WriteString(templ.EscapeString(string(var_8)))
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("\"")
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString(">")
			if err != nil {
				return err
			}
			// Text
			var_9 := `Templ`
			_, err = templBuffer.WriteString(var_9)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</a>")
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</td>")
			if err != nil {
				return err
			}
			// Element (standard)
			_, err = templBuffer.WriteString("<td>")
			if err != nil {
				return err
			}
			// Element (standard)
			_, err = templBuffer.WriteString("<a")
			if err != nil {
				return err
			}
			// Element Attributes
			_, err = templBuffer.WriteString(" href=")
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("\"")
			if err != nil {
				return err
			}
			var var_10 templ.SafeURL = getGoURL(uri)
			_, err = templBuffer.WriteString(templ.EscapeString(string(var_10)))
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("\"")
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString(">")
			if err != nil {
				return err
			}
			// Text
			var_11 := `Go`
			_, err = templBuffer.WriteString(var_11)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</a>")
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</td>")
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</tr>")
			if err != nil {
				return err
			}
		}
		_, err = templBuffer.WriteString("</table>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = io.Copy(w, templBuffer)
		}
		return err
	})
}
