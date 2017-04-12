# HTML Input defaultValue vs. value with jQuery val() vs. attr()

## Problem

最近遇到有人想要把現有的 input tag 使用 .val(value) 改值以後，append 到另一個 form 上，結果發現怎麼 input 呈現的 value 沒變？！

但實驗過後 .attr('value', value) 可以成功！

那這個 attribute 和 value 到底哪裡不一樣呢？

讓我們看看以下的例子

## Example

```html

<form><input name="test" type="text" value="abc" />
<input name="test2" type="text" value="abc" /></form>
```

Javascript

```js
document.getElementsByName('test')[0].value = 123;
document.getElementsByName('test')[0].setAttribute('value', 123);
console.log(document.getElementsByTagName('form')[0].innerHTML);
```

jQuery

```js
$('input[name="test"]').val(123);
$('input[name="test"]').attr('value', 123);
console.log($('form').html());
```

Output

```html
<input name="test" type="text" value="abc" />
<input name="test2" type="text" value="123" />
```

我們可以看到改變 attribute 的話元素本身的 html 也會被改變。
但是改變 value 的話，元素本身的 html 是不會被改變的。

## Explain

為什麼呢？

這要從[Input 的屬性](http://www.w3.org/TR/DOM-Level-2-HTML/html.html#ID-6043025)來看，可以看到裡面有 defaultValue 與 value 兩個關於 value 的值。

在 html 一開始 render 的時候，defaultValue 和 value 會是一樣的。

在 input 的欄位打上你想要的值，或是使用 javascript 改變 input 的 value 時，改變的是它的 value，defaultValue 並不會被改變，它依舊會是一開始被 render 時預設的那個值。

使用 value, val() 方法，改變的是蓋在 input 元素上面的東西，而不會真正改變它的本質。

而因為元素的 html 並沒有真正的被改變，所以使用 html 方法時，拿到的就會是它最原始的樣子。

要改變元素的本質，必須更改元素 value 的 attribute，也就是它的 defaultValue，則需要使用 attribute 相關 method: setAttribute, attr()。

如果想要更瞭解，可以查查 value attribute 與 value property。
以為這一切都沒事了嗎？

不，把 type 改成 hidden，value 和 val() 卻都成功改變了 html...

因為這實在是太 geek 了，所以寫另一篇探討。

[HTML Input hidden defaultValue vs. value with jQuery val() vs. attr()](https://dinos80152.wordpress.com/2015/10/07/html-input-hidden-defaultvalue-vs-value-with-jquery-val-vs-attr/)

### Reference

* [Input Text defaultValue Property@W3Schools](http://www.w3schools.com/jsref/prop_text_defaultvalue.asp)
* [Document Object Model HTML-Interface HTMLInputElement](http://www.w3.org/TR/DOM-Level-2-HTML/html.html#ID-6043025)