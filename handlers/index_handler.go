package handlers

import "github.com/martini-contrib/render"

/**
 * Обработчик запроса на корень сайта
 */
func IndexHandler(rnd render.Render) {
	rnd.HTML(200, "layout", "")
}