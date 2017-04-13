# Laravel - 設計一個好的 Blade Template, 使用 @parent

做網站的都知道，一個 HTML 最佳的讀取順序應該是

1. meta tag
1. css
1. body
1. js

在 Laravel 裡我們會把版面都用到的地方切成 layout 出來，如 基本的 HTML Tag, Meta Tag, CSS, Javascript 等等，而因為 js 要給子頁面用，所以都會把 js 移到上面來。

```html
<!-- layout.blade.php -->
<!DOCTYPE html>
<html>
    <head>
        <title>Laravel 5 Example - @yield('title')</title>
        <meta charset="utf-8">
        <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.5/css/bootstrap.min.css">
        <script src="https://ajax.googleapis.com/ajax/libs/jquery/1.11.3/jquery.min.js"></script>
        <script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.5/js/bootstrap.min.js"></script>
    </head>
    <body>
        <header>
        </header>

        <nav>
        </nav>

        <section>
            @yield('content')
        </section>

        <footer>
        </footer>

    </body>
</html>
```

而子頁面會這樣寫。

```html
<!-- index.blade.php -->
@extends('layout')

@section('title', 'index')

@section('content')
<style>
    div {
        text-align: center;
    }
</style>

<div>
    Home Page
</div>

<script>
    $(function() {
    });
</script>
@endsection
```

但這樣會讓 Laravel 在產生 view 時，css 與 js 在 section 區塊裡面出現，如果頁面又 include 一些區塊 blade 的話，整個產生的 HTML 會變得很糟很難看，body 裡面 html tag, css, js 此起彼落的出現，不僅難看也影響效能。

這時我們可以擅用 @parent 來讓子頁面的 css 與 js 在我們想要的父區塊裡面，又能有 layout 裡那些我們共用的 css, js，完成我們的最佳實踐。(這邊就跟 OOP 裡，子類別可以在 function 裡加上 parent::functionName(); 一樣。)

```html
<!-- layout.blade.php -->
<!DOCTYPE html>
<html>
    <head>
        <title>Laravel 5 Example - @yield('title')</title>
        <meta charset="utf-8">

        @section('style')
        <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.5/css/bootstrap.min.css">
        @show

    </head>
    <body>
        <header>
        </header>

        <nav>
        </nav>

        <section>
            @yield('content')
        </section>

        <footer>
        </footer>

        @section('script')
        <script src="https://ajax.googleapis.com/ajax/libs/jquery/1.11.3/jquery.min.js"></script>
        <script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.5/js/bootstrap.min.js"></script>
        @show

    </body>
</html>
```

內頁加上 style 和 script section，並用 @parent 讓父類別 section 裡的東西繼承下來，而不去覆寫它。

```html
<!-- index.blade.php -->
@extends('layout')

@section('title', 'index')

@section('style')
@parent
<style>
    div {
        text-align: center;
    }
</style>
@endsection

@section('content')
<div>
    Home Page
</div>
@endsection

@section('script')
@parent
<script>
    $(function() {
    });
</script>
@endsection
```

這樣產生的 view 就會是按我們想要的順序出來囉！

```html
<!DOCTYPE html>
<html>
    <head>
        <title>Laravel 5 Example - index</title>
        <meta charset="utf-8">

        <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.5/css/bootstrap.min.css">
        <style>
            div {
                text-align: center;
            }
        </style>

    </head>
    <body>
        <header>
        </header>

        <nav>
        </nav>

        <section>
            <div>
                Home Page
            </div>
        </section>

        <footer>
        </footer>

        <script src="https://ajax.googleapis.com/ajax/libs/jquery/1.11.3/jquery.min.js"></script>
        <script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.5/js/bootstrap.min.js"></script>
        <script>
            $(function() {
            });
        </script>

    </body>
</html>
```

## PS

在官網 @parent 的例子是用 sidebar，可能是沒有立即有感，發現不多人記得有這個東西，所以寫了這篇來提示一下大家可以這樣用。

## Reference

* [Blade Templates - Laravel](http://laravel.com/docs/5.1/blade)