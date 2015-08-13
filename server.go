package main

import (
	"gitlab.dian.fm/livecloud/config-server/pkg/database"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/encoder"

	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

type NewError struct {
	errcode int
	errmsg  string
}

func GetServer(enc encoder.Encoder, db database.DB, parms martini.Params) (int, []byte) {
	fmt.Printf("begin get server\n")
	id, err := strconv.ParseInt(parms["id"], 10, 64)
	al := db.Get(id)
	if err != nil || al == nil {
		msg := fmt.Sprintf("the album with id %s does not exist", parms["id"])
		return http.StatusNotFound, encoder.Must(enc.Encode(
			NewError{errcode: 404, errmsg: msg}))
	}
	return http.StatusOK, encoder.Must(enc.Encode(al))
}

func FindServer(enc encoder.Encoder, db database.DB, r *http.Request) (int, []byte) {
	fmt.Printf("begin find server\n")
	params := r.URL.Query()
	roomId, err := strconv.ParseInt(params.Get("room_id"), 10, 64)
	al := db.Find(roomId)
	if err != nil || al == nil {
		msg := fmt.Sprintf("server witch room_id %s does not exist", params.Get("id"))
		return http.StatusNotFound, encoder.Must(enc.Encode(
			NewError{errcode: 404, errmsg: msg}))
	}
	return http.StatusOK, encoder.Must(enc.Encode(al))
}

var rxExt = regexp.MustCompile(`(\.(?:xml|json))\/?$`)

func MapEncoder(c martini.Context, w http.ResponseWriter, r *http.Request) {
	// Get the format extension
	matches := rxExt.FindStringSubmatch(r.URL.Path)
	ext := ".json"
	if len(matches) > 1 {
		// Rewrite the URL without the format extension
		lentghWithoutExt := len(r.URL.Path) - len(matches[1])
		if strings.HasSuffix(r.URL.Path, "/") {
			lentghWithoutExt--
		}
		r.URL.Path = r.URL.Path[:lentghWithoutExt]
		ext = matches[1]
	}
	// Inject the requested encoder
	switch ext {
	case ".xml":
		c.MapTo(encoder.XmlEncoder{}, (*encoder.Encoder)(nil))
		w.Header().Set("Content-Type", "application/xml")
	default:
		c.MapTo(encoder.JsonEncoder{}, (*encoder.Encoder)(nil))
		w.Header().Set("Content-Type", "application/json")
	}
}

func init() {
}

func main() {
	fmt.Printf("begin \n")

	m := martini.New()
	m.Use(martini.Recovery())
	m.Use(martini.Logger())
	m.Use(MapEncoder)

	router := martini.NewRouter()
	router.Get(`/servers`, FindServer)
	router.Get(`/servers/:id`, GetServer)
	//    router.Post(`/servers`, AddAlbum)
	//    router.Put(`/servers/:id`, UpdateAlbum)
	//    router.Delete(`/servers/:id`, DeleteAlbum)

	db := database.NewServerDB()
	m.MapTo(&db, (*database.DB)(nil))

	m.Action(router.Handle)
	m.Run()

	fmt.Printf("end \n")

}
