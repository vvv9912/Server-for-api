package server

import (
	"context"
	"errors"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"html/template"
	"io"
	"path"
	"serverBot2/internal/api"
	"serverBot2/internal/handler"
)

type Server struct {
	echo *echo.Echo
}
type TemplateRegistry struct {
	templates map[string]*template.Template
}

// Implement e.Renderer interface
func (t *TemplateRegistry) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	tmpl, ok := t.templates[name]
	if !ok {
		err := errors.New("Template not found -> " + name)
		return err
	}
	return tmpl.ExecuteTemplate(w, "base.html", data)
}

func NewServer(storager api.Storager, storagerr api.PostStorager) *Server {
	s := &Server{echo: echo.New()}
	s.echo.Static("/static", "static")

	get := api.GetDB{Storage: storager}

	templates := make(map[string]*template.Template)

	templates["auth.html"] = template.Must(template.ParseFiles(path.Join("static", "templates", "auth.html"), path.Join("static", "templates", "base.html")))
	templates["auth2.html"] = template.Must(template.ParseFiles(path.Join("static", "templates", "auth2.html"), path.Join("static", "templates", "base.html")))
	templates["auth3.html"] = template.Must(template.ParseFiles(path.Join("static", "templates", "auth3.html"), path.Join("static", "templates", "base.html")))
	templates["bd.html"] = template.Must(template.ParseFiles(path.Join("static", "templates", "bd.html"), path.Join("static", "templates", "base.html")))

	s.echo.Renderer = &TemplateRegistry{templates: templates}

	s.echo.GET("/auth", handler.AuthHandler)
	s.echo.GET("/auth2", handler.AuthHandler2)
	s.echo.GET("/auth3", handler.AuthHandler3)
	s.echo.GET("/bd", handler.BdHandler)
	s.echo.GET("", handler.Hand)

	s.echo.POST("/api/post-auth", api.PostAuth)
	s.echo.GET("/api/get-data-db", get.GetDataDb)
	s.echo.GET("/api/get-data", api.GetData)
	s.echo.GET("/api/get-file-bd", get.GetDownloadDB)
	post := api.PostDB{Storage: storagerr}
	s.echo.POST("/api/save-change-bd/change", post.PostChangeBD) //добавить авторизацию
	s.echo.POST("/api/save-change-bd/add", post.PostNewAddBD)
	return s
}

func (s *Server) ServerStart(ctx context.Context, addr string) error {
	//err := s.echo.Start("172.17.0.2:8080")

	err := s.echo.Start(addr)
	defer s.echo.Close()
	return err
}
