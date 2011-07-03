package app

app.go is a simple web framework for use in Google AppEngine. Just copy the app folder to your working folder and import it from your program. That's it. A web application ready to run in no time. Also, app.go comes with a powerful datastore manager to simplify your interactions with BigTable.


Here is the Guestbook example from AppEngine rewritten using app.go

    package hello

    import "app"

    var DB = &app.DB

    type Greeting struct {
        Author  string
        Content string
        Date    int64
    }

    func init() {
      views := app.Views{
        "index" : index,
        "sign"  : sign,
      }
      app.Run(views)
    }

    func index(self app.Context) {
        recs := make([]Greeting, 0, 10)
        qry  := DB.Query("Greeting").Order("-Date").Limit(10)
        DB.Select(qry,&recs)
        self.Render("index",recs)
    }

    func sign(self app.Context) {
        rec := Greeting{
            Author : self.User.Nick,
            Content: self.GetValue("content"),
            Date   : DB.Now(),
        }
        DB.New(&rec)
        self.Redirect("/")
    }

Enjoy!