# Google Custom Search

## Integrate with Self Design Search Bar

search-result.html

```html
<body>
    <div>
        <gcse:searchresults-only></gcse:searchresults-only>
    </div>
</body>
```

index.html

```html
<form id="search-form">
    <input id="search-input" type="text"/>
    <input type="submit" value="submit"/>
</form>
```

index.js

```javascript
$("#search-form").submit(function (event) {
    event.preventDefault();
    location.href = location.protocol + "//" + location.host + "/search?q=" + $("#search-input").val();
});
```

### Reference

* [Add Google to your Bootstrap Search Bar](http://alt-web.com/TUTORIALS/?id=add_google_to_bootstrap_search_bar)

## Reference

* [Custom Search Help](https://support.google.com/customsearch)
* [Custom Search Element Control API](https://developers.google.com/custom-search/docs/element)