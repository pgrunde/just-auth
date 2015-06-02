package server

import "net/http"

func signin(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`
	<form action="/post-session" method="post">
		<div>
			<label for="signin_email">Email:</label>
			<input type="text" id="signin_email">
		</div>
		<div>
			<label for="password">Password:</label>
			<input type="text" id="password">
		</div>
		<div>
			<button type="submit">Sign in</button>
		</div>
	</form>
	`))
}

func createAccount(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`create account stub`))
}
