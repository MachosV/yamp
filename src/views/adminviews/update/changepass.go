package update

import (
	"datastorage"
	"messages"
	"middleware"
	"net/http"
	"views"

	"golang.org/x/crypto/bcrypt"
)

func init() {
	views.GetMux().HandleFunc("/changepass", middleware.WithMiddleware(ChangePass,
		middleware.Time(),
		middleware.NeedsSession(),
		middleware.CsrfProtection(),
		middleware.IsAdmin(),
	))
}

/*
ChangePass function
is used to change the password of an account
if it ever gets lost.
*/
func ChangePass(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	password1 := r.PostFormValue("password1")
	password2 := r.PostFormValue("password2")
	if password1 != password2 {
		messages.SetMessage(r, "Οι 2 κωδικοί χρήστη δεν ταιριάζουν")
		http.Redirect(w, r, "/retrieveuser?id="+username, http.StatusMovedPermanently)
		return
	}
	hash, _ := bcrypt.GenerateFromPassword([]byte(password1), bcrypt.DefaultCost)
	stmt := datastorage.GetDataRouter().GetStmt("update_password")
	_, err := stmt.Exec(hash, username)
	if err != nil {
		messages.SetMessage(r, "Σφάλμα κατά την αλλαγή κωδικού χρήστη")
		http.Redirect(w, r, "/retrieveuser?id="+username, http.StatusMovedPermanently)
		return
	}
	messages.SetMessage(r, "Ο κωδικός τροποποιήθηκε επιτυχώς!")
	http.Redirect(w, r, "/retrieveuser?id="+username, http.StatusMovedPermanently)
}
