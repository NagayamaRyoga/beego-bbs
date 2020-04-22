<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>BBS</title>
    <link rel="stylesheet" href="/static/css/styles.css">
  </head>
  <body>
    <header>
      <nav class="nav">
        <div class="nav__container">
          <a class="nav__brand" href="/">BBS</a>
        </div>
      </nav>
    </header>
    <main class="main">
      <form class="form" action="/" method="POST">
        <input class="form__input" name="author" type="text" placeholder="名前">
        <textarea class="form__textarea" name="body" placeholder="本文"></textarea>
        <button class="form__button" type="submit">投稿</button>
      </form>
      <p>{{.Error}}</p>
      <div class="posts">
        {{ range $post := .Posts }}
        <div class="posts__item">
          <div class="posts__item-header">
            <span class="posts__item-id">{{ $post.ID }}.</span>
            <span class="posts__item-author">{{ $post.Author }}</span>
          </div>
          <div class="posts__item-body">
            <p>{{ $post.Body }}</p>
          </div>
          <div class="posts__item-footer">
            <span class="posts__item-datetime">{{ $post.CreatedAt }}</span>
          </div>
        </div>
        {{end}}
      </div>
    </main>
  </body>
</html>
