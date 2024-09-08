// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.771
package templates

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func Home(title string) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
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
		templ_7745c5c3_Var2 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
			templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
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
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<section><h1>Welcome to my ASCII art generator</h1><p>This is a simple website to convert images to ASCII art.</p><p>It allows you to upload an image and convert it to an ASCII text format.</p><p>At the moment, it supports common image formats like PNG, JPEG, JPG, and WEBP.</p><p>The output is black and white text at the moment. In the future, it will allow you to choose the background color and/or conserve the colors of the original image.</p></section><!-- Form to upload an image and generate ASCII --> <form hx-post=\"/convert-to-ascii\" hx-target=\"#output\" enctype=\"multipart/form-data\"><label for=\"image\">UPLOAD IMAGE <input type=\"file\" id=\"image\" name=\"image\" accept=\"image/*\" alt=\"Upload image for ASCII conversion\" required></label> <button type=\"submit\">Generate ASCII</button></form><!-- Output container to display ASCII art and download button --> <div id=\"output-container\"><div id=\"output\"></div></div><!-- Content editable div to paste image\r\n        <div id=\"paste-container\" contenteditable=\"true\"></div> --> <!-- Script to handle paste event\r\n        <script>\r\n            document.getElementById('paste-container').addEventListener('paste', async (event) => {\r\n                event.preventDefault();\r\n                try {\r\n                    const clipboardItems = await navigator.clipboard.read();\r\n                    for (const clipboardItem of clipboardItems) {\r\n                        if (clipboardItem.kind === 'image') {\r\n                            const blob = await clipboardItem.getType(clipboardItem.types[0]);\r\n                            const file = new File([blob], 'pasted_image.' + clipboardItem.types[0].split('/')[1], {\r\n                                type: clipboardItem.types[0],\r\n                            });\r\n                            // Create a new FormData object to send the pasted image to the server\r\n                            const formData = new FormData();\r\n                            formData.append('image', file);\r\n                            // Send the formData to the server using the same endpoint as the upload form\r\n                            fetch('/convert', {\r\n                                method: 'POST',\r\n                                body: formData,\r\n                            })\r\n                            .then(response => response.text())\r\n                            .then((asciiArt) => {\r\n                                // Display the generated ASCII art\r\n                                document.getElementById('output').innerHTML = asciiArt;\r\n                            })\r\n                            .catch((error) => {\r\n                                console.error('Error generating ASCII art:', error);\r\n                            });\r\n                        }\r\n                    }\r\n                } catch (error) {\r\n                    console.error('Error reading clipboard:', error);\r\n                }\r\n            });\r\n        </script> -->")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = Base(title).Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
