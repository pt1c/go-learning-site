package render

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}

func RenderTemplate(rw http.ResponseWriter, tempName string) {

	templateCache, err := CreateTemplateCache()
	if err != nil {
		log.Fatal("Ошибка при построении кэша")
	}

	template, found := templateCache[tempName]
	if !found {
		log.Fatal("Не найден запрашиваемый шаблон в кэше")
	}

	err = template.Execute(rw, nil)
	if err != nil {
		log.Println("Ошибка запуска шаблона")
		return
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	tmplCache := map[string]*template.Template{}

	pages, err := filepath.Glob("../../templates/*.page.html")
	if err != nil {
		log.Println("Не удалось считать страницы")
		return tmplCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			log.Println("Не удалось спарсить страницы")
			return tmplCache, err
		}

		matches, err := filepath.Glob("../../templates/*.layout.html")
		if err != nil {
			log.Println("Не удалось считать лайауты")
			return tmplCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("../../templates/*.layout.html")
			if err != nil {
				log.Println("Не удалось считать лайауты")
				return tmplCache, err
			}
		}

		tmplCache[name] = ts
	}

	return tmplCache, nil
}
