<!DOCTYPE html>
<html>

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <meta name="description" content="Dino Lai Personal Web Site include Blogs, Tutorials and Code">
    <meta name="keywords" content="Programming,System,Coding,Software Develop">
    <meta name="author" content="Dino Lai">
    <title>{{.Title}} - Dino Lai</title>
    <!-- Facebook Open Graph -->
    <meta property="og:url" content="http://dinolai.com/" />
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
    <link rel="stylesheet" href="/node_modules/github-markdown-css/github-markdown.css">
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
            <div id="navbar" class="navbar-collapse collapse">
                <ul class="nav navbar-nav">
                    <li id="info-link"><a href="/info">Info</a></li>
                    <li id="notes-link"><a href="/notes">Notes</a></li>
                    <li id="toys-link"><a href="/toys">Toys</a></li>
                    <li id="slides-link"><a href="/slides">Slides</a></li>
                    <li id="books-link"><a href="/books">Books</a></li>
                    <li><a href="https://flic.kr/s/aHskVZ1hES" target="_blank">Records</a></li>
                    <li><a href="https://www.xmind.net/share/dinos80152/" target="_blank">MindMaps</a></li>
                </ul>
                <form id="search-form" class="navbar-form navbar-right" role="search">
                    <div class="form-group">
                        <input id="search-input" type="text" class="form-control" placeholder="Search">
                    </div>
                    <button type="submit" class="btn btn-default">Submit</button>
                </form>
            </div>
        </div>
        <!--/.nav-collapse -->
        </div>
    </nav>
    <div class="container">
        <div class="row">
            <div class="col-lg-9">
                <article class="markdown-body">
                    <header>Updated: {{.UpdatedAt}}</header>
                    <main>
                        {{.Content}}
                        <gcse:searchresults-only></gcse:searchresults-only>
                    </main>
                </article>
                <hr>
                <div id="disqus_thread"></div>
            </div>
            <div class="col-lg-3">
                <aside>
                    <section class="panel panel-grey">
                        <div class="panel-heading clickable" data-toggle="collapse" data-target="#notes-list-all" aria-controls="notes-list-all">
                            Notes Categories
                        </div>
                        <ul id="notes-list-all" class="list-group collapse in" aria-expanded="true">
                            <li class="list-group-item">
                                <a href="/notes#development">Development</a>
                            </li>
                            <li class="list-group-item">
                                <a href="/notes#php">PHP</a>
                            </li>
                            <li class="list-group-item">
                                <a href="/notes#laravel">Laravel</a>
                            </li>
                            <li class="list-group-item">
                                <a href="/notes#golang">Golang</a>
                            </li>
                            <li class="list-group-item">
                                <a href="/notes#git">Git</a>
                            </li>
                            <li class="list-group-item">
                                <a href="/notes#database">Database</a>
                            </li>
                            <li class="list-group-item">
                                <a href="/notes#linux">Linux</a>
                            </li>
                            <li class="list-group-item">
                                <a href="/notes#system">System</a>
                            </li>
                            <li class="list-group-item">
                                <a href="/notes#web">Web</a>
                            </li>
                            <li class="list-group-item">
                                <a href="/notes#editor">Editor</a>
                            </li>
                            <li class="list-group-item">
                                <a href="/notes#os">OS</a>
                            </li>
                            <li class="list-group-item">
                                <a href="/notes#others">Others</a>
                            </li>
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
                            <li class="list-group-item">
                                <a href="https://www.linkedin.com/in/dinolai" target="_blank">LinkedIn</a>
                            </li>
                            <li class="list-group-item">
                                <a href="mailto:dinos80152@gmail.com">Mail</a>
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

        s.src = '//dinolai.disqus.com/embed.js';

        s.setAttribute('data-timestamp', +new Date());
        (d.head || d.body).appendChild(s);
    })();

</script>
<noscript>Please enable JavaScript to view the <a href="https://disqus.com/?ref_noscript" rel="nofollow">comments powered by Disqus.</a></noscript>
<script id="dsq-count-scr" src="//dinolai.disqus.com/count.js" async></script>
<!--End Disqus Block-->
<script src="/assets/js/google-search.js"></script>
<script src="https://code.jquery.com/jquery-3.1.1.slim.min.js" integrity="sha384-A7FZj7v+d/sdmMqp/nOQwliLvUsJfDHW+k9Omg/a/EheAdgtzNs3hpfag6Ed950n"
    crossorigin="anonymous"></script>
<!-- Latest compiled and minified JavaScript -->
<script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/js/bootstrap.min.js" integrity="sha384-Tc5IQib027qvyjSMfHjOMaLkfuWVxZxUPnCJA7l2mCWNIpG9mGCD8wGNIcPD7Txa"
    crossorigin="anonymous"></script>
<script src="//cdnjs.cloudflare.com/ajax/libs/highlight.js/9.10.0/highlight.min.js"></script>
<script src="/bower_components/bower-webfontloader/webfont.js"></script>
<script src="/bower_components/snap.svg/dist/snap.svg-min.js"></script>
<script src="/bower_components/underscore/underscore-min.js"></script>
<script src="/bower_components/js-sequence-diagrams/dist/sequence-diagram-min.js"></script>
<script src="https://cdnjs.cloudflare.com/ajax/libs/raphael/2.2.7/raphael.min.js"></script>
<script src="https://cdnjs.cloudflare.com/ajax/libs/flowchart/1.6.6/flowchart.min.js"></script>
<script>
    $(document).ready(function () {
        // enable highlight.js
        $('pre code').each(function (i, block) {
            hljs.highlightBlock(block);
        });

        //use bootstrap css
        $("table").addClass("table table-striped")

        // use jsSequence
        var options = {
            theme: "simple"
        }
        $("pre code.language-sequence").sequenceDiagram(options);

        // use flowchart.js
        document.querySelectorAll("pre code.language-flow").forEach(function (element, index) {
            let md = element.textContent;
            element.innerHTML = "";
            let id = "flowchart" + index;
            element.id = id;
            flowchart.parse(md).drawSVG(id);
        }, this);

        // navbar button active highlight
        $("div#navbar div.row ul.nav li").removeClass("active");
        parentDir = location.pathname.split("/")[1];
        $("#" + parentDir + "-link").addClass("active")

        // integrate google custom search
        $("#search-form").submit(function (event) {
            event.preventDefault();
            location.href = location.protocol + "//" + location.host + "/search?q=" + $("#search-input").val();
        });
    });

</script>
</body>

</html>