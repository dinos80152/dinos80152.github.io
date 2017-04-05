(function() {

    marked.setOptions({
      highlight: function(code) {
          return hljs.highlightAuto(code).value;
        }
    });

    renderMarkDown();
    renderNotesList();

    function renderMarkDown() {
        var markdownFileName = getUrlLastUri() + '.md';
        getUrlContent('/notes/' + markdownFileName, function(content) {
            document.getElementById('content').innerHTML = marked(content);
            setTitle(document.getElementById('content').getElementsByTagName('h1')[0].innerText);
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

    function setTitle(titleSuffix) {
        var title = "DinoLai Personal Page"+ titleSuffix;
        document.title = title;
        document.querySelector('meta[property="og:title"]').content = title;
    }

    function renderNotesList() {
        var jsonFile = '/src/notes.json';
        getUrlContent(jsonFile, function(content) {
            var noteList = JSON.parse(content);
            var ul = document.getElementById('notes-list-all');
            noteList.all.forEach(function(note) {
                var li = document.createElement('li');
                li.className = 'list-group-item';
                li.innerText = note;
                ul.appendChild(li);
            });
        });
    }
})();
