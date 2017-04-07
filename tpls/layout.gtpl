<!DOCTYPE html>
<html>

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <meta name="description" content="Dino Lai Personal Page include blog, tutorial and code">
    <meta name="keywords" content="Programming,System,Coding,Software Develop">
    <meta name="author" content="Dino Lai">
    <title>{{.Title}} - Dino Lai</title>
    <!-- Facebook Open Graph -->
    <meta property="og:url" content="http://dinos80152.github.io/" />
    <meta property="og:type" content="article" />
    <meta property="og:title" content="{{.Title}} - Dino Lai" />
    <meta property="og:description" content="" />
    <meta property="og:image" content="" />
    <!-- End Facebook Open Graph -->
    <!-- Latest compiled and minified CSS -->
    <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap.min.css" integrity="sha384-BVYiiSIFeK1dGmJRAkycuHAHRg32OmUcww7on3RYdg4Va+PmSTsz/K68vbdEjh4u"
        crossorigin="anonymous">
    <!-- Optional theme -->
    <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap-theme.min.css" integrity="sha384-rHyoN1iRsVXV4nD0JutlnGaslCJuC7uwjduW9SVrLvRYooPp2bWYgmgJQIXwl/Sp"
        crossorigin="anonymous">
    <!--<link rel="stylesheet" type="text/css" href="/node_modules/highlight.js/styles/default.css">-->
    <link rel="stylesheet" href="//cdnjs.cloudflare.com/ajax/libs/highlight.js/9.10.0/styles/default.min.css">    
    <link rel="stylesheet" href="/assets/css/layout.css">
    <script src="/assets/js/google-analytics.js"></script>
</head>

<body>
    <nav class="navbar navbar-inverse">
        <div class="container">
            <div class="navbar-header">
                <button type="button" class="navbar-toggle collapsed" data-toggle="collapse" data-target="#navbar" aria-expanded="false"
                    aria-controls="navbar">
                    <span class="sr-only">Toggle navigation</span>
                    <span class="icon-bar"></span>
                    <span class="icon-bar"></span>
                    <span class="icon-bar"></span>
                </button>
                <a class="navbar-brand" href="/">DinoLai Personal Website</a>
            </div>
            <div id="navbar" class="collapse navbar-collapse">
                <div class="row">
                    <ul class="nav nav navbar-nav">
                        <li class="active"><a href="#">Home</a></li>
                        <li><a href="/about">Info</a></li>
                        <li><a href="/notes">Notes</a></li>
                        <li><a href="/toys">Toys</a></li>
                        <li><a href="/slides">Slides</a></li>
                        <li><a href="/books">Books</a></li>
                        <li><a href="/records">Records</a></li>
                        <li><a href="https://www.xmind.net/share/dinos80152/">MindMaps</a></li>
                    </ul>
                    <!--<form class="navbar-form navbar-right" role="search">
                        <div class="form-group">
                            <input type="text" class="form-control" placeholder="Search">
                        </div>
                        <button type="submit" class="btn btn-default">Submit</button>
                    </form>-->
                </div>
            </div>
        </div>
        <!--/.nav-collapse -->
        </div>
    </nav>
    <div class="container">
        <div class="row">
            <div class="col-lg-9">
                <article>
                    <header>{{.LastEdit}}</header>
                    <main>{{.Content}}</main>
                </article>
                <hr>
                <div id="disqus_thread"></div>
            </div>
            <div class="col-lg-3">
                <aside>
                    <section class="panel panel-grey">
                        <div class="panel-heading">Articles</div>
                        <ul id="notes-list-all" class="list-group">

                        </ul>
                    </section>
                    <section class="panel panel-grey">
                        <div class="panel-heading">Tutorial by Code Example</div>
                        <ul id="links-list-all" class="list-group">
                            <li class="list-group-item">
                                <a href="https://github.com/dinos80152/hackme" target="_blank">Hack Me</a>
                            </li>
                            <li class="list-group-item">
                                <a href="https://github.com/dinos80152/php-tutorial" target="_blank">PHP Tutorial</a>
                            </li>
                            <li class="list-group-item">
                                <a href="https://github.com/dinos80152/laravel5-example" target="_blank">Laravel5 Examples</a>
                            </li>
                            <li class="list-group-item">
                                <a href="https://github.com/dinos80152/lairavel" target="_blank">Web Framework Implementation</a>
                            </li>
                        </ul>
                    </section>
                    <section class="panel panel-grey">
                        <div class="panel-heading">Links</div>
                        <ul id="links-list-all" class="list-group">
                            <li class="list-group-item">
                                <a href="https://github.com/dinos80152" target="_blank">Github</a>
                            </li>
                            <li class="list-group-item">
                                <a href="https://www.flickr.com/dinolai" target="_blank">Flickr</a>
                            </li>
                        </ul>
                    </section>
                </aside>
            </div>
        </div>
    </div>
    <footer>
        <p>© Copyright 2015-2017 by Dino Lai. All Rights Reserved.</p>
    </footer>
</body>
<!-- Disqus Block -->
<script>
    var disqus_config = function () {
        this.page.url = location.href; // Replace PAGE_URL with your page's canonical URL variable
        this.page.identifier = location.pathname + location.hash; // Replace PAGE_IDENTIFIER with your page's unique identifier variable
    };
    (function () { // DON'T EDIT BELOW THIS LINE
        var d = document,
            s = d.createElement('script');

        s.src = '//dinolaigithubpersonalwebsite.disqus.com/embed.js';

        s.setAttribute('data-timestamp', +new Date());
        (d.head || d.body).appendChild(s);
    })();

</script>
<noscript>Please enable JavaScript to view the <a href="https://disqus.com/?ref_noscript" rel="nofollow">comments powered by Disqus.</a></noscript>
<script src="https://code.jquery.com/jquery-3.1.1.slim.min.js" integrity="sha384-A7FZj7v+d/sdmMqp/nOQwliLvUsJfDHW+k9Omg/a/EheAdgtzNs3hpfag6Ed950n"
    crossorigin="anonymous"></script>
<!-- Latest compiled and minified JavaScript -->
<script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/js/bootstrap.min.js" integrity="sha384-Tc5IQib027qvyjSMfHjOMaLkfuWVxZxUPnCJA7l2mCWNIpG9mGCD8wGNIcPD7Txa"
    crossorigin="anonymous"></script>
<script src="//cdnjs.cloudflare.com/ajax/libs/highlight.js/9.10.0/highlight.min.js"></script>
<script src="/google-search.js"></script>
<script id="dsq-count-scr" src="//dinolaigithubpersonalwebsite.disqus.com/count.js" async></script>
<script>
    $(document).ready(function () {
        // enable highlight.js
        $('pre code').each(function (i, block) {
            hljs.highlightBlock(block);
        });

        //use bootstrap css
        $("table").addClass("table")
    });
</script>
</body>

</html>