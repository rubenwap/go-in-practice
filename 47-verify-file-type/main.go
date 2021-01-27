/*
Three alternatives
*/

file, header, err:= r.FormFile("file")
contenttype := header.Header["Content-Type"][0]


file, header, err := r.FormFile("file")
extension := filepath.Ext(header.Filename)
type := mime.TypeByExtension(extension)


file, header, err := r.FormFile("file")
buffer := make([]byte, 512)
_, err = file.Read(buffer)
filetype := http.DetectContentType(buffer)
