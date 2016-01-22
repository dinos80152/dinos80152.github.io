(function() {

    marked.setOptions({
      highlight: function(code) {
          return hljs.highlightAuto(code).value;
        }
    });

    renderMarkDown();

    function renderMarkDown() {
        var markdownFileName = getUrlLastUri() + '.md';
        getUrlContent('/notes/' + markdownFileName, function(content) {
            document.getElementById('content').innerHTML = marked(content);

        });
    }

    function getUrlLastUri() {
        var hash = location.hash;
        return hash.split('/').pop();
    }

    function getUrlContent(url, callback) {
        var xhr = new XMLHttpRequest();
        xhr.onload = function() {
            if(callback) {
                callback(this.responseText);
            }
        };
        xhr.open('GET', url, true);
        xhr.send();
    }
})();
