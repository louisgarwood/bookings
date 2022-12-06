package render

import (
	"encoding/gob"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/louisgarwood/bookings/internal/config"
	"github.com/louisgarwood/bookings/internal/models"
)

var session *scs.SessionManager
var testApp config.AppConfig

type myWriter struct{}

func (tw *myWriter) Header() http.Header {
	return http.Header{}
}

func (tw *myWriter) WriteHeader(statusCode int) {

}

func (tw *myWriter) Write(bytes []byte) (int, error) {
	return len(bytes), nil
}

func TestMain(m *testing.M) {

	gob.Register(models.ReservationData{})

	testApp.IsProduction = false

	session = scs.New()
	session.Lifetime = 1 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false

	testApp.Session = session

	app = &testApp

	os.Exit(m.Run())
}
